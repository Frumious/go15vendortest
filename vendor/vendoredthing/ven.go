package vendoredthing

import (
	"fmt"
)

type Data struct {
	Msg   string
	Count int
}

func DoSomething() {
	fmt.Println("I am a vendored something")
}

func GetData() Data {
	return Data{"hello", 2}
}
