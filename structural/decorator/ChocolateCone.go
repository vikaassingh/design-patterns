package decorator

type ChocolateCone struct {
	Decorator IIcecreamDecorator
}

func (c *ChocolateCone) GetIngredientDecorate(decorator IIcecreamDecorator) IIcecreamDecorator {
	c.Decorator = decorator
	return c.Decorator
}

func (c *ChocolateCone) GetPrice() int {
	if c.Decorator == nil {
		return 10
	}

	return c.Decorator.GetPrice() + 10
}

func (c *ChocolateCone) GetDescription() string {
	if c.Decorator == nil {
		return "Chocolate Cone"
	}

	return c.Decorator.GetDescription() + ", Chocolate Cone"
}
