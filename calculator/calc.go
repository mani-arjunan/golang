package calc

import "fmt"

type Test struct {
	Name string
	id   *int // kinda private variable, coz if it is not populated it would be nil, since its a pointer
}

func Calculator() {
	var operation string
	var num1, num2 int

	fmt.Printf("Enter the operation: ")
	fmt.Scanf("%s", &operation)

	fmt.Printf("Performing operation: %s\n", operation)

	if operation != "add" && operation != "sub" {
		str := fmt.Sprintf("Operation %s not supported", operation)
		panic(str)
	}

	fmt.Printf("Enter first Number: ")
	fmt.Scanln(&num1)
	temp1 := 10
	test := Test{Name: "Manikandan", id: &temp1}

	fmt.Println(*test.id, "==")
	fmt.Printf("Enter second Number: ")
	fmt.Scanln(&num2)

	switch operation {
	case "add":
		fmt.Println(num1 + num2)
	case "sub":
		fmt.Println(num1 - num2)
	}
}
