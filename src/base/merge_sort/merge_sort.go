package merge_sort

func Merge_Sort(nums []int )[]int{
	len := len(nums)
	if(len<2){
		return nums
	}
	mid := len/2
	left := Merge_Sort(nums[:mid])
	right :=Merge_Sort(nums[mid:])
	return Merge(left,right)
}
func Merge(left,right []int)[]int{
	res := make([]int,len(left)+len(right))
	i,j,index := 0,0,0
	for {
		if left[i]<right[j]{
			res[index]=left[i]
			i++
			index++
			if(i==len(left)){
				copy(res[index:],right[j:])
				break
			}
		} else{
			res[index]=right[j]
			j++
			index++
			if(j==len(right)){
				copy(res[index:],left[i:])
				break
			}
		}
	}
	return res
	
}