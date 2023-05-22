

func takeFirstN(ctx context.Context, dataSource <-chan interface{}, n int) <-chan interface{} {
    // 1
    takeChannel := make(chan interface{})
    
    // 2
    go func() {
        defer close(takeChannel)

        // 3 
        for i := 0; i< n; i++ {
          select {
            case val, ok := <-dataSource:
              if !ok{
                return
              }
              takeChannel <- val
            case <-ctx.Done():
              return
          }
        }
    }()
    return takeChannel
}

func main() {

  done := make(chan struct{})
  defer close(done)

  // Generates a channel sending integers
  // From 0 to 9
  range10 := rangeChannel(done, 10)

  for num := range takeFirstN(done, range10, 5) {
      fmt.Println(num)
  }
}