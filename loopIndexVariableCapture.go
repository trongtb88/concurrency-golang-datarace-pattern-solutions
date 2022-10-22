package main

func processJob(job int) {
	// do something
}

//Data race due to loop index variable capture
//ORIGINAL
/* func main() {
	jobs := [2]int{1, 2}
	for job := range jobs {
		go func() {
			processJob(job)
		}()
	}
}
*/

// FIX
func main() {
	jobs := [2]int{1, 2}
	for index, _ := range jobs {
		go func(index int) {
			processJob(jobs[index])
		}(index)
	}
}
