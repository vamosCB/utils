package utils

import "strconv"

//RemoveTargetFromIntSlice 切片中移除特定元素
func RemoveTargetFromIntSlice(src []int, target int) []int {
	if len(src) == 0 {
		return src
	}
	for k, v := range src {
		if v == target {
			if k == len(src)-1 {
				src = src[:k]
			} else {
				src = append(src[:k], src[k+1:]...)
			}
			return src
		}
	}
	return src
}

//IntSliceToStringSlice int切片转string切片
func IntSliceToStringSlice(src []int) (dst []string) {
	for _, v := range src {
		dst = append(dst, strconv.Itoa(v))
	}
	return dst
}

//StringSliceContainTarget 字符串切片中是否存在某一字符串
func StringSliceContainTarget(src []string, target string) bool {
	for _, v := range src {
		if v == target {
			return true
		}
	}
	return false
}
