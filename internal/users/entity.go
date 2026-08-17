package users

import (
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
)

const (
	passwordMinLen = 6
	passwordMaxLen = 72
)

var (
	ErrNameRequired     = errors.New("Name is required")
	ErrLoginRequired    = errors.New("Login is required")
	ErrPasswordRequired = errors.New("Password is required")
	ErrPasswordTooShort = errors.New("Password must have at least 6 characters")
	// For a password longer than 72 bytes, the user thinks that this extra 
	// length makes the password stronger, but it makes no difference to bcrypt.
	// Standard web payloads (such as JSON or form data sent via fetch/HTTP) encode 
	// text using UTF-8. In UTF-8, every character in the standard ASCII set (a-z, 
	// A-Z, 0-9, !, @, #, $, %, ^, &, *) encodes directly into 1 byte, so the amount
	// of bytes and web characters are the same.
	ErrPasswordTooLong  = errors.New("Password must have less than 72 characters")
	ErrPasswordContainsInvalidChars = errors.New("Password contains invalid characters")
)

type User struct {
	ID         int64     `json:"id"`
	Name       string    `json:"name"`
	Login      string    `json:"login"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `json:"created_at"`
	ModifiedAt time.Time `json:"modified_at"`
	Deleted    bool      `json:"-"`
	LastLogin  time.Time `json:"last_login"`
}

func isAllowedPasswordChar(ch rune) bool {
	
	switch {
		case ch >= 'a' && ch <= 'z':
			return true
		case ch >= 'A' && ch <= 'Z':
			return true
		case ch >= '0' && ch <= '9':
			return true
		case ch == '!' || ch == '@' || ch == '#' || ch == '$' || ch == '%' || ch == '^' || ch == '&' || ch == '*':
			return true
		default:
			return false
	}
	
}

func validatePasswordCharacters(password string) error {
	
	for _, ch := range password {
		if ! isAllowedPasswordChar(ch) {
			return ErrPasswordContainsInvalidChars
		}
	}
	return nil
	
}

func (u *User) SetHashedOriginalPassword(password string) error {

	if password == "" {
		return ErrPasswordRequired
	}
	if len(password) < passwordMinLen {
		return ErrPasswordTooShort
	}
	if len(password) > passwordMaxLen {
		return ErrPasswordTooLong
	}

	err := validatePasswordCharacters(password)
	if err != nil {
		return err
	}
	
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return errors.New("Error hashing the password")
	}

	u.Password = string(hashedPassword)

	return nil

}

func (u *User) Validate() error {

	if u.Name == "" {
		return ErrNameRequired
	}
	if u.Login == "" {
		return ErrLoginRequired
	}

	return nil

}
