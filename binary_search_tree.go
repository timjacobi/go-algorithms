package main

import (
	"fmt"
)

type Node struct {
	Parent *Node
	Left   *Node
	Right  *Node
	Key    int
}

func NewNode(key int, left, right *Node) *Node {
	node := Node{}
	node.Key = key
	node.Left = left
	node.Right = right

	if left != nil {
		left.Parent = &node
	}

	if right != nil {
		right.Parent = &node
	}

	return &node
}

// Pre-order tree walk walks the tree as follows
// If the node exists
// 1. Print the current node's value
// 2. Recurse on the current node's left child
// 3. Recurse on the current node's right child
func PreOrderTreeWalk(root *Node) {
	if root != nil {
		fmt.Printf("%v ", root.Key)
		PreOrderTreeWalk(root.Left)
		PreOrderTreeWalk(root.Right)
	}
}

// In-order tree walk walks the tree as follows
// If the node exists
// 1. Recurse on the current node's left child
// 2. Print the current node's value
// 3. Recurse on the current node's right child
func InOrderTreeWalk(root *Node) {
	if root != nil {
		InOrderTreeWalk(root.Left)
		fmt.Printf("%v ", root.Key)
		InOrderTreeWalk(root.Right)
	}
}

// Post-order tree walk walks the tree as follows
// If the node exists
// 1. Recurse on the current node's left child
// 2. Recurse on the current node's right child
// 3. Print the current node's value
func PostOrderTreeWalk(root *Node) {
	if root != nil {
		PostOrderTreeWalk(root.Left)
		PostOrderTreeWalk(root.Right)
		fmt.Printf("%v ", root.Key)
	}
}

// Non-recursive tree walk walks the tree as follows
// If the node exists
// 1. Add the node to an auxiliary stack
// 2. While the stack has elements
// 2.1 Pop the top element off the stack and print its value
// 2.2 Push the element's left child onto the stack if it exists
// 2.3 Push the element's right child onto the stack if it exists
func TreeWalkNonRecursive(root *Node) {
	if root == nil {
		return
	}
	stack := []*Node{root}

	for len(stack) > 0 {
		node := stack[0]
		stack = stack[1:]

		fmt.Printf("%v ", node.Key)

		if node.Left != nil {
			stack = append([]*Node{node.Left}, stack...)
		}
		if node.Right != nil {
			stack = append([]*Node{node.Right}, stack...)
		}
	}
}

// Calculates the height of a tree
func TreeHeight(root *Node) int {
	if root == nil {
		return -1
	}

	var maxChildHeight int

	heightLeft := TreeHeight(root.Left)
	heightRight := TreeHeight(root.Right)

	if heightLeft > heightRight {
		maxChildHeight = heightLeft
	} else {
		maxChildHeight = heightRight
	}

	return 1 + maxChildHeight
}

// Non-recursive pre-order tree walk with constant space
// walks the tree as follows
//
// Maintain pointers to the current and previously visited nodes
// Initially we set current to the root and then run the
// algorithm as long as current points to a node
//
// There are 3 scenarios when walking the tree:
//
// 1. The previous node was the current node's parent
// In this case we print the node's value.
// If the node has a left child it will be the next node to
// visit therefore we set
// - previous to current and current to current.left
// Otherwise we just go back up in the tree so we set
// - previous to current and current to current.parent
//
// 2. The previous node was the current node's left child
// Here we need to check if the node has a right child too
// If so we visit it and therefore set
// - previous to current and current to current.right
// If not we just go back up in the tree so we set
// - previous to current and current to current.parent
//
// 3. The previous node was the current node's right child
// In this case we're done with that subtree and can just
// walk back up so we set
// - previous to current and current to current.parent
//
// Eventually we will reach the root node and set
// current to nil which terminates the algorithm
func PreOrderTreeWalkNonRecursiveConstantSpace(node *Node) {
	var prev *Node

	for node != nil {
		if prev == node.Parent {
			prev = node
			fmt.Printf("%v ", node.Key)
			if node.Left == nil {
				node = node.Parent
			} else {
				node = node.Left
			}

		} else if prev == node.Left {
			prev = node
			if node.Right == nil {
				node = node.Parent
			} else {
				node = node.Right
			}
		} else {
			prev = node
			node = node.Parent
		}
	}
}

