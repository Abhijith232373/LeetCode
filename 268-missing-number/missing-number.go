func missingNumber(nums []int) int {
    n:=len(nums)
    totalSum:=n*(n+1)/2
    arrSum :=0
    for _,v :=range nums{
           arrSum+=v
    }
    return totalSum -arrSum

}