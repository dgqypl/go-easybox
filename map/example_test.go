package _map

type User struct {
	Id   int64
	Name string
}

func ExampleSliceToMap() {
	var slice []*User
	u1 := &User{Id: 1, Name: "helen"}
	u2 := &User{Id: 2, Name: "tim"}

	slice = append(slice, u1, u2)

	SliceToMap(slice, func(u *User) int64 {
		return u.Id
	})
}