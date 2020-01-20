package main

import (
  "fmt"
  "math"
)

/**
*In a row of dominoes, A[i] and B[i] represent the
* top and bottom halves of the i-th domino.  (A
* domino is a tile with two numbers from 1 to 6 -
* one on each half of the tile.)

*We may rotate the i-th domino, so that A[i] and B
*[i] swap values.

*Return the minimum number of rotations so that
* all the values in A are the same, or all the
* values in B are the same.

*If it cannot be done, return -1.
*
*Input: A = [2,1,2,4,2,2], B = [5,2,6,2,3,2]
*Output: 2
* 
**/


func minDominoRotations(A []int, B []int) int {
  var equalCount int
  aLength := len(A)
  topCounts := make([]int, 6)
  bottomCounts := make([]int, 6)
  totalNecessary := []int{aLength, aLength, aLength, aLength, aLength, aLength}
  minimum := -1

  for i := 0; i < len(A); i++ {
    if A[i] != B[i] {
      topCounts[A[i]-1]++
      bottomCounts[B[i]-1]++
    } else {
      equalCount++
      totalNecessary[A[i]-1]--
    }
  }

  if equalCount == len(A) {
    return 0
  }

  for i := 0; i < 6; i++ {
    if topCounts[i] + bottomCounts[i] == totalNecessary[i] {
      minCount := int(math.Min(float64(topCounts[i]), float64(bottomCounts[i])))
      if minimum == -1 {
        minimum = minCount
      } else {
        minimum = int(math.Min(float64(minimum), float64(minCount)))
      }

    }
  }
  
  return minimum
}

func main() {
  fmt.Println(minDominoRotations([]int{1,1,1,1,1,1,1,1}, []int{1,1,1,1,1,1,1,1}))
}