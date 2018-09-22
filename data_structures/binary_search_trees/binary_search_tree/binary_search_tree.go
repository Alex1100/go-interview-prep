package binary_search_tree

import (
	"errors"
	"math"
)

type OrderType struct {
	Order_Type string
}

type DataValue struct {
	Value int
}

type BST struct {
	Data   *DataValue
	Left   *BST
	Right  *BST
	Parent *BST
}

func InitEmptyBST() *BST {
	return &BST{
		Data:   nil,
		Left:   nil,
		Right:  nil,
		Parent: nil,
	}
}

func (bst *BST) InitBST(data int) *BST {
	return &BST{
		Data:   &DataValue{Value: data},
		Left:   nil,
		Right:  nil,
		Parent: bst,
	}
}

func CheckOrderType(order_type string) (string, error) {
	if order_type != "pre_order" && order_type != "in_order" && order_type != "post_order" {
		return "", errors.New("Must pass in allowed order_type Enum")
	}

	return order_type, nil
}

func (bst *BST) AddNode(data int) bool {
	if bst.Data == nil {
		bst.Data = &DataValue{Value: data}
		return true
	}

	if data < bst.Data.Value {
		if bst.Left != nil {
			bst.Left.AddNode(data)
		} else {
			bst.Left = bst.InitBST(data)
			return true
		}
	} else if data > bst.Data.Value {
		if bst.Right != nil {
			bst.Right.AddNode(data)
		} else {
			bst.Right = bst.InitBST(data)
			return true
		}
	}

	return false
}

func DFSTraverse(node *BST, ordering string, result []int) []int {
	if ordering == "pre_order" {
		result = append(result, node.Data.Value)

		if node.Left != nil {
			result = DFSTraverse(node.Left, ordering, result)
		}

		if node.Right != nil {
			result = DFSTraverse(node.Right, ordering, result)
		}
	} else if ordering == "in_order" {
		if node.Left != nil {
			result = DFSTraverse(node.Left, ordering, result)
		}

		result = append(result, node.Data.Value)

		if node.Right != nil {
			result = DFSTraverse(node.Right, ordering, result)
		}
	} else if ordering == "post_order" {
		if node.Left != nil {
			result = DFSTraverse(node.Left, ordering, result)
		}

		if node.Right != nil {
			result = DFSTraverse(node.Right, ordering, result)
		}

		result = append(result, node.Data.Value)
	}
	return result
}

func (bst *BST) DepthFirstSearch(order_type string) ([]int, error) {
	order_wanted, err := CheckOrderType(order_type)
	result := make([]int, 0)

	if err != nil {
		return nil, err
	}

	result = DFSTraverse(bst, order_wanted, result)
	return result, nil
}

func (bst *BST) BreadthFirstSearch() []int {
	result := make([]int, 0)

	current := make([]*BST, 0)
	current = append(current, bst)

	for len(current) > 0 {
		next := make([]*BST, 0)

		for i := 0; i < len(current); i++ {
			current_node := current[i]
			result = append(result, current_node.Data.Value)

			if current_node.Left != nil {
				next = append(next, current_node.Left)
			}

			if current_node.Right != nil {
				next = append(next, current_node.Right)
			}
		}

		current = next
	}

	return result
}

func (bst *BST) IsValidBST(min, max float64) bool {
	if bst.Left == nil {
		return true && bst.Right.IsValidBSTUtil(bst.Right, float64(bst.Data.Value)+float64(1), max)
	}

	if bst.Right == nil {
		return true && bst.Left.IsValidBSTUtil(bst.Left, min, float64(bst.Data.Value)-float64(1))
	}

	return (bst.Left.IsValidBSTUtil(bst.Left, min, float64(bst.Data.Value)-float64(1)) && bst.Right.IsValidBSTUtil(bst.Right, float64(bst.Data.Value)+float64(1), max))
}

