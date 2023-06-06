/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var number int
var cfgFile string

// rootCmd represents the base command when called without any subcommands
var card = &cobra.Command{
	Use:   "card",
	Short: "Get card from a poker deck",
	Long: `Application created for initiative in Savage World
	get card from a poker deck`,
	Run: func(cmd *cobra.Command, args []string) {

		if _, err := os.Stat(cfgFile); err != nil {
			_, err := os.Create(cfgFile)
			deck := createDeck()
			deck.shuffle()
			deck.saveToFile(cfgFile)
			if err != nil {
				panic(err)
			}
		}

		if cmd.Flags().Changed("create") {
			deck := createDeck()
			if cmd.Flags().Changed("flush") {
				deck = deck.shuffle()
			}
			deck.saveToFile(cfgFile)
		}

		number, _ := cmd.Flags().GetInt("number")
		getDeck := newDeckFromFile(cfgFile)
		hand, remainingDeck := deal(getDeck, number)
		remainingDeck.saveToFile(cfgFile)

		for i := 0; i < len(hand); i++ {
			println(hand[i])
		}

	},
}

func Execute() {
	err := card.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	card.PersistentFlags().StringVar(&cfgFile, "config", ".deck.txt", "config file (default is .deck.txt)")
	card.Flags().BoolP("create", "c", false, "Create new deck")
	card.Flags().BoolP("flush", "f", false, "flush deck")
	card.Flags().IntVarP(&number, "number", "n", 1, "Get number of card")
}
