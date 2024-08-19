### README.md

# Easy-Cache

Easy-Cache is a flexible and efficient caching library written in Go, designed to handle common data structures such as maps, slices, and structs with advanced features like TTL (Time-To-Live), eviction policies, and event-driven invalidation. The library is optimized for performance and supports native caching without the need for JSON serialization, making it ideal for high-performance applications.

## Features

- **Native Caching**: Supports caching of Go's native data structures (`map`, `slice`, `struct`) without the need for JSON serialization.
- **TTL (Time-To-Live)**: Automatically expire cache entries after a specified duration.
- **Eviction Policies**: Supports Least Recently Used (LRU) and Least Frequently Used (LFU) eviction policies.
- **Event-Driven Invalidation**: Invalidate cache entries based on custom events.
- **Sharded Cache**: Scale cache horizontally by sharding data across multiple caches.

## Installation

To install Easy-Cache, use the following `go get` command:

```sh
go get github.com/onurbaran/easy-cache
```

## Getting Started

### Basic Usage

Here is a basic example of how to use Easy-Cache to cache a simple key-value pair:

```go
package main

import (
    "context"
    "fmt"
    "github.com/onurbaran/easy-cache/pkg/cache"
)

func main() {
    // Create a new cache with default configuration
    config := cache.DefaultConfig()
    c := cache.NewCache(config)

    // Set a value in the cache
    ctx := context.Background()
    c.Set(ctx, "key1", "value1", 1, "default")

    // Get the value from the cache
    data, found, _ := c.Get(ctx, "key1")
    if found {
        fmt.Println("Found:", data)
    } else {
        fmt.Println("Not found")
    }

    // Delete the value from the cache
    c.Delete(ctx, "key1")

    // Try to get the value again after deletion
    data, found, _ = c.Get(ctx, "key1")
    if found {
        fmt.Println("Found:", data)
    } else {
        fmt.Println("Not found after deletion")
    }
}
```

### Caching Maps

Easy-Cache supports caching `map[string]interface{}` natively:

```go
package main

import (
    "context"
    "fmt"
    "github.com/onurbaran/easy-cache/pkg/cache"
)

func main() {
    // Create a MapCache
    config := cache.DefaultConfig()
    mapCache := cache.NewMapCache(config)

    // Cache a map
    ctx := context.Background()
    exampleMap := map[string]interface{}{"key1": "value1", "key2": "value2"}
    mapCache.Set(ctx, "example", exampleMap)

    // Retrieve the map from the cache
    data, found, err := mapCache.Get(ctx, "example")
    if err != nil {
        fmt.Println("Error getting map from cache:", err)
        return
    }
    if found {
        fmt.Printf("Found map in cache: %+v\n", data)
    } else {
        fmt.Println("Map not found in cache")
    }
}
```

### Caching Slices

You can also cache `[]interface{}` slices natively:

```go
package main

import (
    "context"
    "fmt"
    "github.com/onurbaran/easy-cache/pkg/cache"
)

func main() {
    // Create a SliceCache
    config := cache.DefaultConfig()
    sliceCache := cache.NewSliceCache(config)

    // Cache a slice
    ctx := context.Background()
    exampleSlice := []interface{}{"value1", "value2", "value3"}
    sliceCache.Set(ctx, "exampleSlice", exampleSlice)

    // Retrieve the slice from the cache
    data, found, err := sliceCache.Get(ctx, "exampleSlice")
    if err != nil {
        fmt.Println("Error getting slice from cache:", err)
        return
    }
    if found {
        fmt.Printf("Found slice in cache: %+v\n", data)
    } else {
        fmt.Println("Slice not found in cache")
    }
}
```

### Caching Structs with TTL

Structs can be cached with a TTL (Time-To-Live) to automatically expire after a certain period:

