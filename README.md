# Goche

A thread safe caching module for golang with customizable cache eviction strategies

## Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/GrayHat12/goche"
	"github.com/GrayHat12/goche/strategy"
)

func AddWithDelay(a int, b int) int {
	fmt.Printf("Sleeping for %d seconds\n", a+b)
	time.Sleep(time.Duration(a+b) * time.Second)
	return a + b
}

func CachedAddWithDelay(strategyGenerator goche.NewStrategyGenerator[int]) func(int, int) int {

	var cache = goche.NewCache(10, strategyGenerator) // will cache 10 elements at max, and use the provided strategy

	return func(a int, b int) int {

		// Hashing logic
		hash := fmt.Sprintf("%d+%d", a, b)

		// try getting result from cache
		response, err := cache.Get(hash)

		// if response in cache, return cached value
		if err == nil {
			return response
		}

		// else, fetch the live response
		response = AddWithDelay(a, b)

		// store response in cache
		cache.Set(hash, response)
		return response
	}
}

func main() {
    // specifying my strategy here to FIFO, I can implement any other strategy and use that as well.
	cachedAdd := CachedAddWithDelay(strategy.NewFifoStrategy)

	tenplus2 := cachedAdd(10, 2)
	oneplus1 := cachedAdd(1, 1)
	tenplus2_2 := cachedAdd(10, 2)
	oneplus1_2 := cachedAdd(1, 1)

	fmt.Printf("tenplus2 = %d\n", tenplus2)
	fmt.Printf("oneplus1 = %d\n", oneplus1)
	fmt.Printf("tenplus2_2 = %d\n", tenplus2_2)
	fmt.Printf("oneplus1_2 = %d\n", oneplus1_2)
}
```

**Output**
```sh
Sleeping for 12 seconds
Sleeping for 2 seconds
tenplus2 = 12
oneplus1 = 2
tenplus2_2 = 12
oneplus1_2 = 2
```