package dictionary

import (
	"bufio"
	"math/rand"
	"os"
	"time"
)

var words = make([]string, 0, 50)

// load the file
func Load(filename string) error {
	f, err :=os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		words = append(words, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		return err 
	}
	return nil 
}
//return a random word
func PickWord() string {
	rand.NewSource(time.Now().Unix())
	i := rand.Intn(len(words))
	return words[i]
}
