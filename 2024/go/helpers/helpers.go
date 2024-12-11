package helpers

func partition(arr []int, low, high int) ([]int, int){
  pivot := arr[high]
  i := low
  for j := low; j<high; j++{
    if arr[j] < pivot{
      if i != j{
        arr[i], arr[j] = arr[j], arr[i]
      }
      i++
    }
  }

  arr[i], arr[high] = arr[high], arr[i]
  return arr, i
}

func quicksort(arr []int, low, high int) ([]int){
  if low < high{
    var pivot int
    arr, pivot = partition(arr, low, high)
    arr = quicksort(arr, low, pivot-1)
    arr = quicksort(arr, pivot+1, high)
  }

  return arr
}

func QuicksortStart(arr []int) ([]int){
  return quicksort(arr, 0, len(arr)-1)
}
