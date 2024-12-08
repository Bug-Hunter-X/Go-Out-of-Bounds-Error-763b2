func myFunc(a []int) int {
    if len(a) == 0 {
        return 0
    }
    return a[len(a)]
}