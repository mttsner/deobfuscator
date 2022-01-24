package ironbrew

func GetBit(num, start, end int) int {
	mask := ((1<<(end-start+1))-1) << start
	return (num & mask) >> start
}