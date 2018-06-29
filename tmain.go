package main

// import (
// 	"core/Core"
// )

import (
	"fmt"
	"errors"
)

type ICore interface{
	Configure(map[string]string) error
	GetInformation(string)	(string, error) 
}

type MockCore struct{
	version 	string
}

func (self *MockCore) Configure(conf map[string] string) error{
	if v, ok := conf["version"]; ok != true {
		fmt.Println("Cannot find version in conf!")
		return errors.New("Bad conf file")
	}else{
		self.version = v
	}

	return nil
}

func (self *MockCore) GetInformation(key string) (string, error) {
	if key != "version" {
		fmt.Println("Unknown config key : ", key)
		return "", errors.New("Bad parameters")
	}
	
	return self.version, nil
}

func main(){
	core := &MockCore{}

	conf := make(map[string]string)
	
	var C ICore = core
	
	C.Configure(conf)

	conf["version"] = "1.2.3.4"
	C.Configure(conf)

	str, _ := C.GetInformation("XXXXXXX")
	str, _ = C.GetInformation("version")
	fmt.Println("Config value of version is :", str)
}