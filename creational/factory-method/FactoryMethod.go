package factorymethod

import "fmt"

func TestFactoryMethod() {
	mobileFactory := GetMobileFactory()
	laptopFactory := GetLaptopFactory()

	fmt.Println("Mobile Factory Name & Address:" + mobileFactory.GetFactoryName() + " " + mobileFactory.GetFactoryAddress() + ", Laptop Factory Name & Address:" + laptopFactory.GetFactoryName() + " " + laptopFactory.GetFactoryAddress())
}
