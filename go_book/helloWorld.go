package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	);

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

// generate evene numbers - closure
func makeEvenGenerator() func () uint {
	i := uint(0)
	return func() (ret uint) {
	ret = i
	i += 2
	return
	}
}

// Recursion

func factorial(fact uint) uint {
	if fact == 0 {
		return 1
	}
	return fact * factorial(fact-1);
	
}

// Defer, Panic and Recover
/*
* Defer if a special statement that schedules a funcition call to run after function completes
* It is often used to when resources need to be freed in some way. forexample
* when we open a file we need to make sure it is closed later.
*/

func first () {
	fmt.Println("1st")
}

func second () {
	fmt.Println("2nd")
}

// Pointers
func zero(_xptr *int) {
	*_xptr = 0
}

func one(xptr *int) {
	*xptr = 1;
}

func square(x__ *float64) {
	*x__ = *x__ * *x__;
}

// structs and interfaces
func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
	}

	func rectangleArea(x1, y1, x2, y2 float64)float64 {
	l := distance(x1, y1, x1, y2)
	w := distance(x1, y1, x2, y1)
	return l * w
	}

func circleArea(x, y, r float64) float64 {
	return math.Pi * r*r;
}


type Circle struct {
	x, y, r float64;
}

//initialization
var c Circle;
// use new function to initialize too
var c_ = new(Circle);
// let's try to give it value
var  _c = Circle{ x: 0, y: 0, r: 5}
/*Or we can leave off the field names if we know the or-
* der they were defined:
*/
var __c = Circle{0, 0, 5}


   // Let's modify `circleArea` function so that it uses a Circle:
  

func _circleArea(c Circle) float64 {
	return math.Pi * c.r*c.r;
   }
   /*
   *keep in mind that arguments are always copied in GO, so if you modify a variable above, it is not entirely modified in a program
   * So to solve that, we'll use pointer  
   */

   func __circleArea(c *Circle) float64{
	return math.Pi * c.r*c.r;

   }

   // Methods
   /*
   * We have made our program better, but why not make it best with a special function called method
   */

   func (c *Circle) area() float64 { // (c *Circle) => is the receiver
	return math.Pi * c.r*c.r
   }

   // Use method to find the area of a rectangle
   type Rectangle struct {
	x1, y1, x2, y2 float64
   }

   func (r * Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x1, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
   }

   // Embedded types
   type Person struct {
	Name string
   }

   func (p *Person) Talk() string {
	fmt.Println("Hi, my name is ", p.Name);
	return p.Name;

   }

   // we need to create a new android
   type Android struct {
	Person
	Model string
   }

   // Interfaces
   type Shape interface {
	area() float64
   }

   // let's implement the area logic like we did above 
   func totalArea(shapes ...Shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
   }

   // Interfaces are also used as fields
   type MultiShape struct {
	shapes []Shape
   }

   /*
   We can even turn MultiShape itself into a Shape by giv-
   ing it an area method:
  */

  func (m *MultiShape) area() float64 {
	var area float64
	for _, s := range m.shapes {
		area += s.area()
	}
	return area
  }
  /*
  Now a MultiShape can contain Circles, Rectangles or
  even other MultiShapes.
  */

  // Concurrency
  /*
  * Go has rich support for concurrency with goroutines and channels
  */

  func goroutine(n int) {
	for i := 0; i < 10; i++ {
		fmt.Println(n, ":", i)
		// let's add some delay
		amt := time.Duration(rand.Intn(250));
		time.Sleep(time.Millisecond * amt);
	}
  }


  // Channels
  // member 1
  func pinger(ch chan string) {
	for i := 0; ; i++ {
		ch <- "ping"
	}
  }

  func printer(ch chan string) {
	for {
		msg := <- ch
		fmt.Println(msg) // alternatively `fmt.Println(<- c)`
		time.Sleep(time.Second * 1)
	}
  }
   // member 2
  func ponger(ch chan string) {
	for i := 0; ; i++ {
		ch <- "pong"
	}
  }

  // Channel direction
//   in the pinger()
func pinger_(ch chan <- string){}; // sender
func printer_(ch <- chan string){}; // receiver

// select - works like switch but for channels

// Buffered channels
/*It's also possible to pass a second parameter to the
* make function when creating a channel:
*/
ch_ := make(chan int, 1);









