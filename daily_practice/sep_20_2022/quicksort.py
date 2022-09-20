array = [20, 10, 50, 90, 40 , 60, 30]


def partition(arr, low, high):
    pivot = arr[high]
    i = low-1
    for j in range(low, high+1):
        if arr[j] < pivot:
            i = i+1
            (arr[i], arr[j]) = (arr[j], arr[i])


def quickSortArray(arr, low, high):

    pivot_correct_index = partition(arr, low, high)


