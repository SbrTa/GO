package main

import "fmt"

func main() {
	var n int;
	fmt.Println("if else :");
	fmt.Printf("n = ");
	fmt.Scan(&n);
	if n>100{
		fmt.Printf("n = %d >> 100\n",n);
	} else if n<100{
		fmt.Printf("n = %d << 100\n",n);
	}else{
		fmt.Printf("n = %d\n",n);
	}


	fmt.Println("Switch Case:");
	fmt.Printf("n = (1 / 2) ");
	fmt.Scan(&n);
	switch n{
		case 1:{
			fmt.Printf("n = 1\n");
		}
		case 2:{
			fmt.Printf("n = 2\n");
		}
		default:{
			fmt.Printf("None\n");
		}
	}
}
