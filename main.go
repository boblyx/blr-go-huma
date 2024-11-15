/*
* main.go
*
*/

package main

import (
  "context"
  "fmt"
  "net/http"

  "github.com/danielgtaylor/huma/v2"
  "github.com/danielgtaylor/huma/v2/adapters/humachi"
  "github.com/go-chi/chi/v5"

  _ "github.com/danielgtaylor/huma/v2/formats/cbor"
)

// HelloOutput represents a hello response.
type HelloOutput struct {
  Body struct {
    Message string `json:"message" example:"Hello, world!" doc:"My docs are awesome."`
  }
}

func main(){
  router := chi.NewMux()
  api := humachi.New(router, huma.DefaultConfig("blr-go-huma", "0.0.1"))

  huma.Get(api, "/api/v1/hello/{name}", func(ctx context.Context, input *struct {

    Name string `path:"name" maxLength:"30" example:"world" doc: "Name to say hello to."`

  }) (*HelloOutput, error){
    resp := &HelloOutput{}
    resp.Body.Message = fmt.Sprintf("Hello, %s!", input.Name)
    return resp, nil
  })
  fmt.Printf("Listening at %s!\n", "127.0.0.1:8888")
  http.ListenAndServe("127.0.0.1:8888", router)
}
