func buildArray(nums []int) []int {
    arr := make([]int, len(nums))

    for i, _:= range nums {
        arr[i] = nums[nums[i]]
    }
    return arr
}