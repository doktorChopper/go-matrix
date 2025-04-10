package matrix

import (
	"fmt"
	"math"
	"math/rand"
)

type Matrix struct {
    rows    int
    cols    int
    mat     [][]float64
}

func NewMatrix() *Matrix {
    return &Matrix {
        rows:   0,
        cols:   0,
        mat:    nil,
    }
}

func randFloats(min, max float64) float64 {
    return min + rand.Float64() * (max - min)
}

func NewRandomMatrix(n, m int) *Matrix {

    min := rand.Float64() * 10
    max := rand.Float64() * 10

    r := NewMatrixNM(n, m)
    
    for i := 0; i < r.rows; i++ {
        for j := 0; j < r.cols; j++ {
            r.mat[i][j] = math.Pow(-1, float64(rand.Intn(2))) * randFloats(min, max)
        }
    }

    return r
}

func NewMatrixNM(n, m int) *Matrix {

    r := make([][]float64, n)
    for i := range r {
        r[i] = make([]float64, m)
    }

    return &Matrix {
        rows:   n,
        cols:   m,
        mat:    r,
    }
}

func NewMatrixFromSlice(s [][]float64) *Matrix {

    return &Matrix {
        rows:   len(s),
        cols:   len(s[0]),
        mat:    s,
    }
}

func (m *Matrix) Cols() int {
    return m.cols
}

func (m *Matrix) Rows() int {
    return m.rows
}

func (m *Matrix) Transposition() *Matrix {

    mT := NewMatrixNM(m.cols, m.rows)

    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            mT.mat[j][i] = m.mat[i][j]
        }
    }

    return mT
}

func (m *Matrix) Minor(row, col int) *Matrix {
    ret := NewMatrixNM(m.rows - 1, m.cols - 1)

    r := 0
    for i := 0; i < m.rows; i++ {
        if i == row {
            continue
        }
        c := 0
        for j := 0; j < m.cols; j++ {
            if j == col {
                continue
            }
            ret.mat[r][c] = m.mat[i][j]
            c++
        }
        r++
    }

    return ret
}

func (m *Matrix) ScalarMult(v float64) {

    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            m.mat[i][j] *= v
        }
    }
}

func (m *Matrix) ScalarDiv(v float64) {

    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            m.mat[i][j] /= v
        }
    }
}

func (m *Matrix) Sub(o *Matrix) *Matrix {
    
    if m.cols != o.cols && m.rows != o.rows {
        return nil
    }

    ret := NewMatrixNM(m.rows, m.cols)

    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            ret.mat[i][j] = m.mat[i][j] - o.mat[i][j]
        }
    }

    return ret
}

func (m *Matrix) Add(o *Matrix) *Matrix {
    if m.cols != o.cols && m.rows != o.rows {
        return nil
    }

    ret := NewMatrixNM(m.rows, m.cols)

    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            ret.mat[i][j] = m.mat[i][j] + m.mat[i][j]
        }
    }

    return ret
}

func (m *Matrix) Mult(o *Matrix) *Matrix {

    if m.cols != o.rows {
        return nil
    }
    
    ret := NewMatrixNM(m.rows, o.cols)

    for i := 0; i < m.rows; i++ {
        for j := 0; j < o.cols; j++ {
            s := 0.0
            for k := 0; k < o.rows; k++ {
                s += m.mat[i][k] * o.mat[k][j]
            }
            ret.mat[i][j] = s
        }
    }
    return ret
}

func (m *Matrix) GetAt(i, j int) float64 {
    return m.mat[i][j]
}

func (m *Matrix) SetAt(i, j int, v float64) {
    m.mat[i][j] = v
}

func (m *Matrix) Det() (float64, error) {

    if m.cols != m.rows {
        return 0.0, nil
    }

    if m.rows == 1 {
        return m.mat[0][0], nil
    }

    if m.rows == 2 {
        return m.mat[0][0] * m.mat[1][1] - m.mat[0][1] * m.mat[1][0], nil
    }

    // TODO

    s := 0.0
    for i := 0; i < m.rows; i++ {
        minor := m.Minor(0, i)
        minorDet, err := minor.Det()

        if err == nil {
            if i % 2 != 0 {
                s -= m.mat[0][i] * minorDet
            } else {
                s += m.mat[0][i] * minorDet
            }
        }
    }

    return s, nil
}

func (m *Matrix) InverseMatrix() *Matrix {

    d, _ := m.Det()
    ret := m.AdjugateMatrix().Transposition()
    
    ret.ScalarDiv(d)

    return ret
} 

func (m *Matrix) AdjugateMatrix() *Matrix {

    ret := NewMatrixNM(m.rows, m.cols)

    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {

        // TODO check err

            d, _ := m.Minor(i, j).Det()
            ret.mat[i][j] = math.Pow(-1, float64(i) + float64(j)) * d
        }
    }
    return ret
}


func (m *Matrix) Display() {

    for _, r := range m.mat {

        for _, c := range r {
            fmt.Printf("%.2f  ", c)
        }
        fmt.Println()
    }

    fmt.Println()
}


