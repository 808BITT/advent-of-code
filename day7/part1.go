package main

import (
	"fmt"
	"os"
)

func main() {
	// read the input file
	currDir, _ := os.Getwd()
	os.Chdir(currDir)

	fileStream, err := os.Open("input.txt")
	if err != nil {
		fmt.Println("Error reading input file:", err)
		return
	}
	defer fileStream.Close()

	// read the hands line by line
	var hands Input
	for {
		var hand Hand
		var cards string
		var bid int
		_, err := fmt.Fscanf(fileStream, "%s %d\n", &cards, &bid)
		if err != nil {
			if err.Error() == "EOF" {
				break
			}
			fmt.Println("Error reading input file:", err)
			return
		}
		fmt.Println(cards, bid)
		hand = parseHand(cards, bid)
		hands.UnsortedHands = append(hands.UnsortedHands, hand)
	}

	fmt.Println("Unsorted hands:")
	for _, hand := range hands.UnsortedHands {
		fmt.Println(hand)
	}

	// sort the hands
	for i := 1; i <= 7; i++ {
		ties := make([]Hand, 0)
		for _, hand := range hands.UnsortedHands {
			if hand.Rank.Value == i {
				ties = append(ties, hand)
			}
		}
		if len(ties) > 1 {
			sortTies(ties)
			hands.SortedHands = append(hands.SortedHands, ties...)
		} else if len(ties) == 1 {
			for _, hand := range hands.UnsortedHands {
				if hand.Rank.Value == i {
					hands.SortedHands = append(hands.SortedHands, hand)
				}
			}
		}
	}

	fmt.Println("Sorted hands:")
	for _, hand := range hands.SortedHands {
		fmt.Println(hand)
	}

	var winnings int
	for i := 1; i < len(hands.SortedHands)+1; i++ {
		// fmt.Println("Hand", i, ":", hands.SortedHands[i-1].Bid)
		winnings += hands.SortedHands[i-1].Bid * i
	}

	fmt.Println("Winnings:", winnings)

}

func sortTies(ties []Hand) {
	for i := 0; i < len(ties); i++ {
		for j := i + 1; j < len(ties); j++ {
			if ties[i].Cards[0].CardValue > ties[j].Cards[0].CardValue {
				ties[i], ties[j] = ties[j], ties[i]
			} else if ties[i].Cards[0].CardValue == ties[j].Cards[0].CardValue {
				if ties[i].Cards[1].CardValue > ties[j].Cards[1].CardValue {
					ties[i], ties[j] = ties[j], ties[i]
				} else if ties[i].Cards[1].CardValue == ties[j].Cards[1].CardValue {
					if ties[i].Cards[2].CardValue > ties[j].Cards[2].CardValue {
						ties[i], ties[j] = ties[j], ties[i]
					} else if ties[i].Cards[2].CardValue == ties[j].Cards[2].CardValue {
						if ties[i].Cards[3].CardValue > ties[j].Cards[3].CardValue {
							ties[i], ties[j] = ties[j], ties[i]
						} else if ties[i].Cards[3].CardValue == ties[j].Cards[3].CardValue {
							if ties[i].Cards[4].CardValue > ties[j].Cards[4].CardValue {
								ties[i], ties[j] = ties[j], ties[i]
							}
						}
					}
				}
			}
		}
	}
}

type Input struct {
	UnsortedHands []Hand
	SortedHands   []Hand
}

type Card struct {
	CardType  string
	CardValue int
}

type Hand struct {
	Cards []Card
	Bid   int
	Rank  HandRank
}

type HandRank struct {
	Name  string
	Value int
}

func parseHand(cards string, bid int) Hand {
	var hand Hand
	hand.Bid = bid
	for _, card := range cards {
		switch card {
		case 'A':
			hand.Cards = append(hand.Cards, Card{"A", 14})
		case 'K':
			hand.Cards = append(hand.Cards, Card{"K", 13})
		case 'Q':
			hand.Cards = append(hand.Cards, Card{"Q", 12})
		case 'J':
			hand.Cards = append(hand.Cards, Card{"J", 11})
		case 'T':
			hand.Cards = append(hand.Cards, Card{"T", 10})
		case '9':
			hand.Cards = append(hand.Cards, Card{"9", 9})
		case '8':
			hand.Cards = append(hand.Cards, Card{"8", 8})
		case '7':
			hand.Cards = append(hand.Cards, Card{"7", 7})
		case '6':
			hand.Cards = append(hand.Cards, Card{"6", 6})
		case '5':
			hand.Cards = append(hand.Cards, Card{"5", 5})
		case '4':
			hand.Cards = append(hand.Cards, Card{"4", 4})
		case '3':
			hand.Cards = append(hand.Cards, Card{"3", 3})
		case '2':
			hand.Cards = append(hand.Cards, Card{"2", 2})
		}
	}

	// get a count of each card value
	cardCount := make(map[int]int)
	for _, card := range hand.Cards {
		cardCount[card.CardValue]++
	}

	fmt.Println(len(cardCount))
	fmt.Println("Card counts:", cardCount)

	if len(cardCount) == 1 {
		hand.Rank.Name = "Five of a Kind"
		hand.Rank.Value = 7
	}

	if len(cardCount) == 2 {
		if cardCount[hand.Cards[0].CardValue] == 4 ||
			cardCount[hand.Cards[1].CardValue] == 4 {
			hand.Rank.Name = "Four of a Kind"
			hand.Rank.Value = 6
		} else {
			hand.Rank.Name = "Full House"
			hand.Rank.Value = 5
			return hand
		}
	}

	if len(cardCount) == 3 {
		hand.Rank.Name = "Three of a Kind"
		hand.Rank.Value = 4
	}

	if len(cardCount) == 4 {
		hand.Rank.Name = "Two Pairs"
		hand.Rank.Value = 3
	}

	// check for three of a kind
	for _, count := range cardCount {
		if count == 3 {
			hand.Rank.Name = "Three of a Kind"
			hand.Rank.Value = 4
		}
	}

	// check for two pairs
	var pairs int
	for _, count := range cardCount {
		if count == 2 {
			pairs++
		}
	}
	if pairs == 2 {
		hand.Rank.Name = "Two Pairs"
		hand.Rank.Value = 3
	} else if pairs == 1 {
		hand.Rank.Name = "Pair"
		hand.Rank.Value = 2
	}

	// check for high card
	if hand.Rank.Value == 0 {
		hand.Rank.Name = "High Card"
		hand.Rank.Value = 1
	}

	return hand
}

func (h *Hand) String() string {
	return fmt.Sprintf("%v", h.Cards)
}
