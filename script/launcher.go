package main

// simply prototyped launcher script for whole cell -- from upper script to create core, load game, load MT and facade and then give controll to cell itself, good for test

import(
	"os"
	"os/signal"
	"plugin"
	"errors"
	"flag"
	"fmt"
	"time"
	"bufio"
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
	runtimeMT, err := Mantle.NewRuntimeMT(core)
	if err != nil 
	graphicMT, err := Mantle.NewGraphicMT(core)
	audioMT, err := Mantle.NewRuntimeMT(core)
	controllerMT, err := Mantle.NewControllerMT(core)
	// bootstrap for front

	sigChan := make(chan os.Signal, 10)
	signal.Notify(sigChan)

	cmdChan := make(chan []byte, 1)
	go receiveCmd(cmdChan)

	startTime := time.Now()
	for {
		select{
		case <-sigChan:
			fmt.Println("Ctrl + C hit, gracefully quit")
		case cmdLine := <-cmdChan:
			processCmd(cmdLine)
		case <-time.After(5*time.Minute):
			hereNow := time.Now()
			fmt.Println("Elapsed : ", hereNow - startTime)
		}
	}
	fmt.Println("Bye~")
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

func receiveCmd(C chan []byte){
	inputReader := bufio.NewReader(os.Stdin)
	
	for {
		input, err := inputReader.ReadBytes((byte)"\n")
		if err != nil {
			fmt.Println("Error read cmd!", err)
			continue
		}
		inputReader <- input
	}
}

