package main

import (
	"errors"
	"fmt"
	"os"
)

type PizzaBase string

const (
	ThinCrust    PizzaBase = "Thin Crust"
	ThickCrust   PizzaBase = "Thick Crust"
	StuffedCrust PizzaBase = "Stuffed Crust"
)

type PizzaSauce string

const (
	Tomato  PizzaSauce = "Tomato"
	BBQ     PizzaSauce = "BBQ"
	Alfredo PizzaSauce = "Alfredo"
)

type PizzaSize string

const (
	Small  PizzaSize = "Small"
	Medium PizzaSize = "Medium"
	Large  PizzaSize = "Large"
)

type PizzaTopping string

const (
	Pepperoni PizzaTopping = "Pepperoni"
	Mushrooms PizzaTopping = "Mushrooms"
	Olives    PizzaTopping = "Olives"
	Onions    PizzaTopping = "Onions"
	Bell      PizzaTopping = "Bell"
	Peppers   PizzaTopping = "Peppers"
)

type Pizza struct {
	base     PizzaBase
	sauce    PizzaSauce
	size     PizzaSize
	toppings []PizzaTopping
}

func (p Pizza) String() string {
	toppingsStr := ""
	for _, topping := range p.toppings {
		toppingsStr += string(topping) + ", "
	}
	if len(p.toppings) > 0 {
		toppingsStr = toppingsStr[:len(toppingsStr)-2]
	}

	return fmt.Sprintf("Base: %s, Sauce: %s, Size: %s, Toppings: %s",
		p.base, p.sauce, p.size, toppingsStr)
}

type Builder struct {
	pizza Pizza
}

func (b *Builder) NewBuilder() *Builder {
	b.pizza = Pizza{}
	return b
}

func (b *Builder) Base(base PizzaBase) *Builder {
	b.pizza.base = base
	return b
}

func (b *Builder) Sauce(sauce PizzaSauce) *Builder {
	b.pizza.sauce = sauce
	return b
}

func (b *Builder) Topping(topping PizzaTopping) *Builder {
	b.pizza.toppings = append(b.pizza.toppings, topping)
	return b
}

func (b *Builder) Size(size PizzaSize) *Builder {
	b.pizza.size = size
	return b
}

func (b *Builder) GetPizza() (*Pizza, error) {
	if b.pizza.base == "" || b.pizza.sauce == "" || b.pizza.size == "" {
		return nil, errors.New(fmt.Sprintf("Called 'GetPizza()' on unfinished pizza: %v", b.pizza))
	}
	return &b.pizza, nil
}

func main() {
	var b Builder

	b.NewBuilder().Base(ThinCrust).Topping(Pepperoni).Topping(Mushrooms).Topping(Olives).Sauce(Tomato).Size(Large)
	pizza, err := b.GetPizza()
	if err != nil {
		fmt.Printf("Error while getting pizza: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Your Pizza is ready: %v\n", pizza)
}
