package matrix

import (
	"errors"
	"strconv"
	"strings"
)

type Matrix struct {
	data [][]int
}

func New(in string) (*Matrix, error) {
	matrix := &Matrix{}

	rows := strings.Split(in, "\n")
	var lastCols []string

	for i := 0; i < len(rows); i++ {
		cols := strings.Fields(rows[i])

		if lastCols != nil && len(cols) != len(lastCols) {
			return nil, errors.New("uneven rows")
		}

		matrix.data = append(matrix.data, make([]int, len(cols)))

		for j := 0; j < len(cols); j++ {
			val, err := strconv.Atoi(cols[j])

			if err != nil {
				return nil, err
			}

			matrix.data[i][j] = int(val)
		}

		lastCols = cols
	}

	return matrix, nil
}

func (m *Matrix) Rows() [][]int {
	ret := make([][]int, 0)
	for _, row := range m.data {
		cp := make([]int, len(row))
		copy(cp, row)
		ret = append(ret, cp)
	}
	return ret
}

func (m *Matrix) Cols() (ret [][]int) {
	if len(m.data) == 0 {
		return
	}

	ret = make([][]int, len(m.data[0]))

	for i := 0; i < len(ret); i++ {
		ret[i] = make([]int, 0)

		for j := 0; j < len(m.data); j++ {
			ret[i] = append(ret[i], m.data[j][i])
		}
	}

	return
}

// Set places the specified value at the specified location
func (m *Matrix) Set(r int, c int, val int) bool {
	if len(m.data) == 0 {
		return false
	}

	if r > len(m.data)-1 || r < 0 {
		return false
	}

	if c > len(m.data[0])-1 || c < 0 {
		return false
	}

	m.data[r][c] = val

	return true
}
