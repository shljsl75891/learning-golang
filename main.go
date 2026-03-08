package main // Let's go compiler knows that this program is supposed to be run as a standalone program, instead of a library which will be imported in other programs

import (
	"fmt" // Import the formatting package from the standard library
)

func getCords() (x, y int) { // Named returns = These named vairables are automatically declared at the top of the function body. These should be used for documenting short functions only
	x, y = 5, 6;
	return; // Naked return = The latest value of named variables declared at the top of function are returned

	// return 9, 10; ==> Explicit return
}

func main() { // =====> ENTRY POINT OF THE PROGRAM
	fmt.Println("Starting the Textio server")
		// Sad way of declaring variables
	var mySkillIssues int
	fmt.Println(mySkillIssues)

	var declaredStr string;
	fmt.Printf("Declared String: %s\n", declaredStr)

	// Goated way using walrus operator (:=).
	// It declares and assigns value in the same line along with automatic type inference based on value
	aiSkillIssues := 1000

	fmt.Println(aiSkillIssues)
	printHelloWorld();

	const firstName = "Sahil"
	const lastName = "Jassal"

	x, y := getCords();
	fmt.Printf("X: %d, Y: %d\n", x, y)

	learnSlices();
}

func printHelloWorld() {
	defer func () {
		if r := recover(); r != nil {
			fmt.Printf("Something went wrong: %v\n", r)
		}
	}();

	defer fmt.Println("Two times") 
	defer fmt.Println("World") // this function's execution is defer to just before enclosing function returns

	fmt.Println("Hello")

	a := account{
		balance: 25000.0,
	}
	balance,err := performBankingOperation(a, 200.0, "withdrawal")
	fmt.Println("Error: ", err)
	fmt.Printf("Balance after withdrawal: %.1f\n", balance)
}

func learnSlices() {
	arrays := [3]int{1, 2, 3};
	slices := []int{1, 2, 3, 4, 5};

	fmt.Printf("Array is fixed size: %v, but slices are dynamically sized (although just a abstracted underlying array): %v\n", arrays, slices);

	for i, element := range arrays { // similar to for i := 0; i < len(slices); i++, but modern way to write, and element is copy of the value, not reference
		fmt.Printf("array[%d]: %d ", i, element);
	}


	fmt.Printf("\nThe capacity of slice %v is %d and its length is %v\n", slices, cap(slices), len(slices))
	fmt.Println("Appending 6 in slices")
	// The underlying array will be doubled in size after appending one element, making illusion of dynamically sized array
	slices = append(slices, 6);
	fmt.Printf("The capacity of slice %v is %d and its length is %v\n", slices, cap(slices), len(slices))

	var zeroSlice []int;
	fmt.Printf("Zero slice: %v, and is nil: %t. Capacity: %d, Length: %d\n", zeroSlice, zeroSlice == nil, cap(zeroSlice), len(zeroSlice));

	emptySlice := []int{};
	fmt.Printf("Empty slice: %v, and is nil: %t. Capacity: %d, Length: %d\n", emptySlice, emptySlice == nil, cap(emptySlice), len(emptySlice));
}

