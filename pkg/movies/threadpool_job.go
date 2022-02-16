package movies

import "gorm.io/gorm"

//concurrent job

// Job is the concurrent job
func Job(db *gorm.DB) {
	inputChan := make(chan uint, 1000)

	for i := 0; i < threads; i++ {
		go func() {
			job(inputChan, db)
		}()
	}

	for i := 1; i < maxRun; i++ {
		inputChan <- uint(i)
	}

	close(inputChan)
}
