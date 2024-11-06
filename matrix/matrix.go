package matrix

import "errors"

// 定义获取矩阵形状的方法
func (m1 Matrix) GetShape() (s [2]int) {
	s[0] = len(m1)
	s[1] = len(m1[0])
	return
}

// 定义矩阵与标量 c 相乘的运算
func (m1 *Matrix) Mul(c float64) {
	row, col := m1.GetShape()[0], m1.GetShape()[1]
	mTemp := make([][]float64, row, row)
	copy(mTemp, *m1) // 将 m1 矩阵复制给临时变量 mTemp
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			mTemp[i][j] *= c
		}
	}
	copy(*m1, mTemp) // 将临时变量 mTemp 复制回 m1 矩阵
}

// 定义矩阵的初始化方法，r 为矩阵的行数，c 为矩阵的列数
func NewMatrix(r, c int) (mat Matrix, err error) {
	// 行数或列数小于或等于零，返回 err
	if r <= 0 || c <= 0 {
		err = errors.New("rows and columns of the matrix must > 0")
		return
	}
	// 初始化 mat 的行
	mat = make([][]float64, r, r)
	// 此处不使用 for range，是因为我们要改变遍历元素的值
	for i := 0; i < r; i++ {
		mat[i] = make([]float64, c, c)
	}
	return
}

// 定义矩阵的加法运算
func (m1 Matrix) Add(m2 Matrix) (m3 Matrix, err error) {
	if m1.GetShape() != m2.GetShape() {
		err = errors.New("the shape of the two matrix are not equal")
		return
	}
	// 获取矩阵的形状
	r, c := m1.GetShape()[0], m1.GetShape()[1]
	m3, _ = NewMatrix(r, c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m3[i][j] = m1[i][j] + m2[i][j]
		}
	}
	return
}

// 定义矩阵的减法运算
func (m1 Matrix) Minus(m2 Matrix) (m3 Matrix, err error) {
	if m1.GetShape() != m2.GetShape() {
		err = errors.New("the shape of the two matrix are not equal")
		return
	}
	r, c := m1.GetShape()[0], m1.GetShape()[1]
	m3, _ = NewMatrix(r, c)
	for i := 0; i < r; i++ {
		for j := 0; j < c; j++ {
			m3[i][j] = m1[i][j] - m2[i][j]
		}
	}
	return
}

// 定义两个矩阵相乘的运算
func (m1 Matrix) MatMul(m2 Matrix) (m3 Matrix, err error) {
	r1, c1 := m1.GetShape()[0], m1.GetShape()[1]
	r2, c2 := m2.GetShape()[0], m2.GetShape()[1]
	if c1 != r2 {
		err = errors.New("the shape of the two matrix are not match")
		return
	}
	// 初始化计算结果 m3
	m3, _ = NewMatrix(r1, c2)
	for i := 0; i < r1; i++ {
		for j := 0; j < c2; j++ {
			// 取出 m1 的第 i 行
			v1 := RowVector(m1[i])
			v2 := make([]float64, r2, r2)
			for k := 0; k < r2; k++ {
				// 讲 m2 第 j 列的值依次放入 v2 中
				v2[k] = m2[k][j]
			}
			// 利用行向量乘法计算 m3 的矩阵元素
			m3[i][j], _ = v1.Dot(v2)
		}
	}
	return
}

// 定义矩阵的转置运算，本质上就是将矩阵元素的索引值 i 和 j 进行互换
func (m1 Matrix) Transpose() (m2 Matrix) {
	r, c := m1.GetShape()[0], m1.GetShape()[1]
	m2, _ = NewMatrix(c, r)
	for i := 0; i < c; i++ {
		for j := 0; j < r; j++ {
			m2[i][j] = m1[j][i]
		}
	}
	return
}
