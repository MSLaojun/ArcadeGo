package core

type ICore interface{
	Configure(map[string] string) error
	GetInformation(string)	(string, error) 
}