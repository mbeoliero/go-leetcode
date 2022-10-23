/*
 * @lc app=leetcode.cn id=297 lang=golang
 *
 * [297] 二叉树的序列化与反序列化
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

// type Codec struct {
// }

// func Constructor() Codec {
// 	return Codec{}
// }

// func (this *Codec) serialize(root *TreeNode) string {
// 	var queue []*TreeNode

// 	enqueue(&queue, root)

// 	var result string

// 	for len(queue) > 0 {
// 		dequeuedElement := dequeue(&queue)

// 		if dequeuedElement != nil {
// 			result = result + fmt.Sprintf("%v", dequeuedElement.Val) + ","
// 			enqueue(&queue, dequeuedElement.Left)
// 			enqueue(&queue, dequeuedElement.Right)
// 		} else {
// 			result = result + "nil" + ","
// 		}
// 	}

// 	return result[:len(result)-1]
// }

// func (this *Codec) deserialize(data string) *TreeNode {
// 	dataArr := strings.Split(data, ",")

// 	var result []*TreeNode

// 	for _, num := range dataArr {
// 		if num == "nil" {
// 			result = append(result, nil)
// 		} else {
// 			result = append(result, &TreeNode{
// 				Val: toInt(num),
// 			})
// 		}
// 	}

// 	begin := 0

// 	for _, node := range result {
// 		if node != nil {
// 			left := 2*begin + 1
// 			right := 2*begin + 2

// 			if left < len(result) {
// 				node.Left = result[left]
// 			}

// 			if right < len(result) {
// 				node.Right = result[right]
// 			}

// 			begin++
// 		}
// 	}

// 	return result[0]
// }

// func enqueue(queue *[]*TreeNode, newItem *TreeNode) {
// 	if queue == nil {
// 		panic("nil pointer")
// 	}

// 	*queue = append(*queue, newItem)
// }

// func dequeue(queue *[]*TreeNode) *TreeNode {
// 	if queue == nil {
// 		panic("nil pointer")
// 	}

// 	if len(*queue) == 0 {
// 		panic("empty queue")
// 	}

// 	dequeuedElement := (*queue)[0]

// 	*queue = (*queue)[1:len(*queue)]

// 	return dequeuedElement
// }

// func toInt(s string) int {
// 	value, err := strconv.Atoi(s)

// 	if err != nil {
// 		panic(err)
// 	}

// 	return value
// }

type Codec struct {
}

func Constructor() (_ Codec) {
	return
}

// Serializes a tree to a single string.
func (this *Codec) serialize(root *TreeNode) string {
	if root == nil {
		return ""
	}

	const empty = "null"
	res := []string{}
	stack := make([]*TreeNode, 0)
	stack = append(stack, root)

	for len(stack) > 0 {
		head := stack[0]
		stack = stack[1:]

		if head != nil {
			res = append(res, strconv.Itoa(head.Val))
			stack = append(stack, head.Left, head.Right)
		} else {
			res = append(res, empty)
		}
	}

	return strings.Join(res, ",")
}

// Deserializes your encoded data to tree.
func (this *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}

	res := strings.Split(data, ",")
	const empty = "null"

	result := make([]*TreeNode, 0)
	for _, num := range res {
		if num == empty {
			result = append(result, nil)
		} else {
			result = append(result, &TreeNode{
				Val: toInt(num),
			})
		}
	}

	begin := 0

	for _, node := range result {
		if node != nil {
			left := 2*begin + 1
			right := 2*begin + 2

			if left < len(result) {
				node.Left = result[left]
			}

			if right < len(result) {
				node.Right = result[right]
			}

			begin++
		}
	}

	return result[0]
}

func toInt(s string) int {
	value, err := strconv.Atoi(s)

	if err != nil {
		panic(err)
	}

	return value
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */
// @lc code=end

