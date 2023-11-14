package climbing_stairs_70

/**
Let's check 4 steps.

1. 1+1+1+1

2. 2+1+1
3. 1+2+1
4. 1+1+2

5. 2+2

Let's check 5 steps.

1. 1+1+1+1+1

2. 2+1+1+1
3. 1+2+1+1
4. 1+1+2+1
5. 1+1+1+2

6. 2+2+1
7. 2+1+2
8. 1+2+2

Now, let's check 7 steps.

1. 1+1+1+1+1+1+1

1. 2+1+1+1+1+1
2. 1+2+1+1+1+1
3. 1+1+2+1+1+1
4. 1+1+1+2+1+1
5. 1+1+1+1+2+1
6. 1+1+1+1+1+2

1.  2+2+1+1+1
2.  2+1+2+1+1
3.  2+1+1+2+1
4.  2+1+1+1+2
5.  1+2+2+1+1
6.  1+2+1+2+1
7.  1+2+1+1+2
8.  1+1+2+2+1
9.  1+1+2+1+2
10. 1+1+1+2+2

13. 2+2+2+1
13. 2+2+1+2
13. 2+1+2+2
13. 1+2+2+2

*/

func climbStairs(steps int) int {
	memo := map[int]int{}

	return climbStairsRec(steps, memo)
}

func climbStairsRec(steps int, memo map[int]int) int {
	if steps <= 1 {
		return 1
	}

	if _, ok := memo[steps]; !ok {
		memo[steps] = climbStairsRec(steps-1, memo) + climbStairsRec(steps-2, memo)
	}

	return memo[steps]
}
