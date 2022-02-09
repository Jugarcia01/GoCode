package main

import (
	"fmt"
	"sort"	
)

func main() {
  var numbers = []int{2,3,1}

  firstOddIndex := sort.Search(len(numbers), 
  func(i int) bool { 
    return numbers[i]%2 == 1 
  })
  //firstOddIndex is 1
  fmt.Println(firstOddIndex)
}


