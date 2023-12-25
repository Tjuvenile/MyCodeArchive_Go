package fault

import (
	"MyCodeArchive_Go/utils/logging"
	"fmt"
)

type Code string

type Fault struct {
	// Comp indicates the group or family for the fault
	// (e.g. storage, network, etc.) It is used as a prefix
	// when displaying the fault.
	Comp string
	// Code is the unique numeric identifier for known faults.
	Code Code
	//fault level:warning/error
	FaultLevel string
	// DescriptionEn is the main description of the fault. It usually
	// includes the reason for the fault, and therefore it is not
	// necessary to display both DescriptionEn and Reason.
	DescriptionEn string
	DescriptionCh string
}

var (
	ParseData = CreateFault(
		"error",
		"parse data fail",
		"解析JSON数据失败",
	)
	ConnectDB = CreateFault(
		"error",
		"Fail to create db connect session",
		"建立数据库连接会话失败",
	)
	QueryRecord = CreateFault(
		"error",
		"Failed to query task",
		"查询任务记录失败",
	)
	NetworkReachable = CreateFault(
		"error",
		"network is not reachable",
		"网络不可达",
	)
	MatchName = CreateFault(
		"error",
		"The 'name' must start with a letter or a digit, and can only contain letters, digits, hyphens, and underscores",
		"'name' 只能以字母或数字开头，且仅包含字母、数字、中划线、下划线",
	)
	InsertRecord = CreateFault(
		"error",
		"Failed to insert table records",
		"插入表记录失败",
	)
	DeleteRecord = CreateFault(
		"error",
		"Failed to delete table records",
		"删除表记录失败",
	)
	UpdateRecord = CreateFault(
		"error",
		"Failed to update table records",
		"更新表记录失败",
	)
)

// Err 当有error类型时，需要封装一层，返回faultErr，并且打印相关信息
func Err(msg string, err error, faultErr *Fault) *Fault {
	logging.Log.Errorf("%s, err: %v", msg, err)
	return faultErr
}

// Wrap 对Fault的封装
func Wrap(enMessage, chMessage string) *Fault {
	logging.Log.Errorf(enMessage)
	return CreateFault(
		"error",
		fmt.Sprintf("%s", enMessage),
		fmt.Sprintf("%s", chMessage),
	)
}

func CreateFault(faultLevel, descEn, descCh string) *Fault {
	return &Fault{
		FaultLevel:    faultLevel,
		DescriptionEn: descEn,
		DescriptionCh: descCh,
	}
}

func RpcIpEmpty() *Fault {
	logging.Log.Errorf("rpc ip cannot be empty")
	return CreateFault(
		"error",
		fmt.Sprintf("rpc ip cannot be empty"),
		fmt.Sprintf("无法连接rpc"),
	)
}

func InvalidParam(paramName, value string) *Fault {
	logging.Log.Errorf("invalid parameter %s, value: %s", paramName, value)
	return CreateFault(
		"error",
		fmt.Sprintf("invalid parameter %s, value: %s", paramName, value),
		fmt.Sprintf("参数%s不是有效的值, 值: %s", paramName, value),
	)
}

// ParamLen 用来进行参数长度的校验，paramName是参数名称，isLong是来判断此参数是太长了，还是太短了。 以及区间是[low, high]之间
func ParamLen(paramName string, isLong bool, low, high int) *Fault {
	rangeNum := fmt.Sprintf("[%d,%d]", low, high)
	if isLong {
		logging.Log.Errorf("%q param is too long, the length must be %s", paramName, rangeNum)
		return CreateFault(
			"error",
			fmt.Sprintf("%q param is too long, the length must be %s", paramName, rangeNum),
			fmt.Sprintf("%q参数超过长度限制,长度需在%s区间内", paramName, rangeNum),
		)
	} else {
		logging.Log.Errorf("parameter %q is empty", paramName)
		return CreateFault(
			"error",
			fmt.Sprintf("parameter %q is empty", paramName),
			fmt.Sprintf("参数%q为空", paramName),
		)
	}
}

func ParamEmpty(paramName string) *Fault {
	logging.Log.Errorf("%s parameter cannot be empty", paramName)
	return CreateFault(
		"error",
		fmt.Sprintf("parameter %q is empty", paramName),
		fmt.Sprintf("参数%q为空", paramName),
	)
}

// 数量超过最大限制的error信息
func ParamCount(paramName string, high int) *Fault {
	logging.Log.Errorf("%s exceed the maximum limit of %d", paramName, high)
	return CreateFault(
		"error",
		fmt.Sprintf("%s exceed the maximum limit of %d", paramName, high),
		fmt.Sprintf("%s的数量超过了最大限制: %d", paramName, high),
	)
}

func ParseIp(ip string, err error) *Fault {
	logging.Log.Errorf("invalid input ip %s, err: %v", ip, err)
	return CreateFault(
		"error",
		"failed to parse ip address",
		"无法解析ip地址",
	)
}

func NotExist(paramName, name string) *Fault {
	logging.Log.Errorf("%s %q does not exist", paramName, name)
	return CreateFault(
		"error",
		fmt.Sprintf("%s %q does not exist", paramName, name),
		fmt.Sprintf("%s %s不存在", paramName, name),
	)
}

func Existed(paramName, name string) *Fault {
	logging.Log.Errorf("%s %q existed", paramName, name)
	return CreateFault(
		"error",
		fmt.Sprintf("%s %q existed", paramName, name),
		fmt.Sprintf("%s %s已存在", paramName, name),
	)
}

func RpcConnect(addr string) *Fault {
	logging.Log.Errorf("connect %q failed, reply is nil", addr)
	return CreateFault(
		"error",
		fmt.Sprintf("network connection to %q failed", addr),
		fmt.Sprintf("网络连接失败，%q无法访问", addr),
	)
}

func CmdExec(cmd, hostname string) *Fault {
	return CreateFault(
		"error",
		fmt.Sprintf("Execute cmd %s on %s failed", cmd, hostname),
		fmt.Sprintf("在节点%s上执行命令%s失败", hostname, cmd),
	)
}

func CallRpc(nodeName string, funcName string, err error) *Fault {
	return CreateFault(
		"error",
		fmt.Sprintf("call %s %s rpc func faile, err: %v", nodeName, funcName, err),
		fmt.Sprintf("调用%s %s rpc函数失败，err: %v", nodeName, funcName, err),
	)
}

func QueryColEmpty(colName string) *Fault {
	logging.Log.Errorf("query for the %s is empty", colName)
	return CreateFault(
		"error",
		fmt.Sprintf("query for the %s is empty", colName),
		fmt.Sprintf("查询%s为空", colName),
	)
}
