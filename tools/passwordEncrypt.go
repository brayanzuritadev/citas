package tools

import "golang.org/x/crypto/bcrypt"

func PasswordEncrypt(password string) (string, error) {

	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), cost)

	if err != nil {
		return err.Error(), err
	}

	return string(bytes), nil
}