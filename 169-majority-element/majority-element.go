func majorityElement(nums []int) int {
    count :=0
    cantidate:=0

    for _,num:=range nums {
        if count == 0{
        cantidate =num

        }
        if num ==cantidate{
            count++
        }else{
            count--
        }
    }
    return cantidate
}