// Licensed to Elasticsearch B.V. under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Elasticsearch B.V. licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package apmhttp // import "go.elastic.co/apm/module/apmhttp"

import (
	"net/http"
	"strings"
)

// UnknownRouteRequestName returns the transaction name for the server request, req,
// when the route could not be determined.
func UnknownRouteRequestName(req *http.Request) string {
	const suffix = " unknown route"
	var b strings.Builder
	b.Grow(len(req.Method) + len(suffix))
	b.WriteString(req.Method)
	b.WriteString(suffix)
	return b.String()
}

// ServerRequestName returns the transaction name for the server request, req.
func ServerRequestName(req *http.Request) string {
	var b strings.Builder
	b.Grow(len(req.Method) + len(req.URL.Path) + 1)
	b.WriteString(req.Method)
	b.WriteByte(' ')
	b.WriteString(req.URL.Path)
	return b.String()
}

// ClientRequestName returns the span name for the client request, req.
func ClientRequestName(req *http.Request) string {
	var b strings.Builder
	b.Grow(len(req.Method) + len(req.URL.Host) + 1)
	b.WriteString(req.Method)
	b.WriteByte(' ')
	b.WriteString(req.URL.Host)
	return b.String()
}
