package core

type IRunnable interface{
	Run() 		error
	Step(n int) error
	Pause()		error
	Stop()		error	
} 