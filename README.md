# Cell Walker

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