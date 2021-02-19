package deobfuscator 

import (
	"strconv"
)

func getBit(str string, pos *int) int64 {
	len,_ := strconv.ParseInt(str[*pos:*pos+1], 36, 64)
	val,_ := strconv.ParseInt(str[*pos+1:*pos+int(len)+1], 36, 64)
	*pos += int(len) + 1
	return val
} 

func decompress(compressed string) []byte {
	pos := 0
	var code int64 = 256
	dictionary := make(map[int64][]int64)
	currChar := []int64{getBit(compressed, &pos)}
    result := currChar
    for pos < len(compressed) {
		var word []int64
		element := getBit(compressed, &pos)

		if element < 256 {
			word = []int64{element}
		} else if x, ok := dictionary[element]; ok {
			word = x
		} else{
			word = append(currChar, currChar[0])
		}
		result = append(result,  word...)
		
		app := append(currChar, word[0])
		dst := make([]int64, len(app))
		copy(dst, app)

		dictionary[code] = dst
        code++
    
        currChar = word
	}

	// Dirty fix cuz I don't care.
	ret := []byte{}
	for _, b := range result {
		ret = append(ret, byte(b))
	}

    return ret
}