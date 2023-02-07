package appium_cli

import (
	"github.com/imroc/req/v3"
	"testing"
)

func TestDeviceCapabilityModel_CreateSession(t *testing.T) {
	session := &DeviceCapabilityModel{
		Platform:              Android,
		PlatformVersion:       "9.0",
		DeviceName:            "910KPGS2086327",
		AppPackage:            "com.wechat",
		AppActivity:           "com.wechat",
		NewCommandTimeout:     80000,
		AndroidInstallTimeout: 80000,
		AutomationName:        "UiAutomator2",
		SystemPort:            8002,
		Udid:                  "910KPGS2086327",
		NoReset:               false,
		Port:                  4001,
		Client:                req.C(),
	}
	driver, err := session.CreateSession()
	if err != nil {
		return
	}
	driver.SetKeyboardType(&ImeKeyboardModel{
		AppPackage:  "com.sohu.inputmethod.sogou",
		AppActivity: ".SogouIME",
	})
	// Click the element
	driver.CloseSession()
}
