package main

import (
	"fmt"

	"github.com/frumious/go15vendortest2"

	"vendoredthing"
)

func main() {
	var vi1, vi2 interface{}

	fmt.Println("Hi there from main.")
	fmt.Println("About to call vendoredthing.DoSomething()")
	vendoredthing.DoSomething()
	fmt.Println("Done calling vendoredthing.DoSomething()")

	vt1 := vendoredthing.GetData()
	vt2 := go15vendortest2.GetData()

	fmt.Printf("vt1 is %#v\n", vt1)
	fmt.Printf("vt2 is %#v\n", vt2)
	fmt.Printf("vt1 is type %T\n", vt1)
	fmt.Printf("vt2 is type %T\n", vt2)

	vi1 = vt1
	vi2 = vt2

	vt3 := vi1.(vendoredthing.Data)
	vt4 := vi2.(vendoredthing.Data)
	fmt.Println(vt3, vt4)

	fmt.Println("Bye bye")
}
