package util

import (
	"math/rand"
	"strings"
	"time"
)

const alphabet = "abcdefghijklmnopqrstuvwxyz"

func init() {
  rand.Seed(time.Now().UnixNano())
}

// Gererate random Integer between min and max
func RandomInt(min, max int64) int64 {
  return min + rand.Int63n(max - min + 1)
}

// Gererate random String of length n
func RandomString(n int) string {
  var sb strings.Builder

  k := len(alphabet)

  for i := 0; i < n; i++ {
    c := alphabet[rand.Intn(k)]
    sb.WriteByte(c)
  }

  return sb.String()
}

func RandomEmail() string {
  return RandomString(6)+"@testemail.com"
}

func RandomName() string {
  return RandomString(4) + "name"
}

func RandomNick() string {
  return RandomString(5) + "nick"
}

func RandomAge() int64 {
  return RandomInt(18,90)
}

func RandomPass() string {
  return RandomString(8) + "pass" 
}
