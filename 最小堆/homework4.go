package main

import "fmt"

type Node struct {
	Value int
}

// 用于构建结构体切片为最小堆，需要调用down函数
func Init(nodes []Node) {
	// 从第 n/2（向下取整）- 1 个元素开始考虑，因为这个元素之前的元素才有子节点
	for i := len(nodes)/2 - 1; i >= 0; i -- {
		down(nodes,i,len(nodes))
	}
}

// 需要down（下沉）的元素在切片中的索引为i，n为heap的长度，将该元素下沉到该元素对应的子树合适的位置，从而满足该子树为最小堆的要求
func down(nodes []Node, i, n int) {
	parent := i
	left_child := 2 * i + 1
	right_child := 2 * i + 2
	temp := nodes[parent].Value
	for {
		if left_child < n {
			if right_child < n {
				if nodes[left_child].Value < nodes[right_child].Value {
					if nodes[left_child].Value < temp {
						nodes[parent].Value = nodes[left_child].Value
						nodes[left_child].Value = temp
						parent = left_child
						left_child = left_child*2+1
						right_child = left_child*2+2
					} else {
						break
					}
				} else {
					if nodes[right_child].Value < temp {
						nodes[parent].Value = nodes[right_child].Value
						nodes[right_child].Value = temp
						parent = right_child
						left_child = right_child*2+1
						right_child = right_child*2+2
					} else {
						break
					}
				}
			} else {
				if nodes[left_child].Value < temp {
					nodes[parent].Value = nodes[left_child].Value
					nodes[left_child].Value = temp
					parent = left_child
					left_child = left_child*2+1
					right_child = left_child*2+2
				} else {
					break
				}
			}
		} else {
			break
		}
	}
}

// 用于保证插入新元素(j为元素的索引,切片末尾插入，堆底插入)的结构体切片之后仍然是一个最小堆
func up(nodes []Node, j int) {
	child := j
	parent := (j-1)/2
	for{
		// up到堆顶
		if child == 0 {
			break
		}
		// 满足parent节点不大于child节点
		if nodes[parent].Value <= nodes[child].Value {
			break
		}
		temp := nodes[child].Value
		nodes[child].Value = nodes[parent].Value
		nodes[parent].Value = temp
		child = parent
		parent = (parent-1)/2
	}
}

// 弹出最小元素，并保证弹出后的结构体切片仍然是一个最小堆，第一个返回值是弹出的节点的信息，第二个参数是Pop操作后得到的新的结构体切片
func Pop(nodes []Node) (Node, []Node) {
	min := nodes[0]
	nodes = nodes[1:len(nodes)]
	Init(nodes)
	return min, nodes
}

// 保证插入新元素时，结构体切片仍然是一个最小堆，需要调用up函数
func Push(node Node, nodes []Node) []Node {
	nodes = append(nodes,node)
	// 只需要up新插入的最后一个元素即可
	up(nodes,len(nodes)-1)
	return nodes
}

// 移除切片中指定索引的元素，保证移除后结构体切片仍然是一个最小堆
func Remove(nodes []Node, node Node) []Node {
	for i := 0; i < len(nodes); i ++ {
		if nodes[i].Value == node.Value {
			nodes[i].Value = nodes[len(nodes)-1].Value
			nodes = nodes[0:len(nodes)-1]
			Init(nodes)
			break
		}
	}
	return nodes
}

func main() {
	// 构造测试数据
	nodes := []Node {
		Node{3},
		Node{13},
		Node{17},
		Node{2},
		Node{5},
		Node{8},
		Node{22},
		Node{1},
		Node{65},
		Node{50},
		Node{77},
		Node{99},
	}
	// test Init
	for _, element := range nodes {
		fmt.Printf("%d ", element.Value)
	}
	fmt.Printf("\n\n")
	Init(nodes)
	for _, element := range nodes {
		fmt.Printf("%d ", element.Value)
	}
	fmt.Printf("\n\n")
	// test Push
	push_node := Node{0}
	nodes = Push(push_node,nodes)
	for _, element := range nodes {
		fmt.Printf("%d ", element.Value)
	}
	fmt.Printf("\n\n")
	// test Pop
	node_min,nodes := Pop(nodes)
	fmt.Printf("%d\n",node_min.Value)
	for _, element := range nodes {
		fmt.Printf("%d ", element.Value)
	}
	fmt.Printf("\n\n")
	// test Remove
	remove_node := Node{5}
	nodes = Remove(nodes,remove_node)
	for _, element := range nodes {
		fmt.Printf("%d ", element.Value)
	}
	fmt.Printf("\n\n")
}