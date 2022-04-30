package double

import "testing"

type StubSercher struct {
	phone string
}

func (ss StubSercher) Search(people []*Person, firstName string, lastName string) *Person {
	return &Person{
		FirstName: firstName,
		LastName:  lastName,
		Phone:     ss.phone,
	}
}
func TestFindReturnPerson(t *testing.T) {
	phoneNumber := "1234567890"

	phonebook := &Phonebook{}
	phone, _ := phonebook.Find(StubSercher{phoneNumber}, "Jone", "Jone")

	if phone != phoneNumber {
		t.Errorf("got %v, want %v", phone, phoneNumber)
	}
}
