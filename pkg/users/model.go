package users

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"github.com/rs/zerolog/log"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username string
	Password string
	Email    string
}

type UserCreateDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email,omitempty"`
}

type LoginDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email,omitempty"`
}

func (dto UserCreateDTO) ToUser() *User {
	var err error
	user := new(User)
	user.Username = dto.Username
	user.Password, err = hashPassword(dto.Password)
	if err != nil {
		log.Err(err)
		return nil
	}

	user.Email, err = encryptData(dto.Email)
	if err != nil {
		log.Err(err)
		return nil
	}

	return user
}

var bytes = []byte{35, 46, 57, 24, 85, 35, 24, 74, 87, 35, 88, 98, 66, 32, 14, 05}

const secret string = "abc&1*~#^2^#s0^=)^^7%b34"

func encode(b []byte) string {
	return base64.StdEncoding.EncodeToString(b)
}

func decode(s string) []byte {
	data, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		panic(err)
	}
	return data
}

func encryptData(s string) (string, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}
	plainText := []byte(s)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)

	return encode(cipherText), nil
}

func decryptData(s string) (string, error) {
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}
	cipherText := decode(s)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)

	return string(plainText), nil
}

func MigrateModels(db *gorm.DB) error {
	return db.AutoMigrate(&User{})
}