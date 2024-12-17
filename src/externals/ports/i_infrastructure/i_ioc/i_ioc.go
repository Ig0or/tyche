package i_ioc

type IIoc interface {
	LoadProviders()
	Invoke(function func())
}
