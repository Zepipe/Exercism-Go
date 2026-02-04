package blackjack

// ParseCard returns the integer value of a card following blackjack ruleset.
func ParseCard(card string) int {
    
	switch card {
        case "ace":
            return 11
        case "king", "queen", "jack", "ten":
            return 10
        case "nine":
            return 9
        case "eight":
            return 8
        case "seven":
            return 7
        case "six":
            return 6
        case "five":
            return 5
        case "four":
            return 4
        case "three":
            return 3
        case "two":
            return 2
        default:
            return 0
    }
}

// FirstTurn returns the decision for the first turn, given two cards of the
// player and one card of the dealer.
func FirstTurn(card1, card2, dealerCard string) string {
    // Convert cards to their values
    card1Val := ParseCard(card1)
    card2Val := ParseCard(card2)
    dealerCardVal := ParseCard(dealerCard)
    
    total := card1Val + card2Val
    
    // Rule 1: Pair of aces (both cards are "ace")
    if card1 == "ace" && card2 == "ace" {
        return "P" // Split
    }
    
    // Rule 2: Blackjack (total of 21)
    if total == 21 {
        // Check dealer's card for Blackjack
        if dealerCard == "ace" || dealerCard == "king" || 
           dealerCard == "queen" || dealerCard == "jack" || 
           dealerCard == "ten" {
            return "S" // Stand
        }
        return "W" // Automatically win
    }
    
    // Rule 3: Sum in range [17, 20]
    if total >= 17 && total <= 20 {
        return "S" // Stand
    }
    
    // Rule 4: Sum in range [12, 16]
    if total >= 12 && total <= 16 {
        if dealerCardVal >= 7 {
            return "H" // Hit
        }
        return "S" // Stand
    }
    
    // Rule 5: Sum of 11 or lower
    return "H" // Hit
}
