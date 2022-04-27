package main

import (
   "fmt"
   "regexp"
)

func main() {
   s := `"Abraham Lincoln" @en`
   reg := regexp.MustCompile(`"([^"]*)" *@en`)
   res := reg.ReplaceAllString(s, "${1}")
   fmt.Println(res) // Abraham Lincoln
}