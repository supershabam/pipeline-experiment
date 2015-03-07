# pipeline
golang pipeline patterns

### ideas

`//go:generate pipeliner batch(maxSize, maxDuration, string)`
```
func batchString(maxSize int, maxDuration time.Duration, in <-chan string) <-chan []string {
  out := make(chan string)
  go func() {
    defer close(out)
  Start:
    first, active := <-in // only start the timer once we have a first item
    if !active {
      return
    }
    batch := []string{first}
    timeout := time.After(maxDuration)
    for {
      if len(batch) > maxSize {
        out <- batch
        continue Start
      }
      select {
      case <-timeout:
        out <- batch
        continue Start
      case item, active := <-in:
        if !active {
          out <- batch
          return
        }
        batch = append(batch, item)
      }
    }
  }()
  return out
}
```
