func buddyStrings(s string, goal string) bool {
    if len(s) != len(goal) || len(s) < 2 {
        return false
    }
    diff := make([]int,0)
    same := make(map[string]int,len(s))
    for i := 0; i < len(s); i++ {
        if s[i] == goal[i] {
            same[string(s[i])]++
        } else {
            diff = append(diff,i)
        }
    }
    if len(diff) > 2 || len(diff) == 1 {
        return false
    }

    if len(diff) == 2 && s[diff[0]] == goal[diff[1]] && s[diff[1]] == goal[diff[0]] {
        return true
    }

    if len(diff) == 0 {
        for _,v := range same {
            if v >= 2 {
                return true
            }
        }
    }

    return false

}