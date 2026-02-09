package triangle

type Kind int

const (
	_	= iota	// 0
    NaT  		// not a triangle - 1
	Equ 		// equilateral - 2
	Iso 		// isosceles - 3
	Sca 		// scalene - 4
)

func KindFromSides(a, b, c float64) Kind {
	var k Kind

    if a == 0 && b == 0 && c == 0 {
        k = Kind(NaT)
    } else if !(a + b >= c && b + c >= a && a + c >= b) {
        k = Kind(NaT)
    } else if a == b && b == c {
        k = Kind(Equ)
    } else if (a == b && b != c) || (a == c && c != b) || (b == c && c != a) {
        k = Kind(Iso)
    } else {
        k = Kind(Sca)
    }

	return k
}
