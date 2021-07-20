func majorityElement(nums []int) []int {

    cand1,cand2,count1,count2 := 0,0,0,0

    for _, num := range nums {
        if cand1 == num {
            count1++
            continue
        }
        if cand2 == num {
            count2++
            continue
        }
        if count1 == 0 {
            cand1 = num
            count1++
            continue
        }
        if count2 == 0 {
            cand2 = num
            count2++
            continue
        }

        count1--
        count2--
    }

    count1 = 0
    count2 = 0
    
    for _, num := range nums {
        if cand1 == num {
            count1++
        } else if cand2 == num {
            count2++
        }
    }

    result := make([]int,0)
    if count1 > len(nums)/3 {
        result = append(result,cand1)
    }
    if count2 > len(nums)/3 {
        result = append(result,cand2)
    }
    return result
}