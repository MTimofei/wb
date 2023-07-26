// Разработать программу,
// которая в рантайме способна определить тип переменной: int, string, bool, channel
// из переменной типа interface{}.

package main

func main() {

}

func isType(i interface{}) {
	switch i.(type) {
	case int:
		return
	case string:
		return
	case bool:
		return
	case chan int:
		return
	}
	return
}
