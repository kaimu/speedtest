package speedtest

import (
	"fmt"
	"testing"
)

func BenchmarkNetflix(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dl, err := Netflix()
		if err != nil {
			panic(err)
		}
		fmt.Printf("(Fast.com) Download: %.2f Mb/s\n", dl)
	}
}

func BenchmarkOokla(b *testing.B) {
	for n := 0; n < b.N; n++ {
		dl, ul, err := Ookla()
		if err != nil {
			panic(err)
		}
		fmt.Printf("(speedtest.net) Download: %.2f Mb/s, Upload: %.2f Mb/s\n", dl, ul)
	}
}
