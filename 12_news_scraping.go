// https://www.youtube.com/watch?v=ccANcNk8Dac&t=22s&list=PLQVvvaa0QuDeF3hP0wQoSxpkqgRcgxMqX&index=10
package main

import ("fmt"
"net/http"
"io/ioutil"
"encoding/xml")


type SitemapIndex struct {
  // slice of location type
  Locations []string `xml:"sitemap">loc`
// > means location tag under sitemap tag

// type Location struct {
//   Loc string `xml:"loc"`
// }

// func (l Location) String() string {
//   return fmt.Sprintf(l.Loc)
// Sprintf
}

type News struct {
  Titles []string `xml:"url>news>title"`
  Keywords []string `xml:"url>news>Keywords"`
  Locations []string `xml:"url>loc"`
}

// pulling info from the Internet
func main() {
  var s SitemapIndex
  var n News
  //  _ means a variable you don't intend to use after.
  resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  xml.Unmarshal(bytes, &s)
  // fmt.Println(s.Locations)
  for _, Location := range s.Locations {
    resp, _ := http.Get(Location)
    bytes, _ := ioutil.ReadAll(resp.Body)
    xml.Unmarshal(bytes, &n)

  }
}
