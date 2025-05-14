package main

type MyNUmber int

// ~ habilita usar o MyNumber (int) como o tipo generic Number (int|float)
type Number interface {
	~int | ~float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

// Comparable só faz check de igualdade (só tipos compatíveis)
func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Wesley": 1000, "João": 2000, "Maria": 3000}
	m2 := map[string]float64{"Wesley": 1000.2, "João": 2000.3, "Maria": 3000.1}
	m3 := map[string]MyNUmber{"Wesley": 1000, "João": 2000, "Maria": 3000}

	println(Soma(m))
	println(Soma(m2))
	println(Soma(m3))
	println(Compara(10, 10))
}
