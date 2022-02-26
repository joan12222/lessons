package main

import "time"

//给定一个字符串数组
//[“I”,“am”,“stupid”,“and”,“weak”]
//用 for 循环遍历该数组并修改为func main() {
//[“I”,“am”,“smart”,“and”,“strong”]
//}

//func main() {
//	var myArray = [5]string{"I", "am", "stupid", "and", "weak"}
//	//	var mySlice []string
//	for i, v := range myArray {
//		//		mySlice[i] = v
//		if v == "stupid" {
//			myArray[i] = "smart"
//		} else if v == "weak" {
//			myArray[i] = "strong"
//		}
//		println(myArray[i])
//	}
//}

func main() {
	c := make(chan int, 10)
	go func() {
		i := 0
		for ; ; i++ {
			time.Sleep(time.Second)
			c <- i
		}
	}()
	``
	for {
		time.Sleep(time.Second)
		v := <-c
		println(v)
	}
	time.Sleep(2 * time.Second)
}
