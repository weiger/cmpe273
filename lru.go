package main


const CACHE_SIZE int = 3

var mm = make(map[int]int)

var end int = 0
var first int = 0
var second int = 0
var length int = 0 

func Get(key int) int{

	item, ok := mm[key]

	if  ok{
		return item

	}

	return -1
}

func Set(key int, value int){

	if length < CACHE_SIZE{
		if first == 0{
			first = key
		}else if second == 0{
			second = key
		}

		mm[key] = value
		end = key
		length++
	}else{
		delete(mm, first)
		first = second
		second = end
		end = key
		mm[key] = value
		
	}
}
