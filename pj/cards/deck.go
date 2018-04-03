package main

import "fmt"

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
