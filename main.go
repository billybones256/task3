package main

import "fmt"

func main() {
	fmt.Print(SoupMaker(6))
}

func SoupMaker(days int) string {
	menu := [5]string{"щи", "борщ", "щавелевый суп", "овсяный суп", "суп по-чабански"}
	var result string
	for i := 0; i < days; i++ {
		result += menu[i%5] + "\n"
	}
	result = result[:len(result)-1]
	return result
}
