func calPoints(operations []string) int {
	var scores []int
	for _, op := range operations {
		score, err := strconv.Atoi(op)
		switch {
		case err == nil:
			scores = append(scores, score)
		case op == "+":
			scores = append(scores, scores[len(scores) -1] + scores[len(scores)-2])
		case op == "D":
			scores = append(scores, 2 * scores[len(scores) - 1]) 
		case op == "C":
			scores = scores[:len(scores)-1]
		}
	}
	
	var total int
	for _, s := range scores {
		total += s
	}
	return total
}
