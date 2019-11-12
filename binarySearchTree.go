/**
* @FileName   : binarySearchTree
* @Author     : Chen Wang
* @Version    : Go1.13.1 、Windows or Linux
* @Description: 用 Go 实现等价二叉搜索树(二叉排序树,二叉查找树)，用 Go 并发和信道阻塞打印二叉树，检查两个二叉树是否保存了相同序列
* @Time       : 2019/11/11 9:12
* @Software   : GoLand
* @Contact    : 1416825008@qq.com
* @Blog       : https://github.com/JackMin1314/GoLang_Study
* 代 码 仅 限 学 习 使 用，严 禁 商 业 用 途，转 载 请 注 明 出 处
 */

package main

import (
	"fmt"
	"sort"
)

/* Binary Search Tree
它或者是一棵空树，或者是具有下列性质的二叉树：
若它的左子树不空，则左子树上所有结点的值均小于它的根结点的值；
若它的右子树不空，则右子树上所有结点的值均大于它的根结点的值；
它的左、右子树也分别为二叉排序树。
*/
type BSTree struct {
	left  *BSTree
	right *BSTree
	value int
}

// 切片属于引用类型
type arryBSTree []BSTree

// 这里引入两个全局切片指针为了后面checkTrees()函数使用，保存每次生成树的数组地址
var checkTree1, checkTree2 *arryBSTree

// 要想使用sort包，需要先实现sort接口的内置三个方法
func (bsp arryBSTree) Len() int {
	return len(bsp)
}

func (bsarry arryBSTree) Less(i int, j int) bool {
	return bsarry[i].value < bsarry[j].value
}

func (bsp arryBSTree) Swap(i int, j int) {
	var temp = bsp[i].value
	bsp[i].value = bsp[j].value
	bsp[j].value = temp
}

/***输入一个序列实现二叉搜索树***/
// 先初始化value节点,并实现排序(增序)
func initValue(arrBST arryBSTree) (n int, err error) {
	fmt.Printf("请输入一个BSTree序列(长度为%d):", len(arrBST))
	if len(arrBST)%2 == 0 {
		return -1, fmt.Errorf("输入的长度%d,不是奇数", len(arrBST))
	}
	for i := 0; i < len(arrBST); i++ {
		arrBST[i].right = nil
		arrBST[i].left = nil
		n, err = fmt.Scanf("%d", &arrBST[i].value) // 如果输入的不是int则用0代替
	}

	sort.Sort(arrBST) // 实现了三个方法之后才可以用这个sort
	fmt.Println(arrBST)
	return n, err
}

// 根据节点的顺序构建树结构(非递归)，有序节点可以利用规律去构造二叉搜索树;如果节点没有规律需要递归判断构造.
func makeBSTree(arrTrees arryBSTree) *BSTree {
	fmt.Println("开始构建 BSTree 节点关系")
	//var bsTemp BSTree = BSTree{value: 0, left: nil, right: nil} 结构体初始化每个成员都要初始化，不能笼统写等于nil(会报错空指针)
	var tree *BSTree = nil
	// 自下而上构造,除根节点外每层只有两个节点(不是平衡二叉搜索树)
	for i := 0; i < len(arrTrees); {
		if i == 0 {
			arrTrees[i+1].left = &arrTrees[i]
			arrTrees[i+1].right = &arrTrees[i+2]
			tree = &arrTrees[i+1]
			i += 3
		} else {
			arrTrees[i].left = tree
			arrTrees[i].right = &arrTrees[i+1]
			tree = &arrTrees[i]
			i += 2
		}
	}
	fmt.Println("构建完毕...")
	return tree

}

/***遍历二叉搜索树按层次打印输出***/
// 1.（从0层根节点开始）当 tree.left != nil 自上往下输出
func showBST1() string {
	arrBST := make(arryBSTree, 63) // 树节点以奇数为例，创建含有63个节点的
	if (len(arrBST) % 2) == 0 {
		return fmt.Sprintf("err:初始化的树节点个数%d,不是奇数", len(arrBST))
	}
	_, err := initValue(arrBST) // 初始化和排序
	if err != nil {
		return fmt.Sprintf("error: %v,输入个数为%v", err, len(arrBST))
	}
	checkTree1 = &arrBST       // 赋给全局数组指针,为了后面checkTrees使用
	tree := makeBSTree(arrBST) // 非递归构建BST

	for i := 0; tree.left != nil; i++ {
		if i == 0 {
			fmt.Printf("第%d层，节点为%v\n", i, tree.value)
		} else {
			fmt.Printf("第%d层，节点为%v,%v\n", i, tree.left.value, tree.right.value)
			tree = tree.left
		}
	}
	return fmt.Sprintf("打印完毕...")
}

// 2. 借助通道来实现并发打印(goroutine)请注意这里ch的用法->阻塞main线.如果不阻塞，会导致main很快跑完，就结束了，可能bsprinter没来得及及执行
func bstprinter(tree *BSTree, ch chan BSTree) {
	for i := 0; tree.left != nil; i++ {
		if i == 0 {
			fmt.Printf("第%d层，节点为%v\n", i, tree.value)
			ch <- *tree
		} else {
			fmt.Printf("第%d层，节点为%v,%v\n", i, tree.left.value, tree.right.value)
			ch <- *tree
			tree = tree.left
		}
	}
	ch <- BSTree{value: 0, left: nil, right: nil} // 向ch中加数据，如果没有其他goroutine来取走这个数据，那么挂起foo, 直到main函数把0这个数据拿走

}
func showBST2() string {
	arrBST := make(arryBSTree, 63) // 树节点以奇数为例，创建含有63个节点
	if (len(arrBST) % 2) == 0 {
		return fmt.Sprintf("err:初始化的树节点个数%d,不是奇数", len(arrBST))
	}
	_, err := initValue(arrBST) // 初始化和排序
	if err != nil {
		return fmt.Sprintf("error: %v,输入个数为%v", err, len(arrBST))
	}
	checkTree2 = &arrBST // 赋给全局数组指针,为了后面checkTrees使用
	tree := makeBSTree(arrBST)
	// 创建一个缓冲容量64的管道channel
	ch := make(chan BSTree, 70) // ch存节点的value
	go bstprinter(tree, ch)
	<-ch // 从ch中取数据，如果ch中还没有放数据，挂起main线(main调用showBST2())，直到bstprinter()放数据为止

	return fmt.Sprintf("打印完毕...")
}

/***比较两个二叉搜索树是否相同(每层节点value是否对应)***/
// 这里需要说明下:前提是两个树都已经是二叉搜索树。
// 两个思路:一、直接按层遍历节点比较value，见showBST1；二、我们已经实现了从数组转变成一个二叉搜索树，因而转为数组问题处理更简单，见下
func checkTrees(arrTrees1 *arryBSTree, arrTrees2 *arryBSTree) (bool, error) {
	if arrTrees1.Len() != arrTrees2.Len() {
		return false, fmt.Errorf("error:两个树节点个数不一致")
	} else {
		for i := 0; i < arrTrees1.Len(); i++ {
			if (*arrTrees1)[i].value != (*arrTrees2)[i].value {
				return false, fmt.Errorf("error:节点不一致")
			}
		}
		return true, fmt.Errorf("两个树结构一致，是二叉搜索树")
	}
}
func main() {
	var msg = showBST1()
	fmt.Println(msg)
	var msg2 = showBST2()
	fmt.Println(msg2)
	isEqual, err := checkTrees(checkTree1, checkTree2) // 使用这个函数需要先执行上面两个showBST1、showBST2获取[]BSTree
	fmt.Printf("是否相等:%v; %v", isEqual, err)

}
