package main
import (
	"fmt"
	// "encoding/json"
	"net/http"
	"io/ioutil"
	// "github.com/sdputurn/learn-go"
	// "github.com/sdputurn/learn-go/patterns"
	// "hello/patterns"
)
func main () {
	// learn_go.Hello1()
	// pattern_examples.P1()
	// learn_go.main()
	// patterns.P1()
	// a := []int{1,2,3,4,5}
	// d := map[string]map[string]string{"a":{"b":"c"}} //, "n":string(map[string]string{"n1":"1"})}
	// fmt.Println(a, "+++++++++++++",d)
	s := "string" + "-'string"
	fmt.Println(s)
	// var a string = "+++++++++++++++"
	// a:= 'a'
	// fmt.Printf("v - %v , c- %c, #v - %#v, p - %p, T - %T \n", a, a, a,&a,a)
	// j := `{"name":"sandeep", "school": "kvs", "sub":["phy","math"], "extra": {"hosbbies":"linux"}}`
	// type student struct {
	// 	Name string
	// 	School string
	// }
	// var stu student
	// err := json.Unmarshal([]byte(j), &stu)
	// fmt.Println(err)
	// fmt.Printf("%T - %v \n",stu,stu)
	resp, err := http.Get("https://jsonplaceholder.typicode.com/posts/1")
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))


}

