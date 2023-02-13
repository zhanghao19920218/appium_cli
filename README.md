# Appium Client By Golang

> This project just support golang appium(still in beta) 

## Installation
> go get github.com/zhanghao19920218/appium_cli@v0.1.35

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
	Value: "com.wechat.hello:id/editText",
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
```

### StartActivity

> **Start an Android activity by providing package name and activity name**:
>
> * `AppPackage`: The package of application
> * `AppActivity`: The activity of application

```go
sessionId.StartActivity(&StartActivityParam{
		AppPackage:  "com.wechat",
		AppActivity: "com.wechat",
	})
```

### FindElement

> Find the element id by accessibility-id or id
>
> * `AppiumBy`: ID, AccessibilityID, Xpath, Selector
> * `Value`: String value

```go
elementId, err := session.FindElement(&FindElementPoint{
		AppiumBy: ID,
		Value:    "com.another.baidu:id/editText",
	})
```

### ImplicitWait

> Set the amount of time the driver should wait when searching for elements
>
> * millisecond: wait for the seconds to find element. 

```go
session.ImplicitWait(seconds) // This is seconds to wait
```

### ElementActionMov

> Clicks element at its center point. If the element's center point is obscured by another element, an element click intercepted error is returned. If the element is outside the viewport, an element not interactable error is returned. Not all drivers automatically scroll the element into view and may need to be scrolled to in order to interact with it.

```go
session.ElementActionMov(param *FindElementPoint, seconds time.Duration, action ActionType)
```

### TerminateApp

> Terminate the application
>
> * param appId: Android is app package

### FindInputMethods

> Find the keyboard input-methods return

### SetKeyboardType

> Set the keyboard, Like Google Keyboard Or Sogou Keyboard

### GetElementText

### GetNetworkStatus
