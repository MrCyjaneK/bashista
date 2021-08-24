package main

import (
	"github.com/sqweek/dialog"
)

func main() {
	ok := true
	for ok {
		ok = dialog.Message("%s", getQuote()).Title("Here goes your quote").YesNo()
	}
}
