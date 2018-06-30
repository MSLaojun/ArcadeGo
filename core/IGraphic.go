package core

type IGraphic interface{
	GetFrameBuffer() ([]byte, error)
	GetFrameId() (int, error)
}