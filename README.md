== REV.io utility

== USAGE
```
package main

import "github.com/meetme2meat/rev.io/revio"

func main() {
  r := revio.New(username, code, password)
  // create inventory
  inventory := r.Inventory()
  // params is http params
  header := request.Header{
    Authorization: r.Authorization(),
    Accept: "application/json",
    ContentType: "application/json",
  }
  ctx := context.WithValue(context.Background(),"HEADERS", &header)
  response := inventory.Create(ctx, params)
  fmt.Println(response.Code)
  fmt.Println(response.Error)
  fmt.Println(response.Message) 
}
```
