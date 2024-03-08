package leetcode

import (
	"fmt"
	"sort"
)

func canPartitionKSubsets(nums []int, k int) bool {
	sum := 0
	for _, i := range nums {
		sum += i
	}
	if sum%k != 0 {
		return false
	}
	size := len(nums)
	// 倒序排列数据
	sort.Slice(nums, func(i, j int) bool{
		return nums[i] > nums[j]
	})
	total := sum / k
	fmt.Println("total, sum", total, sum)

	// 需要进一步验证，每个集合的元素是否都能满足总和为total
	visited := make(map[int]bool)
	// 深度搜索，
	// 参数：
	// index：当前位置
	// eachSum：每个集合目前的总和，
	// count：当前已经得到的满足条件的集合个数
	//
	var dfs func(int, int,  int) bool
	dfs = func(index int, eachSum int, count int) bool {
		fmt.Println(index, eachSum, count)
		//  得到了K个集合
		if count == k {
			return true
		}
		// 得到满足条件的一个集合
		if  eachSum == total {
			// 继续递归，获取下一个集合：count + 1
			return dfs(size - 1, 0, count + 1)
		}
		// 从index开始，往右搜索可能的元素
		for i:= index ; i >= 0; i-- {
			// 如果被访问，或者和超过total，跳过并继续往右
			if  visited[i] || eachSum + nums[i] > total {
				continue
			}
			// 设置该元素为被访问状态
			visited[i] = true
			// 每次只需要往当前index的右边搜索
			// 继续递归下一个元素，count参数不变
			if dfs(i - 1, eachSum + nums[i], count) {
				return true
			}
			// 如果选择当前的不能满足，将其设置为未被访问
			visited[i] = false
			// 重点： 如果不添加此处，运行将超时
			// 如果剩余集合还有未被访问的元素，但是求和一直都无法等于total，最终得不到K个集合
			// 返回false
			if eachSum == 0 {
				return false
			}
		}
		return false
	}
	return dfs(size - 1, 0, 0)
}
