package main
func main() {
	i := GetValue()

	switch i.(type) {  //不能编译通过，因为type只能用在interface里面
	case int:
		println("int")
	case string:
		println("string")
	case interface{}:
		println("interface")
	default:
		println("unknown")
	}

}

func GetValue() int {
	return 1
}