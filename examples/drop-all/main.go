// Copyright (c) 2018 Northwestern Mutual.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package main

import (
	"flag"

	"go.uber.org/zap"

	"github.com/neel-bp/grammes"
	"github.com/neel-bp/grammes/examples/exampleutil"
)

var (
	// addr is used for holding the connection IP address.
	// for example this could be, "ws://127.0.0.1:8182"
	addr string
)

func main() {
	flag.StringVar(&addr, "h", "", "Connection IP")
	flag.Parse()

	logger := exampleutil.SetupLogger()
	defer logger.Sync()

	if addr == "" {
		logger.Fatal("No host address provided. Please run: go run main.go -h <host address>")
		return
	}

	// Create a new Grammes client with a standard websocket.
	client, err := grammes.DialWithWebSocket(addr)
	if err != nil {
		logger.Fatal("Couldn't create client", zap.Error(err))
	}

	// DropAll will remove all vertices from the graph currently.
	// Essentially blank slating all of our data.
	client.DropAll()

	// Add a vertex to bring the vertex count to 1.
	client.AddVertex("testingvertex")

	count, err := client.VertexCount()
	if err != nil {
		logger.Fatal("Couldn't count vertices", zap.Error(err))
	}

	// Log the amount of vertices that are now on the graph.
	logger.Info("Counted vertices", zap.Int64("count", count))

	// Now drop all of the vertices.
	client.DropAll()

	// Recount the vertices on the graph
	count, err = client.VertexCount()
	if err != nil {
		logger.Fatal("Couldn't count vertices", zap.Error(err))
	}

	// Log the amount of vertices that are now on the graph.
	logger.Info("Counted vertices", zap.Int64("count", count))
}
