# pipeline
golang pipeline patterns

### ideas

`//go:generate pipeliner batch(maxSize, maxDuration, string)`

```
func batchString(maxSize int, maxDuration time.Duration, in <-chan string) <-chan []string {
  out := make(chan []string)
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

`//go:generate pipeliner batch(maxSize, Event)`

```
func batchEvent(maxSize uint, in <-chan Event) <-chan []Event {
  if maxSize == 0 {
    panic("batch size must be greater than zero")
  }
  out := make(chan []Event)
  go func() {
    defer close(out)
  Start:
    batch := []Event{}
    for {
      if len(batch) > maxSize {
        out <- batch
        continue Start
      }
      item, active := <- in
      if !active {
        if len(batch) > maxSize {
          out <- batch
        }
        return
      }
      batch = append(batch, item)
    }
  }()
  return out
}
```

