func countStudents(students []int, sandwiches []int) int {
    counts := []int{0, 0}
    for _, s := range students {
        counts[s]++
    }

    for _, sandwich := range sandwiches {
        // If no student left wants this sandwich, 
        // the rest of the sandwiches can't be processed.
        if counts[sandwich] == 0 {
            break
        }
        counts[sandwich]--
    }

    // The remaining students are the sum of the counts
    return counts[0] + counts[1]
}