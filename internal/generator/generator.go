package generator

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"github.com/mrjxtr-dev/mr-shush/internal/models"
)

const (
	defaultLength = 8
	maxLength     = 50
)

var validStrength = map[string]struct{}{
	"weak":   {},
	"good":   {},
	"strong": {},
}

type Generator struct {
	password *models.Password
}

func New() *Generator {
	return &Generator{
		password: &models.Password{
			Lowercase: "abcdefghijklmnopqrstuvwxyz",
			Uppercase: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Numbers:   "0123456789",
			Symbols:   "!@#$%^&*()-_=+[]{}<>?/",
		},
	}
}

func (g *Generator) GeneratePassword(length int, strength string) (string, error) {
	if length <= 0 {
		length = defaultLength
	}

	if err := g.validateFlags(length, strength); err != nil {
		return "", fmt.Errorf("flags error: %w", err)
	}

	charset := g.buildCharset(strength)
	password, err := g.buildPassword(charset, length)
	if err != nil {
		return "", fmt.Errorf("error building password: %w", err)
	}

	return password, nil
}

// Validate and process flags before generating
func (g *Generator) validateFlags(length int, strength string) error {
	if length > maxLength {
		return errors.New("length must be less than 50")
	}

	if _, ok := validStrength[strength]; !ok {
		return errors.New("invalid strength flag")
	}

	return nil
}

// Build a charset based on strength (weak, good, strong)
// Defaults to "good" if strength flag is not passed
func (g *Generator) buildCharset(strength string) string {
	weakCharset := g.password.Lowercase + g.password.Uppercase
	goodCharset := weakCharset + g.password.Numbers
	strongCharset := goodCharset + g.password.Symbols

	var charset string

	switch strength {
	case "weak":
		charset = weakCharset
	case "strong":
		charset = strongCharset
	default:
		charset = goodCharset
	}

	return charset
}

func (g *Generator) buildPassword(charset string, length int) (string, error) {
	password := make([]byte, length)

	for i := range password {
		n, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			return "", err
		}
		password[i] = charset[n.Int64()]
	}

	return string(password), nil
}
