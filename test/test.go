package test
func Quick_Sort(nums []int,start ,end int)[]int{
	if len(nums)<1{
		return []int{}
	}
	i,j := start,end
	mid := nums[0]
	if start<end{
		for i<j{
			//右找比中小
			if i<j&&nums[j]>=mid{
				j--
			}
			if i<j && nums[j]<mid{
				nums[i]=nums[j]
				i++
			}
			if i<j&&nums[i]<=mid{
				i++
			}
			if i<j&&nums[i]>mid{
				nums[j]=nums[i]
				j--
			}
		}
		nums[i]=mid
		Quick_Sort(nums,start,i-1)
		Quick_Sort(nums,i+1,end)
	}
	return nums
}