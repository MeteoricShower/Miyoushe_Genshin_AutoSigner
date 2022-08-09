package util

import (
	"crypto/md5"
	"fmt"
	"math/rand"
	"time"
)

func GetMD5(str string) string {
	data := []byte(str)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

func GetRandString(len int, seed int64) string {
	bytes := make([]byte, len)
	r := rand.New(rand.NewSource(seed))
	for i := 0; i < len; i++ {
		b := r.Intn(36)
		if b > 9 {
			b += 39
		}
		b += 48
		bytes[i] = byte(b)
	}
	return string(bytes)
}

func GetDs() string {
	currentTime := time.Now().Unix()
	stringRom := GetRandString(6, currentTime)
	stringAdd := fmt.Sprintf("salt=%s&t=%d&r=%s", "dWCcD2FsOUXEstC5f9xubswZxEeoBOTc", currentTime, stringRom)
	stringMd5 := GetMD5(stringAdd)
	return fmt.Sprintf("%d,%s,%s", currentTime, stringRom, stringMd5)
}
