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

package traversal

import (
	"testing"

	"github.com/neel-bp/grammes/query/scope"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMin(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'Min' is called with a Scope", func() {
			result := g.Min(scope.Global)
			Convey("Then result should equal 'g.min(global)'", func() {
				So(result.String(), ShouldEqual, "g.min(global)")
			})
		})

		Convey("When 'Min' is called with no params", func() {
			result := g.Min()
			Convey("Then result should equal 'g.min()'", func() {
				So(result.String(), ShouldEqual, "g.min()")
			})
		})

		Convey("When 'Min' is called with a multiple Scope", func() {
			result := g.Min(scope.Global, scope.Local)
			Convey("Then result should equal 'g.min()'", func() {
				So(result.String(), ShouldEqual, "g.min()")
			})
		})
	})
}
