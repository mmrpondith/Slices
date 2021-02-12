package main 
import "fmt"
func main(){
s:=[...]int{1, 2, 3, 5, 6, }
a:=s[:]
a=append(a,32)
fmt.Println(a)
fmt.Println(a,len(a),s)
}