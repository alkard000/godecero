package variables

import "fmt"

func ShowInt() {
	var intcomun int
	intOf32 := int32(0)
	intOf64 := int64(100)

	fmt.Println("intcomun= ", intcomun)
	fmt.Println("int32= ", intOf32)
	fmt.Println("int64= ", intOf64)
}
