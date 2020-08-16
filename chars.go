package main

type Chars []Char


func (c Chars) Flatten() []int {
	flattened := make([]int, 0)
	for i := 0; i < len(c); i++ {
		flattened = append(flattened, c[i].Flatten()...)
	}
	return flattened
}

func (c Chars) Size() int {
	i := 0
	for _, cc := range c {
		i += cc.Size()
	}
	return i
}
