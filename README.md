# Appium Client By Golang

> This project just support golang appium(still in beta)

## Usage
```go
package appium_cli

// Create the session
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
// Get the session id
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
// Click the element
sessionId.ActionElement(&ActionRequestParam{Element: elementId}, Click)
time.Sleep(1 * time.Second)
// Click the element by location
sessionId.TouchActionByLoc(&ActionChainParams{
		X:        49,
		Y:        1215,
		Duration: 1,
})
// Close the session
sessionId.CloseSession()
```
### `CreateSession`
```go
// Platform: Android || IOS || Mac || Windows
// PlatformVersion: as the appium doc api
session := &DeviceCapabilityModel{
		Platform:              Android,
		PlatformVersion:       "7.0.0",
		DeviceName:            "192.168.58.101:5555",
		AppPackage:            "com.iflytek.inputmethod.imehook",
		AppActivity:           "com.iflytek.inputmethod.imehook.MainActivity",
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
```