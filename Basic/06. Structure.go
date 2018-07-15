package main

import "fmt"

type student struct {
	name string;
	id int;
	cgpa float64;
}

func main(){
	var s[100] student;
	var n int;
	fmt.Printf("Total student : ");
	fmt.Scan(&n);

	var cgpa float64;
	var name string;
	for i:=1; i<=n; i++ {
		fmt.Printf("Student %d:\n",i);
		fmt.Printf("Name : ");
		fmt.Scan(&name);
		fmt.Printf("CGPA : ");
		fmt.Scan(&cgpa);
		s[i].id = i;
		s[i].name = name;
		s[i].cgpa = cgpa;
	}

	for i:=1; i<=n; i++{
		fmt.Printf("ID : %d \nName : %s \nCGPA : %.2f\n\n",s[i].id,s[i].name,s[i].cgpa);
	}
}
