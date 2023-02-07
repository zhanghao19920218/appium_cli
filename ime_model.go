package appium_cli

import (
	"fmt"
	"strings"
)

type ImeKeyboardModel struct {
	AppPackage  string
	AppActivity string
}

func StrConvertImeModel(line string) ImeKeyboardModel {
	ret := strings.Split(line, "/")
	return ImeKeyboardModel{
		AppPackage:  ret[0],
		AppActivity: ret[1],
	}
}

func (model *ImeKeyboardModel) ToString() string {
	return fmt.Sprintf("%s/%s", model.AppPackage, model.AppActivity)
}
