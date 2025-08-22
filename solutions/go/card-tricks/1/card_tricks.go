package cards

func FavoriteCards() []int {
    slice := []int{2,6,9}
    return slice
}

func GetItem(slice []int, index int) int {
    if index >= len(slice)|| index < 0{
        return -1
    }
    return slice[index]
}

func SetItem(slice []int, index, value int) []int {
     if index >= len(slice) || index < 0 {
         slice = append(slice, value)
         return slice
     }
    slice[index] = value
    return slice
}

func PrependItems(slice []int, values ...int) []int {
	return append(values, slice...)
}

func RemoveItem(slice []int, index int) []int {
    if index >= len(slice) || index < 0 {
         return slice
     }
    slice = append(slice[:index],slice[index+1:]...)
    return slice
}
