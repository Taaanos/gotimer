package gotimer

import (
	"log"
	"time"
)

// TimeTrack measures elapsed time of the invoker function - use with defer
func TimeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("function %s took %s", name, elapsed)
}

// TODO option to write to file

