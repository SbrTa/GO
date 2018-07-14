package main

import(
	"fmt"
)
func main()  {

	fmt.Printf("Integer: a & b = ");
	var a,b int;
	fmt.Scan(&a,&b);
	fmt.Printf("%d %d\n",a,b);

	fmt.Printf("Float: x = ");
	var x float64;
	fmt.Scan(&x);
	fmt.Printf("%0.3f\n",x);

	fmt.Printf("String: s = ")
	var s string;
	fmt.Scan(&s);
	fmt.Printf("%s ; length = %d\n",s,len(s));
}