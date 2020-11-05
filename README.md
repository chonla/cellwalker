# Cell Walker

Cell Walker is a go package for virtually traversing Excel cell by cell's name. The package does not actually traverse into a real Excel file.

## Example

```
package main

import (
	"fmt"

	"github.com/chonla/cellwalker"
)

func main() {
	fmt.Println(cellwalker.At("B3").Right().Below().String()) // C4
}
```

## License

[MIT](LICENSE)