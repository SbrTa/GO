package main

import (
	"fmt"
)

func fact(n,x int) int{
	if n==0{
		return x;
	}
	return fact(n-1,x*n);
}

func main(){
	fmt.Printf("Factorial of n (n!) : \n");
	fmt.Printf("n = ");
	var n int;
	fmt.Scan(&n);
	fmt.Printf("n! = %d\n",fact(n,1));
}
