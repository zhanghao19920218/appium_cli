package appium_cli

import (
	"fmt"
	sf "github.com/sa-/slicefunk"
	"strings"
	"time"
)

// CreateSession Create the appium new session to do the test, Get the screen size of device
func (capModel *DeviceCapabilityModel) CreateSession() (deviceModel *DeviceDriverModel, serverErr *AppiumError) {
	// Make the struct to the desired capabilities
	model := &AppiumParameter{DesiredCapabilities: DesiredCapabilities{
		PlatformName:          capModel.Platform.ToString(),
		PlatformVersion:       capModel.PlatformVersion,
		DeviceName:            capModel.DeviceName,
		AppPackage:            capModel.AppPackage,
		AppActivity:           capModel.AppActivity,
		NewCommandTimeout:     capModel.NewCommandTimeout,
		AndroidInstallTimeout: capModel.AndroidInstallTimeout,
		AutomationName:        capModel.AutomationName,
		SystemPort:            capModel.SystemPort,
		Udid:                  capModel.Udid,
		NoReset:               capModel.NoReset,
	}}
	var result SessionResponse
	var errorResult SessionErrorResponse

	client := capModel.Client

	resp, err := client.R().
		SetBody(model).
		SetSuccessResult(&result).
		SetErrorResult(&errorResult).
		Post(fmt.Sprintf("http://127.0.0.1:%d/wd/hub/session", capModel.Port))
	if err != nil {
		serverErr = &AppiumError{
			Message:   err.Error(),
			ErrorCode: CreateSessionError,
		}
		return
	}

	if !resp.IsSuccessState() {
		serverErr = &AppiumError{
			Message:   errorResult.Value.Message,
			ErrorCode: CreateSessionError,
		}
		return
	}
	deviceModel = &DeviceDriverModel{
		SessionId:  result.SessionId,
		Port:       capModel.Port,
		DeviceName: capModel.DeviceName,
		Client:     client}
	return
}

// CloseSession close the appium session
func (driver DeviceDriverModel) CloseSession() (serverErr *AppiumError) {
	resp, err := driver.Client.R().
		Delete(fmt.Sprintf("http://127.0.0.1:%d/wd/hub/session/%s", driver.Port, driver.SessionId))
	if err != nil {
		serverErr = &AppiumError{
			Message:   "Stop Session Error",
			ErrorCode: StopSessionError,
		}
		return
	}

	if !resp.IsSuccessState() {
		serverErr = &AppiumError{
			Message:   "Stop Session Error",
			ErrorCode: StopSessionError,
		}
		return
	}
	return
}

// FindElement
//
//	@Description: Find the element id
//	@receiver driver
//	@param param
//	@return elementId
//	@return serverErr
func (driver DeviceDriverModel) FindElement(param *FindElementPoint) (elementId string, serverErr *AppiumError) {
	var result ElementResponse

	resp, err := driver.Client.R().
		SetBody(&FindElementParam{
			Using: param.GetUsingType(),
			Value: param.Value,
		}).
		SetSuccessResult(&result).
		Post(fmt.Sprintf("http://127.0.0.1:%d/wd/hub/session/%s/element", driver.Port, driver.SessionId))
	if err != nil {
		serverErr = &AppiumError{
			Message:   "Find Element Error",
			ErrorCode: NotFoundElement,
		}
		return
	}

	if !resp.IsSuccessState() {
		serverErr = &AppiumError{
			Message:   "Create Session Error",
			ErrorCode: NotFoundElement,
		}
		return
	}
	elementId = result.Value.ELEMENT
	return
}

// ActionElement take action the device element
func (driver DeviceDriverModel) ActionElement(elementParam *ActionNormalParam, action ActionType) (serverErr *AppiumError) {
	var result SessionResponse
	var body any
	var requestUrl string
	if action == SendKeys {
		body = &SendKeysParam{Text: elementParam.Text}
		requestUrl = fmt.Sprintf("http://127.0.0.1:%d/wd/hub/session/%s/element/%s/value", driver.Port, driver.SessionId, elementParam.Element)
	} else {
		requestUrl = fmt.Sprintf("http://127.0.0.1:%d/wd/hub/session/%s/element/%s/click", driver.Port, driver.SessionId, elementParam.Element)
		body = &ActionRequestParam{Element: elementParam.Element}
	}
	resp, err := driver.Client.R().
		SetBody(body).
		SetSuccessResult(&result).
		Post(requestUrl)
	if err != nil {
		serverErr = &AppiumError{
			Message:   "Action Element Error",
			ErrorCode: ActionElementError,
		}
		return
	}

	if !resp.IsSuccessState() {
		serverErr = &AppiumError{
			Message:   "Create Session Error",
			ErrorCode: ActionElementError,
		}
		return
	}
	return
}

