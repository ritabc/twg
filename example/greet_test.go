package example

import "fmt"

func ExampleHello() {
	greeting := Hello("Jon")
	fmt.Println(greeting)

	// Output:
	// Hello, Jon
}

func ExampleHello_spanish() {
	greeting := Hello("Juan")
	fmt.Println(greeting)

	// Output:
	// Hello, Juan
}

func ExamplePage() {
	checkIns := map[string]bool{
		"Bob":   true,
		"Alice": false,
		"Eve":   false,
		"John":  false,
	}
	Page(checkIns)

	// Unordered output:
	// Paging Alice; please see the front desk to check in.
	// Paging Eve; please see the front desk to check in.
	// Paging John; please see the front desk to check in.
}
