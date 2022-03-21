package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"myapp/doctor"
	"myapp/employee"
	"myapp/packageone"
	"os"
	"sort"
	"strings"
	"time"
)

func main() {
	fmt.Print("Hello World!")
	whatTosay := doctor.Intro()
	fmt.Print(whatTosay)
	reader := bufio.NewReader(os.Stdin)
	// go has only one kind of loop
	for {
		fmt.Print("-> ")
		userInput, _ := reader.ReadString('\n')                // waits until enter is pressed
		userInput = strings.Replace(userInput, "\r\n", "", -1) // -1 replace everywhere you find it
		userInput = strings.Replace(userInput, "\n", "", -1)
		if userInput == "quit" {
			break
		} else {
			fmt.Println(doctor.Response(userInput))
		}
	}
	//play game
	game()

	// Array declaration
	var myString [3]string
	myString[0] = "some value"
	myString[1] = "second value"
	myString[2] = "third value"

	// Structs
	type Car struct {
		NumOfTyres int
		Model      string
		Year       int
	}

	var myCar Car
	myCar.NumOfTyres = 4
	myCar.Model = "ALTO"
	myCar.Year = 1998

	newCar := Car{
		NumOfTyres: 4,
		Model:      "BMW",
		Year:       1998,
	}

	fmt.Printf("My car is %s and it is made in %d ", newCar.Model, newCar.Year)
	learningPointers()
	learnSlicing()

	// learn reciever function calling
	var dog Animal
	dog.Name = "harry"
	dog.Sound = "Whoo!!"
	dog.NumOfLegs = 4
	dog.Says() // function attached to types

	var cat Animal
	cat.Name = "Jane"
	cat.Sound = "Meow!"
	cat.NumOfLegs = 4
	cat.Says()

	// Jst like creating objects for a class in OOP's

	//learn Maps
	learnMaps()

	// Function to pass multiple number of arguements
	sumMany(1, 2, 3, 4, 5, 6, 7) // can take n num of args.

	// learn about interfaces
	/*
		1. List all the functions that interfaces must have
		2. Implement the same method for each object type
		3. Use a common interface to access the methods

	*/
	var myHome Home
	myHome.usage = "Personal"
	myHome.rooms = 4

	var office Commercial
	office.usage = "Work"
	office.rooms = 2

	fmt.Println((&myHome))
	fmt.Println((&office))
	//PrintProperty(&office)

	// call to employee struct
	employee.EmployeeLogic()

}

const prompt = "and press ENTER"

var welcome = "welcome to game!" // package level variable

func game() {
	var welcome = "welcome to game in block !"
	fmt.Println(welcome) // prints block level var

	/*
		general variable declaration
		var fv = 2
		var sn = 3
		var substraction = 4
	*/

	// seed the random number generator
	rand.Seed(time.Now().UnixNano())

	// specifying 8 as max range
	var fv = rand.Intn(8) + 2
	var sn = rand.Intn(8) + 2
	var substraction = rand.Intn(8) + 2
	var answer = fv*sn - substraction
	playThegame(fv, sn, substraction, answer)
	printPackageVar()

}

func playThegame(fv, sn, substraction, answer int) {
	fmt.Print(welcome) // prints package level var

	reader := bufio.NewReader(os.Stdin)
	//display something
	fmt.Println("Think of a number between 1 and 10 :", prompt)
	reader.ReadString('\n')

	fmt.Println("Mutilply the number by", fv)
	reader.ReadString('\n')

	fmt.Println("Now multiply", sn, prompt)
	reader.ReadString('\n')

	fmt.Println("Divide the result by the number you orignal thought of ", prompt)
	reader.ReadString('\n')

	fmt.Println("Now Substract", substraction, prompt)
	reader.ReadString('\n')

	fmt.Println("The answer is ", answer)

}

func printPackageVar() {
	mystr := packageone.PublicVar
	fmt.Println(mystr)
}

func learningPointers() {
	// Pointers are something which points to a location in memory.
	x := 10
	myPointer := &x
	fmt.Println("x is ", x)
	fmt.Println("myPointer location is ", myPointer)
	// change the value correspoding to the address
	// *myPointer = 15
	changeValueOfUsingPointer(&x)
	fmt.Println("x is now ", x)

}

func changeValueOfUsingPointer(num *int) {
	*num = 24 // assigning address a value
}

func learnSlicing() {
	var animals []string
	animals = append(animals, "dog")
	animals = append(animals, "cat")
	animals = append(animals, "tiger")

	fmt.Println(animals)

	for _, each_animal := range animals {
		fmt.Println(each_animal)
	}

	// with index attached to a variable
	for i, x := range animals {
		fmt.Println(i, x)
	}

	fmt.Println(animals[0:2]) // range of elements
	// len(animals) also works
	// to sort a slice
	fmt.Println(sort.StringsAreSorted(animals)) // is sorted or not
	sort.Strings(animals)                       // sorts the string slice
	fmt.Println(DeleteFromSlice(animals, 2))    // breaks the sort
}

func DeleteFromSlice(a []string, i int) []string {
	a[i] = a[len(a)-1] // copy the last element to index i
	a[len(a)-1] = ""   // mark last element as blank
	a = a[:len(a)-1]   // copy the whole slice except last ele which is marked as blank
	return a
}

func learnMaps() {
	// maps are unsorted
	// no insertion order
	intMap := make(map[string]int)
	intMap["one"] = 1
	intMap["two"] = 2
	intMap["three"] = 3
	intMap["four"] = 4

	for key, values := range intMap {
		fmt.Println(key, values)
	}

	// delete a key,value pair
	delete(intMap, "three")
	// to check something exists in a map
	el, ok := intMap["four"]
	if ok {
		fmt.Println(el, "is in map")
	} else {
		fmt.Println(el, "is not in map")
	}
	// override or change value in map
	intMap["two"] = 4

}

func sumMany(nums ...int) int {
	// adds up all the number of integers provided as input
	total := 0

	for _, x := range nums {
		total = total + x
	}

	return total
}

type Animal struct {
	Name      string
	Sound     string
	NumOfLegs int
}

// reciever
// right before function name
// anytime you call this function it will call the type Animal
// and will print the values associated to it

// LINK A FUNCTION TO A TYPE
func (a *Animal) Says() {
	fmt.Printf("A %s says %s", a.Name, a.Sound)
}

// interfaces

type Property interface {
	PropertyType() string
	NoOfRooms() int
}

type Home struct {
	usage string
	rooms int
}

type Commercial struct {
	usage string
	rooms int
}

func (d *Home) PropertyType() string {
	return d.usage
}

func (d *Home) NoOfRooms() int {
	return d.rooms
}

func (d *Commercial) PropertyType() string {
	return d.usage
}

func (d *Commercial) NoOfRooms() int {
	return d.rooms
}

func PrintProperty(prop Property) string {
	propStr := fmt.Sprintf("The property is for usage %s and has %d rooms .", prop.PropertyType(), prop.NoOfRooms())
	return propStr
}
