package data_structures

type mapOperation func(int32) int32
type filterOperation func(int32) bool

func mapInts(op mapOperation, vals []int32) []int32 {
	for i := range vals {
		vals[i] = op(vals[i])
	}
	return vals
}

func filterInts(op filterOperation, vals []int32) []int32 {
	result := []int32{}
	for i := range vals {
		if op(vals[i])==true {
			result = append(result,vals[i])
		}
	}
	return result
}

func concatenate(dest []string, newValues ...string) []string {
	for i := range newValues {
		dest = append(dest,newValues[i])
	}
	return dest
}

func equals(list1 []string, list2 []string) bool {
	if len(list1) != len(list2){
		return false
	}
	for i := range list1 {
		if list1[i] != list2[i]{
			return false
		}
	}
	return true
}

func partialReverse(src []int, from, to int) []int {
	newSlice := []int{}
	for i:=to; i>=from; i-- {
		newSlice = append(newSlice, src[i])
	}
	return newSlice
}
