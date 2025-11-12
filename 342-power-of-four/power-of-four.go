func isPowerOfFour(n int) bool {
    if n<=0{
        return false
    }
    for n%4==0{
        // return true
        n=n/4
    }
    return n==1
}