package main
import ( 
	"fmt"
)
const NMAX int = 100
type transactionType int
const (
	ItemIn	transactionType = iota
	ItemOut
	Null
)
var transType = map[transactionType]string {
	ItemIn:		"item-in",
	ItemOut: 	"item-out",
	Null:		"Null",
}
type date struct {
	da, mo, ye int
}
type time struct {
	hh, mm, ss int
}
type item struct {
	id 		string
	name 		string
	category 	string
	stock 		int
	cost 		float64
}
type transaction struct {
	transactionID	string
	transactionType transactionType
	item		item
	date		date
	time		time
	applied		bool
}
type itemTab [NMAX]item
type transactionData [NMAX]transaction
type itemData struct {
	items itemTab
	isSorted string
}
func main() {
	var ITEMDATA itemData
	var TRANDATA transactionData
	var nData, nTData int
	takeInput(&ITEMDATA,&nData,&TRANDATA,&nTData)
}
func takeInput(IDATA *itemData, n *int, TDATA *transactionData, nT *int) {
	var arg1, arg2, arg3 string
	var stop bool = false
	for !stop {
		fmt.Println("Type action to start (list, write, sort, replace, remove, edit, help, find) :")
		arg1, arg2, arg3 = "", "", ""
		fmt.Scan(&arg1)
		switch arg1 {
		case "write":
			fmt.Println("write items or transactions?:")
			fmt.Scan(&arg2)
			switch arg2 {
			case "items":
				readItemData(IDATA, n)
				fmt.Println("Data updated!")
				printItemData(*IDATA, *n)
			case "transactions":
				readTransactionData(*IDATA,*n,TDATA,nT)
				adjustItemData(IDATA, *n, TDATA, *nT)
				fmt.Println("Transaction Data Update!")
				printTransactionData(*TDATA, *nT)
			}
		case "sort":
			fmt.Println("Sort by? (time, name, stock, cost) :")
			fmt.Scan(&arg2)
			switch arg2 {	
			case "time":
				fmt.Println("Sorting transaction data by time")
				fmt.Println("('asc' for ascending, 'des' for descending.) :")
				fmt.Scan(&arg3)
				switch arg3 {
				case "asc": 
					fmt.Println("Sorted by transaction time ascending")
					insSortDateTimeAsc(TDATA,*nT)
				case "des":
					fmt.Println("Sorted by transaction time descending")
					insSortDateTimeDes(TDATA,*nT)
				}
				printTransactionData(*TDATA,*nT)
			case "name":
				fmt.Println("Sorting item data by name.")
				fmt.Println("('asc' for ascending, 'des' for descending.) :")
				fmt.Scan(&arg3)
				switch arg3 {
				case "asc":
					fmt.Println("Sorted by name ascending")
					selSortNameAsc(IDATA,*n)
				case "des":
					fmt.Println("Sorted by name descending")
					selSortNameDes(IDATA,*n)
				}
			case "stock":
				fmt.Println("Sorting item data by stock.")
				fmt.Println("('asc' for ascending, 'des' for descending.) :")
				fmt.Scan(&arg3)
				switch arg3 {
				case "asc":
					fmt.Println("Sorted by stock ascending")
					selSortStockAsc(IDATA,*n)
				case "des":
					fmt.Println("Sorted by stock descending")
					selSortStockDes(IDATA,*n)
				}
			case "cost":
				fmt.Println("Sorting item data by cost.")
				fmt.Println("('asc' for ascending, 'des' for descending.)")
				fmt.Scan(&arg3)
				switch arg3 {
				case "asc":
					selSortCostAsc(IDATA,*n)
				case "des":
					selSortCostDes(IDATA,*n)
				}
			}
		case "edit":
			editTransaction(IDATA,*n,TDATA,*nT)
		case "replace":
			fmt.Scan(&arg2)
			switch arg2{
			case "item":
				replaceItem(IDATA,n)
				fmt.Println("Data replaced!")
				printItemData(*IDATA, *n)
			case "transaction":
				replaceTransaction(IDATA,*n,TDATA,*nT)
				fmt.Println("Data replaced!")
				printTransactionData(*TDATA,*nT)
			}
		case "find" :
			findItem(*IDATA,*n)
		case "remove":
			fmt.Scan(&arg2)
			switch arg2 {
			case "item":
				removeItemData(IDATA,n)
				printItemData(*IDATA, *n)
			case "transaction":
				removeTransaction(IDATA,n,TDATA,nT)
				printTransactionData(*TDATA,*nT)
			}
		case "list":
			fmt.Scan(&arg2)
			switch arg2 {
			case "items":
				printItemData(*IDATA, *n)
			case "category" :
				printCateg(*IDATA, *n)
			case "transactions":
				printTransactionData(*TDATA,*nT)
			case "transactions-in":
				printTransactionIn(*TDATA,*nT)
			case "transactions-out":
				printTransactionOut(*TDATA,*nT)
			case "transactions-by-time":
				printTransactionByDate(*TDATA, *nT)
			default:
				fmt.Println("Unknown argument.")
			}
		case "help":
			fmt.Println("=================")
			fmt.Println("    WAREHOUSE    ")
			fmt.Println("=================")
			fmt.Println("command structure:")
			fmt.Println("[action] [parameter/data] [option]")
			fmt.Println("write (items/transactions ) - write item(s) or transaction(s) data into respective data array.")
			fmt.Println("sort time (asc/des)  - sort transactions by time in data.")
			fmt.Println("sort name (asc/des)  - sort items by name in data.")
			fmt.Println("sort cost (asc/des)  - sort items by name in data.")
			fmt.Println("sort stock (asc/des) - sort items by stock in data.")
			fmt.Println("list (items/category)- list items in data.")
			fmt.Println("list (transactions/transactions-(in/out/by-time))- list items in data.")
			fmt.Println("edit 	- edit parameter of existing transaction data.")
			fmt.Println("remove 	- remove item from data.")
			fmt.Println("replace 	- replace item in data.")
			fmt.Println("exit 	- exit program and clear data.")
		case "exit":
			fmt.Println("Exiting program")
			stop = true
		default:
			fmt.Println("Option not recognized.")
			fmt.Println("type 'help' for help.")
		}
	}
}
func printItemData(DATA itemData, n int) {
	fmt.Printf("======= ITEMS =======\n")
	fmt.Printf("%-4s %-12s %-12s %-8s %-8s\n", "ID", "Name", "Category","Stock", "Cost")
	for i:=0;i<n;i++ {
		printItem(DATA,i)
	}
}
func printCateg(DATA itemData, n int) {
	var categ string
	fmt.Scan(&categ)
	fmt.Printf("======= ITEMS =======\n")
	fmt.Printf("%-4s %-12s %-12s %-8s %-8s\n","ID","Name","Category","Stock","Cost")
	for i:=0;i<n;i++ {
		if categ == DATA.items[i].category {
			printItem(DATA,i)
		}
	}
}
func printItem(DATA itemData, i int) {
	fmt.Printf("%4s ",DATA.items[i].id)
	fmt.Printf("%-12s ",DATA.items[i].name)
	fmt.Printf("%-12s ",DATA.items[i].category)
	fmt.Printf("%8d %-5.2f\n",DATA.items[i].stock,
				DATA.items[i].cost)
}
func readItem(DATA *itemData, i int) {	
	fmt.Scan(&DATA.items[i].name)
	fmt.Scan(&DATA.items[i].category)
	fmt.Scan(&DATA.items[i].stock)
	fmt.Scan(&DATA.items[i].cost)
}
func readItemData(DATA *itemData, n *int) {
	var stop bool = false
	for !stop {
		fmt.Scan(&DATA.items[*n].id)
		if DATA.items[*n].id == "null" { stop = true }
		if !stop {
			readItem(DATA,*n)
			*n++
		}
	}
	DATA.isSorted = "false"
}
func removeItemData(DATA *itemData, n *int) {
	var remTar string
	var idx int
	fmt.Println("Type item ID to remove")
	fmt.Scan(&remTar)
	if DATA.isSorted == "id" {
		idx = binSearchItemID(*DATA,*n,remTar)
	} else {
		idx = seqSearchItemID(*DATA,*n,remTar)
	}
	if idx > -1 {
		for i:=idx;i<*n;i++ { DATA.items[i] = DATA.items[i+1] }
		*n--
	} else { fmt.Println("Item not found!") }
}
func replaceItem(DATA *itemData, n *int) {
	var tar string
	var idx int
	fmt.Println("Replace item by ID: ")
	fmt.Scan(&tar)
	if DATA.isSorted == "id" {
		idx = binSearchItemID(*DATA,*n,tar)
	} else {
		idx = seqSearchItemID(*DATA,*n,tar)
	}
	fmt.Println("Input new data: ")
	fmt.Println("(ID, name, category, stock, cost.)")
	readItem(DATA,idx)
}
func findItem(DATA itemData, n int) {
	var by, tar string = "",""
	var idx int = 0
	fmt.Println("Find item by? : ")
	fmt.Scan(&by)
	switch by {
	case "name":
		fmt.Println("Find item by Name: ")
		fmt.Scan(&tar)
		if DATA.isSorted == "name" {
			idx = binSearchItemName(DATA,n,tar)
			if idx != -1 { 
				fmt.Println("Item Found!")
				printItem(DATA,idx) 
			} else { fmt.Println("Item Not Found") }
		} else {
			idx = seqSearchItemName(DATA,n,tar)
			if idx != -1 {
				fmt.Println("Item Found!")
				printItem(DATA,idx)
			} else { fmt.Println("Item Not Found") }
		}
	case "id":
		fmt.Println("Find item by ID: ")
		fmt.Scan(&tar)
		if DATA.isSorted == "id" {
			idx = binSearchItemID(DATA,n,tar)
			if idx != -1 {
				printItem(DATA,idx)
			} else { fmt.Println("Item Not Found!") }
		} else {
			idx = seqSearchItemID(DATA,n,tar)
			if idx != -1 {
				printItem(DATA,idx)
			} else { fmt.Println("Item Not Found") }
		}
	}
}
func readTransactionData(SOURCE itemData, nI int, DATA *transactionData, n *int) {
	var stop bool = false
	var idx int
	for !stop {
		fmt.Scan(&DATA[*n].item.id)
		if DATA[*n].item.id == "null" { stop = true }
		idx = seqSearchItemID(SOURCE, nI, DATA[*n].item.id)
		if idx != -1 && !stop{
			fmt.Println("Input transaction for this item:")
			printItem(SOURCE,idx)
			DATA[*n].item = SOURCE.items[idx]
			fmt.Println("Input in order:")
			fmt.Println("Transaction ID, Date(day, month, year), Time (hours, minutes, seconds).")
			fmt.Scan(&DATA[*n].transactionID)
			inputTransactionTime(DATA,*n)
			fmt.Println("Stock in/out: (use - for out, + for in)")
			fmt.Scan(&DATA[*n].item.stock)
			if DATA[*n].item.stock > 0 {
				DATA[*n].transactionType = ItemIn 
			} else {
				DATA[*n].transactionType = ItemOut
			}
			DATA[*n].applied = false
			*n++
		} else if idx == -1 && !stop {
			fmt.Println("Item ID not found!")
		}
	}
}
func replaceTransaction(SOURCE *itemData, sn int, DATA *transactionData, n int) {
	var target, temp string
	var tgtIdx, idx, formerIdx int = -1,0,0
	var valid,tIDOK bool = false, false
	for !tIDOK {
		fmt.Println("Enter Transaction ID to replace ('cancel' to cancel) :")
		fmt.Scan(&target)
		if target == "cancel" { tIDOK = true }
		tgtIdx = seqSearchTransactionID(DATA,n,target)
		if tgtIdx != -1 && !tIDOK{
			fmt.Println("Transaction Found!")
			printTransaction(*DATA,tgtIdx)
			tIDOK = true
		} else if !tIDOK { fmt.Println("Item not found") }
	}
	for !valid && tIDOK {
		fmt.Println("Enter Item ID to replace ('cancel' to cancel) :")
		fmt.Scan(&temp)
		if temp == "cancel" { valid = true }
		idx = seqSearchItemID(*SOURCE,sn,temp)
		if idx != -1 && !valid {
			fmt.Println("Item Found!")
			fmt.Println("Replacing with the following item:")
			printItem(*SOURCE,idx)
			formerIdx = seqSearchItemID(*SOURCE,sn,DATA[tgtIdx].item.id)
			SOURCE.items[formerIdx].stock += DATA[tgtIdx].item.stock * (-1)
			DATA[tgtIdx].item = SOURCE.items[idx]
			valid = true
		} else if idx == -1 && !valid { fmt.Println("Item Not Found") }
	}
	if tgtIdx != -1 {
		inputTransactionTime(DATA,tgtIdx)
		fmt.Scan(&DATA[tgtIdx].item.stock)
		if DATA[tgtIdx].item.stock > 0 {
			DATA[tgtIdx].transactionType = ItemIn 
		} else {
			DATA[tgtIdx].transactionType = ItemOut
		}
		DATA[tgtIdx].applied = false
	}
	adjustItemData(SOURCE,sn,DATA,n)
}
func removeTransaction(SOURCE *itemData, sn *int, DATA *transactionData,n *int) {
	var target string
	var tgtIdx, formerIdx int = -1,0
	var stop bool = false
	for !stop && tgtIdx == -1 {
		fmt.Scan(&target)
		tgtIdx = seqSearchTransactionID(DATA,*n,target)
		if tgtIdx == -1 { fmt.Println("Transaction entry not found.") }
	}
	if tgtIdx != -1 {
		formerIdx = seqSearchItemID(*SOURCE,*sn,DATA[tgtIdx].item.id)
		SOURCE.items[formerIdx].stock += DATA[tgtIdx].item.stock * (-1)
		for i:=0;i<*n;i++ { DATA[i] = DATA[i+1] }
		fmt.Println("Data removed!")
		*n--
	} else { fmt.Println("Transaction entry not found!") }
}
func editTransaction(SOURCE *itemData, sn int, DATA *transactionData,n int) {
	var choice, target string
	var tgtIdx int = -1
	var stop,stopChoice bool = false, false
	for !stop {
		fmt.Println("Enter Transaction ID to replace ('cancel' to cancel) :")
		fmt.Scan(&target)
		if target == "cancel" { stop = true }
		tgtIdx = seqSearchTransactionID(DATA,n,target)
		if tgtIdx != -1 && !stop {
			fmt.Println("Transaction Found!")
			printTransaction(*DATA,tgtIdx)
			stop = true
		} else if !stop { fmt.Println("Item not found") }
	}
	for !stopChoice {
		fmt.Println("Type parameter to edit ")
		fmt.Println("(item,stock,date,time,'cancel' to cancel) :")
		fmt.Scan(&choice)
		switch choice {
		case "item":
			var idx,formerIdx,stockTemp int
			var temp string
			var valid bool = false
			for !valid {
				formerIdx = seqSearchItemID(*SOURCE,sn,DATA[tgtIdx].item.id)
				fmt.Println("Enter Item ID to replace ('cancel' to cancel) :")
				fmt.Scan(&temp)
				if temp == "cancel" { valid = true }
				idx = seqSearchItemID(*SOURCE,sn,temp)
				if idx != -1 && !valid {
					fmt.Println("Item Found!")
					fmt.Println("Replacing with the following item:")
					printItem(*SOURCE,idx)
					SOURCE.items[formerIdx].stock += DATA[tgtIdx].item.stock * (-1)
					stockTemp = DATA[tgtIdx].item.stock
					DATA[tgtIdx].item = SOURCE.items[idx]
					DATA[tgtIdx].item.stock = stockTemp 
					SOURCE.items[idx].stock += DATA[tgtIdx].item.stock
					valid = true

				} else if idx == -1 && !valid { fmt.Println("Item Not Found") }
			}
			stopChoice = true
		case "stock":
			var idx int
			var temp int
			fmt.Scan(&temp)
			idx = seqSearchItemID(*SOURCE,sn,DATA[tgtIdx].item.id)
			SOURCE.items[idx].stock += (DATA[tgtIdx].item.stock * (-1))
			DATA[tgtIdx].item.stock = temp
			if DATA[tgtIdx].item.stock > 0 {
				DATA[tgtIdx].transactionType = ItemIn 
			} else {
				DATA[tgtIdx].transactionType = ItemOut
			}
			DATA[tgtIdx].applied = false
			adjustItemData(SOURCE,sn,DATA,n)
			stopChoice = true
		case "datetime":
			inputTransactionTime(DATA,tgtIdx)
			stopChoice = true
		case "cancel":
			stopChoice = true
		default:
			fmt.Println("Choice not recognized, try again.")
		}	
	}
}
func inputTransactionTime(DATA *transactionData,i int) {
	fmt.Scan(&DATA[i].date.da)
	fmt.Scan(&DATA[i].date.mo)
	fmt.Scan(&DATA[i].date.ye)
	fmt.Scan(&DATA[i].time.hh)
	fmt.Scan(&DATA[i].time.mm)
	fmt.Scan(&DATA[i].time.ss)
	dateAdjust(&DATA[i].date.da,
			&DATA[i].date.mo,
			&DATA[i].date.ye)
	timeAdjust(&DATA[i].time.hh,
			&DATA[i].time.mm,
			&DATA[i].time.ss)
}
func printTransaction(DATA transactionData, i int) {
	fmt.Printf("%6s ",DATA[i].transactionID)
	fmt.Printf("%4s ",DATA[i].item.id)
	fmt.Printf("%-10s ",DATA[i].item.name)
	fmt.Printf("%-10s ",DATA[i].item.category)
	fmt.Printf("%8d %-5.2f ",DATA[i].item.stock,
				DATA[i].item.cost)
	fmt.Print(DATA[i].transactionType)
	fmt.Printf(" %2d-",DATA[i].date.da)
	fmt.Printf("%2d-",DATA[i].date.mo)
	fmt.Printf("%4d ",DATA[i].date.ye)
	fmt.Printf("%2d-",DATA[i].time.hh)
	fmt.Printf("%2d-",DATA[i].time.mm)
	fmt.Printf("%2d\n",DATA[i].time.ss)
}
func printTransactionData(DATA transactionData, n int) {
	fmt.Printf("%s | %-4s | %-20s | %-12s | %-8s | %-8s | %s |\n", "Trans.ID","ID", "Item Information","Stock",
								"Trans.Type","Date", "Time")
	for i:=0;i<n;i++ {
		printTransaction(DATA,i)
	}
}
func printTransactionIn(DATA transactionData, n int) {
	for i:=0;i<n;i++ {
		if DATA[i].item.stock > 0 { printTransaction(DATA,i) }
	}
}
func printTransactionOut(DATA transactionData, n int) {
	for i:=0;i<n;i++ {
		if DATA[i].item.stock < 0 { printTransaction(DATA,i) }
	}
}
func printTransactionByDate(DATA transactionData, n int) {
	var upLimD, loLimD date
	var upLimT, loLimT time
	fmt.Println("Enter date lower limit for search: ")
	fmt.Scan(&loLimD.da)
	fmt.Scan(&loLimD.mo)
	fmt.Scan(&loLimD.ye)
	fmt.Println("Enter date upper limit for search: ")
	fmt.Scan(&upLimD.da)
	fmt.Scan(&upLimD.mo)
	fmt.Scan(&upLimD.ye)
	fmt.Println("Enter time lower limit for search: ")
	fmt.Scan(&loLimT.hh)
	fmt.Scan(&loLimT.mm)
	fmt.Scan(&loLimT.ss)
	fmt.Println("Enter time upper limit for search: ")
	fmt.Scan(&upLimT.hh)
	fmt.Scan(&upLimT.mm)
	fmt.Scan(&upLimT.ss)
	for i:=0;i<n;i++ {
		if compareDateGRT(DATA[i].date ,loLimD ,DATA[i].time, loLimT) && compareDateLSR(DATA[i].date, upLimD,DATA[i].time, upLimT) {
			printTransaction(DATA,i)
		}
	}
}
// 	SEARCHING
func seqSearchItemID(DATA itemData, n int, target string) int {
	for i:=0;i<n;i++ {
		if target == DATA.items[i].id { return i }
	}
	return -1
}
func binSearchItemID(DATA itemData, n int, target string) int {
	var l,r int = 0, n-1
	var m int = (r+l)/2
	for l <= r {
		switch{
		case target == DATA.items[m].id:
			return m
		case target == DATA.items[l].id:
			return l
		case target == DATA.items[r].id:
			return r
		case target > DATA.items[m].id:
			l = m + 1
		case target < DATA.items[m].id:
			r = m - 1
		default:
			return m
		}
		m = (l+r)/2
	}
	return -1
}
func seqSearchItemName(DATA itemData, n int, target string) int {
	for i:=0;i<n;i++ {
		if target == DATA.items[i].name { return i }
	}
	return -1
}
func binSearchItemName(DATA itemData, n int, target string) int {
	var l,r int = 0, n-1
	var m int = (r+l)/2
	for l<=r {
		switch{
		case target == DATA.items[m].name:
			return m
		case target == DATA.items[l].name:
			return l
		case target == DATA.items[r].name:
			return r
		case target > DATA.items[m].name:
			l = m + 1
		case target < DATA.items[m].name:
			r = m - 1
		default:
			return m
		}
		m = (l+r)/2
	}
	return -1
}
func seqSearchTransactionID(DATA *transactionData, n int, target string) int {
	for i:=0;i<n;i++ {
		if target == DATA[i].transactionID { return i }
	}
	return -1
}
//	ADJUSTMENTS
func dateAdjust(da,mo,ye *int) {
	if *da > 30 {
		*mo += (*da-(*da%30))/30
		*da =  *da%30
	}
	if *mo > 12 {
		*ye += (*mo-(*mo%12))/12
		*mo =  *mo%12
	}
}
func timeAdjust(hh,mm,ss *int) {
	if *ss >= 60 {
		*mm += (*ss-(*ss%60))/60
		*ss =  *ss%60
	}
	if *mm >= 60 {
		*hh += (*mm-(*mm%60))/60
		*mm =  *mm%60
	}
	if *hh >= 24 { *hh = *hh%24 }
}
func adjustItemData(IDATA *itemData, nI int, TDATA *transactionData, nT int) {
	for i:=0;i<nI;i++ {
		for j:=0;j<nT;j++ {
			if TDATA[j].item.id == IDATA.items[i].id && TDATA[j].applied == false {
				IDATA.items[i].stock += TDATA[j].item.stock
				TDATA[j].applied = true
			}
		}
	}
}
func (ss transactionType) String() string {
	return transType[ss]
}
//	SWAP
func swap(DATA *itemData, x,y int) {
	var temp item
	temp = DATA.items[x]
	DATA.items[x] = DATA.items[y]
	DATA.items[y] = temp
	DATA.isSorted = "false"
}
func swapTrans(DATA *transactionData, x,y int) {
	var temp transaction
	temp = DATA[x]
	DATA[x] = DATA[y]
	DATA[y] = temp
}
//	COMPARATORS
func compareDateLSR(x,y date, xt, yt time) bool {
	if x.ye < y.ye { 
		return true 
	} else  if x.ye == y.ye {
		if x.mo < y.mo { 
			return true 
		} else if x.mo == y.mo {
			if x.da < y.da { 
				return true 
			} else {
				compareTimeLSR(xt,yt)
			}
		}
	}
	return false
}
func compareDateGRT(x,y date, xt, yt time) bool {
	if x.ye > y.ye { 
		return true 
	} else  if x.ye == y.ye {
		if x.mo > y.mo { 
			return true 
		} else if x.mo == y.mo {
			if x.da > y.da { 
				return true 
			} else {
				compareTimeGRT(xt,yt)
			}
		}
	}
	return false
}
func compareTimeLSR(x,y time) bool {
	if x.hh < y.hh { 
		return true 
	} else  if x.hh == y.hh {
		if x.mm < y.mm { 
			return true 
		} else if x.mm == y.mm {
			if x.ss < y.ss { return true }
		}
	}
	return false
}
func compareTimeGRT(x,y time) bool {
	if x.hh > y.hh { 
		return true 
	} else  if x.hh == y.hh {
		if x.mm > y.mm { 
			return true 
		} else if x.mm == y.mm {
			if x.ss > y.ss { return true }
		}
	}
	return false
}
//	SORTING
func selSortNameAsc(DATA *itemData, n int) {
	var i, j, minima int
	for i=0;i<n;i++ {
		minima = i
		for j = i+1;j<n;j++ {
			if DATA.items[j].name < DATA.items[minima].name { minima = j }
		}
		swap(DATA, i, minima)
	}
	DATA.isSorted = "name"
}
func selSortNameDes(DATA *itemData, n int) {
	var i, j, max int
	for i=0;i<n;i++ {
		max = i
		for j = i+1;j<n;j++ {
			if DATA.items[j].name > DATA.items[max].name { max = j }
		}
		swap(DATA, i, max)
	}
	DATA.isSorted = "name"
}
func selSortStockAsc(DATA *itemData, n int) {
	var i, j, minima int
	for i=0;i<n;i++ {
		minima = i
		for j = i+1;j<n;j++ {
			if DATA.items[j].stock < DATA.items[minima].stock { minima = j }
		}
		swap(DATA, i, minima)
	}
	DATA.isSorted = "stock"
}
func selSortStockDes(DATA *itemData, n int) {
	var i, j, max int
	for i=0;i<n;i++ {
		max = i
		for j = i+1;j<n;j++ {
			if DATA.items[j].stock > DATA.items[max].stock { max = j }
		}
		swap(DATA, i, max)
	}
	DATA.isSorted = "stock"
}
func selSortCostAsc(DATA *itemData, n int) {
	var i, j, minima int
	for i=0;i<n;i++ {
		minima = i
		for j = i+1;j<n;j++ {
			if DATA.items[j].cost < DATA.items[minima].cost { minima = j }
		}
		swap(DATA, i, minima)
	}
	DATA.isSorted = "cost"
}
func selSortCostDes(DATA *itemData, n int) {
	var i, j, max int
	for i=0;i<n;i++ {
		max = i
		for j = i+1;j<n;j++ {
			if DATA.items[j].cost > DATA.items[max].cost { max = j }
		}
		swap(DATA, i, max)
	}
	DATA.isSorted = "cost"
}
func insSortDateTimeAsc(DATA *transactionData, n int) {
	var p, i int = 1,0
	var temp transaction
	for p < n {
		i=p
		temp=DATA[p]
		for i > 0 && compareDateLSR(temp.date,DATA[i-1].date,temp.time,DATA[i-1].time){
			swapTrans(DATA,i,i-1)
			i--
		}
		DATA[i] = temp
		p++
	}
}
func insSortDateTimeDes(DATA *transactionData, n int) {
	var p, i int = 1,0
	var temp transaction
	for p < n {
		i=p
		temp=DATA[p]
		for i > 0 && compareDateGRT(temp.date,DATA[i-1].date,temp.time,DATA[i-1].time){	
			swapTrans(DATA,i,i-1)
			i--
		}
		DATA[i] = temp
		p++
	}
}
