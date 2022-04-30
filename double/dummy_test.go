package double

import "testing"

type DummySercher struct {
}

func (ds DummySercher) Search(people []*Person, firstName string, lastName string) *Person {
	return &Person{}
}

func TestFindError(t *testing.T) {
	phonebook := &Phonebook{}
	want := ErrMissingArgs
	_, got := phonebook.Find(DummySercher{}, "", "")

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}
}
