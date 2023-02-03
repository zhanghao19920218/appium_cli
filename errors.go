package appium_cli

type AppiumErrorType int64

const (
	CreateSessionError AppiumErrorType = iota
	StopSessionError
	NotFoundElement
	ActionElementError
	TouchActionError
	StartActivityError
	Others
)

type AppiumError struct {
	Message   string          `json:"message"`
	ErrorCode AppiumErrorType `json:"error_code"`
}

func (sErr *AppiumError) Error() string {
	return sErr.Message
}