func main () {

	x_s := []float64{1,2,3,4}
	avg := m.Average(x_s);
	
	// select statement in channels
	c1 := make(chan string)
	c2 := make(chan string)

	// start communication
	go func() {
		for {
			c1 <- "From 1"
			time.Sleep(time.Second * 2)
		}
	}()
	go func() {
		for {
			c2 <- "From 2"
			time.Sleep(time.Second * 3)
		}
	}()

	go func() {
		for {
			select {
			case msg1 := <- c1 : fmt.Println(msg1)
			case msg2 := <- c2 : fmt.Println(msg2)
			}
		}
	}()

	go func() {
		for {
			select {
			case msg_1 := <- c1 : fmt.Println("Message 1", msg_1)
			case msg_2 := <- c2 : fmt.Println("Message 2", msg_2)
			case <- time.After(time.Second) : fmt.Println("Timeout",)
			// default: fmt.Println("Nothing ready")
			} // The default case happens immediately if none of the channels are ready.
		}
	}()

	var _input string
	fmt.Scanln(&_input)

	/*
	* time.After creates a channel and after the given dura-
    * tion will send the current time on it.
	*/

	//channels
	var ch chan string = make(chan string)

	go pinger(ch)
	go ponger(ch)
	go printer(ch)


	var _input_ string
	fmt.Scanln(&_input_)



	// Handle concurrency - goroutines
	go goroutine(0)
	var input_ string
	fmt.Scanln(&input_);

	// Let's make 10 goroutines in our program
	for i := 0; i < 10; i++ {
		go goroutine(i)
	}
	var input__ string
	fmt.Scanln(&input__, "\n")

    c_ := Circle{0, 0, 5}
    r_ := Rectangle{0, 0, 10, 20}
	// implement interface logic to find area as we did in previous approaches
	fmt.Println("Used interface to find total area of circle and rectangle =>",totalArea(&c_, &r_))

	// acess embedded types
	_a := new(Android)
	_a.Person.Talk();

	//But we can also call any Person methods directly on the Android:
	__a := new(Android)
	__a.Talk();

	// Used method to calculate area of a rectangle
	r := Rectangle{0, 0, 10, 10}
	fmt.Println("Used method to find area of a rectangle =>",r.area())

	c := Circle{0, 0, 5}
	// used method to find the area of a circle
	fmt.Println("Used method to find area of a circle =>",c.area())

	// used pointer to find area of circle - pointer implementation
	
	fmt.Println("Used best way - pointer got find area of a circle =>",__circleArea(&c))

	_c := Circle {0, 0, 5}
	fmt.Println("Used structs to find area of a circle - Unefficient way =>",_circleArea(_c))

	// distance, rectangleArea && circleArea
	var rx1, ry1 float64 = 0, 0
	var rx2, ry2 float64 = 10, 10
	var cx, cy, cr float64 = 0, 0, 5

	fmt.Println("Rectangle area - tidious way => ",rectangleArea(rx1, ry1, rx2,ry2))
	fmt.Println("Circle area - tidious way => ",circleArea(cx, cy, cr))

	/*
	* The above approach to solve the problem is tedious  - let's implement the use of structs and interfaces to make things easy
	*
	* Structs
	* See above - implementation logic
	*/
	
// Fields
   /*We can access fields with the `.` operator*/
   fmt.Println(c.x, c.y, c.r);
   c.x = 10
   c.y = 5
   c.r = 1


   


	/*
	x__ := 1.5
	fmt.Println(square(&x__))
	*/

	// new
	xptr := new(int)
	one(xptr)
	fmt.Println(*xptr);

	// pointers
	_xptr := 5
	zero(&_xptr);
	fmt.Println(_xptr) // 0


	// defer
	// let's open and close a file with defer
	// f, _ := os.Open(filename)
	// defer f.Close()

	// Panic && recover 
	/* Panic() instantly causes runtime error 
	* Recover() is used to handle the runtime error caused by panic() 
	*/

	/*
	panic("PANIC")
	str := recover();
	fmt.Println(str)

	// to get call recovered, we have to pair it with defer
	defer func() {
		str := recover();
		fmt.Println(str)
	} ()
	panic("PANIC")

	*/

	defer second(); // waits for first
	first(); // executes first

	// it does like this - difference comes to durations
	first();
	second();

	// generate even numbers
	nextEven := makeEvenGenerator();
	fmt.Println("EvenGenerator -: ")
	fmt.Println(nextEven()); // 0
	fmt.Println(nextEven()); // 2
	fmt.Println(nextEven()); // 4

	n := 0
	increment := func () int {
		n++
		return n;
	}
	fmt.Println("\nIncrement count -:")
	fmt.Println(increment()); // 1
	fmt.Println(increment()); // 2
	fmt.Println(increment()); // 3

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
		cc = 3
		d = 4
	)
	fmt.Println(a,b,cc,d);

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