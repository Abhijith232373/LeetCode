func twoSum(nums []int, target int) []int {
    indexmap:=make(map[int]int)
    for i,num:=range nums{
        complement:=target - num
        if idx,found:=indexmap[complement];found{
            return []int{idx,i}
        }
        indexmap[num]=i
    }
    return nil
}