func (bst *BST) IsValidBSTUtil(node *BST, min, max float64) bool {
	left := true
	right := true

	if node.Right == nil && node.Left != nil {
		right = true && node.Left.IsValidBSTUtil(node.Left, min, float64(node.Data.Value)-float64(1))
	}

	if node.Left == nil && node.Right != nil {
		left = true && node.Right.IsValidBSTUtil(node.Right, float64(node.Data.Value)+float64(1), max)
	}

	if float64(node.Data.Value) < min || float64(node.Data.Value) > max {
		return false
	}

	if node.Right == nil || node.Left == nil {
		return left && right
	} else if node.Left != nil && node.Right != nil {
		return (node.Left.IsValidBSTUtil(node.Left, min, float64(node.Data.Value)-float64(1)) && node.Right.IsValidBSTUtil(node.Right, float64(node.Data.Value)+float64(1), max))
	}

	return true
}

func (bst *BST) FindInOrderSuccessor(val int) *BST {
	var result *BST
	if bst.Data == nil {
		return bst
	}

	node := bst

	for node.Data.Value != val {
		if val < node.Data.Value {
			node = node.Left
		} else if val > node.Data.Value {
			node = node.Right
		}
	}

	if node.Right != nil {
		if node.Right.Left != nil && node.Right.Left.Left == nil {
			return node.Right.Left
		} else if node.Right.Left != nil && node.Right.Left.Left != nil {
			result = node.Right.Left
			for result.Left != nil {
				result = result.Left
			}
		}
	} else {
		root_input_node := node
		target_input_node := node.Parent
		for target_input_node.Data.Value <= root_input_node.Data.Value {
			if target_input_node.Parent != nil {
				target_input_node = target_input_node.Parent
			} else {
				return node
			}
		}

		return target_input_node
	}

	return result
}

func (bst *BST) GetMinDepth() int {
	if bst.Right == nil && bst.Left == nil {
		return 1
	}

	if bst.Left == nil && bst.Right != nil {
		return bst.Right.GetMinDepth() + 1
	}

	if bst.Right == nil && bst.Left != nil {
		return bst.Left.GetMinDepth() + 1
	}

	left := bst
	right := bst

	if bst.Left != nil {
		left = bst.Left
	}

	if bst.Right != nil {
		right = bst.Right
	}

	return int(math.Min(float64(left.GetMinDepth()), float64(right.GetMinDepth()))) + 1
}

func MaxSum(node *BST, result float64) float64 {
	if node != nil {
		result = float64(MaxSum(node.Left, result))
		left_sum := math.Max(float64(0), result)
		result = float64(MaxSum(node.Right, result))
		right_sum := math.Max(float64(0), result)
		result = math.Max(result, left_sum+float64(node.Data.Value)+right_sum)
		result = math.Max(left_sum, right_sum) + float64(node.Data.Value)
		return result
	}

	return float64(0)
}

func (bst *BST) GetMaxPathSum() int {
	return int(MaxSum(bst, math.Inf(-1)))
}

func (bst *BST) ContainsNode(val int) bool {
	current_node := bst

	for current_node != nil {
		if current_node.Data.Value == val {
			return true
		} else if current_node.Data.Value >= val {
			if current_node.Data.Value == val {
				return true
			}

			if current_node.Left == nil {
				return false
			} else {
				current_node = current_node.Left
			}
		} else if current_node.Data.Value < val {
			if current_node.Right == nil {
				return false
			} else {
				current_node = current_node.Right
			}
		}
	}

	return false
}

func (bst *BST) FindNode(val int) (*BST, error) {
	current_node := bst

	for current_node != nil {
		if current_node.Data.Value == val {
			return current_node, nil
		} else if current_node.Data.Value >= val {
			if current_node.Left != nil {
				current_node = current_node.Left
			} else {
				return bst, (errors.New("Node does not exist"))
			}
		} else if current_node.Data.Value < val {
			if current_node.Right != nil {
				current_node = current_node.Right
			} else {
				return bst, (errors.New("Node does not exist"))
			}
		}
	}

	return bst, (errors.New("Node does not exist"))
}

func (bst *BST) FindMax() *BST {
	right_leaf := bst.Right
	if right_leaf != nil {
		for right_leaf.Right != nil {
			right_leaf = right_leaf.Right
		}

		return right_leaf
	}

	return right_leaf
}

func (bst *BST) FindMin() *BST {
	left_leaf := bst.Left
	if left_leaf != nil {
		for left_leaf.Left != nil {
			left_leaf = left_leaf.Left
		}

		return left_leaf
	}

	return left_leaf
}
