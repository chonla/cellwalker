package cellwalker

import (
	"fmt"
	"math"
)

// Sheet returns sheet boundary
func Sheet() *Range {
	return newRange(1, 1, RowsLimit, ColumnsLimit)
}

// ColumnIndexToName converts column index to default excel name
func ColumnIndexToName(id int) string {
	name := ""
	dividend := id
	modulo := 0

	for dividend > 0 {
		modulo = (dividend - 1) % 26
		name = fmt.Sprintf("%c%s", rune(modulo+'A'), name)
		dividend = (dividend - modulo) / 26
	}
	return name
}

// ColumnNameToIndex converts default excel column name to index, 1-based index
// name must be uppercase start from A, B, C, ..., Z, AA, AB, ... ZZ, AAA, ..., ZZZ, ...
func ColumnNameToIndex(name string) int {
	index := 0
	for colCharIndex, colCharLen := 0, len(name); colCharIndex < colCharLen; colCharIndex++ {
		charNum := int(name[colCharIndex]-'A') + 1
		digitNum := (colCharLen - (colCharIndex + 1))
		columnWeight := int(math.Pow(26.0, float64(digitNum)))
		columnValue := charNum * columnWeight
		index += columnValue
	}
	return index
}
