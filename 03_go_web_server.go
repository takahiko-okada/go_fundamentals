// https://www.youtube.com/watch?v=-QmdZ7821wA

package main

import ("fmt"
        "net/http")

// argument  datatype(customed one), requrest(* means reading through the request)
func index_handler(w http.ResponseWriter, r *http.Request) {
// func Fprintf(w io.Writer, format string, a ...interface{}) (n int, err error)
// Fprintf formats according to a format specifier and writes to w. It returns the number of bytes written and any write error encountered.
  fmt.Fprintf(w, "Whoa, Go is neat!")
}

func about_handler(w http.ResponseWriter, r *http.Request) {
  fmt.Fprintf(w, "Expert web design by Taka")
}

func main() {
                  // root, and action
  http.HandleFunc("/", index_handler)
                  // port, serverstuff, for now, nil

  http.HandleFunc("/about", about_handler)
  http.ListenAndServe(":8000", nil)
}

// Go doesn't have Class



