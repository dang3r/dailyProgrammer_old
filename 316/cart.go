package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"
)

// Prices for each tour
var costs = map[string]float64{
	"OH": 300.00,
	"BC": 110.00,
	"SK": 30.00,
}

type cart struct {
	order []string
	tours map[string]int
	rules []rule
}

func (c *cart) AddTour(tour string) {
	c.tours[tour] += 1
}

func (c *cart) Total() float64 {
	total := 0.00
	for tour, size := range c.tours {
		total += costs[tour] * float64(size)
	}
	for _, rule := range c.rules {
		total -= rule(c.tours)
	}
	return total
}

func (c *cart) String() string {
	output := ""
	for _, item := range c.order {
		output += item + ", "
	}
	output += "= "
	output += fmt.Sprintf("%.2f", c.Total())
	return output
}

// Given a cart's inventory, return the discount applied for your rule
type rule func(tours map[string]int) float64

func main() {
	rules := []rule{
		func(tours map[string]int) float64 {
			num := tours["OH"]
			return float64(num/3) * costs["OH"]
		},
		func(tours map[string]int) float64 {
			numOH := tours["OH"]
			numSK := tours["SK"]
			return math.Min(float64(numSK), float64(numOH)) * costs["SK"]
		},
		func(tours map[string]int) float64 {
			if tours["BC"] > 4 {
				return 20.00 * float64(tours["BC"])
			}
			return 0.00
		},
	}
	carts := make([]cart, 0, 0)

	// Read input
	reader := bufio.NewReader(os.Stdin)
	for {
		text, err := reader.ReadString('\n')
		if err != nil {
			break
		}

		text = text[:len(text)-1]
		items := make(map[string]int)
		order := make([]string, 0, 0)
		for _, item := range strings.Split(text, " ") {
			order = append(order, item)
			items[item] += 1
		}
		c := cart{
			order: order,
			rules: rules,
			tours: items,
		}
		carts = append(carts, c)
	}

	// Output
	fmt.Println("Items\t\tTotal")
	for _, c := range carts {
		fmt.Println(&c)
	}
}
