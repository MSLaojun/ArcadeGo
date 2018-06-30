package mantle

import (
	"errors"
	Core "github.com/MSLaojun/ArcadeGo/core"
)

type AudioMT struct{
	coreAudio		Core.IAudio
}

func NewAudioMT(core interface{}) (*AudioMT, error){
	audioMT := &AudioMT{}
	coreAudio, ok := core.(Core.IAudio)
	if ok != true {
		// logger.log
		return nil, errors.New("Error to load audio component from this core.")
	}
	audioMT.coreAudio = coreAudio
	return audioMT, nil
}

func (self *AudioMT)EnableAudio() error{
	return errors.New("Not implemented!")
}

func (self *AudioMT)DisableAudio() error{
	return errors.New("Not implemented!")
}