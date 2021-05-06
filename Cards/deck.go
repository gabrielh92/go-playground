package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	values := []string{"Ace", "Two", "Three", "Four", "Five", "Six",
		"Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	cards := deck{}

	for _, suit := range suits {
		for _, value := range values {
			cards = append(cards, value+" of "+suit)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, num int) (deck, deck) {
	return d[:num], d[num:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)

	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(255)
	}

	return deck(strings.Split(string(bs), ", "))
}

func (d deck) shuffle() {
	seed := rand.NewSource(time.Now().UnixNano())
	rng := rand.New(seed)
	for i := range d {
		newIdx := rng.Intn(len(d) - 1)
		d[i], d[newIdx] = d[newIdx], d[i]
	}
}
