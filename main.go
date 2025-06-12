package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go <amt> <fromC> <toC>")
		return
	}
	amt := os.Args[1]
	fromC := strings.ToUpper(os.Args[2])
	toC := strings.ToUpper(os.Args[3])

	amount, err := valid(amt, fromC, toC)
	if err != nil {
		fmt.Println(err)
		return
	}

	res, err := covert(amount, fromC, toC)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(" %.2f  %s  is equal to  %.2f %s \n", amt, fromC, res, toC)
}

func valid(amt, fromC, toC string) (float64, error) {
	amount, err := strconv.ParseFloat(amt, 64)
	if err != nil {
		return 0, err
	}
	if amount <= 0 {
		return 0, fmt.Errorf("amount must be greater than 0")
	}
	if !isvalidPair(fromC, toC) {
		return 0, fmt.Errorf("Currency not found")
	}
	return amount, nil
}
func isvalidPair(fromC, toC string) bool {
	if fromC == "USD" && (toC == "INR" || toC == "EUR" || toC == "JPY") {
		return true
	}
	if fromC == "EUR" && (toC == "INR" || toC == "USD" || toC == "JPY") {
		return true
	}
	if fromC == "INR" && (toC == "USD" || toC == "EUR" || toC == "JPY") {
		return true
	}
	if fromC == "JPY" && (toC == "INR" || toC == "EUR" || toC == "USD") {
		return true
	}
	return false
}
func covert(amount float64, fromC, toC string) (float64, error) {
	var rate float64
	if fromC == "USD" && toC == "INR" {
		rate = 83.12
	} else if fromC == "USD" && toC == "EUR" {
		rate = 0.92
	} else if fromC == "USD" && toC == "JPY" {
		rate = 155.60
	} else if fromC == "EUR" && toC == "USD" {
		rate = 1.09
	} else if fromC == "EUR" && toC == "INR" {
		rate = 90.43
	} else if fromC == "EUR" && toC == "JPY" {
		rate = 169.00
	} else if fromC == "INR" && toC == "USD" {
		rate = 0.012
	} else if fromC == "INR" && toC == "EUR" {
		rate = 0.011
	} else if fromC == "INR" && toC == "JPY" {
		rate = 1.87
	} else if fromC == "JPY" && toC == "USD" {
		rate = 0.0064
	} else if fromC == "JPY" && toC == "EUR" {
		rate = 0.0059
	} else if fromC == "JPY" && toC == "INR" {
		rate = 0.53
	} else {
		return 0, fmt.Errorf("conversion not possible beteween %s and %s", fromC, toC)
	}
	return amount * rate, nil
}
