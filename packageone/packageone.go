package packageone

// lower case letter says that package is not available outside the package
// var privateVar = "I am private"

// Captial first letter says that it is available outside package
var PublicVar = "I am public"

// private function
/* func notExported() {
	fmt.Println(privateVar)

}
*/

// public function
func Exported() {

}
