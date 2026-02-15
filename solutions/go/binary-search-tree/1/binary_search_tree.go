package binarysearchtree

type BinarySearchTree struct {
	left  *BinarySearchTree
	data  int
	right *BinarySearchTree
}

// NewBst creates and returns a new BinarySearchTree.
func NewBst(i int) *BinarySearchTree {
	return &BinarySearchTree{nil, i, nil}
}

// Insert inserts an int into the BinarySearchTree.
// Inserts happen based on the rules of a binary search tree
func (bst *BinarySearchTree) Insert(i int) {
    if i <= bst.data {
        // Go left (including duplicates)
        if bst.left == nil {
            bst.left = NewBst(i)
        } else {
            bst.left.Insert(i) // RECURSIVELY insert into left subtree
        }
    } else {
        // Go right
        if bst.right == nil {
            bst.right = NewBst(i)
        } else {
            bst.right.Insert(i) // RECURSIVELY insert into right subtree
        }
    }
}

// SortedData returns the ordered contents of BinarySearchTree as an []int.
func (bst *BinarySearchTree) SortedData() []int {
    result := []int{}
    
    if bst.left != nil {
        result = append(result, bst.left.SortedData()...)
    }
    
    result = append(result, bst.data)
    
    if bst.right != nil {
        result = append(result, bst.right.SortedData()...)
    }
    
    return result
}