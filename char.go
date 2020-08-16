package main

type Char [][]int

func (c Char) Flatten() []int {
	flattened := make([]int, c.Size())

	for rowNum := 0; rowNum < len(c); rowNum++ {
		for colNum := 0; colNum < len(c[rowNum]); colNum++ {
			flattened[(c.Height()*colNum)+rowNum] = c[rowNum][colNum]
		}
	}
	return flattened
}

func (c Char) Width() int {
	if len(c) == 0 || len(c[0]) == 0 {
		return 0
	}

	return len(c[0])
}

func (c Char) Height() int {
	return len(c)
}

func (c Char) Size() int {
	return c.Height() * c.Width()
}