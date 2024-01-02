package decorator

import "fmt"

func TestDecorator() {
	icecream := &OrangeCone{}
	conTop := icecream.GetIngredientDecorate(&ChocolateCone{}).GetIngredientDecorate(&VenillaScoop{}).GetIngredientDecorate(&Topping{}).GetIngredientDecorate(&Cashew{}).GetIngredientDecorate(&Almond{}).GetIngredientDecorate(&Raisins{})
	conTop.GetIngredientDecorate(&Peanuts{}).GetIngredientDecorate(&VenillaScoop{})

	fmt.Println("Price:", icecream.GetPrice(), ", Description:", icecream.GetDescription())
}
