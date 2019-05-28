package multiple_pointers

import "testing"

var atdTestData = [][]rune{
	{'1', '2', '3'},
	{'1', '2', '2'},
	{'a', 'b', 'c', 'a'},
}

var atdTestResult = []bool{false, true, true}

func TestAreThereDuplicate(t *testing.T) {
	for i, v := range atdTestData {
		res := AreThereDuplicate(v...)
		if atdTestResult[i] != res {
			t.Errorf("%v : expect %v, got %v", v, atdTestResult[i], res)
		}
	}
}

func BenchmarkAreThereDuplicate(b *testing.B) {
	for i := 0; i < b.N; i++ {
		AreThereDuplicate(atdTestData[2]...)
	}
}
