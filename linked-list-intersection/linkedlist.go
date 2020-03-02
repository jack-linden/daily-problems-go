package linkedlist

import "fmt"
type Node struct {
    val int
    next * Node
}

type LinkedList struct {
    head * Node
    len int
}

func(l * LinkedList) String() string {
    str := fmt.Sprintf("{ len = %d, [", l.len)
    curr := l.head
    for curr != nil {
        str += fmt.Sprintf("{%d} ", curr.val)
        curr = curr.next
    }

    str += "]"
    return str
}

func(l * LinkedList) Append(other * LinkedList) {
    if other.len <= 0 {
        return
    }
    if l.head == nil {
        l.head = other.head;
    } else {
        curr := l.head
        for curr.next != nil {
            curr = curr.next;
        }
        curr.next = other.head
    }
    l.len += other.len
}

func NewLinkedList(vals[] int)(list LinkedList) {
    if len(vals) == 0 {
        return LinkedList {len: 0}
    }
    var node Node = Node {val: vals[0]}
    curr := & node
    for i := 1;
    i < len(vals);
    i++{
        next := Node {val: vals[i]}
        curr.next = &next
        curr = curr.next
    }

    return LinkedList {head: &node,len: len(vals)}
}

func FindIntersection(M * LinkedList, N * LinkedList)(result * Node) {

    lenM, lenN, absDiff := M.len, N.len, abs(M.len - N.len)
    mCurr, nCurr := M.head, N.head

    if lenM > lenN {
        for;
        absDiff > 0;
        absDiff--{
            mCurr = mCurr.next
        }
    } else {
        for;
        absDiff > 0;
        absDiff--{
            nCurr = nCurr.next
        }
    }

    for mCurr != nCurr {
        mCurr = mCurr.next
        nCurr = nCurr.next
    }

    return mCurr
}

func abs(x int)(int) {
    if x < 0 {
        return -x;
    }
    return x;
}
