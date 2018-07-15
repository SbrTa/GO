package main

import "fmt"
import s "strings"

var p = fmt.Println

func main(){
	fmt.Printf("Sring:");
	var str string;
	fmt.Scan(&str);
	fmt.Printf("String: %s\n",str);
	fmt.Printf("Length: %d\n",len(str));
	fmt.Printf("Print string characterwise reverse: ")
	for i:=len(str)-1; i>=0;i-- {
		fmt.Printf("%c",str[i]);
	}
	fmt.Println();

	p("Contains:  ", s.Contains("test", "es"))
	p("Count:     ", s.Count("test", "t"))
	p("HasPrefix: ", s.HasPrefix("test", "te"))
	p("HasSuffix: ", s.HasSuffix("test", "st"))
	p("Index:     ", s.Index("test", "e"))
	p("Join:      ", s.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", s.Repeat("a", 5))
	p("Replace:   ", s.Replace("foo", "o", "0", -1))
	p("Replace:   ", s.Replace("foo", "o", "0", 1))
	p("Split:     ", s.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", s.ToLower("TEST"))
	p("ToUpper:   ", s.ToUpper("test"))
	p()
}
