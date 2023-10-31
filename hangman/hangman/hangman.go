package hangman

import (
	"fmt"
	"strings"
)

type Game struct {
	State string //Game state
	Letters []string // letters in the world to find
	FoundLetters []string // good guesses
	UsedLetters []string // used letters
	TurnsLetf int // remaining attempts
}

//* LE MOT DEVINER | LE NOMBRE DE TOUR INIAL

func New(turns int, word string) (*Game, error) {
	if len(word) < 3 {
		return nil, fmt.Errorf("Word '%s' must be at least 3 characters. got =%v", word, len(word))
	}
	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(letters); i++ {
	found[i] = "_"
	}

	g := &Game {
		State: "",
		Letters: letters,
		FoundLetters: found,
		UsedLetters: []string{},
		TurnsLetf: turns,
	} 
	return g, nil
}
// Lettre proposer par le joueur | état de la partie
func (g *Game) MakeAGuess(guess string) {
	 guess = strings.ToUpper(guess)

	 switch g.State {
	case "won", "lost":
		return
	}

	if letterInWord(guess, g.UsedLetters){
		g.State = "alreadyGuessed"							
	} else if letterInWord(guess, g.Letters){
		g.State = "goodGuess"										
		g.RevealLetter(guess)										

		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}
	} else {
		g.State = "badGuess"
		g.LoseTurn(guess)

		if g.TurnsLetf <= 0 {
			g.State ="lost"
		}
	}
}

// ajoute la et les lettres du mot 
func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, l := range g.Letters  {
		if l == guess {
			g.FoundLetters[i] = guess
		}
	}
}
// tour de jeux fini et ajout dans les lettre utilisé
func (g *Game) LoseTurn(guess string){
	g.TurnsLetf--
	g.UsedLetters = append(g.UsedLetters, guess)
}

// trouver toute les lettres du mot
func hasWon(letters []string, foundLetters []string) bool {
	for i := range letters {
		if letters[i] != foundLetters[i]{
			return false
		}
	}
	return true
}



// la lettre dans le mot ?  
func letterInWord(guess string, letters []string) bool {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}
	return false
}

 
