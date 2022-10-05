package optional

import "testing"

func TestNew(t *testing.T) {
	type User struct {
		Name string
		Age  int
	}

	op1 := New(User{})

	r := op1.OrElse(User{
		Name: "tester",
		Age:  12,
	})

	var user User
	if r == user {
		t.Errorf("error")
	}

	t.Logf("%T %+v", r, r)
}
