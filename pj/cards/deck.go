package main

import "fmt"

// Create a new type of 'deck', a slice of strings
// this means, type 'deck' borrows all the features
// from string type in Go bu default.
type deck []string

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
