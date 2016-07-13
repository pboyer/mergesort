package mergesort

func Sort(a []int) []int {
    if len(a) < 2 {
        return a
    }

    da, db := divide(a)
    return merge(Sort(da), Sort(db))
}

func divide(a []int) ([]int, []int){
    mid := len(a) / 2
    return a[0:mid], a[mid:len(a)]
}

func merge(a []int, b []int) []int {
    n := make([]int, 0, len(a) + len(b))
    for {
        if len(a) == 0 && len(b) == 0 {
            return n
        }

        if len(a) != 0 {
            if len(b) == 0 {
                n = append(n, a[0])
                a = a[1:len(a)]
                continue
            }
            
            if a[0] < b[0] {
                n = append(n, a[0])
                a = a[1:len(a)]
            } else {
                n = append(n, b[0])
                b = b[1:len(b)]
            }
            continue
        }

        if len(b) != 0 {
            n = append(n, b[0])
            b = b[1:len(b)]
            continue
        }
    }
}