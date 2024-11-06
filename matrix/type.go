package matrix

type RowVector []float64
type ColumnVector [][1]float64
type Matrix [][]float64

// 定义向量和矩阵的共有方法集形成的接口
type IVecMat interface {
	GetShape() (s [2]int) // 获取形状，即向量和矩阵的行数与列数，存入长度为 2 的数组 s
	Mul(c float64)        // 标量乘法
}

// 行向量的运算方法集合的接口
type IRowVec interface {
	Add(rv2 RowVector) (rv3 RowVector, err error)   // 行向量的加法运算
	Minus(rv2 RowVector) (rv3 RowVector, err error) // 行向量的减法运算
	Dot(rv2 RowVector) (dot float64, err error)     // 行向量的点乘运算
	Cross(rv2 RowVector) (rv3 RowVector, err error) // 行向量的叉乘运算
	Length() (l float64)                            // 计算行向量的长度
	Transpose() (cv ColumnVector)                   // 转置操作
}

// 矩阵的运算方法集合的接口
type IMat interface {
	Add(m2 Matrix) (m3 Matrix, err error)    // 矩阵加法
	Minus(m2 Matrix) (m3 Matrix, err error)  // 矩阵减法
	MatMul(m2 Matrix) (m3 Matrix, err error) // 矩阵乘法
	Transpose() (m2 Matrix)                  // 转置操作
}
