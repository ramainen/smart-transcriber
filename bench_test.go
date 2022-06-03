package transcriber

//go test -bench=. bench_test.go.
//BEFORE: 2 mins  08:08:11 - 08:10:11
/*
BEFORE:

BenchmarkTranscriberSpeed-4         1131           1002516 ns/op

AFTER:
BenchmarkTranscriberSpeed-4         3400            303452 ns/op

*/
import (
	"testing"
)

func BenchmarkTranscriberSpeed(b *testing.B) {

	transcriber := NewTranscriber()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		transcriber.Transcribe("bosch")
	}

}
