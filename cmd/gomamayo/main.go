/*
 */
package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/jugesuke/gomamayo"
)

const stdinPrefix string = ">> "

func init() {
	fmt.Println("This is go-mamayo")
	fmt.Println("Type \"q\" to exit.")
	fmt.Print(stdinPrefix)
}

func main() {
	g, err := gomamayo.Init()
	if err != nil {
		os.Exit(1)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if scanner.Text() == "q" {
			os.Exit(0)
		} else {
			r := g.Analyze(scanner.Text())
			// fmt.Println(r.IsGoMamayo)
			if r.IsGomamayo {
				fmt.Printf("%d項%d次のゴママヨです\n", r.Ary, r.Degree)
			} else {
				fmt.Println("ゴママヨではありません")
			}
			fmt.Print(stdinPrefix)
		}
	}
}
