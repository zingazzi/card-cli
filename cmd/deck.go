package cmd

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type IDeck []string

func (d IDeck) shuffle() IDeck {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	for i := range d {
		n := r.Intn(len(d) - 1)
		d[i], d[n] = d[n], d[i]
	}
	return d
}

func createDeck() IDeck {
	cards, shades := getCards()
	deck := IDeck{}
	for _, shade := range shades {
		for _, card := range cards {
			deck = append(deck, card+" di "+shade)
		}
	}
	deck = append(deck, "Jolly", "Jolly")
	return deck
}

func getCards() (IDeck, IDeck) {
	cards := IDeck{"Asso", "2", "3", "4", "5", "6", "7", "8", "9", "10", "Jack", "Regina", "Re"}
	shades := IDeck{"Cuori", "Quadri", "Fiori", "Picche"}
	return cards, shades
}

func deal(deck IDeck, handSize int) (IDeck, IDeck) {
	return deck[:handSize], deck[handSize:]
}

func (d IDeck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d IDeck) saveToFile(filename string) error {
	err := ioutil.WriteFile(filename, []byte(d.toString()), 0666)
	if err != nil {
		switch err {
		case os.ErrInvalid:
			fmt.Println("Invalid argument")
		case os.ErrPermission:
			fmt.Println("Permission denied")
		case os.ErrNotExist:
			fmt.Println("File does not exist")
		default:
			fmt.Println(err)
		}
	}
	return err
}

func newDeckFromFile(filename string) IDeck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		// Option #1 - log the error and return a call to newDeck()
		// Option #2 - log the error and entirely quit the program
		println("Error: ", err)
		os.Exit(1)
	}
	s := strings.Split(string(bs), ",")
	return IDeck(s)
}
