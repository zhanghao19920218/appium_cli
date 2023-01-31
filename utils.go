package appium_cli

func (platform PlatformType) ToString() string {
	var ret string
	switch platform {
	case Android:
		ret = "Android"
	case IOS:
		ret = "IOS"
	case Mac:
		ret = "Mac"
	case Windows:
		ret = "Windows"
	}
	return ret
}
