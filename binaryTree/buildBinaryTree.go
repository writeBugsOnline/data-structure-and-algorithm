package main

import (
	"strings"
)

// 递归方法：
//
// 对于任意一棵树而言，根节点总是前序遍历的第一个节点。
// 前序遍历：[根节点，[左子树所有的节点]，[右子树所有的节点]]
// 中序遍历：[[左子树所有的节点]，根节点，[右子树所有的节点]]
type Node struct{
	left    *Node
	right   *Node
	data    string
}

func buildTree(preorder []string, inorder []string) *Node {
	//若preorder的长度为0，说明这是棵空的二叉树
	if len(preorder)==0{
		return nil
	}
	//根据二叉树前序遍历的第一个节点是根节点的特性，可以定义一个只有根节点的二叉树;然后再递归遍历左右子树来填充这棵二叉树。
	root:=&Node{
		nil,
		nil,
		preorder[0],
	}
	//遍历中序序列中的节点，找到根节点在中序中位置
	i:=0
	for ;i<len(inorder);i++{
		if inorder[i]==preorder[0]{
			break
		}
	}

	//根据中序遍历的特点：[[左子树]，根节点，[右子树]]来确定左右子树的起始位置和结束位置
	//遍历左子树:
	//前序参数中：第一个是根节点，所以起始位置从1开始，结束位置是中序序列i前面节点的个数，+1是因为slice截取原则：左闭右开；中序参数中：起始位置就是从头开始，到i位置结束。
	root.left=buildTree(preorder[1:len(inorder[:i])+1],inorder[:i])
	//遍历右子树
	//前序参数是：从[中序序列i前面节点的个数+1]开始，到最后结束；中序参数是：从i位置的后一个开始到最后结束
	root.right=buildTree(preorder[len(inorder[:i])+1:],inorder[i+1:])

	return root
}

func NewNode(data string) *Node {
	return &Node{
		left: nil,
		right: nil,
		data: data,
	}
}

func BT2(pre []string, mid []string) *Node {
	if len(pre)==0{
		return nil
	}

	root := NewNode(pre[0])

	i :=0

	for ;i<len(mid);i++{
		if pre[0] == mid[i]{
			break
		}
	}

	root.right=BT2(pre[1:len(mid[:i])+1],mid[:i])
	root.left=BT2(pre[len(mid[:i])+1:],mid[i+1:])

	return root
}



func TestBT(){
	pre :=strings.Split("A B D H E I C F J K G"," ")
	mid :=strings.Split("D H B E I A J F K C G"," ")

	tree :=BT2(pre,mid)

	PreTravel(tree)
	println()
	MidTravel(tree)
}

func main() {
	TestBT()
}

func PreTravel(node  *Node)  {
	if node == nil{
		return
	}
	print(node.data)
	PreTravel(node.left)
	PreTravel(node.right)
}

func MidTravel(node *Node)  {
	if node == nil{
		return
	}

	PreTravel(node.left)
	print(node.data)
	PreTravel(node.right)

}