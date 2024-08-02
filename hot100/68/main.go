package main

import (
	"fmt"
	"math"
)

func main() {
	//[2,3,4,5,6,7]
	res := findMedianSortedArrays([]int{2, 3, 4, 5, 6, 7}, []int{1})
	fmt.Println(res)
}
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	len1, len2 := len(nums1), len(nums2)
	start1, start2 := -1, -1
	findTwo := (len1+len2)%2 == 0
	var mid int
	if findTwo {
		mid = (len1 + len2) / 2
	} else {
		mid = (len1+len2)/2 + 1
	}
	for start1 < len1-1 && start2 < len2-1 && mid != 1 {
		var v1, v2 int
		// 说明不够中位数
		if start1+mid/2 < len1 {
			v1 = start1 + mid/2
		} else {
			v1 = len1 - 1
		}
		if start2+mid/2 < len2 {
			v2 = start2 + mid/2
		} else {
			v2 = len2 - 1
		}
		if nums1[v1] <= nums2[v2] {
			mid -= v1 - start1 // 减去已经遍历的元素
			start1 = v1        // 移动指针
		} else {
			mid -= v2 - start2

			start2 = v2
		}
	}
	// nums1 遍历完了
	if start1 == len1-1 {
		if findTwo {
			return float64(nums2[mid+start2]+nums2[mid+1+start2]) / 2

		}
		return float64(nums2[mid+start2])
	}
	// nums2 遍历完了
	if start2 == len2-1 {
		if findTwo {
			return float64(nums1[mid+start1]+nums1[mid+1+start1]) / 2
		}
		return float64(nums1[mid+start1])
	}
	// 此时说明mid已经处于中间了，
	if findTwo {
		// 偶数  如果start1<start2 说明start1更小
		if nums1[start1+1] < nums2[start2+1] {
			// start1+1 == len1-1 说明start1已经遍历完了，无需在进行 start下一个 和 start2 作比较
			if start1+1 == len1-1 {
				return float64(nums1[1+start1]+nums2[start2+1]) / 2
			} else {

				return float64(
					float64(nums1[1+start1])+
						math.Min(float64(nums1[start1+2]), float64(nums2[start2+1])),
				) / 2
			}
		} else {
			if start2+1 == len2-1 {
				return float64(nums1[1+start1]+nums2[start2+1]) / 2
			} else {
				return float64(
					float64(nums2[1+start2])+
						math.Min(float64(nums2[start2+2]), float64(nums1[start1+1])),
				) / 2
			}

		}
	} else {
		return math.Min(float64(nums1[start1+1]), float64(nums2[start2+1]))
	}
}
