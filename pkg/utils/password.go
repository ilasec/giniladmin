package utils

import (
	"crypto/md5"
	"encoding/hex"
	"golang.org/x/crypto/bcrypt"
)

func Password(password, salt string) string {
	m := md5.New()
	m.Write([]byte(password + salt))
	pwd := hex.EncodeToString(m.Sum(nil))
	return pwd
}

func PasswordByBcrypt(password string) (ret string, err error) {
	hashedPassword, hashErr := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if hashErr != nil {
		err = hashErr
		return
	}
	ret = string(hashedPassword)
	return
}

func PasswordCheckByBcrypt(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
