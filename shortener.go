package shortener

import (
	"log"
	"strings"
)

type Generator struct {
	Base      []string
	MaxDigits int
}

func (g *Generator) Encode(id int64) string {
	if id > g.MaxId() {
		log.Fatalf("number %d to encode is bigger than the maximum allowed (%d)", id, g.MaxId())
	}
	aux := id
	ret := ""
	n := int64(g.N())
	l := g.L()
	for digit := g.MaxDigits - 1; digit >= 0; digit-- {
		a := aux % n
		ret += g.Base[digit%l][a:a+1]
		aux = (aux - a) / n
	}
	return ret
}

func (g *Generator) Decode(alias string) int64 {
	if len(alias) > g.MaxDigits {
		log.Fatalf("alias %s to decode is bigger than the max bits allowed (%d)", alias, g.MaxDigits)
	}
	var x, id int64
	n := int64(g.N())
	x = 1
	for digit := 0; digit < g.MaxDigits; digit++ {
		c := alias[digit:digit+1]
		v := strings.Index(g.Base[g.MaxDigits - digit - 1], c)
		id += int64(v) * x
		x *= n
	}
	return id
}

func pow(x, y int) int64 {
	if y == 0 {
		return 1
	}
	if x == 0 {
		return 0
	}
	X := int64(x)
	res := X
	for i := 1; i < y; i++ {
		res *= X
	}
	return res
}

func (g *Generator) N() int {
	return len(g.Base[0])
}

func (g *Generator) L() int {
	return len(g.Base)
}

func (g *Generator) MaxId() int64 {
	var nant, n int
	for i := 0; i < len(g.Base); i++ {
		n = len(g.Base[i])
		// just to unsure all strings have the same length
		if i > 0 && n != nant {
			log.Fatal("error")
		}
		nant = n
	}
	return pow(n, g.MaxDigits)
}

var AliasGenerator = &Generator{MaxDigits: 8,
	Base: []string{
		"VaRCScevfNjk9HDGlmTobPq7rIsYtux5yzAgBE4pFihUJnK3LMdO1Q6XZ",
		"JK3LMNtuvx5yzABCDE4FGHIOP1QRSTUV6XYZabcdefghijk9lmnopq7rs",
		"STUV6XYZabcdefghijk9lmK3LMNnopq7rstuvOP1QRx5yzABCDE4FGHIJ",
		"OP1QRk9lmnopq7rsSTdzABCDEuJK3LMN4FGHefghijtvx5yUV6XYZabcI",
		"lmnopDE4FGHefghijtuJK3LMNvx5q7rsSTdzABCIOPyUV6XYZabc1QRk9",
		"vx5yzABCDE4FGHUV6XYZk9lmnopq7rsOP1QRSTdefghijtuabcIJK3LMN",
		"TdzABCIOPyUV6XYZabc1QRk9lmnopDE4FGHefghijtuJK3LMNvx5q7rsS",
		"QRSTdefghijtuabcIJK3LMNvx5yzABCDE4FGHUV6XYZk9lmnopq7rsOP1",
		"4FGHefghijtOP1QRk9lmnopq7rsSTdzABCDEuJK3LMNvx5yUV6XYZabcI",
		"q7rsSTdzABCIOPyUV6XYZabc1QRk9lmnopDE4FGHefghijtuJK3LMNvx5",
	},
}

