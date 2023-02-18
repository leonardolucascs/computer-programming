package main

import "fmt"

func main(){
    // Sort slice by passing anonymous function to sort.Slice for custom sorting
    a := []int{5, 3, 4, 7, 8, 9}
    sort.Slice(a, func(i, j int) bool {
        return a[i] < a[j]
    })
    for _, v := range a {
        fmt.Println(v)
    }


    // Sort int Slice by sort.IntSlices()

    // Sort strings by sort    
}



