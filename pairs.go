package main
import "fmt"


func Solution(str string) []string {

	// fmt.Println(str)
	// for i,elem := range str{
	// 	fmt.Println(str[elem])
	// 	fmt.Print(i)
	// }

	 str1:=[] string{}
	 tempstr:=""

	for i, c := range str {
		
		if i%2==0{
			
			//fmt.Println(tempstr)
			tempstr=""
		}else{
			tempstr=tempstr+string(c)
			fmt.Println(i,tempstr)
		}
	}
	return str1
}


func main(){


fmt.Println(Solution("Test8"))
}