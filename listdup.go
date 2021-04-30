package main

import (
  "fmt"
)

func diff(a []string, b []string) []string {
    var m map[string]bool
    m = make(map[string]bool, len(b))
    for _, s := range b {
        m[s] = false
    }
    var diff []string
    for _, s := range a {
        if _, ok := m[s]; !ok {
            diff = append(diff, s)
            continue
        }


    }


    return diff
}

func main() {
    
l1 := []string{"a", "b", "c", "d","a"}
l2 := []string{"a", "c", "e"}
fmt.Println(diff(l1,l2))
}