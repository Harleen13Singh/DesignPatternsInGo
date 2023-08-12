package prototype

import "time"

type Connection interface {
  // ConnectionHealthCheck tells the health of the connection
  ConnectionHealthCheck() bool
  // ConnectionInfo returns the basic info of the connection like
  // 1. what type of connection is this (eg. db connection)
  ConnectionInfo() string
  // returns the deep copy of the connection
  Clone() Connection
  // aborts the connection and hence cannot be used after abortion
  AbortConnection()
}

func NewConnection() Connection{
  // to denote that creating a new connection is an expensive operation
  time.Sleep(5 * time.Second)
  return NewDb()
}