package rest

import (
	"container/heap"
	"strconv"
	"test-fizz-buzz/dto"
)

func FizzBuzz(firstString string, secondString string, firstNumber int, secondNumber int, limit int) (result []string) {
	for i := 1; i <= limit; i++ {

		if i%firstNumber == 0 && i%secondNumber == 0 {
			result = append(result, firstString+secondString)
		} else if i%firstNumber == 0 {
			// Multiple of int1
			result = append(result, firstString)
		} else if i%secondNumber == 0 {
			// Multiple of int2
			result = append(result, secondString)
		} else {
			// Neither, so print the number itself
			result = append(result, strconv.Itoa(i))
		}
	}
	return result
}

type kv struct {
	Key   string
	Value int
}

func getHeap(m map[string]int) *KVHeap {
	h := &KVHeap{}
	heap.Init(h)
	for k, v := range m {
		heap.Push(h, kv{k, v})
	}
	return h
}

type KVHeap []kv

func (h KVHeap) Len() int           { return len(h) }
func (h KVHeap) Less(i, j int) bool { return h[i].Value > h[j].Value }
func (h KVHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *KVHeap) Push(x interface{}) {
	*h = append(*h, x.(kv))
}

func (h *KVHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func GetRequestParams(str []string) dto.Request {
	firstNumber, _ := strconv.Atoi(str[2])
	secondNumber, _ := strconv.Atoi(str[3])
	limit, _ := strconv.Atoi(str[4])
	request := dto.Request{
		FirstString:  str[0],
		SecondString: str[1],
		FirstNumber:  firstNumber,
		SecondNumber: secondNumber,
		Limit:        limit,
	}
	return request
}
