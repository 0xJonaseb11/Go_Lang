package main

import "fmt";

func main () {
	fmt.Println("Hello world");

	// numbers
	fmt.Println("1 + 1 = ", 1 + 1);
	fmt.Println("1.0 + 1.0 = ", 1.0 + 1.0);
	fmt.Println("32132 × 42452 = ", 32132 * 42452);
	
	// strings
	fmt.Println(len("hello world"));
	// access individual characters
	fmt.Println("hello world"[1]);
	// string concatenation
	fmt.Println("hello " + "world");

	// booleans
	fmt.Println(true && true);
	fmt.Println(true && false);
	fmt.Println(true || true);
	fmt.Println(true || false);
	fmt.Println(!true);

	// implement the use of variables
	var greeting string = "Hello world";
	fmt.Println(greeting);
	// alternatively
	var x string
	x = "first"
	fmt.Println(x);
	x = "second"
	fmt.Println(x);
	x = x + " third";
	// alternatively
	x += " fourth"
	fmt.Println(x)

	// check for equity
	var x1 string = "hello";
	var y1 string = "world";
	fmt.Println(x1 == y1); // prints false since strings are not the same

	// Creating a variable with a starting value 
	var num = 10;
	// num := 5
	fmt.Println(num);

	// support for defining multiple variables
	var (
		a = 1
		b = 2
		c = 3
		d = 4
	)
	fmt.Println(a,b,c,d);

	// getting user input
	fmt.Println("Please enter a number:\n");

	var input float64

	fmt.Scanf("You entered %f", &input);

	fmt.Println("You entered ", input)

	output := input * 2;

	fmt.Println("Output:\n", output);

	// Control statements
	// for
	i := 1
	for i <= 10 {
		fmt.Println(i);
		i++;
	}
	// alternatively
	for i:= 1 ; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, ":Even");
		} else {
			fmt.Println(i, ":Odd");
		}
	}

	nbr :=  3

	// if statement
	if nbr == 0 {
		fmt.Println("zero");
	} else if nbr == 1 {
		fmt.Println("One");
	} else if nbr == 2 {
		fmt.Println("Two");
	} else if nbr == 3 {
		fmt.Println("Three");
	} else {
		fmt.Println("Invalid input")
	}


}

func prompt() {

	fmt.Println("Please enter a number:\n");
	var input float64

	fmt.Scanf("%f", &input);

	fmt.Println("You entered %f", &input)

	output := input * 2;

	fmt.Println("Output:\n", output);
}