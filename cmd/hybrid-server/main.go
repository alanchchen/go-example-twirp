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
	"flag"
	"fmt"
	"log"
	"net"
	"net/http"

	api "github.com/alanchchen/go-example-twirp/api"
	"github.com/soheilhy/cmux"
	"google.golang.org/grpc"
)

var (
	port int
)

func init() {
	flag.IntVar(&port, "port", 8000, "The server listening port")
	flag.Parse()
}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	m := cmux.New(lis)

	// Match connections in order:
	grpcListener := m.Match(cmux.HTTP2HeaderField("content-type", "application/grpc"))
	httpListener := m.Match(cmux.HTTP1Fast())

	service := &exampleServer{}

	s := grpc.NewServer()
	api.RegisterPetstoreServiceServer(s, service)

	go func() {
		if err = s.Serve(grpcListener); err != nil {
			log.Fatal("grpc: ", err)
		}
	}()

	twirpHandler := api.NewPetstoreServiceServer(service, nil)
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: twirpHandler,
	}

	go func() {
		if err = srv.Serve(httpListener); err != nil {
			log.Fatal("http: ", err)
		}
	}()

	if err = m.Serve(); err != nil {
		log.Fatal("cmux: ", err)
	}
}
