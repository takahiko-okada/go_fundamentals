// https://www.youtube.com/watch?v=ccANcNk8Dac&t=22s&list=PLQVvvaa0QuDeF3hP0wQoSxpkqgRcgxMqX&index=10
package main

import ("fmt"
"net/http"
"io/ioutil")

// pulling info from the Internet
func main() {
  //  _ means a variable you don't intend to use after.
  resp, _ := http.Get("https://www.washingtonpost.com/news-sitemap-index.xml")
  bytes, _ := ioutil.ReadAll(resp.Body)
  string_body := string(bytes)
  fmt.Println(string_body)
  resp.Body.Close()
}
