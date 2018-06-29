package mantle

import (
	"errors"
	"os"
	Core "github.com/MSLaojun/ArcadeGo/core"
)

type RuntimeMT struct{
	coreRuntime		Core.IRuntime
	// game getter
	// record getter
}

func NewRuntime(core interface{}) (*RuntimeMT, error){
	runtimeMT := &RuntimeMT{}
	coreRuntime, err := core.(IRuntime)
	if err != nil {
		// logger.log
		return nil, errors.New("Error to load runtime from this core.")
	}
	runtimeMT.coreRuntime = coreRuntime
	return runtime, nil
}

func (self *RuntimeMT)Run() error{
	return self.coreRuntime.Run()
}

func (self *RuntimeMT)Step(n int) error{
	return self.coreRuntime.Step(n)
}

func (self *RuntimeMT)Pause() error{
	return self.coreRuntime.Pause()
}

func (self *RuntimeMT)Stop() error{
	return self.coreRuntime.Stop()
}

func (self *RuntimeMT)LoadGame() error{
	f , err := os.OpenFile("./game.image", os.O_CREATE | os.O_WRONLY)
	defer f.Close()
	if err != nil {
		return errors.New("Error to load game image.")
	}
	err = self.coreRuntime.LoadGameImage(f)
	if err != nil {
		return errors.New("Error to load game image.")
	}
	return nil
}

// simply implemented save record
func (self *RuntimeMT)SaveRecord() error{
	f , err := os.OpenFile("./game.sav", os.O_CREATE | os.O_WRONLY)
	defer f.Close()
	if err != nil {
		return errors.New("Error to save game record.")
	}
	err = self.coreRuntime.SaveRecord(f)
	if err != nil {
		return errors.New("Error to save game record.")
	}
	return nil
}

func (self *RuntimeMT)LoadRecord() error{
	f , err := os.OpenFile("./game.sav", os.O_CREATE | os.O_WRONLY)
	defer f.Close()
	if err != nil {
		return errors.New("Error to load game record.")
	}
	err = self.coreRuntime.LoadRecord(f)
	if err != nil {
		return errors.New("Error to load game record.")
	}
	return nil
}

func (self *RuntimeMT)Run() error{
	return self.coreRuntime.Run()
}

// func RunSync(){
// 	// when run sync, the runtime will wait for command
// 	// and wil go according to command 
// 	// then give caller the picture / vars
// }
// func RunAsync(){
// 	// when run async, the runtime will not wait for anyone 
// 	// caller will only get the "current" state of runtime
// }
// func ChangeRunMode(){

// }


