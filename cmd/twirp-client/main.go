// MIT License
//
// Copyright (c) 2018 Alan Chen
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	api "github.com/alanchchen/go-example-twirp/api"
)

var (
	host string
	port int
)

func init() {
	flag.StringVar(&host, "host", "http://127.0.0.1", "The server host")
	flag.IntVar(&port, "port", 8000, "The server port")
	flag.Parse()
}

func main() {
	c := api.NewPetstoreServiceProtobufClient(fmt.Sprintf("%s:%d", host, port), &http.Client{})

	r, err := c.AddPet(context.Background(), &api.AddPetRequest{
		Pet: &api.AddPetRequest_PetMessage{
			Name: "Chubby",
		},
	})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("AddPet: %s", r)
}
