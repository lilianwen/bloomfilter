package bloomFilter

import (
	"bufio"
	"fmt"
	"os"
	"testing"
)



func TestBloomFuncs(t *testing.T) {
	var elemNum uint64 = 10000000000
	var errRate  = 0.0001
	var bloomSize = CalBloomSize(elemNum, errRate)
	fmt.Println("bloom size in bit:",bloomSize)//1917011675474
	fmt.Println("bloom size in Gbyte:",bloomSize/1024/1024/1024/8)

	var hashFuncNum = CalHashFuncNum(elemNum, bloomSize)
	fmt.Println("hash functions number:",hashFuncNum)

	errRate = CalErrRate(elemNum, bloomSize, hashFuncNum)
	fmt.Println("error rate:",errRate)
}

func TestBloomFilter(t *testing.T) {
	file, err := os.Open("word-list-large.txt")
	if err != nil {
		t.Error(err)
		return
	}
	defer file.Close()
	var wordlist []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		wordlist = append(wordlist, scanner.Text())
	}
	var elemNum = uint64(len(wordlist))
	var errRate = 0.0001
	bloomSize := CalBloomSize(elemNum, errRate)
	hashFuncNum := CalHashFuncNum(elemNum,bloomSize)

	filter := NewFilter(elemNum, bloomSize, hashFuncNum, errRate)
	filter.Init()
	for _,elem := range wordlist {
		filter.Add([]byte(elem))
	}

	var testcases = []struct {
		Elem string
		Want bool
	}{
		{"hello", true},
		{"zoo", false},
		{"word", true},
		{"alibaba", false},
	}

	for _,oneCase := range testcases {
		got := filter.IsContain([]byte(oneCase.Elem))
		if got != oneCase.Want {
			if got {
				t.Error("should not contain", oneCase.Elem)
			} else {
				t.Error("contain", oneCase.Elem)
			}
			t.Error("got != want")
			return
		}
	}
}
