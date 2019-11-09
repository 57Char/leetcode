package algorithms

import "sync"

// https://leetcode.com/problems/design-circular-deque/

type MyCircularDeque struct {
	mu   *sync.Mutex
	size int
	data []int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{
		mu:   &sync.Mutex{},
		size: k,
		data: make([]int, 0, k),
	}
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	this.mu.Lock()
	defer this.mu.Unlock()
	if len(this.data) == this.size {
		return false
	}
	this.data = append([]int{value}, this.data[:]...)
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	this.mu.Lock()
	defer this.mu.Unlock()
	if len(this.data) == this.size {
		return false
	}
	this.data = append(this.data, value)
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	this.mu.Lock()
	defer this.mu.Unlock()
	if len(this.data) == 0 {
		return false
	}
	this.data = this.data[1:]
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	this.mu.Lock()
	defer this.mu.Unlock()
	if len(this.data) == 0 {
		return false
	}
	this.data = this.data[:len(this.data)-1]
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	this.mu.Lock()
	defer this.mu.Unlock()
	if len(this.data) == 0 {
		return -1
	}
	return this.data[0]
}


/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	this.mu.Lock()
	defer this.mu.Unlock()
	if len(this.data) == 0 {
		return -1
	}
	return this.data[len(this.data)-1]
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	this.mu.Lock()
	defer this.mu.Unlock()
	return len(this.data) == 0
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	this.mu.Lock()
	defer this.mu.Unlock()
	return len(this.data) == this.size
}

// ====================================================

type MyCircularDequeV2 struct {
	size  int
	len   int
	front int
	rear  int
	data  []int
}


/** Initialize your data structure here. Set the size of the deque to be k. */
func ConstructorV2(k int) MyCircularDequeV2 {
	return MyCircularDequeV2{
		size:  k,
		front: -1,
		rear:  -1,
		data:  make([]int, 0, k),
	}
}


/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDequeV2) InsertFront(value int) bool {
	if this.IsFull() {
		return false
	}
	this.data = append([]int{value}, this.data[:]...)
	this.len++
	this.front = value
	this.rear = this.data[this.len-1]
	return true
}


/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDequeV2) InsertLast(value int) bool {
	if this.IsFull(){
		return false
	}
	if this.IsEmpty() {
		this.front = value
	}
	this.data = append(this.data, value)
	this.len++
	this.rear = value
	return true
}


/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDequeV2) DeleteFront() bool {
	if this.IsEmpty() {
		return false
	}
	this.data = this.data[1:]
	this.len--
	if this.IsEmpty() {
		this.front = -1
		this.rear = -1
	} else {
		this.front = this.data[0]
	}
	return true
}


/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDequeV2) DeleteLast() bool {
	if this.IsEmpty() {
		return false
	}
	this.data = this.data[:this.len-1]
	this.len--
	if this.IsEmpty() {
		this.front = -1
		this.rear = -1
	} else {
		this.rear = this.data[this.len-1]
	}
	return true
}


/** Get the front item from the deque. */
func (this *MyCircularDequeV2) GetFront() int {
	return this.front
}


/** Get the last item from the deque. */
func (this *MyCircularDequeV2) GetRear() int {
	return this.rear
}


/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDequeV2) IsEmpty() bool {
	return this.len == 0
}


/** Checks whether the circular deque is full or not. */
func (this *MyCircularDequeV2) IsFull() bool {
	return this.len == this.size
}


/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
