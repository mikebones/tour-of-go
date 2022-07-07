package contains_duplicate

import "testing"

func TestContainsDuplicateFalse(t *testing.T) {
	got := ContainsDuplicate([]int{1,2,3,4,5})
	want := false
	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}
}

func TestContainsDuplicateTrue(t *testing.T) {
	got := ContainsDuplicate([]int{1,2,1,4,5})
	want := true
	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}
}

func TestDuplicate3True(t *testing.T) {
	got := ContainsDuplicate([]int{1,1,1,3,3,4,3,2,4,2})
	want := true
	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}
}
