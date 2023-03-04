package controllers

import (
	"math/rand"
	"password_generator/global"
	"strconv"

	"github.com/labstack/echo/v4"
)

var Length int
var Uppercase bool
var Lowercase bool
var Numbers bool
var Symbols bool

func Home(c echo.Context) error {
	var err error
	var counter int = 0
	checkPos := [4]bool{false, false, false, false}
	global.Password = ""

	if c.FormValue("length") != "" {
		Length, err = strconv.Atoi(c.FormValue("length"))
		if err != nil {
			return err
		}

		if Length < 8 {
			Length = 8
		}

	} else {
		Length = 8
	}

	if len(c.FormValue("uppercase")) == 0 {
		Uppercase = false
	} else {
		Uppercase = true

		checkPos[0] = true
		counter++
	}

	if len(c.FormValue("lowercase")) == 0 {
		Lowercase = false
	} else {
		Lowercase = true

		checkPos[1] = true
		counter++
	}

	if len(c.FormValue("numbers")) == 0 {
		Numbers = false
	} else {
		Numbers = true

		checkPos[2] = true
		counter++
	}

	if len(c.FormValue("symbols")) == 0 {
		Symbols = false
	} else {
		Symbols = true

		checkPos[3] = true
		counter++
	}

	generatePassword(counter, checkPos)

	c.JSON(200, "generated password: " + global.Password)

	return nil
}

func generatePassword(max int, pos [4]bool) {

	for i := 0; i < Length; i++ {
		field := randField(max)

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
			if Uppercase {
				global.Password += string(global.Uppercase[rand.Intn(len(global.Uppercase))])
			} else {
				i--
			}
		case 1:
			if Lowercase {
				global.Password += string(global.Lowercase[rand.Intn(len(global.Lowercase))])
			} else {
				i--
			}
		case 2:
			if Numbers {
				global.Password += string(global.Numbers[rand.Intn(len(global.Numbers))])
			} else {
				i--
			}
		case 3:
			if Symbols {
				global.Password += string(global.Symbols[rand.Intn(len(global.Symbols))])
			} else {
				i--
			}

		}

	}

}

func randField(max int) int {
	values := make([]int, 30)

	for i := 0; i < 30; i++ {
		values[i] = rand.Intn(max)
	}

	return values[rand.Intn(30)]
}