// TouchActionByLoc make the location x and y to action
func (driver DeviceDriverModel) TouchActionByLoc(coordinate Coordinate) (serverErr *AppiumError) {
	// Create the parameters
	var actions []ActionRequestChain
	actions = append(actions, ActionRequestChain{
		Type:     "pointerMove",
		Duration: coordinate.GetDuration(),
		X:        coordinate.GetPosition().X,
		Y:        coordinate.GetPosition().Y,
	})
	actions = append(actions, ActionRequestChain{
		Type:   "pointerDown",
		Button: 0,
	})
	actions = append(actions, ActionRequestChain{
		Type:   "pointerUp",
		Button: 0,
	})
	pointerParam := ActionRequestParams{
		PointerType: "touch",
	}
	actionParam := ActionsRequest{
		Actions:    actions,
		Parameters: pointerParam,
		Id:         "finger1",
		Type:       "pointer",
	}
	var actionTemp []ActionsRequest
	actionTemp = append(actionTemp, actionParam)
	requestParams := &ActionRequestArr{
		Actions: actionTemp,
	}
	var result SessionResponse

	resp, err := driver.Client.R().
		SetBody(requestParams).
		SetSuccessResult(&result).
		Post(fmt.Sprintf("http://127.0.0.1:%d/wd/hub/session/%s/actions", driver.Port, driver.SessionId))
	if err != nil {
		serverErr = &AppiumError{
			Message:   "Touch Action Error",
			ErrorCode: TouchActionError,
		}
		return
	}

	if !resp.IsSuccessState() {
		serverErr = &AppiumError{
			Message:   "Touch Action Error",
			ErrorCode: TouchActionError,
		}
		return
	}
	return
}

// StartActivity
//
//	@Description: Start the activity of the another application
//	@receiver driver
//	@param param
//	@return serverErr
func (driver DeviceDriverModel) StartActivity(param *StartActivityParam) (serverErr *AppiumError) {
	var result SessionResponse

	resp, err := driver.Client.R().
		SetBody(param).
		SetSuccessResult(&result).
		Post(fmt.Sprintf("http://127.0.0.1:%d/wd/hub/session/%s/appium/device/start_activity", driver.Port, driver.SessionId))
	if err != nil {
		serverErr = &AppiumError{
			Message:   "Start Activity Error",
			ErrorCode: StartActivityError,
		}
		return
	}

	if !resp.IsSuccessState() {
		serverErr = &AppiumError{
			Message:   "Start Activity Error",
			ErrorCode: StartActivityError,
		}
		return
	}
	return
}

// ImplicitWait
//
//	@Description: Set the amount of time the driver should wait when searching for elements
//	@receiver driver
//	@param seconds
//	@return serverErr
func (driver DeviceDriverModel) ImplicitWait(seconds time.Duration) (serverErr *AppiumError) {
	var result SessionResponse

	resp, err := driver.Client.R().
		SetBody(&ImplicitWaitParam{
			Seconds: int(seconds / time.Millisecond),
		}).
		SetSuccessResult(&result).
		Post(fmt.Sprintf("http://127.0.0.1:%d/wd/hub/session/%s/timeouts/implicit_wait", driver.Port, driver.SessionId))
	if err != nil {
		serverErr = &AppiumError{
			Message:   "ImplicitWait Time Out",
			ErrorCode: ImplicitWaitError,
		}
		return
	}

	if !resp.IsSuccessState() {
		serverErr = &AppiumError{
			Message:   "ImplicitWait Time Out",
			ErrorCode: ImplicitWaitError,
		}
		return
	}
	return
}

