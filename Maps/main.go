package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#FF0000",
		"green": "#00FF00",
		"blue":  "#0000FF",
	}
	printMap(colors)
}

func printMap(c map[string]string) {
	for k, v := range c {
		fmt.Println(k, v)
	}
}
