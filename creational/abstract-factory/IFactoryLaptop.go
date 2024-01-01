package abstractfactory

type IFactoryLaptop interface {
	SetName(string)
	SetBrand(string)
	GetName() string
	GetBrand() string
}
