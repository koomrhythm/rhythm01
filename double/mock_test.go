package double

import "testing"

type MockSercher struct {
	phone        string
	methodToCall map[string]bool
	ps           []*Person
}

func (ms *MockSercher) Search(people []*Person, firstName string, lastName string) *Person {
	ms.methodToCall["Search"] = true

	if ms.ps == nil {
		return &Person{
			FirstName: firstName,
			LastName:  lastName,
			Phone:     ms.phone,
		}
	}
	if ms.ps != nil {
		return ms.ps[0]
	}
	return nil
}

func (ms *MockSercher) ExpectToCall(method string) {
	if ms.methodToCall == nil {
		ms.methodToCall = make(map[string]bool)
	}
	ms.methodToCall[method] = false
}

func TestFindCallAndReturnPersonUsingMock(t *testing.T) {
	phoneNumber := "1234567890"

	mock := &MockSercher{
		phone: phoneNumber,
		ps:    []*Person{},
	}
	mock.ExpectToCall("Search")

	phonebook := &Phonebook{}
	phone, _ := phonebook.Find(mock, "Jone", "Jone")

	if phone != phoneNumber {
		t.Errorf("got %v, want %v", phone, phoneNumber)
	}

	for k, v := range mock.methodToCall {
		if !v {
			t.Errorf("Expected method '%s' to be called", k)
		}
	}
}
