package main

import (
	"fmt"
)

/* 
	
	Topic: RECIEVERS
	
	GO: Things in parenthesis before function name
	https://www.linkedin.com/pulse/go-things-parenthesis-before-function-name-shivam-chaurasia/

*/

type User struct {
	Name string
	Pet  []string
}

// newPet is a method of type *User, can be invoked by any object of type *User, in this case p2
// p2.newPet(), it's called receiver
func (p2 *User) newPet() {
	fmt.Println(*p2, "underlying value of p2 before")
	p2.Pet = append(p2.Pet, "Lucy")
	fmt.Println(*p2, "underlying value of p2 after")
}
func main() {
	u := User{Name: "Anna", Pet: []string{"Bailey"}} // this time we'll generate a pointer!
	fmt.Println(u, "u before")
	p := &u // Let's make a pointer to u!
	p.newPet()
	fmt.Println(u, "u after")
	// Does Anna have a new pet now? Yes!
}


/*
	Topic: WHY & IS NEED FOR SCAN TO MODIFY VALUE OF A VARIABLE
	
	It needs to change the variable. Since all arguments in C are passed by value you need to
	pass a pointer if you want a function to be able to change a parameter.

	Here's a super-simple example showing it in C++:


	void nochange(int var) {
	    // Here, var is a copy of the original number. &var != &value
	    var = 1337;
	}
	void change(int *var) {
	    // Here, var is a pointer to the original number. var == &value
	    // Writing to `*var` modifies the variable the pointer points to
	    *var = 1337;
	}
	int main() {
	    int value = 42;
	    nochange(value);
	    change(&value);
	    return 0;
	}


*/