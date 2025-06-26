package profile

import (
	"strings"

	"github.com/jmoiron/sqlx"

	"golang.org/x/crypto/bcrypt"
)

type Profile struct {
	ID           int    `db:"id" json:"id"`
	Email        string `db:"email" json:"email"`
	Name         string `db:"name" json:"name"`
	PasswordHash string `db:"password_hash" json:"-"`
	CreatedAt    string `db:"created_at" json:"created_at,omitempty"`
	UpdatedAt    string `db:"updated_at" json:"updated_at,omitempty"`
}

func (p *Profile) SetEmail(value string) {
	p.Email = strings.ToLower(value)
}

func (p *Profile) SetPassword(value string) {
	p.PasswordHash = p.hash(value)
}

func (p *Profile) CheckPassword(value string) bool {
	return p.hash(value) == p.PasswordHash
}

func (p *Profile) hash(value string) string {
	bytes, _ := bcrypt.GenerateFromPassword([]byte(value), bcrypt.DefaultCost)

	return string(bytes)
}

type Usecases struct {
	db *sqlx.DB
}

func NewUsecases(db *sqlx.DB) *Usecases {
	return &Usecases{db: db}
}
