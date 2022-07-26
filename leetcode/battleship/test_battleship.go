package battlehsip

import "testing"

func testGetHitProbability(t *testing.T)  {
	got := getHitProbability(2,3,[][]int32{{0,0,1},{1,0,1}})
	want := 0.5
	if got != want {
		t.Errorf("Got %v want %v", got, want)
	}
}
