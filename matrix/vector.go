package matrix

import (
	"errors"
	"math"
)

// 获取行向量的维度
func (rv1 RowVector) GetShape() (s [2]int) {
	s[0] = 1
	s[1] = len(rv1)
	return
}

// 计算行向量与标量相乘
// 这里使用 *RowVector 指针类型的变量 rv1 作为指针接收者
// 在方法运行完成后改变 rv1 自身
func (rv1 *RowVector) Mul(c float64) {
	l := rv1.GetShape()[1]          // 获取 rv1 的维度，也就是 []float64 切片的长度
	rvTemp := make([]float64, l, l) // 创建一个 rvTemp 变量，用于复制和存放 rv1 指针变量所指向的底层 []float 切片
	copy(rvTemp, *rv1)              // 复制 rv1 指针对应的切片，赋值给 rvTemp 变量
	for i := 0; i < l; i++ {        // 遍历 rvTemp 切片，令每个元素乘以 c
		rvTemp[i] *= c
	}
	copy(*rv1, rvTemp) // 讲 rvTemp 复制给 *rv1 所对应的底层切片，从而改变 rv1
}

// 根据给定向量长度 l 初始化向量， l 需要大于 0，否则返回 err
func NewRowVector(l int) (rv1 RowVector, err error) {
	if l <= 0 { // 若小于或等于 0，则返回 err
		err = errors.New("the dimension of row vector must > 0")
		return
	}
	rv1 = make([]float64, l, l)
	return
}

// 定义行向量的加法运算，判断两个行向量的形状是否相同
func (rv1 RowVector) Add(rv2 RowVector) (rv3 RowVector, err error) {
	if rv1.GetShape()[1] != rv2.GetShape()[1] {
		err = errors.New("the dimensions of the two vectors are not equal")
		return
	}
	rv3, _ = NewRowVector(rv1.GetShape()[1])
	for i, v := range rv1 {
		rv3[i] = v + rv2[i]
	}
	return
}

// 定义行向量的减法运算，判断两个行向量的形状是否相同
func (rv1 RowVector) Minus(rv2 RowVector) (rv3 RowVector, err error) {
	if rv1.GetShape()[1] != rv2.GetShape()[1] { // 若行向量的维度不同，则返回 err
		err = errors.New("the dimensions of the two vectors are not equal")
		return
	}
	rv3, _ = NewRowVector(rv1.GetShape()[1])
	for i, v := range rv1 {
		rv3[i] = v - rv2[i]
	}
	return
}

// 定义行向量的点乘运算
func (rv1 RowVector) Dot(rv2 RowVector) (dot float64, err error) {
	if rv1.GetShape()[1] != rv2.GetShape()[1] {
		err = errors.New("the dimensions of the two vectors are not equal")
		return
	}
	for i, v := range rv1 {
		dot += v * rv2[i]
	}
	return
}

// 由于大于三维的向量叉乘运算较为复杂
// 这里简单期间， 仅实现二维和三维向量的叉乘
func (rv1 RowVector) Cross(rv2 RowVector) (rv3 RowVector, err error) {
	if rv1.GetShape()[1] != rv2.GetShape()[1] {
		err = errors.New("the dimensions of the two vectors are not equal")
		return
	}
	dim := rv1.GetShape()[1]
	if dim != 2 && dim != 3 {
		err = errors.New("only 2 or 3 dimensions row vector is supported")
		return
	}
	rv3, _ = NewRowVector(3)
	switch dim {
	case 2:
		// 两个二维向量进行叉乘运算，得到的是一个与两个二维向量垂直的向量
		// 可以将其想象为两个XY平面上的向量做叉乘，得到沿Z轴方向的向量
		rv3[0] = 0
		rv3[1] = 0
		rv3[2] = rv1[0]*rv2[1] - rv1[1]*rv2[0]
	case 3:
		// 两个三维向量进行叉乘运算，计算公式如下:
		// XxY=[x2*y3-x3*y2, x3*y1-x1*y3, x1*y2-x2*y1]
		rv3[0] = rv1[1]*rv2[2] - rv1[2]*rv2[1]
		rv3[1] = rv1[2]*rv2[0] - rv1[0]*rv2[2]
		rv3[2] = rv1[0]*rv2[1] - rv1[1]*rv2[0]
	}
	return
}

// 定义计算行向量的长度（模）方法
func (rv1 RowVector) Length() (l float64) { // l 的初始值默认是 0.0
	for _, v := range rv1 {
		l += v * v
	}
	l = math.Sqrt(l)
	return
}

// 定义行向量的转置运算，输出结果为列向量
func (rv1 RowVector) Transpose() (cv ColumnVector) {
	// Go 语言会对 ColumnVector 和 [][1]float64 两个类型进行隐式转换
	cv = make([][1]float64, rv1.GetShape()[1], rv1.GetShape()[1])
	for i, v := range rv1 {
		cv[i][0] = v
	}
	return
}
