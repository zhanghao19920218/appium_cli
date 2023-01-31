package appium_cli

import (
	"fmt"
	"github.com/imroc/req/v3"
	"testing"
	"time"
)

func TestDeviceCapabilityModel_CreateSession(t *testing.T) {
	session := &DeviceCapabilityModel{
		Platform:              Android,
		PlatformVersion:       "7.0.0",
		DeviceName:            "192.168.58.101:5555",
		AppPackage:            "com.wechat",
		AppActivity:           "com.wechat",
		NewCommandTimeout:     8000,
		AndroidInstallTimeout: 10000,
		AutomationName:        "UiAutomator2",
		SystemPort:            8002,
		Udid:                  "192.168.58.101:5555",
		NoReset:               false,
		Port:                  4001,
		Client:                req.C(),
	}
	sessionId, err := session.CreateSession()
	if err != nil {
		fmt.Println(err.Message)
		return
	}
	fmt.Println(sessionId.SessionId)
	// Find the element
	elementId, err := sessionId.FindElement(&FindElementParam{
		Using: "id",
		Value: "com.iflytek.inputmethod.imehook:id/editText",
	})
	if err != nil {
		fmt.Println(err.Message)
		return
	}
	fmt.Println(elementId)
	sessionId.ActionElement(&ActionRequestParam{Element: elementId}, Click)
	time.Sleep(1 * time.Second)
	sessionId.TouchActionByLoc(&ActionChainParams{
		X:        49,
		Y:        1215,
		Duration: 1,
	})
	// Click the element
	sessionId.CloseSession()
}
