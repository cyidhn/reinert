package main

import (
   "fmt"
   "regexp"
)

func main() {
   s := "Abraham Lincoln 43 4 3 65 http:www.google.fr"
   reg := regexp.MustCompile(`http\S+`)
   res := reg.ReplaceAllString(s, "${1}")
   fmt.Println(res) // Abraham Lincoln
}