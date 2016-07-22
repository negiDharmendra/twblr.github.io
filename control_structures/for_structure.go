package control_structures

func collatzChainLength(n int) int {
	i := 0
	for n!=1 {
		if n&1 {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}
		i++
	}
	return  i
}
