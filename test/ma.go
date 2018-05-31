package main

import "fmt"


type g2048 [4][4]int

func (t *g2048) MirrorV() {
    tn := new(g2048)
    for i, line := range t {
		fmt.Println(line)
        for j, num := range line {
            tn[len(t)-i-1][j] = num
        }
    }
    *t = *tn
}


func (t *g2048) Print() {
    for _, line := range t {
        for _, number := range line {
            fmt.Printf("%2d ", number)
        }
        fmt.Println()
    }
    fmt.Println()
    tn := g2048{{1, 2, 3, 4}, {5, 8}, {9, 10, 11}, {13, 14, 16}}
    *t = tn

}


func main() {
	fmt.Println("origin")
    t := g2048{{1, 2, 3, 4}, {5, 8}, {9, 10, 11}, {13, 14, 16}}
	t.MirrorV();
}
