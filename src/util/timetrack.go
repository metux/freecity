package util

import "time"
import "log"

func TimeTrack(start time.Time, name string) {
    elapsed := time.Since(start)
    log.Printf("%s took %s\n", name, elapsed)
}
