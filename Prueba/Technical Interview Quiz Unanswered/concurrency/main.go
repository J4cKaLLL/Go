package main

import (
	"log"
)

const target = 100

func main() {
	dataMap := map[int]int{}

	for i := 0; i < 10; i++ {
		dataMap[i] = 0

		incrementFunc := func(index int, incrementAmount int) {
			for {
				if dataMap[index] == target {
					return
				}

				dataMap[index] = dataMap[index] + incrementAmount
			}
		}

		func() {
			go incrementFunc(i, 1)
			go incrementFunc(i, -1)
			go incrementFunc(i, 1)
		}()
	}


	/// Each val in the map must equal the target
	for key, val := range dataMap {
		if val != target {
			log.Fatalf("val %d at key %i does not equal target %d", val, key, target)
		}
	}
}