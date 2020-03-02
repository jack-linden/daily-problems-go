package linkedlist

import (
    "testing"
)

func TestAppend(t * testing.T) {

    M := NewLinkedList([] int {1, 2, 3})
    N := NewLinkedList([] int {4, 5, 6})
    M.Append( &N)
    Expected := NewLinkedList([] int {1, 2, 3, 4, 5, 6})

    if M.String() != Expected.String() {
        t.Errorf("Appended list was incorrect, got: %v but wanted %v", M, Expected)
    }
}

func TestFindIntersection_EqualLengthLists(t * testing.T) {
    Shared := NewLinkedList([] int {8, 11})
    M := NewLinkedList([] int {1, 2, 3})
    N := NewLinkedList([] int {4, 5, 6})
    M.Append(&Shared)
    N.Append(&Shared)

    Observed := FindIntersection( &M, &N)
    Expected := Shared.head

    if Observed != Expected {
        t.Errorf("Response was incorrect, got: %v but wanted %v", Observed, Expected)
    }
}

func TestFindIntersection_UnequalLengthLists_1(t * testing.T) {
    Shared := NewLinkedList([] int {8, 11})
    M := NewLinkedList([] int {3})
    N := NewLinkedList([] int {4, 5, 6})
    M.Append(&Shared)
    N.Append(&Shared)

    Observed := FindIntersection( &M, &N)
    Expected := Shared.head

    if Observed != Expected {
        t.Errorf("Response was incorrect, got: %v but wanted %v", Observed, Expected)
    }
}

func TestFindIntersection_UnequalLengthLists_2(t * testing.T) {
    Shared := NewLinkedList([] int {8, 11})
    M := NewLinkedList([] int {})
    N := NewLinkedList([] int {4, 5, 6})
    M.Append(&Shared)
    N.Append(&Shared)

    Observed := FindIntersection( & M, & N)
    Expected := Shared.head

        if Observed != Expected {
        t.Errorf("Response was incorrect, got: %v but wanted %v", Observed, Expected)
    }
}
