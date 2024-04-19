package players

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Player struct {
	gorm.Model
	Username string
	Password string
	CharName string
}

type APIPlayer struct {
	Username string
	CharName string
}

func newPlayer(usr string, pass string, name string) Player {
	p := Player{
		Username: usr,
		Password: pass,
		CharName: name,
	}
	return p
}

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func CreatePlayer() Player {
	reader := bufio.NewReader(os.Stdin)

	usr, _ := getInput("Enter a new UserName: ", reader)
	pass, _ := getInput("Enter a new Password: ", reader)
	name, _ := getInput("Enter a new Name: ", reader)

	hash, _ := HashPassword(pass)

	p := newPlayer(usr, hash, name)
	fmt.Println("Created bew Character")

	return p
}
