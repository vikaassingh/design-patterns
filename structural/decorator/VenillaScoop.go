package decorator

type VenillaScoop struct {
	Decorator IIcecreamDecorator
}

func (c *VenillaScoop) GetIngredientDecorate(decorator IIcecreamDecorator) IIcecreamDecorator {
	c.Decorator = decorator
	return c.Decorator
}

func (c *VenillaScoop) GetPrice() int {
	if c.Decorator == nil {
		return 10
	}

	return c.Decorator.GetPrice() + 10
}

func (c *VenillaScoop) GetDescription() string {
	if c.Decorator == nil {
		return "Venilla Scoop"
	}

	return c.Decorator.GetDescription() + ", Venilla Scoop"
}
