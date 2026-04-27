func countStudents(students []int, sandwiches []int) int {
	var unable int
	for len(sandwiches) > 0 && unable != len(sandwiches)  {
		if students[0] == sandwiches[0] {
			sandwiches = sandwiches[1:]
			unable = 0
		} else {
			unable++
			students = append(students, students[0])
		}
		students = students[1:]
	}
    return unable
}