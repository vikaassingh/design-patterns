package abstractfactory

type Google struct{}

func (a *Google) GetMobile() IFactoryMobile {
	return &GoogleMobile{}
}

func (a *Google) GetLaptop() IFactoryLaptop {
	return &GoogleLaptop{}
}
