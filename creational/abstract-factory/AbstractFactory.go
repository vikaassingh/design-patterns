package abstractfactory

import "fmt"

func TestAbstractFactory() {
	var apple Apple
	var google Google

	appleLaptop := apple.GetLaptop()
	appleLaptop.SetBrand("Macbook")
	appleMobile := apple.GetMobile()
	appleMobile.SetBrand("Iphone")

	googleLaptop := google.GetLaptop()
	googleLaptop.SetName("Google")
	googleMobile := google.GetMobile()
	googleMobile.SetName("Google")

	fmt.Println(appleLaptop.GetBrand(), appleMobile.GetBrand(), googleLaptop.GetName(), googleMobile.GetName())
}
