package main
//
//import (
//	//"fmt"
//	"fmt"
//	"net/http"
//)
//
//func main() {
//	//addr, _ := net.ResolveUDPAddr("udp", "127.0.0.1:8080")
//	//fmt.Println(addr.Port)
//	//conn, _ := net.DialUDP("udp", nil, addr)
//	//defer conn.Close()
//	//conn.Write([]byte("hello"))
//	//a := []int {1,2,3,4,5}
//	//b := a[:3]
//	//c := make([]int, len(b))
//	//copy(c,b)
//	//c[2] = 3
//	//fmt.Println(b)
//	//fmt.Println(c)
//	http.HandleFunc("/", newHandler)
//
//	_ = http.ListenAndServe(":2345", nil)
//}
//func newHandler(w http.ResponseWriter, r *http.Request) {
//	x := r.URL.Query().Get("a")
//	fmt.Fprintf(w, "HELLO %s", x)
//	//// if only one expected
//	//param1 := r.URL.Query().Get("param1")
//	//if param1 != "" {
//	//	// ... process it, will be the first (only) if multiple were given
//	//	// note: if they pass in like ?param1=&param2= param1 will also be "" :|
//	//}
//	//
//	//// if multiples possible, or to process empty values like param1 in
//	//// ?param1=&param2=something
//	//param1s := r.URL.Query()["param1"]
//	//if len(param1s) > 0 {
//	//	// ... process them ... or you could just iterate over them without a check
//	//	// this way you can also tell if they passed in the parameter as the empty string
//	//	// it will be an element of the array that is the empty string
//	//}
//}
//
////func aux() {
////	go func() {
////		defer wg.Done()
////		for i := 0; i <= 10; i++ {
////			fmt.Println(i)
////		}
////	}()
////}
