package sum_of_two_integers_371

func getSum(a, b int) int {
	for a != 0 {
		/** Where do both bytes have overlap? */
		carry := a & b

		/** We can't "add" two bits where both bytes have overlap. We need to use XOR. */
		b = a ^ b

		/** 0010 + 0010 == 0100. We have to move the carry one to the left. */
		a = carry << 1

		/**
		In the next iteration, we'll check again whether there's overlap between the tmpSum (b)
		and the left-shifted carry (a).
		If there is no overlap, we can cancel the loop:
		1111_0000
		0000_1111
		XOR will yield 1111_1111 which is the sum. No carry needed.
		*/
	}

	return b
}
