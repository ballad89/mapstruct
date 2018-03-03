package mapstruct

import (
	"errors"
	"reflect"
	"testing"
)

type User struct {
	Email      string
	FirstName  string
	ID         string
	LastName   string
	ProfilePic string
	UserName   string
	Verified   bool
}

type Artist struct {
	ID       string
	Name     string
	Verified bool
}

func TestMapInterfaceToStruct(t *testing.T) {

	testData := []struct {
		in  map[string]interface{}
		in2 interface{}
		out error
	}{
		{map[string]interface{}{"email": "test@gmail.com", "first_name": "husayn", "last_name": "arrah"}, &User{}, nil},
		{map[string]interface{}{"first_name": 3}, &User{}, errors.New("value types did not match for field FirstName, expected string, got int")},

		{map[string]interface{}{"name": "husayn", "user_id": "123-456"}, &Artist{}, nil},
	}

	for _, v := range testData {

		err := MapInterfaceToStruct(v.in, v.in2)

		if err != v.out && (v.out == nil || err == nil || err.Error() != v.out.Error()) {
			t.Error(err)
		}
	}

}

func TestStructToMapInterface(t *testing.T) {

	testData := []struct {
		in  *User
		out map[string]interface{}
	}{
		{&User{
			Email:     "test@gmail.com",
			FirstName: "husayn",
			LastName:  "arrah",
			Verified:  true,
			ID:        "123-456",
		}, map[string]interface{}{"email": "test@gmail.com", "first_name": "husayn", "last_name": "arrah"}},
		{&User{
			Email:     "test@gmail.com",
			FirstName: "husayn",
			LastName:  "arrah",
			Verified:  true,
			ID:        "123-456",
		}, map[string]interface{}{"email": "test@gmail.com", "first_name": "husayn",
			"last_name": "arrah", "id": "123-456"}},
	}

	for _, v := range testData {

		ret, err := StructToMapInterface(v.in)

		if err != nil {
			t.Error(err)
		}

		for key, value := range v.out {
			if value != ret[key] {
				t.Errorf("expected %v for key %v but got %v", v.out[key], key, value)
			}
		}

	}

}

func TestMergeMaps(t *testing.T) {
	testData := []struct {
		left  map[string]interface{}
		right map[string]interface{}
		out   map[string]interface{}
	}{
		{
			map[string]interface{}{"1": 1, "2": 2},
			map[string]interface{}{"3": 3, "4": 4},
			map[string]interface{}{"1": 1, "2": 2, "3": 3, "4": 4},
		},
		{
			map[string]interface{}{"1": 1, "2": 2},
			map[string]interface{}{"3": 3, "4": 4, "5": 5, "6": 6},
			map[string]interface{}{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6},
		},
		{
			map[string]interface{}{"3": 3, "4": 4, "5": 5, "6": 6},
			map[string]interface{}{"1": 1, "2": 2},
			map[string]interface{}{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6},
		},
		{
			map[string]interface{}{"3": 3, "4": 4, "5": 5, "6": 6},
			map[string]interface{}{"3": "override", "2": 2},
			map[string]interface{}{"2": 2, "3": "override", "4": 4, "5": 5, "6": 6},
		},
	}

	for _, v := range testData {
		ret := MergeMaps(v.left, v.right)

		if eq := reflect.DeepEqual(ret, v.out); !eq {
			t.Errorf("expected %v but got %v", v.out, ret)
		}

	}
}
