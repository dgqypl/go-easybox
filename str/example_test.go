package str

import "fmt"

type User struct {
	Id   int64
	Name string
}

func ExampleFieldValueJoin() {
	var slice []*User
	u1 := &User{Id: 1, Name: "helen"}
	u2 := &User{Id: 2, Name: "tim"}

	slice = append(slice, u1, u2)

	result, _ := FieldValueJoin(slice, ",", func(u *User) string {
		return u.Name
	})

	fmt.Println(result) // "helen,tim"
}
