package core

import(
	"io"
)
type IRuntime interface{
	Run() 		error
	Step(n int) error
	Pause()		error
	Stop()		error	

	LoadGameImage(io.Reader)	error
	
	SaveGameRecord(io.Writer)	error
	LoadGameRecord(io.Reader)	error
} 