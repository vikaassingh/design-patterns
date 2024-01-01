package abstractfactory

type IFactoryMobile interface {
	SetName(string)
	SetBrand(string)
	GetName() string
	GetBrand() string
}
