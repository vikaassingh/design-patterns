package decorator

type IIcecreamDecorator interface {
	GetPrice() int
	GetDescription() string
	GetIngredientDecorate(IIcecreamDecorator) IIcecreamDecorator
}
