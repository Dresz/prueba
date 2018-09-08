package main
import (
  "fmt"
  "net/http"
	"io/ioutil"
  "strings"
  "google.golang.org/appengine"
)

func handler(w http.ResponseWriter, r *http.Request) {
  url := "https://graph.facebook.com/v3.1/685148168507218/feed?"

	payload := strings.NewReader("message=Alerta%20la%20Casa%20se%20esta%20quemando&access_token=EAADNGF3mxwMBAC8VNUbnyrihiXpM2nQ5m8avo3LdZCPSZCnpy469HSwOVkoFaz9S4S5L0w8I7w4y4xCgZAylgQ1B7ukislASxeq1zOVsohmFUWGoi42K8DWRP4j8fh8Ht3ZCOg2uYIBWJEvQKTl1ZA8ZCBTaCgRCSkK4qqtdlEOQZDZD")

	req, err := http.NewRequest("POST", url, payload)

  if err != nil {
    fmt.Fprint(w, "Fallo req!")
    return
  }
req.Header.Add("content-type", "application/x-www-form-urlencoded")
req.Header.Add("cache-control", "no-cache")
  client := &http.Client{}
  resp, err :=  client.Do(req)
  
  if err != nil {
    fmt.Fprint(w, "Fallo resp!")
    return
  }
  defer resp.Body.Close()
  body, err := ioutil.ReadAll(resp.Body)

  fmt.Printf("%s\n", string(body))
	fmt.Fprint(w, "Post enviado!")
}
func main() {
  http.HandleFunc("/", handler) // Set endpoint handler
  appengine.Main() // Start the server
}

//EAADNGF3mxwMBAC8VNUbnyrihiXpM2nQ5m8avo3LdZCPSZCnpy469HSwOVkoFaz9S4S5L0w8I7w4y4xCgZAylgQ1B7ukislASxeq1zOVsohmFUWGoi42K8DWRP4j8fh8Ht3ZCOg2uYIBWJEvQKTl1ZA8ZCBTaCgRCSkK4qqtdlEOQZDZD
//EAADNGF3mxwMBACbC1RdceaKTO4gE6z3sp9zv4lZCZA0BYdtPwulhkmgVOZB3JbFqmdqpWwiLg0bLnReq3uuJUdIMGHcZAH0Qqo62Vt2ZBQ42UiZCSzhZBP3ApgRJFBdSLwMKKmd5T2p5CXZCd7JJjVD7jurdmNtlOnYZD
/*

curl -i -X POST \
 "https://graph.facebook.com/v3.1/685148168507218/feed?message=Alerta%20la%20Casa%20se%20esta%20quemando&access_token=EAADNGF3mxwMBAC8VNUbnyrihiXpM2nQ5m8avo3LdZCPSZCnpy469HSwOVkoFaz9S4S5L0w8I7w4y4xCgZAylgQ1B7ukislASxeq1zOVsohmFUWGoi42K8DWRP4j8fh8Ht3ZCOg2uYIBWJEvQKTl1ZA8ZCBTaCgRCSkK4qqtdlEOQZDZD"
 */