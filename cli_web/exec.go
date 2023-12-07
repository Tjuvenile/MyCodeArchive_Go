package cli_web

import (
	"MyCodeArchive_Go/logging"
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"strings"
)

// Cmd 需要系统有这个命令，才能去执行
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
