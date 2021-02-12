/*Slice:
is a dynamically-sized, flexible view into the elements of an array
A slice is formed by specifying two indices, a low and high bound, separated by a colon: indicator

When slicing, you may omit the high or low bounds to use their defaults instead
*/
/*package main 
import "fmt"
func main(){
s:=[...]int{1, 2, 3, 5, 6, }
a:=s[:]
a=append(a,32)
fmt.Println(a)
fmt.Println(a,len(a),s)
}
*/
/*package main 
import "fmt"
func main(){
var students [3] string
students [0]="Dinajpur Rahul"
students [1]="Rangpur Shishir"
students [2]="Dhaka Mostak"
s :=students [0:2+1]
*/
package main
import "fmt"
func main(){
var students []string
students=append(students,"Dinajpur Rahul","Rangpur Shishir","Dhaka Mostin")
fmt.Println(students,len(students))
fmt.Printf("%T",students)
}