package util

import (
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"math/rand"
	"regexp"
	"time"

	beego "github.com/beego/beego/v2/server/web"
)

func PrintApiPath() {
	tree := beego.PrintTree()
	methods := tree["Data"].(beego.M)
	for k, v := range methods {
		vv := v.(*[][]string)

		for _, vvv := range *vv {
			fmt.Printf("\033[32m%s\033[0m => %v\n", k, fmt.Sprintf("%-30s", vvv[0]))
			// fmt.Printf("\033[32m%s\033[0m => %v %v %v\n", k, fmt.Sprintf("%-30s", vvv[0]), fmt.Sprintf("%-20s", vvv[1]), fmt.Sprintf("%-30s", vvv[2]))
		}
	}
}

func VaildateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
	return re.MatchString(email)
}

func VaildatePhone(phone string) bool {
	re := regexp.MustCompile(`^(\+?86)?-?1[3-9]\d{9}$`)
	return re.MatchString(phone)
}

var (
	Key = []byte("1234567890abcdef")
	Iv  = make([]byte, aes.BlockSize)
)

func Encrypt(plaintext []byte) (*[]byte, error) {

	block, err := aes.NewCipher(Key)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(block, Iv)
	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)
	return &ciphertext, nil
}

func Decrypt(byt []byte) (*[]byte, error) {
	block, err := aes.NewCipher(Key)
	if err != nil {
		return nil, err
	}
	stream := cipher.NewCTR(block, Iv)
	plaintext := make([]byte, len(byt))
	stream.XORKeyStream(plaintext, byt)
	return &plaintext, nil
}

const charsetLetterLowerCase = "abcdefghijklmnopqrstuvwxyz"
const charsetLetterUpperCase = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
const charsetNumber = "0123456789"

func RandString(length int, ditc string) string {

	r := rand.New(rand.NewSource(time.Now().UnixNano())) // 每次生成不同结果
	b := make([]byte, length)
	for i := range b {
		b[i] = ditc[r.Intn(len(ditc))]
	}
	return string(b)
}

func RandStringLetter(length int) string {
	return RandString(length, charsetLetterLowerCase+charsetLetterUpperCase)
}
func RandNumber(length int) string {
	return RandString(length, charsetNumber)
}
