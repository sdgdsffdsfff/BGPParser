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

package utils

import "errors"

type Queue struct {
	val         []interface{}
	front, rare int
	size        int
	capacity    int
}

func NewQueue() *Queue {
	// default size: 2
	return &Queue{
		val:      make([]interface{}, 2),
		front:    0,
		rare:     0,
		size:     0,
		capacity: 2,
	}
}

func (q *Queue) Push(x interface{}) {
	if q.size == q.capacity {
		buf := NewQueue()
		buf.val = make([]interface{}, q.capacity*2)
		size := q.size
		for i := 0; i < size; i++ {
			buf.val[i] = q.val[q.front]
			q.front = (q.front + 1) % q.capacity
		}
		q.val = buf.val
		q.front = 0
		q.rare = q.size
		q.capacity *= 2
	}
	q.val[q.rare] = x
	q.rare = (q.rare + 1) % q.capacity
	q.size++
	return
}

func (q *Queue) Pop() (x interface{}, err error) {
	if q.IsEmpty() {
		return nil, errors.New("Queue is empty")
	}
	x = q.val[q.front]
	q.front = (q.front + 1) % q.capacity
	q.size--
	return x, nil
}

func (q *Queue) IsEmpty() bool {
	return q.size == 0
}

func (q *Queue) Size() int {
	return q.size
}
