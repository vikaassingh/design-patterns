package factorymethod

type LaptopFactory struct {
	Factory
}

func GetLaptopFactory() IFactoryMethod {
	var laptopFactory LaptopFactory

	laptopFactory.SetFactoryName("Macbook")
	laptopFactory.SetFactoryAddress("abc")

	return &laptopFactory
}
