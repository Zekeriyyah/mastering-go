package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	deckCards := deck{}

	cardSuit := []string{"Spade", "Diamond", "Hearts", "Club"}
	cardValues := []string{"Ace", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "King", "Queen"}

	for _, suit := range cardSuit {
		for _, value := range cardValues {
			deckCards = append(deckCards, value+" of "+suit)
		}
	}
	return deckCards
}

func (d deck) print() {
	for index, card := range d {
		fmt.Println(index, card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), "\n")
}

func (d deck) saveToFile(filename string) error {
	f, err := os.Create(filename)

	if _, err2 := f.Write([]byte(d.toString())); err2 != nil {
		fmt.Println("Error: ", err2)
		os.Exit(1)
	}

	return err
}

func newDeckFromFile(filename string) deck {
	byteSlice, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(byteSlice), "\n")

	return deck(s)

}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())

	r := rand.New(source)
	for i := range d {
		newIndex := r.Intn(len(d))
		d[i], d[newIndex] = d[newIndex], d[i]
	}
}
