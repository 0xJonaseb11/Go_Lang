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
	} else if i == 4 {
		fmt.Println("Four")
	} else if i == 5 {
		fmt.Println("Five")
	} else {
		fmt.Println("Invalid input")
	}

	// Opt to use switch case instead
	switch i {
	  case 0: fmt.Println("Zero");
	  case 1: fmt.Println("One");
	  case 2: fmt.Println("Two");
	  case 3: fmt.Println("Three");
	  case 4: fmt.Println("Four");
	  case 5: fmt.Println("Five");
	default: fmt.Println("Invalid input");
	}

	// Arrays 
	// syntx - 
	var array[5] int;
	array[4] = 100;
	fmt.Println(array);

	var marks[5] float64
	marks[0] = 98;
	marks[1] = 45;
	marks[2] = 76;
	marks[3] = 89;
	marks[4] = 100;

	// also shortens the array - shorter syntax
	_marks := [7] float64 { 65,87,98,77,99,55,33 }
	fmt.Println(_marks);

	var total float64 = 0;

	for i := 0; i < 5; i++ {
		total += marks[i];
	}
	fmt.Println(total / 5)

	/*  To encounter for every change that can occur,
	* we need to use the actual length of the array 
	* rather than hard coding it
    
	*/

	for i := 0; i < len(marks); i++ {
		total += marks[i];
	}
	fmt.Println(total / float64(len(marks)));

	// use a specail form of for loop
	for _/*means we don't need this iterator variable*/, value := range marks {
		total += value;
	}
	fmt.Println(total / float64(len(marks)));

	// slices - segment of an array -- dynamic array
	var slice [] float64 
	fmt.Println("sample slice", slice);
	// create a slice - use built-in make() function
	slice_ := make([] float64, 5, 10);
	fmt.Println(slice_)
	// where 5 is the length of the array and 10 is the capacity of the underlying array which the slice points to.

	/*
	* Another way to create slices is the use of [low:high] expression
	*/
	arr := [9] float64 {1,2,3,4,5,6,7,8,9}
	sliced := arr[4:8]
	sliced_len := arr[0:len(arr)]
	fmt.Println(sliced);
	fmt.Println(sliced_len); // for convenience

	// slice functions
	// append()
	slice1 := []int{1,2,3}
	slice2 := append(slice1,4,5,6);
	fmt.Println("Before appending: ", slice1 , "After appending:", slice2 );

	// copy - copies content of slice1 into slice2
	slice1_ := []int{1,2,3}
	slice2_ := make([]int, 3)
	copy(slice2_, slice1_)
	fmt.Println("Before copying slice1 into slice2:", slice1_, "After copying slice1", slice2_);

	// Maps
	var map map[string]int


	
	
	





}

func prompt() {

	fmt.Println("Please enter a number:\n");
	var input float64

	fmt.Scanf("%f", &input);

	fmt.Println("You entered %f", &input)

	output := input * 2;

	fmt.Println("Output:\n", output);
}