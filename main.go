package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {

	// args
	length := flag.Int("l", 20, "Length of the password to generate.")
	count := flag.Int("c", 1, "How many passwords to generate.")
	alphabet := flag.String("a", "d", "Alphabet(s) to use: d=default, l=lower, u=upper, s=special, n=numbers.")
	flag.Parse()

	// genrate passwords
	dict := getAlphabet(*alphabet)

	i := 0
	for i < *count {
		i++
		fmt.Println(generatePassword(dict, *length))
	}
}

func getAlphabet(alphabet string) string {

	dict := &strings.Builder{}

	// lower
	if strings.ContainsAny(alphabet, "dl") {
		dict.WriteString("abcdefghijklmnopqrstuvwxyz")
	}
	// upper
	if strings.ContainsAny(alphabet, "du") {
		dict.WriteString("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	}
	// special
	if strings.ContainsAny(alphabet, "ds") {
		dict.WriteString("!@#$%&*?_+=")
	}
	// numbers
	if strings.ContainsAny(alphabet, "dn") {
		dict.WriteString("1234567890")
	}

	return dict.String()
}

func generatePassword(dict string, length int) string {

	random := rand.New(rand.NewSource(time.Now().UnixNano()))
	password := &strings.Builder{}

	i := 0
	for i < length {
		i++
		idx := random.Intn(len(dict))
		password.WriteByte(dict[idx])
	}

	return password.String()
}
