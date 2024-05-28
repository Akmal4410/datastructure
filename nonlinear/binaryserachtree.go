package nonlinear

import (
	"fmt"
	"os"

	"github.com/akmal4410/datastructure/utils"
)

func binarySerchTreeFunc() {
	var choice, temp int
	bst := binarySerchTree{}
TryBST:
	fmt.Println("Choose Binary Search Tree :\n1.InOrder \n2.PreOrder \n3.PostOrder \n4.Insert \n5.Contains \n6.Delete \n7.Go Back \n8.Exit")
	choice = utils.ScanValue(choice)
	switch choice {
	case 1:
		bst.inOrder()
		goto TryBST
	case 2:
		bst.preOrder()
		goto TryBST
	case 3:
		bst.postOrder()
		goto TryBST
	case 4:
		fmt.Println("Enter the value to insert")
		temp = utils.ScanValue(temp)
		if temp == -1 {
			goto TryBST
		}
		bst.insert(temp)
		goto TryBST
	case 5:
		fmt.Println("Enter the value to check if Exists")
		temp = utils.ScanValue(temp)
		if temp == -1 {
			goto TryBST
		}
		exits := bst.contains(temp)
		fmt.Printf("%d exists : %t\n", temp, exits)
		goto TryBST
	case 6:
		fmt.Println("Enter the value to delete")
		temp = utils.ScanValue(temp)
		if temp == -1 {
			goto TryBST
		}
		bst.remove(temp)
		goto TryBST
	case 7:
		break
	case -1, 8:
		os.Exit(0)
	default:
		goto TryBST
	}
}

type tNode struct {
	data                int
	leftNode, rightNode *tNode
}

type binarySerchTree struct {
	root *tNode
}

func (tree *binarySerchTree) insert(data int) {
	newNode := &tNode{data: data}
	currentNode := tree.root
	if currentNode == nil {
		tree.root = newNode
		return
	}
	for {
		if data < currentNode.data {
			if currentNode.leftNode == nil {
				currentNode.leftNode = newNode
				break
			} else {
				currentNode = currentNode.leftNode
			}
		} else {
			if currentNode.rightNode == nil {
				currentNode.rightNode = newNode
				break
			} else {
				currentNode = currentNode.rightNode
			}
		}
	}
}

func (tree *binarySerchTree) contains(data int) bool {
	currentNode := tree.root
	for currentNode != nil {
		if data < currentNode.data {
			currentNode = currentNode.leftNode
		} else if data > currentNode.data {
			currentNode = currentNode.rightNode
		} else {
			return true
		}
	}
	return false
}

func (tree *binarySerchTree) remove(data int) {
	tree.removeHelper(data, tree.root, nil)

}

func (tree *binarySerchTree) removeHelper(data int, currentNode *tNode, parentNode *tNode) {
	for currentNode != nil {
		if data < currentNode.data {
			parentNode = currentNode
			currentNode = currentNode.leftNode
		} else if data > currentNode.data {
			parentNode = currentNode
			currentNode = currentNode.rightNode
		} else {
			if currentNode.leftNode != nil && currentNode.rightNode != nil {
				currentNode.data = tree.getMinimumValue(currentNode.rightNode)
				tree.removeHelper(currentNode.data, currentNode.rightNode, currentNode)
			} else {
				if parentNode == nil {
					if currentNode.rightNode == nil {
						tree.root = currentNode.leftNode
					} else {
						tree.root = currentNode.rightNode
					}
				} else {
					if parentNode.leftNode == currentNode {
						if currentNode.rightNode == nil {
							parentNode.leftNode = currentNode.leftNode
						} else {
							parentNode.rightNode = currentNode.rightNode
						}
					} else {
						if currentNode.rightNode == nil {
							parentNode.rightNode = currentNode.leftNode
						} else {
							parentNode.rightNode = currentNode.rightNode
						}
					}
				}
			}
			break
		}
	}
}

func (tree *binarySerchTree) getMinimumValue(currentNode *tNode) int {
	if currentNode.leftNode == nil {
		return currentNode.data
	} else {
		return tree.getMinimumValue(currentNode.leftNode)
	}
}

func (tree *binarySerchTree) inOrder() {
	fmt.Println("IN ORDER==========>")
	tree.inOrderHelper(tree.root)
	fmt.Println()
}

func (tree *binarySerchTree) inOrderHelper(currentNode *tNode) {
	if currentNode != nil {
		tree.inOrderHelper(currentNode.leftNode)
		fmt.Printf("%d ", currentNode.data)
		tree.inOrderHelper(currentNode.rightNode)
	}
}

func (tree *binarySerchTree) preOrder() {
	fmt.Println("PR EORDER==========>")
	tree.preOrderHelper(tree.root)
	fmt.Println()
}

func (tree *binarySerchTree) preOrderHelper(currentNode *tNode) {
	if currentNode != nil {
		fmt.Printf("%d ", currentNode.data)
		tree.preOrderHelper(currentNode.leftNode)
		tree.preOrderHelper(currentNode.rightNode)
	}
}

func (tree *binarySerchTree) postOrder() {
	fmt.Println("POST ORDER==========>")
	tree.postOrderHelper(tree.root)
	fmt.Println()
}

func (tree *binarySerchTree) postOrderHelper(currentNode *tNode) {
	if currentNode != nil {
		tree.postOrderHelper(currentNode.leftNode)
		tree.postOrderHelper(currentNode.rightNode)
		fmt.Printf("%d ", currentNode.data)
	}
}
