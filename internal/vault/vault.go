package vault

import (
	"database/sql"

	"github.com/mrjxtr-dev/mr-shush/internal/models"
)

type Vault struct {
	*models.PasswordData
	*sql.DB
}

func New(password, name string) *Vault {
	return &Vault{
		&models.PasswordData{
			Password: password,
			Name:     name,
		},
		&sql.DB{},
	}
}

func (v *Vault) Store() error {
	return nil
}