// ElementActionMov
//
//		@Description: Clicks element at its center point. If the element's center point is obscured by another element,
//		an element click intercepted error is returned. If the element is outside the viewport, an element not interactable
//		error is returned. Not all drivers automatically scroll the element into view and may need to be scrolled to in order to interact with it.
//		@receiver driver
//		@param param the find way of the element
//	 @param seconds the seconds to implicitWait
//	 @param action take the action
//		@return serverErr
func (driver DeviceDriverModel) ElementActionMov(param *FindElementPoint, seconds time.Duration, action ActionType, sendKeys string) (elementId string, serverErr *AppiumError) {
	// 1. Confirm to find element
	if seconds != 0 {
		serverErr = driver.ImplicitWait(seconds)
	} else {
		serverErr = driver.ImplicitWait(5)
	}
	if serverErr != nil {
		return
	}
	// 2. Find the element
	elementId, serverErr = driver.FindElement(param)
	if serverErr != nil {
		return
	}
	// 3. Touch or move the element
	serverErr = driver.ActionElement(&ActionNormalParam{
		Element: elementId,
		Text:    sendKeys,
	}, action)
	return
}

// GetAttribute
//
//	@Description: Get the element attribute
//	@receiver driver
//	@param param
//	@param elementId
//	@return serverErr
func (driver DeviceDriverModel) GetAttribute(param *AttributeModel, element *FindElementPoint) (value AttributeInterface, elementId string, serverErr *AppiumError) {
	// 1. Find the element
	elementId, serverErr = driver.FindElement(element)
	if serverErr != nil {
		return
	}

	var result AttributeResponse

	resp, err := driver.Client.R().
		SetSuccessResult(&result).
		Get(fmt.Sprintf("http://127.0.0.1:%d/wd/hub/session/%s/element/%s/attribute/%s", driver.Port, driver.SessionId, elementId, param.GetAttributeStr()))
	if err != nil {
		serverErr = &AppiumError{
			Message:   "Not Found the attribute",
			ErrorCode: NotFoundAttribute,
		}
		return
	}

	if !resp.IsSuccessState() {
		serverErr = &AppiumError{
			Message:   "Not Found the attribute",
			ErrorCode: NotFoundAttribute,
		}
		return
	}
	value = &AttributeRetModel{Value: result.Value}
	return
}

// TerminateApp terminate the application
func (driver DeviceDriverModel) TerminateApp(appId string) (ret bool, serverErr *AppiumError) {
	var result TerminateResponse
	resp, err := driver.Client.R().
		SetBody(&AppPropParam{
			AppId: appId,
		}).
		SetSuccessResult(&result).
		Post(fmt.Sprintf("http://127.0.0.1:%d/wd/hub/session/%s/appium/device/terminate_app", driver.Port, driver.SessionId))
	if err != nil {
		serverErr = &AppiumError{
			Message:   "Terminate the app error",
			ErrorCode: TerminalAppError,
		}
		return
	}

	if !resp.IsSuccessState() {
		serverErr = &AppiumError{
			Message:   "Terminate the app error",
			ErrorCode: TerminalAppError,
		}
		return
	}
	ret = result.Value
	return
}

// FindInputMethods Get the input-method
func (driver DeviceDriverModel) FindInputMethods() (imeKeyboards []ImeKeyboardModel, err *AppiumError) {
	args := []string{
		"-s",
		driver.DeviceName,
		"shell",
		"ime",
		"list",
		"-s",
	}
	out, err := GetOutPutString("adb", args)
	if err != nil {
		return
	}
	devices := strings.Split(strings.TrimSpace(out), "\n")
	imeKeyboards = sf.Map(devices, func(item string) ImeKeyboardModel {
		return StrConvertImeModel(item)
	})
	fmt.Println(imeKeyboards)
	return
}

// SetKeyboardType Set the keyboard
func (driver DeviceDriverModel) SetKeyboardType(imeKeyboard *ImeKeyboardModel) (err *AppiumError) {
	args := []string{
		"-s",
		driver.DeviceName,
		"shell",
		"ime",
		"enable",
		imeKeyboard.ToString(),
	}
	err = NoOutPutString("adb", args)
	if err != nil {
		return
	}
	args = []string{
		"-s",
		driver.DeviceName,
		"shell",
		"ime",
		"set",
		imeKeyboard.ToString(),
	}
	err = NoOutPutString("adb", args)
	return
}
