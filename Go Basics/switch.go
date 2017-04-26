package main 
import "fmt"

func main(){

	a:="Reporting from switch demo\n"
	fmt.Print(a)
	count:=2
	
	switch count {
	case 1:
			fmt.Println("Count is 1")
	case 2:
			fmt.Println("Count is 2")
	default:
			fmt.Println("Count is ",count)
	
	}
	
}
