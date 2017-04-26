package main
import "fmt"

func main(){
	if 5%2 == 0 {
		fmt.Println("5 is even")
	} else {
		fmt.Println("5 is odd")
	}	
	
	var a, b int=5, 10
	fmt.Println("Value of a",a,"Value of b",b)
	
	if a%2==0 {
		fmt.Println("a==b")
	} else {
		fmt.Println("a!=b")
	}

}

