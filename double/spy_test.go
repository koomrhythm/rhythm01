package double

import "testing"

type SpySercher struct {
	phone  string
	called bool
}

func (ss *SpySercher) Search(people []*Person, firstName string, lastName string) *Person {
	ss.called = true
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone,
	}
}
func TestFindCheckCalledReturnPerson(t *testing.T) {
	phoneNumber := "1234567890"

	spy := &SpySercher{
		phone: phoneNumber,
	}

	phonebook := &Phonebook{}
	phone, _ := phonebook.Find(spy, "Jone", "Jone")

	if !spy.called {
		t.Errorf("Expected Search to be called")
	}

	if phone != phoneNumber {
		t.Errorf("got %v, want %v", phone, phoneNumber)
	}
}