```go
package main

import (
    "context"
    "fmt"
    "time"
    "github.com/onurbaran/easy-cache/pkg/cache"
)

type User struct {
    Name  string
    Email string
    Age   int
}

func main() {
    // Create a StructCache with a 5-second TTL
    config := cache.DefaultConfig()
    config.BaseTTL = 5 * time.Second
    structCache := cache.NewStructCache(config)

    // Cache a struct
    ctx := context.Background()
    user := User{Name: "Alice", Email: "alice@example.com", Age: 30}
    err := structCache.Set(ctx, "user1", user)
    if err != nil {
        fmt.Println("Error setting struct in cache:", err)
        return
    }

    // Retrieve the struct from the cache
    data, found, _ := structCache.Get(ctx, "user1")
    if found {
        fmt.Printf("Initially found struct in cache: %+v\n", data.(User))
    } else {
        fmt.Println("Struct not found in cache initially")
    }

    // Wait for the TTL to expire
    time.Sleep(6 * time.Second)

    // Check the cache after TTL expiration
    data, found, _ = structCache.Get(ctx, "user1")
    if !found {
        fmt.Println("Struct has expired and is no longer in cache")
    } else {
        fmt.Printf("Struct still found in cache after TTL: %+v\n", data.(User))
    }
}
```

### Event-Driven Invalidation

Easy-Cache supports event-driven cache invalidation, allowing you to remove or update cache entries based on custom events.

```go
package main

import (
    "context"
    "fmt"
    "github.com/onurbaran/easy-cache/pkg/cache"
    "github.com/onurbaran/easy-cache/pkg/event"
)

type User struct {
    Name  string
    Email string
    Age   int
}

type InvalidateEventListener struct {
    cache *cache.StructCache
}

func (l *InvalidateEventListener) OnEvent(e event.Event) {
    ctx := context.Background()
    key := e.Data.(string)
    l.cache.Delete(ctx, key)
    fmt.Println("Invalidated struct with key:", key)
}

func main() {
    // Create an Event Manager and attach it to the cache
    eventManager := event.NewEventManager()
    config := cache.DefaultConfig()
    config.EventManager = eventManager
    structCache := cache.NewStructCache(config)

    // Cache a struct
    ctx := context.Background()
    user := User{Name: "Alice", Email: "alice@example.com", Age: 30}
    structCache.Set(ctx, "user1", user)

    // Register an event listener
    listener := &InvalidateEventListener{cache: structCache}
    eventManager.RegisterListener("invalidateUser", listener)

    // Trigger the invalidation event
    eventManager.TriggerEvent(event.Event{Name: "invalidateUser", Data: "user1"})

    // Check the cache after event
    data, found, _ := structCache.Get(ctx, "user1")
    if !found {
        fmt.Println("Struct successfully invalidated and deleted from cache")
    } else {
        fmt.Println("Struct still found in cache:", data)
    }
}
```

## Advanced Features

### Eviction Policies

Easy-Cache supports LRU (Least Recently Used) and LFU (Least Frequently Used) eviction policies to manage cache size.

#### LRU Eviction

```go
config := cache.DefaultConfig()
config.EvictionPolicy = &cache.LRUEviction{}
```

#### LFU Eviction

```go
config := cache.DefaultConfig()
config.EvictionPolicy = &cache.LFUEviction{}
```

### Sharded Cache

Easy-Cache allows you to shard your cache across multiple instances for improved scalability.

```go
config := cache.DefaultConfig()
config.NumShards = 8
shardedCache := cache.NewShardedCache(config)
```

## TODO

- Distributed cache
- Gossip Protocol Implementation
- Avro Serialization/Deserialization
- Adaptive TTL Management


## Conclusion

Easy-Cache is a flexible caching library designed to handle various data structures efficiently in Go. With features like native caching, TTL, eviction policies, and event-driven invalidation, it provides a  solution for caching in high-performance applications.

Feel free to explore the examples provided and integrate Easy-Cache into your Go projects. For more information, check out the source code and additional documentation.

## License

This project is licensed under the MIT License.

---

