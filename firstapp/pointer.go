package main
import "fmt"

func main(){
	s := "Hello world"

	p := &s
	fmt.Println(p)
	fmt.Println(*p)

	*p = "Hello go"
	fmt.Println(s)
}