package main

import "fmt"

var chars map[rune]Char
var (
	g, e, t, a, l, i, f, space Char
)

//we represent chars in 7*x
func init() {
	chars = map[rune]Char{}

	g = Char{
		{1, 2, 2, 2, 2},
		{2, 1, 1, 1, 1},
		{2, 1, 1, 1, 1},
		{2, 1, 1, 1, 1},
		{2, 1, 1, 2, 2},
		{2, 1, 1, 1, 2},
		{1, 2, 2, 2, 2},
	}
	chars['g'] = g

	e = Char{
		{2, 2, 2, 2},
		{2, 1, 1, 1},
		{2, 1, 1, 1},
		{2, 2, 2, 1},
		{2, 1, 1, 1},
		{2, 1, 1, 1},
		{2, 2, 2, 2},
	}
	chars['e'] = e

	t = Char{
		{2, 2, 2, 2, 2},
		{1, 1, 2, 1, 1},
		{1, 1, 2, 1, 1},
		{1, 1, 2, 1, 1},
		{1, 1, 2, 1, 1},
		{1, 1, 2, 1, 1},
		{1, 1, 2, 1, 1},
	}

	chars['t'] = t
	space = Char{
		{1},
		{1},
		{1},
		{1},
		{1},
		{1},
		{1},
	}

	chars[' '] = space

	a = Char{
		{1, 1, 2, 1, 1},
		{1, 2, 1, 2, 1},
		{2, 1, 1, 1, 2},
		{2, 1, 1, 1, 2},
		{2, 2, 2, 2, 2},
		{2, 1, 1, 1, 2},
		{2, 1, 1, 1, 2},
	}

	chars['a'] = a

	l = Char{
		{2, 1, 1, 1, 1},
		{2, 1, 1, 1, 1},
		{2, 1, 1, 1, 1},
		{2, 1, 1, 1, 1},
		{2, 1, 1, 1, 1},
		{2, 1, 1, 1, 1},
		{1, 2, 2, 2, 2},
	}

	chars['l'] = l

	i = Char{
		{2},
		{1},
		{2},
		{2},
		{2},
		{2},
		{2},
	}

	chars['i'] = i

	f = Char{
		{2, 2, 2, 2},
		{2, 1, 1, 1},
		{2, 1, 1, 1},
		{2, 2, 2, 1},
		{2, 1, 1, 1},
		{2, 1, 1, 1},
		{2, 1, 1, 1},
	}

	chars['f'] = f
}

func string2Matrices(str string) Chars {
	runes := []rune(str)
	chrs := make([]Char, 0)
	for _, r := range runes {
		if c, ok := chars[r]; ok && c.Height() == 7{
			chrs = append(chrs, c)
			if r!= ' ' {
				chrs = append(chrs, space)
			}
		} else {
			fmt.Println("unknown or invalid char :", r)
		}
	}
	return chrs
}
