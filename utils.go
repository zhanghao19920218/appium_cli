package appium_cli

import (
	"context"
	"fmt"
	"os/exec"
	"runtime"
	"strings"
	"time"
)

func (platform PlatformType) ToString() string {
	var ret string
	switch platform {
	case Android:
		ret = "Android"
	case IOS:
		ret = "IOS"
	case Mac:
		ret = "Mac"
	case Windows:
		ret = "Windows"
	}
	return ret
}

func GetOutPutString(commandShell string, commandList []string) (info string, error *AppiumError) {
	out, err := exec.Command(commandShell, commandList...).Output()
	if err != nil {
		error = &AppiumError{
			Message:   "Get shell output error",
			ErrorCode: OsShellError,
		}
		return
	}
	info = string(out)
	return
}

func NoOutPutString(commandShell string, commandList []string) (error *AppiumError) {
	_, err := exec.Command(commandShell, commandList...).Output()
	if err != nil {
		error = &AppiumError{
			Message:   "Get shell output error",
			ErrorCode: OsShellError,
		}
		return
	}
	return
}

func GetAdbOutputString(commandShell string, commandList []string) (error *AppiumError) {
	cmd := exec.Command(commandShell, commandList...)

	// 执行命令并忽略输出
	err := cmd.Start()
	if err != nil {
		error = &AppiumError{
			Message:   "Get shell output error",
			ErrorCode: OsShellError,
		}
		return
	}
	return
}

// KillLoopCmd
// @Note: this function can not kill subprocess, e,g
// "python3 main.py"
func KillLoopCmd(commandShell string, commandList []string) (ret bool, error *AppiumError) {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	process := exec.CommandContext(ctx, commandShell, commandList...)

	processOutBytes, _ := process.Output()

	result := string(processOutBytes)

	pingList := strings.Split(result, "\n")

	if len(pingList) > 1 {
		ret = true
	}

	//if err != nil {
	//	return
	//}

	if ctx.Err() == context.DeadlineExceeded {
		fmt.Println("Timeout")
	}
	return
}

func GetAdbPath() string {
	if runtime.GOOS == "windows" {
		adbPath, err := exec.LookPath("adb")
		if err != nil {
			fmt.Println("找不到 adb 命令：", err)
			return "D:\\AndroidSDK\\android-sdk_r24.4.1-windows\\android-sdk_r24.4.1-windows\\android-sdk-windows\\platform-tools\\adb.exe"
		}
		return adbPath
	}
	return "adb"
}
