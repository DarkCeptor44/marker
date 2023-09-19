# Marker

The `marker` package provides functionality for recording moments and their elapsed times.

## Installation

```bash
go get github.com/DarkCeptor44/marker
```

## How to use

```go
func main(){
    m := marker.New()

    m.Mark("Event 1")
    generateSlice(1000000, 10000000)
    err := m.Done() // properly handle errors if wanted
    if err != nil {
        os.Exit(1)
    }

    m.Mark("") // empty name generates one in the format "Event n+1" where n is the length of the moments slice
    generateSlice(1000000, 50000000)
    _ = m.Done() // dont handle errors if wanted

    m.Mark("Event 3")
    generateSlice(1000000, 100000000)
    m.Done() // dont even think about errors at all (clean looking)

    result, _ := m.String()
    fmt.Printf("Moments:\n%s", result)
}

// generates int slices of varying length and range
func generateSlice(m, n int) []int {
    result := []int{}
    for i := 0; i < n; i++ {
        result = append(result, rand.Intn(m))
    }
    return result
}

/*
Moments:
Event 1: 343.3161ms
Event 2: 1.3645229s
Event 3: 2.575335s
*/
```

## Testing

```bash
$ go test -v
=== RUN   TestMarker
    marker_test.go:30: Moments:
        Event 1: 343.3161ms
        Event 2: 1.3645229s
        Event 3: 2.575335s
--- PASS: TestMarker (4.28s)
PASS
ok      github.com/DarkCeptor44/marker  4.616s
```
