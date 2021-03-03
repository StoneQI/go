package main


func for1() int{
	a := 0
	for i := 0; i<5; i=i+1{
		a = a + i
	}
	return a
}

func main(){
	var aa = func(){ 
		a := 0
		for i := 0; i<5; i=i+1{
			a = a + i
		}
	}
	aa()
	_ = for1()
}