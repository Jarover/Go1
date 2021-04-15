package fibonachy

var Res map[int64]int64


// Рекурсивный алгоритм 
func Fibre(n int64) int64 {
	if n < 2 { return 1 }
	return Fibre(n-2) + Fibre(n-1)
}

// Сохраняем результат в map

func Fibmap(n int64) int64 {
	if n < 2 { return 1 }
	_, ok := Res[n]
	if !ok {
		Res[n] = Fibmap(n-2) + Fibmap(n-1)
	}
	return Res[n]
}


func init() {
	Res = make(map[int64]int64)
}