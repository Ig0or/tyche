package i_ioc_config

type IIocConfig interface {
	LoadProviders()
	Invoke(function func())
}
