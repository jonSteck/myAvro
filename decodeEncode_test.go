package myAvro

import "testing"

func TestEncodeDecode(t *testing.T) {
	test := []string{"myFirstTest", "myFirstTest2", ""} 
	for i := 0; i < 2; i++ {
		got := MyDecode(MyEncode(test[i]))
		if test[i] != got {
			t.Errorf("%s == %s, want == got", test[i], got)
		}
	}
}