package ioc_interface

type IocInterface interface {
	LoadProviders()
	Invoke(function any)
}
