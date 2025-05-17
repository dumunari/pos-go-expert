package tax

import (
	"testing"
)

// go test -bench=. -benchmem
// go test -bench=. -benchmem -count=10
// go test -bench=. -benchmem -count=10 -benchtime=3s
// go test -bench=. -benchmem -run=^$
// go test -bench=. -benchmem -run=^$ -cpuprofile=cpu.out
// go test -bench=. -benchmem -run=^$ -cpuprofile=cpu.out -memprofile=mem.out
// # Abrir prompt interativo
// go tool pprof cpu.out|mem.out
// # Abrir interface web
// go tool pprof -http=:8080 cpu.out|mem.out
// Para aplicações web, é possível expor endpoints HTTP (via `net/http/pprof`)
// para coletar perfis em tempo real, inclusive para heap, CPU, goroutines, etc

func BenchmarkCalculateTax(b *testing.B) {
	income := 1000.0
	for i := 0; i < b.N; i++ {
		CalculateTax(income)
	}
}

func BenchmarkCalculateTaxBadly(b *testing.B) {
	income := 1000.0
	for i := 0; i < b.N; i++ {
		CalculateTaxBadly(income)
	}
}
