package main
import "fmt"

func main() {
	const a = 42
	var b  int = a
	fmt.Println(b)

	const c = iota
	fmt.Println(c)

	const (
		d = 2*5 //10
		e //10
		f = iota //2
		g
		h = 10* iota
	)
	fmt.Println(d, e, f, g, h)
}