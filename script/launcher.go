package main

// simply prototyped launcher script for whole cell -- from upper script to create core, load game, load MT and facade and then give controll to cell itself, good for test

import(
	"os"
	"os/signal"
	"plugin"
	"errors"
	"flag"
	"fmt"
	Core "github.com/MSLaojun/ArcadeGo/core"
	Mantle "github.com/MSLaojun/ArcadeGo/mantle"
)

var platformPluginMap map[string]string
func init(){
	platformPluginMap["nes"] = "nes.so"
}
var platformName = flag.String("platform", "null", "platform name")
var imageName = flag.String("image", "null", "platform name")
var recordName = flag.String("record", "null", "platform name")

func main(){
	flag.Parse()

	// bootstrap for core
	core, err := loadPlugin(*platformName)
	if err != nil {
		fmt.Println("Error loading core plugin.")
		os.Exit(1)
	}
	
	// bootstrap for MT -- let MT do game status handling
	
	// bootstrap for front

	sigChan := make(chan os.Signal, 10)
	signal.Notify(sigChan)
	
}
func loadPlugin(plat string) (interface{}, error){
	pluginFileName, ok := platformPluginMap[plat]
	if ok != true{
		return nil, errors.New("Unsupported platform : " + plat)
	}
	pdll, err := plugin.Open(pluginFileName)
	if err != nil{
		return nil, err
	}
	coreinst, err := pdll.Lookup("CoreInstance")
	return coreinst, err
}
func waitForQuit(){
	
}