package gostash

import (
    "fmt"
    "net"
    "os"
    "time"
    "regexp"
    "strings"
)


type Gostash struct {
  Host  string
  Port  int
  Index_prefix  string
  Write_file  string
  Start string
  End string
  Query String
  Tags  String
  Esconn  Connection
}

func NewGostash(host string, port int, index_prefix string, write_file string, start string, end string, query string, tags string) Gostash {
  gs := new(Gostash)
  gs.Host = host
  gs.Port = port
  gs.Index_prefix = index_prefix
  gs.Write_file = write_file
  gs.Start = start
  gs.End = end
  gs.Query = query
  gs.Tags = tags
  gs.Esconn = goes.NewConnection(gs.Host, gs.Port)
  
  return gs
}

func (gs *Gostash) RunQuery() {

}

func GetIndices(start string, end string) string {

}
