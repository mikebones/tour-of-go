package battlehsip

import "fmt"

// Write any import statements here

func getHitProbability(R int32, C int32, G [][]int32) float64 {
  // r rows
  // c columns
  // 0 or more battleships on the grid
  // each occuping signel distinct cell
  // i'th row from top
  // j'th row from the left
  // either
  // G[i][j] = 1 (contains)
  // G[i][j] = 0 (does not contain)
  // r * c possible cells
  fmt.Println(len(G))
  for i := 0 ; i < len(G); i++ {

  }
  return 0.0
}
