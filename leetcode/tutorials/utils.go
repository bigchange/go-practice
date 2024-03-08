/*
 * @Author: Jerry You
 * @CreatedDate: 2023/3/15
 * @Last Modified by: lolaliva
 * @Last Modified time: 2023/3/15 16:59
 */

package tutorials

import "fmt"

func PrintArr(arr []int) {
	for _, r := range arr {
		fmt.Print(fmt.Sprintf("%v ", r))
	}
	fmt.Println("")
}
