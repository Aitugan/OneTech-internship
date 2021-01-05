package main

import (
	"fmt"
	"strconv"
	"math/rand"
)

type Card struct {
	Nominal string
	Suit string
}

type Deck struct {
	Cards []*Card
}

var(
	Suits = []string{"Heart","Diamond","Spade","Club"}
	Faces = []string{"Jack", "Queen", "King", "Ace"}
)

func (deck *Deck) Shuffle() {
	for i := 0; i < 52; i++ {
		i1 := rand.Intn(52)
		i2 := rand.Intn(52)
		deck.Cards[i1], deck.Cards[i2] = deck.Cards[i2], deck.Cards[i1]
	}
}

func (deck *Deck) ShowDeck() {
	for _,v := range deck.Cards {
		fmt.Println(v)
	}	
}

func GenerateDeck() *Deck {
	deck := &Deck{} 
	for i := 0; i < 4; i++ {
		for j := 2; j < 11; j++ {
			card := &Card{strconv.Itoa(j), Suits[i]}
			deck.Cards = append(deck.Cards,card)
		}
		for j := 0; j < 4; j++ {
			card := &Card{Faces[j], Suits[i]}
			deck.Cards = append(deck.Cards,card)
		}
	}
	return deck
}

func main() {
	deck := GenerateDeck()
	deck.ShowDeck()
	deck.Shuffle()
	fmt.Println("__________")
	deck.ShowDeck()
}