// Tree search searches the tree for a specific key value
// We traverse the tree constanly comparing the current
// key to the value we're searching for.
//
// If the key is greater we traverse the left child tree
// as it contains only values smaller than the current key
//
// If the key is smaller we traverse the right child tree
// as it contains only values greater than the current key
func TreeSearch(root *Node, key int) *Node {
	for root != nil && key != root.Key {
		if key < root.Key {
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return root
}

// To find the minimum of the tree we traverse down the
// left child tree until we hit a leaf
func TreeMinimum(root *Node) int {
	for root.Left != nil {
		root = root.Left
	}
	return root.Key
}

func TreeMinimumRecursive(root *Node) int {
	if root.Left == nil {
		return root.Key
	}

	return TreeMinimumRecursive(root.Left)
}

// To find the maximum of the tree we traverse down the
// right child tree until we hit a leaf
func TreeMaximum(root *Node) int {
	for root.Right != nil {
		root = root.Right
	}
	return root.Key
}

func TreeMaximumRecursive(root *Node) int {
	if root.Right == nil {
		return root.Key
	}

	return TreeMinimumRecursive(root.Right)
}

// We can use the BST attribute to find the
// successor of a node
// We need to cover the following 3 scenarios
//
// 1. The node has a right subtree
// In this case we just need to find the
// minimum of the right subtree
//
// 2. The node has no right subtree and
// is the left child of its parent
// In this case the successor is the parent
//
// 3. The node has no right subtree and
// is the right child of its parent
// In this case we need to walk up the until
// we find a node that is the left child of
// its parent and return the parent's key
func TreeSuccessor(root *Node) int {
	if root.Right != nil {
		return TreeMinimum(root.Right)
	}
	y := root.Parent
	for y != nil && root == y.Right {
		root = y
		y = y.Parent
	}
	return y.Key
}

// We can use the BST attribute to find the
// predecessor of a node
// We need to cover the following 3 scenarios
//
// 1. The node has a left subtree
// In this case we just need to find the
// maximum of the right subtree
//
// 2. The node has no left subtree and
// is the right child of its parent
// In this case the successor is the parent
//
// 3. The node has no right subtree and
// is the right child of its parent
// In this case we need to walk up the until
// we find a node that is the right child of
// its parent and return the parent's key
func TreePredecessor(root *Node) int {
	if root.Left != nil {
		return TreeMaximum(root.Left)
	}
	p := root.Parent
	for p != nil && root == p.Left {
		root = p
		p = p.Parent
	}
	return p.Key
}

func TreeInsert(root *Node, node *Node) {
	var y *Node
	x := root

	for x != nil {
		y = x
		if node.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	node.Parent = y

	if node.Key < y.Key {
		y.Left = node
	} else {
		y.Right = node
	}

}

func TreeInsertRecursive(root *Node, node *Node) *Node {
	if root == nil {
		return node
	}
	if node.Key < root.Key {
		root.Left = TreeInsertRecursive(root.Left, node)
		node.Parent = root.Left
	} else {
		root.Right = TreeInsertRecursive(root.Right, node)
		node.Parent = root.Right
	}

	return root
}

func TreeDelete(root *Node, key int) *Node {
	if root == nil {
		return nil
	} else if key < root.Key {
		root.Left = TreeDelete(root.Left, key)
	} else if key > root.Key {
		root.Right = TreeDelete(root.Right, key)
	} else {
		// We found the node to delete, let's handle the cases
		if root.Left == nil && root.Right == nil {
			// Case 1: Node has no children
			// => Delete node
			root = nil
		} else if root.Left == nil {
			// Case 2a: Node has right child
			// => Replace node with its right child
			root = root.Right
		} else if root.Right == nil {
			// Case 2b: Node has left child
			// => Replace node with its left child
			root = root.Left
		} else {
			// Case 3: Node has two children
			// => Find minimum key in right subtree,
			// set the node's key to the minimum key,
			// delete node with minimum key
			min := TreeMinimum(root.Right)
			root.Key = min
			root.Right = TreeDelete(root.Right, min)
		}
	}
	return root
}

func main() {
	// Set up the following tree
	//
	//            10
	//          /    \
	//         6     20
	//        / \   /  \
	//       5  8  15  22
	root := NewNode(10,
		NewNode(6,
			NewNode(5, nil, nil),
			NewNode(8, nil, nil)),
		NewNode(20,
			NewNode(15, nil, nil),
			NewNode(22, nil, nil)))

	fmt.Println("Performing pre-order tree walk")
	PreOrderTreeWalk(root)
	fmt.Print("\n \n")

	fmt.Println("Performing in-order tree walk")
	InOrderTreeWalk(root)
	fmt.Print("\n \n")

	fmt.Println("Performing post-order tree walk")
	PostOrderTreeWalk(root)
	fmt.Print("\n \n")

	fmt.Println("Performing tree walk using a stack")
	TreeWalkNonRecursive(root)
	fmt.Print("\n \n")

	fmt.Println("Performing pre-order tree walk using pointers \nto current and previous node")
	PreOrderTreeWalkNonRecursiveConstantSpace(root)
	fmt.Print("\n \n")

	fmt.Println("Searching node with key 20")
	fmt.Printf("%+v\n", *TreeSearch(root, 20))
	fmt.Println()

	fmt.Printf("Tree height is: %v \n", TreeHeight(root))
	fmt.Printf("Tree minimum is: %v \n", TreeMinimum(root))
	fmt.Printf("Tree maximum is: %v \n", TreeMaximum(root))
	fmt.Printf("Successor of key 8 is: %v \n", TreeSuccessor(TreeSearch(root, 8)))
	fmt.Printf("Predecessor of key 15 is: %v \n", TreePredecessor(TreeSearch(root, 15)))

	fmt.Print("Adding node with key 13, resulting tree is: ")
	TreeInsertRecursive(root, NewNode(13, nil, nil))
	InOrderTreeWalk(root)
	fmt.Println()


	fmt.Print("Deleting node with key 15, resulting tree is: ")
	TreeDelete(root, 15)
	InOrderTreeWalk(root)
}
