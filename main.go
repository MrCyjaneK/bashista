package main

import (
	"math/rand"
	"time"

	"github.com/sqweek/dialog"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	ok := true
	for ok {
		ok = dialog.Message("%s", getQuote()).Title("Here goes your quote").YesNo()
	}
}
