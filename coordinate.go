package appium_cli

// Coordinate
// @Description: Interface for coordinate of the element
// GetX: Get the coordinate of x
// GetY: Get the coordinate of y
// GetDuration: Get the d
type Coordinate interface {
	GetPosition() *ActionChainParams
	GetDuration() int64
}
