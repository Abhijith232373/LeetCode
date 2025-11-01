

func lengthOfLastWord(s string) int {
    str:=strings.Fields(s)
    if len(str)==0{
        return 0
    }
    return len(str[len(str)-1])
}
