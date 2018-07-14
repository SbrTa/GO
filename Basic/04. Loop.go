package main

import "fmt"

func main(){
	// uses of "break"  &  "continue" as like C/C++
	fmt.Println("For Loop: ");
	var n int;
	fmt.Printf("Sum of 1 to ? ");
	fmt.Scan(&n);
	s:=0;
	for i:=1; i<=n; i++{
		s+=i;
	}
	fmt.Printf("1+...+%d = %d\n",n,s);

	// no while loop
	fmt.Println("No while loop. But for can be used as like while in C/C++:");
	fmt.Printf("Prime numbers from 1 to ? ");
	fmt.Scan(&n);
	var x[10000] int;
	i:=2;
	for i<=n{
		j:=2;
		p:=1;
		for j*j<=i{
			if i%j==0{
				p = 0;
			}
			j++;
		}
		x[i] = p;
		i++;
		//fmt.Println(i);
	}
	for i:=2;i<=n;i++{
		if x[i]==1{
			fmt.Printf("%d ",i);
		}
	}
}
