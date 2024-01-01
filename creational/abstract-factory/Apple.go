package abstractfactory

type Apple struct{}

func (a *Apple) GetMobile() IFactoryMobile {
	return &AppleMobile{}
}

func (a *Apple) GetLaptop() IFactoryLaptop {
	return &AppleLaptop{}
}
