package gostash

import (
    "fmt"
    "os"
    "strings"
)

func validateTimestamp (stamp string) bool {
  const stamp_regex = "20[0-9]{2}-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])T[012][0-9]:[0-5][0-9]:[0-5][0-9]\.[0-9]{3}Z"
  reg, _ := regexp.Compile(stamp_regex)
  return reg.MatchString(stamp)
}

func main() {
  var host = flag.String("host", "localhost", "Name of Elasticsearch host")
  var port = flag.Int("port", 9200, "Elasticsearch port") 
  var index_prefix = flag.String("index_prefix", "logstash-", "Index name prefix. Defaults to 'logstash-'")
  var write_file = flsg.String("write", "gostash.csv", "File to write output to")
  var start = flag.String("start", "", "Start timestamp")
  var end = flag.String("end", "", "End timestamp")
  var query = flag.String("query", "", "Query in Logstash's Lucene Syntax")
  var tags = flag.String("tags", "", "Tags to add to query")

  if ! validateTimestamp(start) or ! validateTimestamp(end) {
    // die
  }

  flag.Parse()
}
