# Goche

A thread safe caching module for golang with customizable cache eviction strategies

## Installation

```sh
go get github.com/GrayHat12/goche
```

## Capabilities
```go
package main

...

func main() {
	// Initialising a new cache
	cache := goche.NewCache(<size>, <strategy>)
	
	// adding value to cache
	cache.Set(<key>, <value>)

	// getting a value from cache
	val, err := cache.Get(<key>)

	// force removing a value from cache
	cahce.Remove(<key>)

	// Creating a cache wrapper of a function
	cachedFn := goche.FunctionDecorator(actualFn, <size>, <strategy>, <hashFunction>)
	// using this to call the function
	val = cachedFn.Call(<args>)
	// using this to force remove key from cache
	cachedFn.Cache.Remove(<key>)
}

```

## Example

```go
package main

import (
	"fmt"
	"time"

	"github.com/GrayHat12/goche"
	"github.com/GrayHat12/goche/strategy"
)

type Numbers struct {
	a int
	b int
}

func AddWithDelay(numbers Numbers) int {
	fmt.Printf("Sleeping for %d seconds\n", numbers.a+numbers.b)
	time.Sleep(time.Duration(numbers.a+numbers.b) * time.Second)
	return numbers.a + numbers.b
}


func main() {
    // create a wrapper for AddWithDelay here specifying the max cache size, an eviction strategy and a hashing function
	cachedAdd := goche.FunctionDecorator(AddWithDelay, 10, strategy.NewFifoStrategy, func(numbers Numbers) string {
		return fmt.Sprintf("%d+%d", numbers.a, numbers.b)
	})

	tenplus2 := cachedAdd.Call(Numbers{a: 10, b: 2})
	oneplus1 := cachedAdd.Call(Numbers{a: 1, b: 1})
	tenplus2_2 := cachedAdd.Call(Numbers{a: 10, b: 2})
	oneplus1_2 := cachedAdd.Call(Numbers{a: 1, b: 1})

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


## Implementing a strategy

Create a new struct that implements the following Interface

```go
type StrategyInterface[K cmp.Ordered, T any] interface {
	Set(*Cache[K, T], K, T)
	Get(*Cache[K, T], K) (T, bool)
	Remove(*Cache[K, T], K)
}
```

That's all you need to use your custom strategy. Refer to [FIFO Strategy Implementation](./strategy/fifo.go) for reference.
> NOTE: Locks are managed on the Cache level and strategies don't have to worry about those