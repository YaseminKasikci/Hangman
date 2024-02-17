package hangman

import "fmt"


func DrawWelcome() {
	fmt.Println(`
	_______  _        _______  _______  _______  _       
	|\     /|(  ___  )( (    /|(  ____ \(       )(  ___  )( (    /|
	| )   ( || (   ) ||  \  ( || (    \/| () () || (   ) ||  \  ( |
	| (___) || (___) ||   \ | || |      | || || || (___) ||   \ | |
	|  ___  ||  ___  || (\ \) || | ____ | |(_)| ||  ___  || (\ \) |
	| (   ) || (   ) || | \   || | \_  )| |   | || (   ) || | \   |
	| )   ( || )   ( || )  \  || (___) || )   ( || )   ( || )  \  |
	|/     \||/     \||/    )_)(_______)|/     \||/     \||/    )_)
																																 	
	`)
}
// display the current state of the scaffold and the state of the game (where we are)
func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLetf)
	drawState(g, guess)

}

// the state of the scaffold
func drawTurns(l int) {
	var draw string
	switch l {
	case 0:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |   / \
  _|_
 |   |______
 |          |
 |__________|
		`
	case 1:
		draw = `
    ____
   |    |      
   |    o      
   |   /|\     
   |    |
   |    
  _|_
 |   |______
 |          |
 |__________|
		`
	case 2:
		draw = `
    ____
   |    |      
   |    o      
   |    |
   |    |
   |     
  _|_
 |   |______
 |          |
 |__________|
		`
	case 3:
		draw = `
    ____
   |    |      
   |    o      
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 4:
		draw = `
    ____
   |    |      
   |      
   |      
   |  
   |  
  _|_
 |   |______
 |          |
 |__________|
		`
	case 5:
		draw = `
    ____
   |        
   |        
   |        
   |   
   |   
  _|_
 |   |______
 |          |
 |__________|
		`
	case 6:
		draw = `
    
   |     
   |     
   |     
   |
   |
  _|_
 |   |______
 |          |
 |__________|
		`
	case 7:
		draw = `
  _ _
 |   |______
 |          |
 |__________|
		`
	case 8:
		draw = `

		`
	}
	fmt.Println(draw)
}
// The state of the game

func drawState(g *Game, guess string) {
//letter found
	fmt.Println("Guessed: ")
	drawLetters(g.FoundLetters)
//letter use
	fmt.Println("Used: ")
	drawLetters(g.UsedLetters)

	// display depending on the state of the game
	switch g.State {
	case "goodGuess":
		fmt.Println("Good guess !")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' was already used\n", guess)
		case "badGuess":
			fmt.Printf("Bad guess !, '%s'is  not in the word\n", guess)
		case "lose":
			fmt.Printf("You lost !  the word was: ")
			drawLetters(g.Letters)
		case "won":
			fmt.Printf("YOU WON ! the word was: ")
			drawLetters(g.Letters)
		
	}
}
// loop over the letters found and already used
func drawLetters(l []string) {
		for _, c := range l {
			fmt.Printf("%v ", c)
		}
		fmt.Println()
}
