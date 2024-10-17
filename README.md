# Cell Walker

[![codecov](https://codecov.io/github/chonla/cellwalker/graph/badge.svg?token=N2WFUOD9OB)](https://codecov.io/github/chonla/cellwalker) [![Go Report Card](https://goreportcard.com/badge/github.com/chonla/cellwalker)](https://goreportcard.com/report/github.com/chonla/cellwalker) [![go reference](https://camo.githubusercontent.com/b1b2c65c10b852ce2d33a873f1adbc216daf4679968a207ad6c895dd011a4cd9/68747470733a2f2f676f646f632e6f72672f6769746875622e636f6d2f626d617570696e2f676f2d68746d6c7574696c3f7374617475732e737667)](https://pkg.go.dev/github.com/chonla/cellwalker)

Cell Walker is a go package for virtually traversing Excel cell by cell's name. The package does not actually traverse into a real Excel file.

## Example

```go
package main

import (
	"fmt"

	"github.com/chonla/cellwalker"
)

func main() {
	// Walk from a cell to other cell
	fmt.Println(cellwalker.At("B3").Right().Below().String()) // C4

	// Jump from a cell to other cell
	fmt.Println(cellwalker.At("C2").ColumnOffset(5).RowOffset(10).String()) // H12

	// Too far jump from a cell to other cell will hit the limit of boundary
	fmt.Println(cellwalker.At("ZZZ2").ColumnOffset(5).RowOffset(10).String()) // XFD12

	// Range walking apply other boundary to walker
	fmt.Println(cellwalker.Within("C2:H3").At("C3").Right().Below().String()) // D3

	// Too far jump in a new boundary
	fmt.Println(cellwalker.Within("C2:H3").At("ZZZ2").ColumnOffset(5).RowOffset(10).String()) // XFD12

	// Range traversal
	result1 := cellwalker.Within("B3:E5").At("C4") // Define range and initial cell position
	fmt.Println(result1.String()) // C4
	result2 := result1.Tour()
	fmt.Println(result2.String()) // D4
	result3 := result2.Tour()
	fmt.Println(result3.String()) // E4
	result4 := result3.Tour()
	fmt.Println(result4.String()) // B5
	result5 := result4.Tour()
	fmt.Println(result5.String()) // C5
	result6 := result5.Tour()
	fmt.Println(result6.String()) // D5
	result7 := result6.Tour()
	fmt.Println(result7.String()) // E5
	result8 := result7.Tour()
	fmt.Println(result8 == nil) // true
}
```

## License

[MIT](LICENSE)