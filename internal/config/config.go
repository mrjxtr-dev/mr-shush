package config

const (
	DefaultLength = 8
	MaxLength     = 50
)

var ValidStrength = map[string]struct{}{
	"weak":   {},
	"good":   {},
	"strong": {},
}
