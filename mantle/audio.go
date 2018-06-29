package mantle


// 
// Audio middle layer
// 

type AudioMiddleLayer struct{
	emuCore  	*core.Core

	audioEnabled bool
}

func NewAudioMiddleLayer(core *core.Core, audioEnabled bool = true) (*AudioMiddleLayer, error){
	return nil, 
}

func (self *AudioMiddleLayer) EnableAudio() error{
	return nil
}

func (self *AudioMiddleLayer) DisableAudio() error{
	return nil
}

func (self *AudioMiddleLayer) GetAudioData() error{
	return nil
}

func (self *Core) SetAudioChannel(channel chan float32) {
	Core.APU.channel = channel
}

func (self *Core) SetAudioSampleRate(sampleRate float64) {
	if sampleRate != 0 {
		// Convert samples per second to cpu steps per sample
		Core.APU.sampleRate = CPUFrequency / sampleRate
		// Initialize filters
		Core.APU.filterChain = FilterChain{
			HighPassFilter(float32(sampleRate), 90),
			HighPassFilter(float32(sampleRate), 440),
			LowPassFilter(float32(sampleRate), 14000),
		}
	} else {
		Core.APU.filterChain = nil
	}
}