package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	str "strings"

	ex "currency/exchange"
	cli "currency/cli"
)

var exchanges ex.Exchanges = ex.Exchanges{}


func main() {
	cliMenu()
}


func checkCurrencyStr(currency string) bool {
	return regexp.MustCompile(`^[a-zA-Z]+$`).MatchString(currency)
}


func listAllRates() {
	cli.ClearScreen(true)
	for _, rate := range exchanges.Rates {
		fmt.Printf("\n%6s - %-6s -> %7f", rate.Source, rate.Target, rate.Rate)
	}
	cli.EnterToContinue()
}

func getCurrencyPairs() (source, target string) {
	for {
		fmt.Println("Write your currency pairs to find(e.g EUR, US)")
		fmt.Printf("Currency From -> ")
		fmt.Scanf("%s", &source)
		fmt.Printf("Currency %s to -> ", source)
		fmt.Scanf("%s", &target)
		if checkCurrencyStr(source) && checkCurrencyStr(target) {
			source = str.ToUpper(source)
			target = str.ToUpper(target)
			return
		} else {
			cli.ClearScreen(true)
			fmt.Println("Invalid currency format!! Try again.\n")
		}
	}
}

func findExchangeRate() int {
	cli.ClearScreen(true)
	result := 1
	source, target := getCurrencyPairs()
	exRate, found := exchanges.FindExchangeRate(source, target)

	if found{
		fmt.Printf("\t-> %f", exRate.Rate)
	} else {
		fmt.Println("\nCouldn't find any result.")
		result = -1
	}
	cli.EnterToContinue()
	return result
}

func getExchangeRate() (newRate float32) {
	for {
		var rateStr string
		fmt.Printf("Write Exchange Rate(e.g 1,34) -> ")
		fmt.Scanf("%s", &rateStr)
		if newRate, err := strconv.ParseFloat(rateStr, 32); err == nil && newRate > 0 {
			return float32(newRate)
		}
	}
}

func addNewExchanceRate()  {
	cli.ClearScreen(true)
	source, target := getCurrencyPairs()
	rateVal := getExchangeRate()
	newRate := ex.ExchangeRate{Source: source, Target: target, Rate: rateVal}
	exchanges.CreateOrUpdateRate(newRate)
}


func printMenu() {
	cli.ClearScreen(true)
	fmt.Println("1-) Find Exchange Rate")
	fmt.Println("2-) List All Exchange Rates")
	fmt.Println("3-) Add Exchange Rate")
	fmt.Println("0-) Exit")
}

func cliMenu() {
	var choice int

	for {
		printMenu()
		fmt.Scanf("%d", &choice)

		switch choice {
		case 1:
			findExchangeRate()
		case 2:
			listAllRates()
		case 3:
			addNewExchanceRate()
		case 0:
			os.Exit(0)
		}
	}
}
