package sum_of_two_integers_371

/**
* 0000_0010 // 2
* 0000_0011 // 3
* 0000_0101 // 5
*
* 1111_1110 // -2
* 0000_0011 // 3
* 0000_0101 // 5
 */

var signBit = 1 ^ -1

func getSum(a, b int) int {
	ors := a | b
	sumSignBit := signBit & ors

	ands := a & b << 1
	xors := (a ^ b) & sumSignBit

	return ands | xors
}
