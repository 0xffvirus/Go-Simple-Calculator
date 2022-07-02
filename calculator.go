// Author 6ViRuS
// github : 6virus
// Language : Go 

package main
import "fmt"

func main() {
	var opt int
	fmt.Println("[1] + [2] - [3] * [4] / ")
	fmt.Printf("Choose : ")
	fmt.Scanln(&opt)
	if opt == 1 {
		var x,y float64
		fmt.Printf("First Num : ")
		fmt.Scanln(&x)
		fmt.Printf("Second Num : ")
		fmt.Scanln(&y)
		fmt.Println("Result :",sum(x,y))
	} else if opt == 2 {
		var x,y float64
		fmt.Printf("First Num : ")
		fmt.Scanln(&x)
		fmt.Printf("Second Num : ")
		fmt.Scanln(&y)
		fmt.Println("Result :",min(x,y))
	} else if opt == 3 {
		var x,y float64
		fmt.Printf("First Num : ")
		fmt.Scanln(&x)
		fmt.Printf("Second Num : ")
		fmt.Scanln(&y)
		fmt.Println("Result :",mul(x,y))
	} else if opt == 4 {
		var x,y float64
		fmt.Printf("First Num : ")
		fmt.Scanln(&x)
		fmt.Printf("Second Num : ")
		fmt.Scanln(&y)
		fmt.Println("Result :",dev(x,y))
	} else {
		fmt.Println("Please Choose a Number")
	}

}

func sum(x float64,y float64) float64 {
	return x + y
}
func min(x float64,y float64) float64 {
	return x - y
}
func mul(x float64,y float64) float64 {
	return x * y
}
func dev(x float64,y float64) float64 {
	return x / y
}