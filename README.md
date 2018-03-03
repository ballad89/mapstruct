MapStruct
==========

A library with utility functions for converting structs to `map[string]interface{}` and vice versa, relies on `reflection`.

`MapInterfaceToStruct`: Converts an unstructured map to a concrete type.

`StructToMapInterface`: Converts a concrete type to unstructured map.

`MergeMaps`: Merges two maps.

```Go
//main.go
package main

import (
	"fmt"

	"github.com/ballad89/mapstruct"
)

type user struct {
	Email      string
	FirstName  string
	ID         string
	LastName   string
	ProfilePic string
	UserName   string
	Verified   bool
}

func main() {
	u := &user{}
	in := map[string]interface{}{"email": "test@gmail.com", "first_name": "husayn", "last_name": "arrah"}
	mapstruct.MapInterfaceToStruct(in, u)

	fmt.Printf("%+v\n", u)

	out, err := mapstruct.StructToMapInterface(u)

	if err != nil {
		panic(err)
	}

	fmt.Println(out)

	left := map[string]interface{}{"3": 3, "4": 4, "5": 5, "6": 6}
	right := map[string]interface{}{"3": "override", "2": 2}

	ret := mapstruct.MergeMaps(left, right)
	fmt.Println(ret)
}

```
go run main.go

&{Email:test@gmail.com FirstName:husayn ID: LastName:arrah ProfilePic: UserName: Verified:false}
map[profile_pic: user_name: verified:false email:test@gmail.com first_name:husayn id: last_name:arrah]
map[5:5 6:6 2:2 3:override 4:4]
```
