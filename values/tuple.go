package values

type Tuple2[T1, T2 any] struct {
	Val1 T1
	Val2 T2
}

func Make2[T1, T2 any](v1 T1, v2 T2) Tuple2[T1, T2] {
	return Tuple2[T1, T2]{
		Val1: v1,
		Val2: v2,
	}
}
