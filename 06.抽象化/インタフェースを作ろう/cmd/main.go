package main

import "fmt"

type stringer interface {
	String() string
}

// Int is an int
type Int int

func (i Int) String() string {
	return fmt.Sprintf("%d", i)
}

// Func is a func
type Func func(s string)

func (f Func) String() string {
	return fmt.Sprintf("%#v", f)
}

// User is a user
type User struct {
	name string
	age  int32
}

func (u User) String() string {
	return fmt.Sprintf("my name is %s(%d y.o.)", u.name, u.age)
}
func main() {
	var sint stringer = Int(64)
	var sfunc stringer = Func(func(s string) {
		println(s)
	})
	var suser stringer = User{"tom", 20}

	print(sint)
	print(sfunc)
	print(suser)
}

func print(s stringer) {
	switch s.(type) {
	case Int:
		fmt.Printf("Int# %T: %s\n", s, s.String())
	case Func:
		fmt.Printf("Func# %T: %s\n", s, s.String())
	case User:
		fmt.Printf("User# %T: %s\n", s, s.String())
	default:
		fmt.Println("default")
	}
}
