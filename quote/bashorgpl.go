package quote

import (
	_ "embed"
	"math/rand"
	"strings"
)

//go:embed bashorgpl.txt
var quotes string

func BashOrgPl() string {
	qs := strings.Split(quotes, "\n%\n")
	return qs[rand.Intn(len(qs)-1)]
}
