func unique(v *[]string, cur string, left int, right int) {
	if left < 0 || right < left {
		return
	}
	if left == 0 && right == 0 {
		*v = append(*v, cur)
	} else {
		unique(v, cur+"(", left-1, right)
		unique(v, cur+")", left, right-1)
	}
}

func generateParenthesis(n int) []string {
	v := make([]string, 0)
	unique(&v, "", n, n)
	return v
}