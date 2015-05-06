//calc.go
package main

import "os"
import "fmt"
import "../pkg/simplemath"//指明包名的相对路径
import "strconv"

var Usage = func() {
	fmt.Println("USAGE: clac command [arguments]...")
	fmt.Println("\nThe commands are:\n\tadd\tAddition of two values.\n\tsqrt\tSquare root of a non_nective value")
}

func main() {
	args := os.Args
	if args == nil || len(args) < 2 {
		Usage()
		return
	}
	fmt.Printf("%s\n",args[0]);
	fmt.Printf("%s\n",args[1]);
	switch args[1] {
	case "add":
		if len(args) != 4 {
			fmt.Println("USAGE:calc add <integer1><integer2>")
			return
		}
		v1, err1 := strconv.Atoi(args[2])
		v2, err2 := strconv.Atoi(args[3])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE:calc add <integer1><integer2>")
			return
		}
		ret := simplemath.Add(v1, v2)
		fmt.Println("Result: ", ret)
	case "sqrt":
		if len(args) != 2 { //参数不为2个：可执行文件名+参数
			fmt.Println("USAGE: calc sqrt <integer?")
			return
		}
		v, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("USAGE: calc sqrt<integer>")
			return
		}
		ret := simplemath.Sqrt(v)
		fmt.Println("Result: ", ret)
	default:
		Usage()
	}
	v := make([]int, 20);
	fmt.Println(v);

	m := make(map[string]int, 0)
	m["one"] = 1
	m["two"] = 2
	m["three"] = 3
	fmt.Println(m)//直接一次性全部printf
	for key, value := range m{
		fmt.Printf("%s->%d\n",key,value);
	}
	fmt.Printf("first key:%s value:%d\n","one", m["one"])
}
