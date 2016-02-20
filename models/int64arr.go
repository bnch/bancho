package models

// Stolen from https://groups.google.com/forum/#!topic/golang-nuts/tyDC4S62nPo

type uint64arr []uint64

func (a uint64arr) Len() int {
	return len(a)
}
func (a uint64arr) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
func (a uint64arr) Less(i, j int) bool {
	return a[i] < a[j]
}
