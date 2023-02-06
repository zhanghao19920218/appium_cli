package appium_cli

import "strconv"

type AttributeRetModel struct {
	Value string
}

func (model *AttributeRetModel) ToString() string {
	return model.Value
}

func (model *AttributeRetModel) ToBool() bool {
	ret, err := strconv.ParseBool(model.Value)
	if err != nil {
		return false
	}
	return ret
}
