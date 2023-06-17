package stuff

import(
	"fmt"
)

func prime(number int)bool{
	for i:=2;i<number;i++{
		if number%i == 0 {
			return false
		}
	}
	return true
}

func Sequence(a int, b int){
	for i:=a; i<=b; i++{
		if prime(i){
			fmt.Printf("%d is prime number\n", i )
		}else{
			fmt.Printf("%d is not prime number\n",i)
		}
	}
}
