package mantle

import (
	core "github.com/deepnes/core"
)
// 
// Audio middle layer
// 

type GraphicsMiddleLayer struct{
	emuCore  	*core.Core

	graphicsEnabled bool
}

func NewAudioMiddleLayer(core *core.Core, graphicsEnabled bool = true) (*AudioMiddleLayer, error){
	return nil, 
}

func (self *AudioMiddleLayer) EnableGraphics() error{
	return nil
}

func (self *AudioMiddleLayer) DisableGrahics() error{
	return nil
}

func (self *AudioMiddleLayer) GetScreenShot() error{
	return nil
}


func (self *Core) Buffer() *image.RGBA {
	return Core.PPU.front
}

func (self *Core) BackgroundColor() color.RGBA {
	return Palette[Core.PPU.readPalette(0)%64]
}