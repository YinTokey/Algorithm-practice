def binary_search(find,lst):
    low = 0
    high = len(lst)
    while low <= high:
        mid = (low + high) // 2
        if lst[mid] == find:
            return mid
        elif lst[mid] > find:
            high = mid -1
        else:
            low = mid + 1

    return -1


