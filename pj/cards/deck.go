package main

import (
  "fmt"
  "strings"
  "io/ioutil"
  "os"
)

// Create a new type of 'deck', a slice of strings
// this means, type 'deck' borrows all the features
// from string type in Go bu default.
type deck []string

func newDeck() deck {
  cards := deck{}

  cardSuits := []string{"Spades", "Diamonds", "Hearts", "Clubs"}
  cardValues := []string{"Ace", "Two", "Three", "Four"}

  for _, suit := range cardSuits {
    for _, value := range cardValues {
      cards = append(cards, value+" of "+suit)
    }
  }
  return cards
}

// (d deck) is a receiver, of type deck
// it means d(item"s"), which is cards in type deck will be accessed and
// printed
// by convention, the items variable is the letter from the type
// in this case, d, because type deck.
func (d deck) print() {
  for i, card := range d {
    fmt.Println(i, card)
  }
}

// 2 arguments, d deck and handSize int
func deal(d deck, handSize int) (deck, deck) {
  return d[:handSize], d[handSize:]
}

func (d deck) toString() string {
  return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
  return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func newDeckFromFile(filename string) deck {
  bs, err := ioutil.ReadFile(filename)
  if err != nil {
    fmt.Println("Error:", err)
    os.Exit(1)
  }

  s := strings.Split(string(bs), ",") //Ace of Spades, Two of Spades, ...
  return deck(s)
}
