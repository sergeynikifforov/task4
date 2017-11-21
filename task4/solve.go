func RemoveEven(a []int)[]int {
    s:=make([]int,0,1000)
    for _,v := range a {
        if(v%2==1){
            s = append(s,v)
        }
    }
    return s
}
func PowerGenerator(b int) (func() int){
    res := 1
    return func() (ret int){
        ret = res*b
        res = ret
        return
    }

}
func DifferentWordsCount(c string) int {
    str := ``
    x := make(map[string]bool)
    for i := 0; i < len(c); i++ {
        for ;(i < len(c) && unicode.IsLetter(rune(c[i]))) ; i++{
            str = str + string(unicode.ToLower(rune(c[i])))
        }
        if(len(str) > 0){
            x[str] = true;
        }
        str = ``
    }
    return len(x)
}
