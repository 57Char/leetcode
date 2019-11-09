package algorithms

import (
	"testing"
)

import (
	"github.com/stretchr/testify/assert"
)

func TestMyCircularDeque(t *testing.T) {
	// Test case from from https://leetcode.com/problems/design-circular-deque/
	//MyCircularDeque circularDeque = new MyCircularDeque(3); // set the size to be 3
	//circularDeque.insertLast(1);			// return true
	//circularDeque.insertLast(2);			// return true
	//circularDeque.insertFront(3);			// return true
	//circularDeque.insertFront(4);			// return false, the queue is full
	//circularDeque.getRear();  			// return 2
	//circularDeque.isFull();				// return true
	//circularDeque.deleteLast();			// return true
	//circularDeque.insertFront(4);			// return true
	//circularDeque.getFront();			    // return 4

	circularDeque := Constructor(3)
	assert.Equal(t, true, circularDeque.InsertLast(1))   // return true
	assert.Equal(t, true, circularDeque.InsertLast(2))   // return true
	assert.Equal(t, true, circularDeque.InsertFront(3))  // return true
	assert.Equal(t, false, circularDeque.InsertFront(4)) // return false, the queue is full
	assert.Equal(t, 2, circularDeque.GetRear())                // return 2
	assert.Equal(t, true, circularDeque.IsFull())              // return true
	assert.Equal(t, true, circularDeque.DeleteLast())          // return true
	assert.Equal(t, true, circularDeque.InsertFront(4))  // return true
	assert.Equal(t, 4, circularDeque.GetFront())               // return 4

	circularDequeV2 := Constructor(3)
	assert.Equal(t, true, circularDequeV2.InsertLast(1))   // return true
	assert.Equal(t, true, circularDequeV2.InsertLast(2))   // return true
	assert.Equal(t, true, circularDequeV2.InsertFront(3))  // return true
	assert.Equal(t, false, circularDequeV2.InsertFront(4)) // return false, the queue is full
	assert.Equal(t, 2, circularDequeV2.GetRear())                // return 2
	assert.Equal(t, true, circularDequeV2.IsFull())              // return true
	assert.Equal(t, true, circularDequeV2.DeleteLast())          // return true
	assert.Equal(t, true, circularDequeV2.InsertFront(4))  // return true
	assert.Equal(t, 4, circularDequeV2.GetFront())               // return 4

	// other test cases

	//`
	//["MyCircularDeque","insertFront","deleteLast","deleteLast","insertLast","getRear","insertLast","isFull","getRear","isFull","deleteLast","isEmpty","getFront","isEmpty","insertFront","isEmpty","insertLast","getRear","getFront","deleteLast","deleteLast","insertLast","insertLast","deleteLast","getFront","insertLast","isEmpty","getFront","insertFront","insertLast","insertFront","deleteLast","getRear","getRear","insertFront","insertLast","deleteFront","getFront","getFront","insertLast","getFront","getFront","insertFront","getFront","deleteFront","deleteFront","deleteFront","insertLast","getRear","insertFront","isFull","insertLast","deleteLast","getRear","deleteLast","getRear","insertFront","deleteLast","insertLast","insertLast","getFront","getRear","insertFront","insertLast","insertLast","insertFront","getRear","getRear","deleteLast","insertFront","getFront","insertFront","insertLast","getRear","insertFront","insertLast","deleteFront","isEmpty","getRear","insertLast","insertFront","getFront","getRear","insertFront","insertLast","deleteLast","deleteLast","isFull","insertLast","deleteLast","deleteFront","insertFront","getRear","isFull","deleteFront","getFront","getRear","isEmpty","getFront","deleteLast","getRear","insertLast"]
	//[[71],[47],[],[],[66],[],[72],[],[],[],[],[],[],[],[53],[],[15],[],[],[],[],[9],[87],[],[],[98],[],[],[33],[40],[21],[],[],[],[69],[14],[],[],[],[60],[],[],[15],[],[],[],[],[97],[],[66],[],[23],[],[],[],[],[30],[],[54],[18],[],[],[50],[24],[24],[25],[],[],[],[0],[],[86],[31],[],[55],[55],[],[],[],[18],[49],[],[],[53],[44],[],[],[],[46],[],[],[58],[],[],[],[],[],[],[],[],[],[61]]
	//
	//[null,true,true,false,true,66,true,false,72,false,true,false,66,false,true,false,true,15,53,true,true,true,true,true,53,true,false,53,true,true,true,true,98,98,true,true,true,21,21,true,21,21,true,15,true,true,true,true,97,true,false,true,true,97,true,60,true,true,true,true,30,18,true,true,true,true,24,24,true,true,0,true,true,31,true,true,true,false,55,true,true,49,18,true,true,true,true,false,true,true,true,true,55,false,true,49,55,false,49,true,31,true]
	//`

	//`
	//["MyCircularDeque","insertFront","getRear","insertFront","getRear","insertLast","getFront","getRear","getFront","insertLast","deleteLast","getFront"]
	//[[3],[9],[],[9],[],[5],[],[],[],[8],[],[]]
	//
	//[null,true,9,true,9,true,9,5,9,false,true,9]
	//`
}
