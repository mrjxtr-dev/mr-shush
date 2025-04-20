package models

type PasswordCharset struct {
	Lowercase string
	Uppercase string
	Numbers   string
	Symbols   string
}

type PasswordData struct {
	Password string
	Name     string
}
