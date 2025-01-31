package main

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	merge := make([]int, len(nums1)+len(nums2))
	i1 := 0
	i2 := 0

	for index := 0; index < len(merge); index++ {
		if i1 < len(nums1) && i2 < len(nums2) {
			if nums1[i1] < nums2[i2] {
				merge[index] = nums1[i1]
				i1++
			} else {
				merge[index] = nums2[i2]
				i2++
			}
		} else if i1 < len(nums1) {
			merge[index] = nums1[i1]
			i1++
		} else {
			merge[index] = nums2[i2]
			i2++
		}
	}

	if len(merge)%2 == 1 {
		return float64(merge[len(merge)/2])
	} else {
		return (float64(merge[len(merge)/2]) + float64(merge[len(merge)/2-1])) / 2.0
	}
}
