package base_tool

import (
	"MyCodeArchive_Go/utils/fault"
	"MyCodeArchive_Go/utils/logging"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

/* 记录命令相关的函数 */

// ExecLocalCommand Cmd 需要系统有这个命令，才能去执行
//
// exec.Command.run会生成一个执行"mv"命令的新进程，产生额外的开销
// cmd.Run()及cmd.OutPut()本身都会执行cmd.Wait()，它会等待命令执行完之后，再进行下一步。 Run()只会返回程序执行本身的错误，OutPut()会返回执行后的输出
// Run()返回的程序执行的错误和命令执行的错误是两码事，需要通过cmd.StdErr来获取标准错误，cmd.Stdout来获取标准输出
// （需要注意有些时候，程序的错误会打到标准输出里，这取决于对方报错的方式。比如在python中，如果用print把错误打印出来，就会被stdOut捕捉，如果用error打印出来，就会被stdErr捕捉）
// 如果错误被打印到了标准输入中，除非能改变源头报错的方式，否则只能通过字符串匹配来判断是否有错误
//
// Stderr和Stdout只是来设置标准输出和标准错误，但它不会自动捕获命令的输出，如果需要自动捕获，需要用cmd.StdoutPipe()
// 意味着如果你直接打印cmd.Stdout是不会得到值的，它只是设置，标准输出往stdout里去存
//
// buffer是一个动态数组，实现了一个可变大小字节的字节缓冲区，方便接收数据
//
// example: 第一个参数为我们要执行的命令，后边为它的参数。
// ret, err = ExecLocalCommand("ganesha_conf", fmt.Sprintf("set EXPORT Export_ID %s", req.ExportID))
func ExecLocalCommand(command string, args string) (map[string]string, error) {
	cmd := exec.Command(command, strings.Split(args, " ")...)
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	err := cmd.Run()

	outStr := stdout.String()
	errStr := stderr.String()
	ret := map[string]string{"stdout": outStr, "stderr": errStr}
	logging.Log.Infof("execute command %s %s, stdout:%s, stderr:%s, cmd:%s", command, args, outStr, errStr, cmd)
	if err != nil {
		logging.Log.Errorf("run command %s %s failed, \n %v", command, args, err)
	}
	return ret, err
}

// MvCmd 需要系统有mv这个命令，可能在所有非类Unix操作系统上都不可用，捕捉错误也需要通过cmd本身来捕获
//
// 如果dstFile已经存在的话，会报错
func MvCmd() {
	_, _ = ExecLocalCommand("mv", fmt.Sprintf("%s %s", "srcPath", "dstPath"))
}

// Rename go的内置函数，是跨平台的，只要支持Go，就能支持它
// 直接返回错误，直接进行函数调用，不需要像cmd那样启动新进程
//
// 它要求源文件和目标文件必须要在同一个文件系统下，否则会遇到invalid cross-device link问题。
// 如果需要在不同文件系统下移动文件，可以通过复制和删除的方式。或者使用filepath.Walk来进行遍历文件并逐个移动
func Rename() {
	srcFile := "old_file.txt"
	dstFile := "new_file.txt"
	err := os.Rename(srcFile, dstFile)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}

// PrintFailedMsg cli中如果失败了，通过这个函数进行打印失败原因
func PrintFailedMsg(reason string, jsonFormat bool) {
	if jsonFormat {
		errInfo := map[string]string{
			"Result": "Failed",
			"Reason": reason,
		}
		output, _ := json.MarshalIndent(errInfo, "", "\t")
		fmt.Println(string(output))
	} else {
		fmt.Printf("Result: Failed\nReason: %s\n", reason)
	}
}

// CheckPort 检测某个端口是否被正在被监听
func CheckPort(ip string, port int) *fault.Fault {
	cmd := exec.Command("nc", "-zv", ip, strconv.Itoa(port))

	output, err := cmd.CombinedOutput()
	logging.Log.Infof("exec nc -zv %s %s, result: %s", ip, port, output)
	if err != nil {
		return fault.Err(fmt.Sprintf("network connection failed for IP %s and port %s", ip, port), err, fault.CmdExec("nc -zv", ip))
	}
	return nil
}

// AutoWiki 自动生成大包wiki
func AutoWiki(day, pkgVersion, pkgTime, ceastorBranch string) {
	// 替换的变量
	packageName := fmt.Sprintf("CeaStor_3.2.1-%s_debug-%s_x86_64.tar.gz、CeaStor_3.2.1-%s_debug-%s_aarch64.tar.gz", pkgVersion, pkgTime, pkgVersion, pkgTime)
	tag := fmt.Sprintf("tag-CeaStor-3.2.1-%s-%s", pkgVersion, pkgTime)
	x86Package := fmt.Sprintf("CeaStor_3.2.1-%s_debug-%s_x86_64.tar.gz", pkgVersion, pkgTime)
	// 原始文本
	rawText := `ceastor出包成功
%s 前合入
包名: %s
底座: %s
tag: %s
71: /ceastor/CI_master
取包: 
smbclient -c 'get 01_ceastor/01_master/3.2.1/ceastor-3.2.1-%s-debug/amd64/%s' //samba.cestc.cn/存储开发部/ -U %s -t 120
取包报错: Connection to samba.cestc.cn failed
解决: vi /etc/hosts     # 加上: 10.32.43.2 samba.cestc.cn
`
	// 替换变量
	formattedText := fmt.Sprintf(rawText,
		day,
		packageName,
		ceastorBranch,
		tag,
		pkgVersion,
		fmt.Sprintf("%s %s", x86Package, x86Package),
		"sdp%sdp@123",
	)
	// 指定文件名
	fileName := "text"
	// 打开文件，如果文件不存在则创建新文件，以覆盖的方式写入
	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	// 将文本写入文件
	_, err = file.WriteString(formattedText)
	if err != nil {
		fmt.Println("Error writing to file:", err)
	} else {
		fmt.Println("Text written to file successfully.")
	}
}
