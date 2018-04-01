// https://www.youtube.com/watch?v=ccANcNk8Dac&t=22s&list=PLQVvvaa0QuDeF3hP0wQoSxpkqgRcgxMqX&index=10
package main

import ("fmt"
"net/http"
"io/ioutil"
"encoding/xml")


type SitemapIndex struct {
  // slice of location type
  Locations []Location `xml:"sitemap"`
}

type Location struct {
  Loc string `xml:"loc"`
}

func (l Location) String() string {
  return fmt.Sprintf(l.Loc)
  // Sprintf
}

// pulling info from the Internet
func main() {
  //  _ means a variable you don't intend to use after.
  resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  resp.Body.Close()

  var s SitemapIndex
  xml.Unmarshal(bytes, &s)
  fmt.Println(s.Locations)
}
