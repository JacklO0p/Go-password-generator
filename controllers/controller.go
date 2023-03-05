package controllers

import (
	"math/rand"
	"password_generator/global"
	"strconv"

	"github.com/labstack/echo/v4"
)

var (
	Uppercase = []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")
	Lowercase = []rune("abcdefghijklmnopqrstuvwxyz")
	Numbers   = []rune("0123456789")
	Symbols   = []rune("!@#$%^&*()_+{}:\"<>?,./;'[]\\|")
)

func Home(c echo.Context) error {
	var counter int
	checkPos := [4]bool{}
	global.Password = ""

	length, err := strconv.Atoi(c.FormValue("length"))
	if err != nil || length < 8 {
		length = 8
	}

	if c.FormValue("uppercase") != "" {
		global.Uppercase = Uppercase
		checkPos[0] = true
		counter++
	}

	if c.FormValue("lowercase") != "" {
		global.Lowercase = Lowercase
		checkPos[1] = true
		counter++
	}

	if c.FormValue("numbers") != "" {
		global.Numbers = Numbers
		checkPos[2] = true
		counter++
	}

	if c.FormValue("symbols") != "" {
		global.Symbols = Symbols
		checkPos[3] = true
		counter++
	}

	if counter == 0 {
		return c.JSON(200, "no options selected")
	}

	generatePassword(length, counter, checkPos)

	return c.JSON(200, "generated password: "+global.Password)
}

func generatePassword(length, max int, pos [4]bool) {
	runes := make([]rune, length)

	for i := 0; i < length; i++ {
		field := randomField(max)

		for j := 0; j < 4; j++ {
			if pos[j] && field == 0 {
				field = j
				break
			} else {
				if pos[j] && field > 0 {
					field--
				}
			}
		}

		switch field {
		case 0:
			if len(Uppercase) > 0 {
				runes[i] = Uppercase[rand.Intn(len(Uppercase))]
			} else {
				i--
			}
		case 1:
			if len(Lowercase) > 0 {
				runes[i] = Lowercase[rand.Intn(len(Lowercase))]
			} else {
				i--
			}
		case 2:
			if len(Numbers) > 0 {
				runes[i] = Numbers[rand.Intn(len(Numbers))]
			} else {
				i--
			}
		case 3:
			if len(Symbols) > 0 {
				runes[i] = Symbols[rand.Intn(len(Symbols))]
			} else {
				i--
			}
		}
	}

	global.Password = string(runes)
}

func randomField(max int) int {
	values := make([]int, 100)
	
	for i := 0; i < 100; i++ {
		values[i] = rand.Intn(max)
	}

	return values[rand.Intn(len(values))]
}