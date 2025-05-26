package judge

import (
	"bytes"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"time"
)

func Judge(sourceCode, inputTest, expectedOutput string) string {
	// 创建临时C源文件
	sourceFile, err := ioutil.TempFile("", "*.c")
	if err != nil {
		return "Error creating source file"
	}
	defer os.Remove(sourceFile.Name()) // 确保函数结束时删除临时文件

	// 写入C代码到源文件
	if _, err := sourceFile.WriteString(sourceCode); err != nil {
		return "Error writing source code"
	}
	sourceFile.Close() // 关闭文件以确保写入完成

	// 编译C源代码
	executable := sourceFile.Name() + ".out"
	cmdCompile := exec.Command("gcc", sourceFile.Name(), "-o", executable)
	if err := cmdCompile.Run(); err != nil {
		return "Compilation Error"
	}
	defer os.Remove(executable) // 确保函数结束时删除编译产物

	// 运行编译后的程序
	cmdRun := exec.Command(executable)
	cmdRun.Stdin = strings.NewReader(inputTest)
	var output bytes.Buffer
	cmdRun.Stdout = &output

	// 设置执行时间限制
	done := make(chan error, 1)
	go func() {
		done <- cmdRun.Run()
	}()

	// 检查程序执行是否超时
	select {
	case <-time.After(2 * time.Second): // 比如设置2秒的超时时间
		cmdRun.Process.Kill()
		return "Time Limit Exceeded"
	case err := <-done:
		if err != nil {
			return "Runtime Error"
		}
	}

	// 比较程序输出和预期输出（统一换行符后进行比较）
	outputString := strings.Replace(output.String(), "\r\n", "\n", -1)
	expectedOutput = strings.Replace(expectedOutput, "\r\n", "\n", -1)
	if outputString == expectedOutput {
		return "Accepted"
	} else {
		return "Wrong Answer"
	}
}
