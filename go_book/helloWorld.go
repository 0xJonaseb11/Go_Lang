package main

import "fmt";

func average(xs []float64) float64 {
	// panic("Not implemented"); // built-in function to cause runtime-Error

	total := 0.0;
	for _, v := range xs {
		total += v;
	}
	return total / float64(len(xs));
}

// can also return multiple values
func f () (int, int) {
	
	return 5,6;
}

// Variadic functions
/*
*By using ... before the type name of the last parame-
ter you can indicate that it takes zero or more of those
parameters. In this case we take zero or more ints.
*/

func add(args ...int) int {
	total := 0;

	for _, v := range args {
		total += v;
	}

	return total
}




func main () {

	/** Closure
    * It is possible to create functions inside of functions
    */

	add_ := func(x, y int) int {
		return x + y;
	}
	fmt.Println("Using closure to add we get", add_(12,12));


	// implementing variadic funcns with interface{}
	// fmt.Println(a ...interface{} ) (n int, err error )

	// applying variadic funcns logic ot slices too
	xs_ := []int{1,2,3,4,}
	fmt.Println("Variadically added slice xs_ :=",add(xs_...))

	// add - variadic
	fmt.Println("Using Variadic funcn to add := ", add(1,2,3,4,5));
	// add();
/*
	// handle error
	(x, err := f())
	// handle success
	(x, ok := f())
*/
	x_, y_ := f();
	fmt.Println("x_ is: ", x_ ,"and y_ is: ", y_);

	xs := [] float64 {12,345,56,78,23}
	fmt.Println("The average of xs is: ", average(xs));

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
	slice2_ := make([]int, 3) // 3 is the length of slice
	copy(slice2_, slice1_)
	fmt.Println("Before copying slice1 into slice2:", slice1_, "After copying slice1", slice2_);

	// Maps
	// var mapped map[string]int
	mapped_ := make(map[string]int)
	mapped_["key"] = 10
	fmt.Println(mapped_["key"])

	// delete an element with built-in delete() method
	delete(mapped_, "key");
	fmt.Println(mapped_["key"]) // outputs 0 to indicate deletion

	// example program
	elements := make(map[string]string);
	elements["H"] = "Hydrogen";
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["He"]); // outputs Helium

	// check for zero value - usual way
	fmt.Println(elements["Un"] == "");// outputs true indicating it doesn't exist
	fmt.Println(elements["H"] == ""); // outputs false - indicating it exists

	// check for zero value - Go way
	if name, ok := elements["Un"]; ok {
		fmt.Println(name, ok);
	}

	// make our program complex - add room temps
	elements_ := map[string] map[string] string {
		"H" : map[string] string {
			"name": "Hydrogen",
			"state": "gas",
		},
		"He": map[string] string {
			"name": "Helium",
			"state": "gas",
		},
		"Li": map[string] string {
			"name": "Lithium",
			"state": "gas",
		} ,
		"Be": map[string] string {
			"name": "Beryllium",
			"state": "solid",
		},
		"B": map[string] string {
			"name": "Boron",
			"state": "solid",
		} ,
		"C": map[string] string {
			"name": "Carbon",
			"state": "solid",
		},
		"N": map[string] string {
			"name": "Nitrogen",
			"state": "gas",
		} ,
		"O": map[string] string {
			"name": "Oxygen",
			"state": "gas",
		},
		"F": map[string] string {
			"name": "Fluorine",
			"state": "gas",
		} ,
		"Ne": map[string] string {
			"name": "Neon",
			"state": "gas",
		},
	}

	if el, ok := elements_["Li"]; ok {
		fmt.Println("Selected element", el["name"], "of state", el["state"]);
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

func elements() {
	elements := make(map[string]string);
	elements["He"] = "Hydrogen";
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	fmt.Println(elements["He"]);

}