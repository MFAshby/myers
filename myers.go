package myers

// Myer's diff algorithm in golang
// Ported from https://blog.robertelder.org/diff-algorithm/

type OpType int

const (
	OpDelete OpType = iota
	OpInsert
)

type Op struct {
	OpType OpType      // Insert or delete, as above
	OldPos int         // Position in the old list of item to be inserted or deleted
	NewPos int         // Position in the _new_ list of item to be inserted
	Elem   interface{} // Actual value to be inserted or deleted
}

//  Returns a minimal list of differences between 2 lists e and f
//  requring O(min(len(e),len(f))) space and O(min(len(e),len(f)) * D)
//  worst-case execution time where D is the number of differences.
func Diff(e, f []interface{}, equals func(interface{}, interface{}) bool) []Op {
	return diffInternal(e, f, equals, 0, 0)
}

func diffInternal(e, f []interface{}, equals func(interface{}, interface{}) bool, i, j int) []Op {
	N := len(e)
	M := len(f)
	L := N + M
	Z := 2*min(N, M) + 2
	if N > 0 && M > 0 {
		w := N - M
		g := make([]int, Z)
		p := make([]int, Z)

		hMax := ((L/2 + toInt(L%2 != 0)) + 1)
		for h := 0; h < hMax; h++ {
			for r := 0; r < 2; r++ {
				var c, d []int
				var o, m int
				if r == 0 {
					c = g
					d = p
					o = 1
					m = 1
				} else {
					c = p
					d = g
					o = 0
					m = -1
				}
				kMin := -(h - 2*max(0, h-M))
				kMax := h - 2*max(0, h-N) + 1
				for k := kMin; k < kMax; k += 2 {
					var a int
					if k == -h || k != h && c[absmod((k-1), Z)] < c[absmod((k+1), Z)] {
						a = c[absmod((k+1), Z)]
					} else {
						a = c[absmod((k-1), Z)] + 1
					}
					b := a - k
					s, t := a, b

					for a < N && b < M && equals(e[(1-o)*N+m*a+(o-1)], f[(1-o)*M+m*b+(o-1)]) {
						a, b = a+1, b+1
					}
					c[absmod(k, Z)] = a
					z := -(k - w)
					if absmod(L, 2) == o && z >= -(h-o) && z <= h-o && c[absmod(k, Z)]+d[absmod(z, Z)] >= N {
						var D, x, y, u, v int
						if o == 1 {
							D = 2*h - 1
							x = s
							y = t
							u = a
							v = b
						} else {
							D = 2 * h
							x = N - a
							y = M - b
							u = N - s
							v = M - t
						}
						if D > 1 || (x != u && y != v) {
							return append(diffInternal(e[0:x], f[0:y], equals, i, j), diffInternal(e[u:N], f[v:M], equals, i+u, j+v)...)
						} else if M > N {
							return diffInternal(make([]interface{}, 0), f[N:M], equals, i+N, j+N)
						} else if M < N {
							return diffInternal(e[M:N], make([]interface{}, 0), equals, i+M, j+M)
						} else {
							return make([]Op, 0)
						}
					}
				}
			}
		}
	} else if N > 0 {
		res := make([]Op, N)
		for n := 0; n < N; n++ {
			res[n] = Op{OpDelete, i + n, -1, e[n]}
		}
		return res
	} else {
		res := make([]Op, M)
		for n := 0; n < M; n++ {
			res[n] = Op{OpInsert, i, j + n, f[n]}
		}
		return res
	}
	panic("Should never hit this!")
}

func max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func toInt(b bool) int {
	if b {
		return 1
	} else {
		return 0
	}
}

/**
 * The remainder op in python always matches the sign of the _denominator_
 * In golang it matches the sign of the numerator.
 * See https://en.wikipedia.org/wiki/Modulo_operation#Variants_of_the_definition
 */
func abs(x int) int {
	if x < 0 {
		return -x
	} else {
		return x
	}
}
func absmod(x, y int) int {
	return abs(x % y)
}

// Convenient wrapper for string lists
func DiffStr(e, f []string) []Op {
	e1, f1 := make([]interface{}, len(e)), make([]interface{}, len(f))
	for i, ee := range e {
		e1[i] = ee
	}
	for i, fe := range f {
		f1[i] = fe
	}
	return Diff(e1, f1, func(s1, s2 interface{}) bool {
		return s1 == s2
	})
}
