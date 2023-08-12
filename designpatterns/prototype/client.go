package prototype

import "fmt"

func ClientCode() {
  connectionForServiceA := NewConnection()
  connectionForServiceB := connectionForServiceA.Clone()
  fmt.Println(connectionForServiceA)
  fmt.Println(connectionForServiceB)
  // aborting a connection for serviceA does not hampers the connection for serviceB since cloning is done via deep copy
  connectionForServiceA.AbortConnection()
  fmt.Println(connectionForServiceA)
  fmt.Println(connectionForServiceB)
}

