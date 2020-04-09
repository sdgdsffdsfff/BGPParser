// MIT License

// Copyright (c) 2019 Yuefeng Zhu

// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:

// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.

// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package analysis

import (
	"sort"
	"strings"
)

func (a AspathList) Len() int {
	return len(a)
}

func (a AspathList) Less(i, j int) bool {
	_aspath1 := len(strings.Split(a[i], " "))
	_aspath2 := len(strings.Split(a[j], " "))
	if _aspath1 != _aspath2 {
		return _aspath1 < _aspath2
	}
	return a[i] < a[j]
}

func (a AspathList) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}

func (b *BGPInfo) SortASpathBySize() {
	sort.Sort(b.Aspath)
	b.isSorted = true
	for _, v := range b.Aspath {
		b.Aspath2str += v
	}
	return
}
