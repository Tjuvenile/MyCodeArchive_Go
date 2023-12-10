package rpc

import (
	"MyCodeArchive_Go/db"
	"MyCodeArchive_Go/fault"
	"MyCodeArchive_Go/logging"
	"context"
	"errors"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"time"
)

type ganeshaRpc struct {
	serverAddr string
	client     GaneshaClient
	connection *grpc.ClientConn
}

const GaneshaRpcServerPort = 50007

func GaneshaSvc(funcName string, callReq interface{}) *fault.Fault {
	logging.Log.Infof("exec rpc func %s, callReq:%+v", funcName, callReq)
	manageIps, err := db.ListRpcIp()
	if err != nil {
		return err
	}
	if len(manageIps) == 0 {
		return fault.RpcIpEmpty()
	}

	for _, manageIp := range manageIps {
		client, err := GetRpcClient(manageIp)
		if err != nil {
			return err
		}
		// 调用rpc service相关功能
		err = client.CallSvc(funcName, callReq)
		if err != nil {
			client.CloseConnection()
			return err
		}
		client.CloseConnection()
	}
	return err
}

func GetRpcClient(addr string) (*ganeshaRpc, *fault.Fault) {
	var opts []grpc.DialOption
	opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))

	// 如果访问不到这个ip，就进行重试访问
	retry := 0
	for retry < 15 {
		conn, err := grpc.Dial(fmt.Sprintf("%s:%d", addr, GaneshaRpcServerPort), opts...)
		if err != nil {
			logging.Log.Errorf("connect ganesha rpc server %s:%d failed, %v", addr, GaneshaRpcServerPort, err)
			retry += 1
			time.Sleep(2 * time.Second)
			continue
		}
		return &ganeshaRpc{serverAddr: addr, client: NewGaneshaClient(conn), connection: conn}, nil
	}
	logging.Log.Errorf("retry %d to connect rpc server failed", retry)
	return nil, fault.NetworkReachable
}

func (rpc *ganeshaRpc) CallSvc(funcName string, req interface{}) *fault.Fault {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Minute)
	defer cancel()

	logging.Log.Infof("start exec call ganesha svc function, funcName: %s, param: %+v", funcName, req)
	err := rpc.checkReq(funcName, req)
	if err != nil {
		logging.Log.Errorf("failed to parse req parameters, %s, %s, %v", funcName, rpc.serverAddr, err)
		return fault.ParseData
	}

	reply := &BaseReply{}
	switch funcName {
	case "Example":
		reply, err = rpc.client.AddExportConf(ctx, req.(*AddExportConfRequest))
	}
	// 当网络不可达时，比如连接某个ip，但pod正在重启，或者rpc超时，就会导致reply为nil。但其实此时的err是有值的。
	// 因为需要对error进行分别判断，函数里需要用到reply.Result等参数，如果为nil会直接空指针，所以提前判断一下。
	if reply == nil {
		return fault.RpcConnect(rpc.serverAddr)
	}

	return rpc.checkError(reply.Result, reply.ErrDescription, funcName, err)
}

func (rpc *ganeshaRpc) checkReq(funcName string, req interface{}) (err error) {
	var ok bool
	switch funcName {
	case "Example":
		_, ok = req.(*AddExportConfRequest)
	default:
		return errors.New(fmt.Sprintf("not found the rpc func name %q", funcName))
	}
	if !ok {
		return errors.New("the request parameters do not match")
	}
	return nil
}

func (rpc *ganeshaRpc) checkError(result, errDsp, funcName string, err error) *fault.Fault {
	if err != nil {
		logging.Log.Errorf("%s on node %s fail, %v", funcName, rpc.serverAddr, err)
		return fault.CmdExec(funcName, rpc.serverAddr)
	}

	// 有可能并没有报错，但是rpc服务端那边返回的状态码是错误的，所以进行分别判断一下
	if result == "Fail" {
		logging.Log.Errorf("command execution failed.%s result is %s, ErrDescription: %v", funcName, result, errDsp)
		return fault.CmdExec(funcName, rpc.serverAddr)
	}
	return nil
}

func (rpc *ganeshaRpc) CloseConnection() {
	if err := rpc.connection.Close(); err != nil {
		logging.Log.Errorf("close connection to ganesha rpc server(%s) failed", rpc.serverAddr)
	}
}

// SyncGaneshaConfig 通过rpc stream，来进行同步conf
func SyncGaneshaConfig(ctx context.Context, activeIp string, ipsToBeAdded []string) *fault.Fault {
	oldNodeRpc, faultErr := GetRpcClient(activeIp)
	if faultErr != nil {
		logging.Log.Errorf("connect ganesha rpc server %s failed, %v", activeIp, faultErr)
		return faultErr
	}
	defer oldNodeRpc.CloseConnection()

	logging.Log.Infof("start read local config")
	for _, newNodeIp := range ipsToBeAdded {
		// get the old node stream to read `ganesha.conf` file.
		fileStream, err := oldNodeRpc.client.ReadLocalConfig(ctx, &ReadLocalConfigRequest{})
		if err != nil {
			logging.Log.Errorf("call ReadLocalConfig function failed, %v, %s", err, newNodeIp)
			return fault.CallRpc(newNodeIp, "ReadLocalConfig", err)
		}

		// Get a new node stream to send old node's `ganesha.conf` file.
		newNodeRpc, faultErr := GetRpcClient(newNodeIp)
		if err != nil {
			logging.Log.Errorf("call ganesha rpc server %s failed, %v", newNodeIp, faultErr)
			return faultErr
		}
		newNodeStream, err := newNodeRpc.client.SaveConfigToLocal(ctx)
		if err != nil {
			logging.Log.Errorf("call SaveConfigToLocal function failed, %v", err)
			newNodeRpc.CloseConnection()
			return fault.CallRpc(newNodeIp, "SaveConfigToLocal", err)
		}

		// `recv data from old node` and `send data to new node`
		for {
			lineData, err := fileStream.Recv()
			if err == io.EOF {
				logging.Log.Infof("data recv over, close send")
				break
			}
			if err != nil {
				logging.Log.Errorf("data recv failed, %v", err)
				newNodeRpc.CloseConnection()
				return fault.CallRpc(newNodeIp, "SaveConfigToLocal", err)
			}

			err = newNodeStream.Send(&SaveConfigToLocalRequest{Data: lineData.GetData()})
			if err != nil {
				logging.Log.Errorf("Send data to server failed, %v", err)
				newNodeRpc.CloseConnection()
				return fault.CallRpc(newNodeIp, "Stream.Send", err)
			}
		}

		// 接收来自服务端的响应
		for {
			logging.Log.Infof("start close send and recv response")
			_, err := newNodeStream.CloseAndRecv()
			if err == io.EOF {
				logging.Log.Infof("close send successfully, %v", err)
				break
			}
			if err != nil {
				logging.Log.Errorf("Close send stream failed, %v", err)
				newNodeRpc.CloseConnection()
				return fault.CallRpc(newNodeIp, "Stream.Send", err)
			}
		}
		newNodeRpc.CloseConnection()
	}
	logging.Log.Infof("sync ganesha conf successfully")
	return nil
}
