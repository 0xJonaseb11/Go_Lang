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
}