/*
 * @Author: Jerry You
 * @CreatedDate: 2023/3/16
 * @Last Modified by: lolaliva
 * @Last Modified time: 2023/3/16 15:37
 */

package exercism

import "errors"

func Distance(a, b string) (int, error) {
	la := len(a)
	lb := len(b)
	err := errors.New("wrong parameters")
	if la != lb {
		return 0, err
	}
	c := 0
	for i, item := range a {
		if b[i] != uint8(item) {
			c++
		}
	}
	return c, nil
}