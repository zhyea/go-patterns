package main

import "fmt"

func main() {

	veggiePizza := &veggieMania{}

	//Add cheese topping
	veggiePizzaWithCheese := &cheeseTopping{
		pizza: veggiePizza,
	}

	//Add tomato topping
	veggiePizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: veggiePizzaWithCheese,
	}

	fmt.Printf("Price of veggieMania pizza with tomato and cheese topping is %d\n", veggiePizzaWithCheeseAndTomato.getPrice())

	peppyTofuPizza := &peppyTofu{}

	//Add cheese topping
	peppyTofuPizzaWithCheese := &cheeseTopping{
		pizza: peppyTofuPizza,
	}

	fmt.Printf("Price of peppyTofu with tomato and cheese topping is %d\n", peppyTofuPizzaWithCheese.getPrice())

}
