package generator

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"github.com/mrjxtr-dev/mr-shush/internal/config"
	"github.com/mrjxtr-dev/mr-shush/internal/models"
)

type Generator struct {
	*models.PasswordCharset
}

func New() *Generator {
	return &Generator{
		&models.PasswordCharset{
			Lowercase: "abcdefghijklmnopqrstuvwxyz",
			Uppercase: "ABCDEFGHIJKLMNOPQRSTUVWXYZ",
			Numbers:   "0123456789",
			Symbols:   "!@#$%^&*()-_=+[]{}<>?/",
		},
	}
}

// Generate a password based on length and strength
func (g *Generator) GeneratePassword(length int, strength string) (string, error) {
	if length <= 0 {
		length = config.DefaultLength
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

// Validate and process flags(length, strength) before generating
func (g *Generator) validateFlags(length int, strength string) error {
	if length > config.MaxLength {
		return errors.New("length must be less than 50")
	}

	if _, ok := config.ValidStrength[strength]; !ok {
		return errors.New("invalid strength flag")
	}

	return nil
}

// Build a charset based on strength (weak, good, strong)
// Defaults to "good" if strength flag is not passed
func (g *Generator) buildCharset(strength string) string {
	weakCharset := g.Lowercase + g.Uppercase
	goodCharset := weakCharset + g.Numbers
	strongCharset := goodCharset + g.Symbols

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

// Build a password based on charset and length
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
