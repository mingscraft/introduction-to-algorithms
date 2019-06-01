package reverse

import "testing"

// reverse('awesome') // 'emosewa'
// reverse('rithmschool') // 'loohcsmhtir'

var testData = [][]string{
	{"awesome", "emosewa"},
	{"rithmschool", "loohcsmhtir"},
}

func TestReverse(t *testing.T) {
	for _, v := range testData {
		res := Reverse(v[0])

		if res != v[1] {
			t.Errorf("reverse of %s: expect %s, got %s", v[0], v[1], res)
		}
	}
}

func BenchmarkReverse(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Reverse(testData[1][0])
	}
}
