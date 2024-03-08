package leetcode

// 主入口
func canBeEqual(target []int, arr []int) bool {
 for i:= 0; i < len(target); i++ {
 	  // 按target的元素顺序依次寻找对应在arr中的元素
 		pos := findTargetItem(arr, i, target[i])
 		// 如果没有找到那么直接返回false
 		if pos == -1 {
 			return false
		}
		// 交换[i,pos]之间子数组的位置
		// arr i位置元素复原
		changePos(arr, i, pos)
	}
	// 能走完循环，那么就是可以复原的
	return true
}

func changePos(arr []int, start, end int) {
	for start < end {
		arr[start], arr[end] = arr[end], arr[start]
		start++
		end--
	}
}

func findTargetItem(arr []int, start, item int) int {
	for i:= start; i < len(arr); i++ {
		if  arr[i] == item {
			return i
		}
	}
	return -1
}
