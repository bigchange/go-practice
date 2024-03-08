package leetcode


func minSteps(n int) int {
	current := 1
	if current == n {
		return  0
	}
	//  first step must be 'copy'
	step := 1
	// copy the default current 'A'
	copied := 1
	for {
	 // AFTER PASTER
	 newStatus := current + copied
	 if newStatus == n {
		 return  step + 1
	 }
	 // newStatus： 当前A的数量，加上最近一次copy操作复制的数量
	 // 如果n能被 newStatus整除
	 // 更新当前A的数量，同时更新最近一次copy的数量
	 // 只有在能够被n整除的时候更新copy数量，才可能刚好到达n，否则就会超出n的大小
	 // 例如： current = 1， copied = 1
	 // newStatus = 2,
	 // 情况一： 当 n % newStatus == 0的时候
	 // 可以有以下两种选择：
	 // 1. 一直只进行paste就能到数量n，但是可能不是最少step数量
	 // 2. 考虑 n % (newStatus * 2) == 0, 那么我先paste一次，然后在copy一次，到数量n的时候，step数只会和选择第一种情况下相等或者甚至更少
	 // 其他情况：
	 //  如果 newStatus = 2  不能被n整除，那么就可以使用上次copy的数量继续在paste一次， 使得newStatus = 3
	 // newStatus = 4, newStatus = 5
	 //  如此循环判断
	 //  在每个newStatus的状态下进行情况一下的判断逻辑即可得到最终结果
	 if n % newStatus == 0 {
		 current = newStatus
		 copied = current
		 // paste一次，copy一次
		 step += 2
	 } else {
		 // 只更新当前A的数量
		 // 最近一次copy的数量保持不变
		 current = newStatus
		 // once paste
		 step++
	 }
	}
}
