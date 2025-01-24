package main
import "fmt"

func main(){
	var a string
	a = "foo"

	fmt.Println(a)

	var b int = 99
	fmt.Println(b)

	d := 3.145 // defaulr float 64
	fmt.Println(d)

	var e int = int(d)
	fmt.Println(e)

	var f int8 = 54
	var g int16 = f

	var g int33 = f

}
