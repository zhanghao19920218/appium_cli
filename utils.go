package appium_cli

import (
	"os/exec"
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
