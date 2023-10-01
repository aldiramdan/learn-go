package src

import (
	"math/rand"
	"strings"
	"time"
)

func GenPass(password, level string) string {
	var s string
	var char int

	if level == "low" {
		char = 6
	} else if level == "mid" {
		char = 12
	} else if level == "strong" {
		char = 18
	} else {
		char = 24
	}

	r := []rune(password)

	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < len(password); i++ {
		c := string(r[i])
		if rand.Float64() < 0.5 {
			l := strings.ToLower(c)
			s = s + l
		} else {
			u := strings.ToUpper(c)
			s = s + u
		}
	}

	m := []rune(s)
	p := char - len(m)
	for i := 0; i < p; i++ {
		min := 33
		max := 126
		g := rand.Intn(max-min+1) + min
		m = append(m, rune(g))

	}
	result := string(m)
	return result
}

func PassGen(password, level string) string {
	var s string
	var g string
	var char int

	lower := "abcdefghijklmnopqrstuvwxyz"
	upper := strings.ToUpper(lower)
	symbols := "!@#$%^&*()_+"
	numbers := "1234567890"
	all := lower + upper + symbols + numbers

	if level == "low" {
		char = 6
	} else if level == "mid" {
		char = 12
	} else if level == "strong" {
		char = 18
	} else {
		char = 24
	}

	r := []rune(password)
	q := rand.NewSource(time.Now().UnixNano())
	o := rand.New(q)
	for i := 0; i < len(password); i++ {
		c := string(r[i])
		if rand.Float64() < 0.5 {
			l := strings.ToLower(c)
			s = s + l
		} else {
			u := strings.ToUpper(c)
			s = s + u
		}
	}
	m := []rune(s)
	p := char - len(password)
	for i := 0; i < p; i++ {
		g += string(all[o.Intn(len(all))])
	}
	result := string(m) + g
	return result
}
