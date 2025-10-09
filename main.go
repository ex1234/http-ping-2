package main
import ("fmt";"net/http";"os";"time")
func main(){
  if len(os.Args)<2{fmt.Println("Usage: ping <url>");return}
  c:=http.Client{Timeout:5*time.Second}
  r,err:=c.Get(os.Args[1]); if err!=nil{fmt.Println("ERROR:",err);return}
  fmt.Println("Status:",r.StatusCode)
}
