package factorymethod

type MobileFactory struct {
	Factory
}

func GetMobileFactory() IFactoryMethod {
	var mobileFactory MobileFactory

	mobileFactory.SetFactoryName("Iphone")
	mobileFactory.SetFactoryAddress("xyz")

	return &mobileFactory
}
