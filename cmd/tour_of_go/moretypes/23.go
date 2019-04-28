/*
 * @author: Jinghua.Yao
 * @email : staugusto91@gmail.com
 */

package main

import (
	"strings"

	"golang.org/x/tour/wc"
)

func WordCount(s string) map[string]int {
	mp := make(map[string]int, 0)
	for _, str := range strings.Fields(s) {
		mp[str] += 1
	}
	return mp
}

func main() {
	wc.Test(WordCount)
}
