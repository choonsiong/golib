/*
MIT License
Copyright (c) 2020 Lee Choon Siong
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/

package compare

import "bytes"

// Deprecated: 1.0.22
// CompareTwoByteSlices returns true if both byte slices are same, else returns false.
func CompareTwoByteSlices(bs1 []byte, bs2 []byte) bool {
	return bytes.Equal(bs1, bs2)

	//if len(bs1) != len(bs2) {
	//	return false
	//}
	//
	//for i := 0; i < len(bs1); i++ {
	//	if bs1[i] != bs2[i] {
	//		return false
	//	}
	//}
	//
	//return true
}
