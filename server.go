package main
 
import "net/http"
import "log"
import "encoding/json"
 
func main() {
  bytes := make([]byte, 1024*1024)
  for i, _ := range bytes{
    bytes[i] = 100
  }
 
  http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
   // w.Write(bytes)
	a := map[string]interface{}{
		"foo": "bar",
		"bar": "baz",
		"nest": map[string]interface{}{
			"a": "b",
		},
	}

	var b []byte
	for i := 0; i < 10; i++ {
		b, _ = json.Marshal(a)
		w.Write(b)
	}
	//fmt.Fprint(w, "Hello world!")
  })
 
  log.Println("Server running at http://127.0.0.1:8000/")
  http.ListenAndServe(":8000", nil)
}
