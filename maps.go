package main
import "fmt"

func main()  {
	 
	// var x map[string]int
	// x := make(map[string]int)
	// x["a"] = 10
	// fmt.Println(x["b"])
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 12
	fmt.Println(m)

	v1 := m["k1"]
	fmt.Println(v1)

	fmt.Println("len:", len(m))

	delete(m, "k1")
	fmt.Println("len2:", len(m))

	_,abc := m["k1"]
	fmt.Println("abc ",abc)

	n := map[string]int {"foo ":1, "bar ":3 }
	fmt.Println(n)
}