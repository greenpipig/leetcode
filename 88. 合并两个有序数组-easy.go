package main

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 && n > 0 {
		copy(nums1, nums2)
		return
	}
	if n == 0 {
		return
	}
	nilsN := n - 1
	nilsM := m - 1
	small := 0
	if nums2[0] > nums1[0] {
		small = nums1[0]
	} else {
		small = nums2[0]
	}
	for i := m + n - 1; i > 0; i-- {
		if nilsN == -1 || nilsM == -1 {
			break
		}
		if nums1[nilsM] <= nums2[nilsN] {
			nums1[i] = nums2[nilsN]
			nilsN--
		} else {
			nums1[i] = nums1[nilsM]
			nums1[nilsM] = small
			nilsM--
		}
	}
	if nilsN > 0 {
		for i := nilsN; i > 0; i-- {
			nums1[i] = nums2[i]
		}
	}
}

func main() {
	merge([]int{1, 2, 3, 0, 0}, 3, []int{1, 2}, 2)
}
