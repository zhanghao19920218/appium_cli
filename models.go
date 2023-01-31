package appium_cli

import "github.com/imroc/req/v3"

type PlatformType int64

// Appium platforms support Android, IOS, Mac, Windows...
const (
	Android PlatformType = 0
	IOS                  = 1
	Mac                  = 2
	Windows              = 3
)

type ActionType int64

// Different Action have different actions
const (
	Click ActionType = iota
	Press
)

type DeviceInfoResponse struct {
	DeviceScreenSize string `json:"deviceScreenSize"`
}

type SessionResponse struct {
	SessionId string             `json:"sessionId"`
	Status    int64              `json:"status"`
	Value     DeviceInfoResponse `json:"value,omitempty"`
}

type ValueErrorMsg struct {
	Message string `json:"message"`
}

type SessionErrorResponse struct {
	SessionResponse
	Value ValueErrorMsg `json:"value"`
}

// DesiredCapabilities create the appium desired capability
type DesiredCapabilities struct {
	PlatformName          string `json:"platformName"`
	PlatformVersion       string `json:"platformVersion"`
	DeviceName            string `json:"deviceName"`
	AppPackage            string `json:"appPackage"`
	AppActivity           string `json:"appActivity"`
	NewCommandTimeout     uint64 `json:"newCommandTimeout"`
	AndroidInstallTimeout uint64 `json:"androidInstallTimeout"`
	AutomationName        string `json:"automationName"`
	SystemPort            uint64 `json:"systemPort"`
	Udid                  string `json:"udid"`
	NoReset               bool   `json:"noReset"`
	//App                   string `json:"app"`
}

// DeviceCapabilityModel user start the device capability
type DeviceCapabilityModel struct {
	Platform              PlatformType
	PlatformVersion       string `json:"platformVersion"`
	DeviceName            string `json:"deviceName"`
	AppPackage            string `json:"appPackage"`
	AppActivity           string `json:"appActivity"`
	NewCommandTimeout     uint64 `json:"newCommandTimeout"`
	AndroidInstallTimeout uint64 `json:"androidInstallTimeout"`
	AutomationName        string `json:"automationName"`
	SystemPort            uint64 `json:"systemPort"`
	Udid                  string `json:"udid"`
	NoReset               bool   `json:"noReset"`
	Port                  int64
	Client                *req.Client
}

type AppiumParameter struct {
	DesiredCapabilities DesiredCapabilities `json:"desiredCapabilities"`
}

type DeviceDriverModel struct {
	SessionId string
	Client    *req.Client
	Port      int64
}

// FindElementParam request post the parameters
type FindElementParam struct {
	Using string `json:"using"`
	Value string `json:"value"`
}

// ElementResponse get the element response
type ElementResponse struct {
	SessionId string            `json:"sessionId"`
	Status    int64             `json:"status"`
	Value     ElementValueModel `json:"value"`
}

type ElementValueModel struct {
	ELEMENT string `json:"ELEMENT"`
}

type ActionRequestParam struct {
	Element string `json:"element"`
}

type ActionChainParams struct {
	X        int64
	Y        int64
	Duration int64
}

type ActionRequestChain struct {
	Type     string `json:"type"`
	Duration int64  `json:"duration,omitempty"`
	X        int64  `json:"x,omitempty"`
	Y        int64  `json:"y,omitempty"`
	Button   int64  `json:"button,omitempty"`
}

type ActionRequestParams struct {
	PointerType string `json:"pointerType"`
}

type ActionsRequest struct {
	Actions    []ActionRequestChain `json:"actions"`
	Parameters ActionRequestParams  `json:"parameters"`
	Id         string               `json:"id"`
	Type       string               `json:"type"`
}

type ActionRequestArr struct {
	Actions []ActionsRequest `json:"actions"`
}
