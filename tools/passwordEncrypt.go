package tools

import "golang.org/x/crypto/bcrypt"

func PasswordEncrypt(password string) (string, error) {

	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return err.Error(), err
	}

	return string(bytes), nil
}
