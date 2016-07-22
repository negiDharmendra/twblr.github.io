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
	return nil
}

func concatenate(dest []string, newValues ...string) []string {
	return nil
}

func equals(list1 []string, list2 []string) bool {
	return false
}

func partialReverse(src []int, from, to int) []int {
	return nil
}
