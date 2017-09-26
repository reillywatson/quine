package main

import "fmt"

var quine = `package main

import "fmt"

var quine = 
var quote = string(96)

func main() {
	fmt.Println(quine[:40] + quote + quine + quote + "\n" + quine[41:])
}`
var quote = string(96)

func main() {
	fmt.Println(quine[:40] + quote + quine + quote + "\n" + quine[41:])
}
