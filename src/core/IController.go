package core

type IController interface{
	// Hot key instructions
	GetHotKeyInfo()	(string, error)

	// Hot key map
	GetHotKeyMap()	(map[string]string, error)
	SetHotKeyMap(map[string]string)	error

	// Controll to hot key
	SetControll(map[string]string)	error
}