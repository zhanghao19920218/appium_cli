package appium_cli

import (
	"fmt"
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
		SessionId: result.SessionId,
		Port:      capModel.Port,
		Client:    client}
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
		SetBody(param).
		SetSuccessResult(&FindElementParam{
			Using: param.GetUsingType(),
			Value: param.Value,
		}).
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
func (driver DeviceDriverModel) ActionElement(elementParam *ActionRequestParam, action ActionType) (serverErr *AppiumError) {
	var result SessionResponse

	resp, err := driver.Client.R().
		SetBody(elementParam).
		SetSuccessResult(&result).
		Post(fmt.Sprintf("http://127.0.0.1:%d/wd/hub/session/%s/element/%s/click", driver.Port, driver.SessionId, elementParam.Element))
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
