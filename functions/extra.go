package functions

import "fmt"

func TestPrintFunction() {
	fmt.Println("testPrintFunction FROM FUNCTIONS PACKAGE")
}

var TestVariable int = 8429043824230

func main() {
	fmt.Println("WILL NOT EXECUTE MAIN FUNCTION WHEN IMPORTING")
}
