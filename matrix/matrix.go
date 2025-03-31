package matrix

import "fmt"

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

func (m *Matrix) At(i, j int) float64 {
    return m.mat[i][j]
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



