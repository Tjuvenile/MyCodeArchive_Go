package rpc

import (
	"MyCodeArchive_Go/utils/tool"
	"bufio"
	"errors"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"io"
	"log"
	"net"
	"os"
)

var (
	port = flag.Int("port", 29895, "The server port")
)

const (
	ConfPath        = "/etc/ganesha/ganesha.conf"
	StorageConfPath = "/etc/storage/nfs/ganesha.conf"
	StdOut          = "stdout"
	StdErr          = "stderr"
)

func (rpc *ganeshaRpc) ReadLocalConfig(req *ReadLocalConfigRequest, svc Ganesha_ReadLocalConfigServer) error {
	log.Printf("start read local config and send the config info to client...")
	conf, err := os.Open(StorageConfPath)
	if err != nil {
		log.Printf("open file failed...")
		return err
	}
	defer conf.Close()

	reader := bufio.NewReader(conf)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			log.Printf("The config file read over")
			break
		}
		if err != nil {
			log.Printf("read config file failed")
			return err
		}

		err = svc.Send(&ReadLocalConfigRes{Data: line})
		if err != nil {
			log.Printf("send config stream failed")
			return err
		}
	}
	log.Printf("config file send successfully")
	return err
}

func (rpc *ganeshaRpc) SaveConfigToLocal(svc Ganesha_SaveConfigToLocalServer) error {
	log.Printf("start config sync")

	file, err := os.OpenFile(ConfPath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		log.Printf("OpenFile error: %v", err)
		return err
	}
	defer file.Close()

	for {
		stream, err := svc.Recv()
		if err == io.EOF {
			log.Printf("recv data over")
			break
		}

		if err != nil {
			log.Printf("recv data from client failed: %v", err)
			return err
		}

		_, err = file.Write([]byte(stream.Data))
		if err != nil {
			log.Printf("write data to conf failed: %v", err)
			return err
		}
	}

	log.Printf("move conf from %s to %s", ConfPath, StorageConfPath)
	ret, err := tool.ExecLocalCommand("mv", fmt.Sprintf("%s %s", ConfPath, StorageConfPath))
	if err != nil {
		log.Printf("move conf from %s to %s failed, %v", ConfPath, StorageConfPath, err)
		return errors.New(ret[StdErr])
	}

	log.Printf("config sync successfully")
	return err
}

func main() {
	flag.Parse()

	logFile, err := os.OpenFile("/var/log/storage/nfs/ganesha_rpc.log",
		os.O_CREATE|os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("open log file failed")
		os.Exit(1)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Ldate | log.Ltime)

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)
	RegisterGaneshaServer(grpcServer, &UnimplementedGaneshaServer{})
	log.Printf("listening on port %v", *port)
	grpcServer.Serve(lis)
}
