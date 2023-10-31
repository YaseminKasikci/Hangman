package hangman

import "fmt"

//* affichage de Bonjour
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
// afficher l'état courant de l'échaffaut et l'état de la partie (où est ce qu'on en est)
func Draw(g *Game, guess string) {
	drawTurns(g.TurnsLetf)
	drawState(g, guess)

}

// l'état de l'échaffaut 
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
// L'état de la partie

func drawState(g *Game, guess string) {
	//lettre trouvé
	fmt.Println("Guessed: ")
	drawLetters(g.FoundLetters)
//lettre utiliser
	fmt.Println("Used: ")
	drawLetters(g.UsedLetters)

	// affichage en fonctionn de l'état de la partie
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
// boucle sur les lettres trouvé et déjà utilisée
func drawLetters(l []string) {
		for _, c := range l {
			fmt.Printf("%v ", c)
		}
		fmt.Println()
}