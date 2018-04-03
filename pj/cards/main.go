package main

func main() {

  cards := newDeck()
  // deal returns 2 values, and needs to be assigned to to var
  hand, remainingCards := deal(cards, 5)

  hand.print()
  remainingCards.print()
}


