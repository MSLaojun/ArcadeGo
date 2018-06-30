package mantle

import (
	"errors"
	Core "github.com/MSLaojun/ArcadeGo/core"
)

type ControllerMT struct{
	coreController	Core.IController
	enabled			bool
}

func NewControllerMT(core interface{}) (*ControllerMT, error){
	controllerMT := &ControllerMT{
		enabled : true,
	}
	coreController, ok := core.(Core.IController)
	if ok != true {
		// logger.log
		return nil, errors.New("Error to load controller component from this core.")
	}
	controllerMT.coreController = coreController
	return controllerMT, nil
}

func (self *ControllerMT)EnableController() error{
	if (self.enabled == false){
		self.enabled = true
	}
	return nil
}

func (self *ControllerMT)DisableController() error{
	if (self.enabled == true){
		self.enabled = false
	}
	return nil
}

func (self *ControllerMT)SetControll(map[string]string)	error{
	return nil
}

func (self *ControllerMT)GetHotKeyInfo()	(string, error){
	return "", nil
}

func (self *ControllerMT)GetHotKeyMap()	(map[string]string, error){
	return nil, nil
}

func (self *ControllerMT)SetHotKeyMap(map[string]string)	error{
	return nil
}