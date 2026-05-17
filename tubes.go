package main
import "fmt"

const NMAX int = 100

type timeTable struct {
	inHH, inMM, inSS		int
	outHH, outMM, outSS	int
}
type item struct {
	Name	string
	ID		string
	Stock int
	Cost	int
	Time	timeTable
}
type itemArr[NMAX]item

func readData(item *itemArr, nData *int, sortedStatus *bool) {
	var i 		int
	var stop	bool
	i = *nData
	stop = false
	for i < NMAX && !stop {
		fmt.Scan(&item[i].Name)
		if item[i].Name == "none" {
			stop = true
		} else {
			fmt.Scan(&item[i].ID)
			fmt.Scan(&item[i].Stock)
			fmt.Scan(&item[i].Cost)
			fmt.Scan(&item[i].Time.inHH, &item[i].Time.inMM, &item[i].Time.inSS)
			fmt.Scan(&item[i].Time.outHH, &item[i].Time.outMM, &item[i].Time.outSS)
			i = i + 1
			*sortedStatus = false 
		}
	}
	*nData = i
}
func replData(item *itemArr, nData int, sortedStatus *bool) {
	var target int

}
