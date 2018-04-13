package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//import "strconv"

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

type area struct {
	W   int
	H   int
	out []string
}

func (a *area) fromOutside() {
	for i, v := range a.out {
		if i == 0 || i == a.H-1 {
			a.out[i] = strings.Replace(v, ".", "X", -1)
		} else {
			runes := []rune(v)
			if string(runes[0]) == "." {
				v = "X" + v[1:]
			}
			if string(runes[a.W-1]) == "." {
				v = v[:a.W-1] + "X"
			}
			a.out[i] = v
		}
	}
}

func (a *area) NeartoX(i, j int) {
	if 0 <= i && i < a.H && 0 <= j && j < a.W {
		r := string([]rune(a.out[i])[j])
		//fmt.Println(r)
		if r == "." {
			a.out[i] = a.out[i][:j] + "X" + a.out[i][j+1:]
			a.NeartoX(i-1, j)
			a.NeartoX(i, j-1)
			a.NeartoX(i+1, j)
			a.NeartoX(i, j+1)
		}
	}
}

func (a *area) Surround(i, j int) {
	if 0 <= i && i < a.H && 0 <= j && j < a.W {
		r := string([]rune(a.out[i])[j])
		if r == "X" {
			a.out[i] = a.out[i][:j] + "z" + a.out[i][j+1:]
		}
	}

}

func (a *area) PrintOut() {
	for _, v := range a.out {
		//a.out[i] = strings.Replace(v, "X", ".", -1)
		//fmt.Println(a.out[i])
		fmt.Println(strings.Replace(v, "X", ".", -1))

	}
}

func (a *area) StartAttck() {
	a.fromOutside()
	for i, v := range a.out {
		for j := range v {
			r := string([]rune(v)[j])
			if r == "X" {
				a.NeartoX(i-1, j)
				a.NeartoX(i, j-1)
				a.NeartoX(i+1, j)
				a.NeartoX(i, j+1)
			}
		}
	}

	for i, v := range a.out {
		for j := range v {
			r := string([]rune(v)[j])
			if r == "B" {
				a.Surround(i-1, j)
				a.Surround(i, j-1)
				a.Surround(i+1, j)
				a.Surround(i, j+1)
				a.Surround(i-1, j-1)
				a.Surround(i+1, j-1)
				a.Surround(i+1, j+1)
				a.Surround(i-1, j+1)

			}
		}
	}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var W, H int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &W, &H)
	out := make([]string, H)
	for i := 0; i < H; i++ {
		scanner.Scan()
		ROW := scanner.Text()
		out[i] = ROW
	}
	Outarea := &area{W, H, out}
	//Outarea.fromOutside()
	Outarea.StartAttck()
	Outarea.PrintOut()
}
