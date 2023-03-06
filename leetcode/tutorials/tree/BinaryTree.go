package tree

// 总结
// 层遍历： 队列
// 前中后序遍历： 栈
// 中序： 一直压入左子树节点，直到null， 弹出栈顶，出栈顺序就是中序遍历结果，然后转向右边节点，如果右节点为空，当前子树完成，否则发现新子树。
// 前序： 根入栈，栈不为空，弹出栈顶，然后先右节点后左节点入栈， 直到栈为空。 出栈顺序就是前序遍历结果。
// 后序： 根入栈， 栈不为空，弹出栈顶， 然后先左后右节点入栈， 直到栈为空。 出栈的逆序就是后序遍历的结果。
// 层序： 根入队列， 队列不为空， 出队列，根据队列的当前长度，移除所有元素，移除的元素的子节点需要再做入队列的操作，如此循环，直到队列为空。
// 性质： 二叉树搜索树的中序遍历是有序的。
