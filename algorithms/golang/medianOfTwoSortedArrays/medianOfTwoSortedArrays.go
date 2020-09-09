// Source :https://leetcode.com/problems/median-of-two-sorted-arrays/submissions/
// Author : gaohua
// Date	  : 2020-09-09

/***************************************************
* 
* 
* Given two sorted arrays nums1 and nums2 of size m and n respectively, return the median of the two sorted arrays.

* Follow up: The overall run time complexity should be O(log (m+n)).

* Example 1:
* Input: nums1 = [1,3], nums2 = [2]
* Output: 2.00000
* Explanation: merged array = [1,2,3] and median is 2.
*
* Example 2:
* Input: nums1 = [1,2], nums2 = [3,4]
* Output: 2.50000
* Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
*
* Example 3:
* Input: nums1 = [0,0], nums2 = [0,0]
* Output: 0.00000
*
* Example 4:
* Input: nums1 = [], nums2 = [1]
* Output: 1.00000
*
* Example 5:
* Input: nums1 = [2], nums2 = []
* Output: 2.00000
* 
* Constraints:
* nums1.length == m
* nums2.length == n
* 0 <= m <= 1000
* 0 <= n <= 1000
* 1 <= m + n <= 2000
* -106 <= nums1[i], nums2[i] <= 106
*
***************************************************/


package medianOfTwoSortedArrays 

import(
	"math"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	if len(nums1) > len(nums2){
		return findMedianSortedArrays(nums2,nums1)
	}
	var result float64 = 0

	low, high := 0, len(nums1)
	for low <= high{
		partitionX := (high - low) / 2 + low
		partitionY := (len(nums1) + len(nums2) + 1) / 2 - partitionX

		maxLeftX := math.MinInt32
		if partitionX > 0 {
			maxLeftX = nums1[partitionX - 1]
		}		
		minRightX := math.MaxInt32
		if partitionX < len(nums1){
			minRightX = nums1[partitionX]
		}

		maxLeftY := math.MinInt32
		if partitionY > 0{
			maxLeftY = nums2[partitionY - 1]
		}
		minRightY := math.MaxInt32
		if partitionY < len(nums2){
			minRightY = nums2[partitionY]
		}

		if  maxLeftX <= minRightY && maxLeftY <= minRightX{
			if (len(nums1) + len(nums2)) % 2 == 0{
				result = (math.Max(float64(maxLeftX),float64(maxLeftY)) + math.Min(float64(minRightX),float64(minRightY))) / 2
				break
			}else{
				result = math.Max(float64(maxLeftX),float64(maxLeftY)) 
				break
			}
		}else if maxLeftX > minRightY {
			high = partitionX - 1
		}else{
			low = partitionX + 1
		}
	}
	return result
} 

