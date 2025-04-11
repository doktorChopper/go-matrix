package matrix

import (
	"errors"
	"fmt"
	"math"
	"reflect"
)

type Matrix struct {
    rows    int
    cols    int
    mat     [][]interface{}
}

func NewMatrix() *Matrix {
    return &Matrix {
        rows:   0,
        cols:   0,
        mat:    nil,
    }
}

func NewMatrixNM(n, m int) *Matrix {

    r := make([][]interface{}, n)
    for i := range r {
        r[i] = make([]interface{}, m)
    }

    return &Matrix {
        rows:   n,
        cols:   m,
        mat:    r,
    }
}

func NewMatrixFromSlice(s [][]interface{}) *Matrix {

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

func (m *Matrix) ScalarMult(v float64) error {

    if reflect.TypeOf(m.mat).Name() != "float64" {
        return errors.New("wrong type") 
    }

    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            m.mat[i][j] = m.mat[i][j].(float64) * v
        }
    }

    return nil
}

func (m *Matrix) ScalarDiv(v float64) error {

    if reflect.TypeOf(m.mat).Name() != "float64" {
        return errors.New("wrong type") 
    }

    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            m.mat[i][j] = m.mat[i][j].(float64) / v
        }
    }

    return nil
}

func (m *Matrix) Sub(o *Matrix) (*Matrix, error) {
    
    if m.cols != o.cols && m.rows != o.rows {
        return nil, errors.New("error")
    }

    if reflect.TypeOf(m.mat).Name() != "float64" &&
        reflect.TypeOf(o.mat).Name() != "float64" &&
        reflect.TypeOf(m.mat).Name() != reflect.TypeOf(o.mat).Name() {
        return nil, errors.New("error")
    }

    ret := NewMatrixNM(m.rows, m.cols)

    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            ret.mat[i][j] = m.mat[i][j].(float64) - o.mat[i][j].(float64)
        }
    }

    return ret, nil
}

func (m *Matrix) Add(o *Matrix) (*Matrix, error) {
    if m.cols != o.cols && m.rows != o.rows {
        return nil, errors.New("error")
    }

    if reflect.TypeOf(m.mat).Name() != "float64" &&
        reflect.TypeOf(o.mat).Name() != "float64" &&
        reflect.TypeOf(m.mat).Name() != reflect.TypeOf(o.mat).Name() {
        return nil, errors.New("error")
    }

    ret := NewMatrixNM(m.rows, m.cols)

    for i := 0; i < m.rows; i++ {
        for j := 0; j < m.cols; j++ {
            ret.mat[i][j] = m.mat[i][j].(float64) + m.mat[i][j].(float64)
        }
    }

    return ret, nil
}

func (m *Matrix) Mult(o *Matrix) (*Matrix, error) {

    if m.cols != o.rows {
        return nil, errors.New("error")
    }

    if reflect.TypeOf(m.mat).Name() != "float64" &&
        reflect.TypeOf(o.mat).Name() != "float64" &&
        reflect.TypeOf(m.mat).Name() != reflect.TypeOf(o.mat).Name() {
        return nil, errors.New("error")
    }
    
    ret := NewMatrixNM(m.rows, o.cols)

    for i := 0; i < m.rows; i++ {
        for j := 0; j < o.cols; j++ {
            s := 0.0
            for k := 0; k < o.rows; k++ {
                s += m.mat[i][k].(float64) * o.mat[k][j].(float64)
            }
            ret.mat[i][j] = s
        }
    }
    return ret, nil
}

func (m *Matrix) GetAt(i, j int) (float64, error) {

    if reflect.TypeOf(m.mat).Name() != "float64" {
        return -1, errors.New("wrong type") 
    }

    return m.mat[i][j].(float64), nil
}

func (m *Matrix) SetAt(i, j int, v float64) {
    m.mat[i][j] = v
}

func (m *Matrix) Det() (float64, error) {

    if reflect.TypeOf(m.mat).Name() != "float64" {
        return -1, errors.New("wrong type") 
    }

    if m.cols != m.rows {
        return 0.0, nil
    }

    if m.rows == 1 {
        return m.mat[0][0].(float64), nil
    }

    if m.rows == 2 {
        return m.mat[0][0].(float64) * m.mat[1][1].(float64) - m.mat[0][1].(float64) * m.mat[1][0].(float64), nil
    }

    // TODO

    s := 0.0
    for i := 0; i < m.rows; i++ {
        minor := m.Minor(0, i)
        minorDet, err := minor.Det()

        if err == nil {
            if i % 2 != 0 {
                s -= m.mat[0][i].(float64) * minorDet
            } else {
                s += m.mat[0][i].(float64) * minorDet
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

        for j := range r {
            fmt.Printf("%.2f  ", r[j].(float64))
        }
        fmt.Println()
    }

    fmt.Println()
}


