package mantle

import (
	"errors"
	Core "github.com/MSLaojun/ArcadeGo/core"
)

type GraphicMT struct{
	coreGraphic		Core.IGraphic
}

func NewGraphicMT(core interface{}) (*GraphicMT, error){
	graphicMT := &GraphicMT{}
	coreGraphic, ok := core.(Core.IGraphic)
	if ok != true {
		// logger.log
		return nil, errors.New("Error to load graphic component from this core.")
	}
	graphicMT.coreGraphic = coreGraphic
	return graphicMT, nil
}

func (self *GraphicMT)EnableGraphic() error{
	return errors.New("Not implemented!")
}

func (self *GraphicMT)DisableGraphic() error{
	return errors.New("Not implemented!")
}

// Get infomation of frame
func (self *GraphicMT)GetFrameBuffer() ([]byte, error){
	return self.coreGraphic.GetFrameBuffer()
}

func (self *GraphicMT)GetFrameId() (int, error){
	return self.coreGraphic.GetFrameId()
}