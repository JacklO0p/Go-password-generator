package controllers

import (
	"fmt"
	"math/rand"
	"password_generator/global"
	"strconv"
	"time"

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

		counter++
	}

	if len(c.FormValue("lowercase")) == 0 {
		Lowercase = false
	} else {
		Lowercase = true

		counter++
	}

	if len(c.FormValue("numbers")) == 0 {
		Numbers = false
	} else {
		Numbers = true

		counter++
	}

	if len(c.FormValue("symbols")) == 0 {
		Symbols = false
	} else {
		Symbols = true

		counter++
	}

	generatePassword(counter)

	return nil
}

func generatePassword(max int) {

	for i := 0; i < Length; i++ {
		field := randField(max)

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

	fmt.Println("Password: " + global.Password)
}

func randField(max int) int {
	values := make([]int, 30)

	for i := 0; i < 30; i++ {
		rand.Seed(time.Now().UnixNano())
		values[i] = rand.Intn(max)
	}

	return values[rand.Intn(30)]
}
