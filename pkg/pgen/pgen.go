/*
Package pgen is a password generator tool to create random passwords
The idea of this package came from https://github.com/Luzifer/password
*/
package pgen

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

// customError represents the errors in this package
type customError string

func (e customError) Error() string {
	return string(e)
}

const (
	// MinimumLength the minimum length of a password
	MinimumLength = 8
	// ErrMinLengthRequired represents the minimum length error
	ErrMinLengthRequired = customError("minimum length error " + string(MinimumLength))
	// CharacterNumerics all numbers
	CharacterNumerics = "0123456789"
	// CharacterLetters all letters
	CharacterLetters = "abcdefghijklmnopqrstuvwxyz"
	// CharacterSpecials all special characters
	CharacterSpecials = "!#$%&()*+,-_./:;=?@[]^{}~|"
)

// Generator provides methods for generating secure passwords
type Generator struct {
	// CharacterTables maps the allowed characters (numerics, letters, specials)
	CharacterTables map[string]string
	// BadCharacters define the characters that will be ignored
	BadCharacters []string
}

// NewGenerator initializes a new generator
func NewGenerator() *Generator {
	rand.Seed(time.Now().UnixNano())
	return &Generator{
		CharacterTables: map[string]string{
			"numerics": CharacterNumerics,
			"letters":  CharacterLetters,
			"specials": CharacterSpecials,
		},
		BadCharacters: []string{"I", "l", "0", "O", "B", "8", "o"},
	}
}

// GeneratePassword generates a new password with a given length and
// optional special characters in it, if length < MinimumLength will return
// the ErrMinLengthRequired error
func (s *Generator) GeneratePassword(length int, special bool) (string, error) {

	if length < MinimumLength {
		return "", ErrMinLengthRequired
	}

	universe := ""
	if len(s.CharacterTables["letters"]) > 0 {
		universe = strings.Join([]string{
			s.CharacterTables["letters"],
			strings.ToUpper(s.CharacterTables["letters"]),
		}, "")
	}

	if len(s.CharacterTables["numerics"]) > 0 {
		universe = strings.Join([]string{
			universe,
			s.CharacterTables["numerics"],
		}, "")
	}

	if special && len(s.CharacterTables["specials"]) > 0 {
		universe = strings.Join([]string{universe, s.CharacterTables["specials"]}, "")
	}

	password := ""
	for {
		char := string(universe[rand.Intn(len(universe))])
		if strings.Contains(strings.Join(s.BadCharacters, ""), char) {
			continue
		}

		password = fmt.Sprintf("%s%s",
			password,
			char,
		)
		if length == len(password) {
			break
		}
	}
	return password, nil
}
