package main


func for1() int{
	a := 0
	for i := 0; i<5; i=i+1{
		a = a + i
	}
	return a
}

func main(){
	_ = for1()
}