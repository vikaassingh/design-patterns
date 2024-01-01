package factorymethod

type IFactoryMethod interface {
	SetFactoryName(name string)
	SetFactoryAddress(addr string)
	GetFactoryName() string
	GetFactoryAddress() string
}
