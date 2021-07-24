func rec(ans *[]string, cur string, left, right int) {

	if left == 0 && right == 0 {
		*ans = append(*ans, cur)
		return
	}

	if right < left {
		return
	}

	if left > 0 {
		a := cur + "("
		rec(ans, a, left-1, right)
	}
	if right > 0 {
		b := cur + ")"
		rec(ans, b, left, right-1)
	}
}

func generateParenthesis(n int) []string {
	ans := make([]string, 0)
	rec(&ans, "", n, n)
	return ans
}