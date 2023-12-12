package main
import (
	"fmt"
	"os"
	"bufio"
)

func main () {
	fmt.Print("Enter your name: ");

	// Read user input from the console
	scanner := bufio.NewScanner(os.Stdin);
	scanner.Scan();
	name := scanner.Text();

	// Print a personalized greeting
	fmt.Printf("Hello, %s! Welcome to the world Of Go Programming.\n", name);

}