// Copyright 2014 The oauth2 Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// +build appengine

// App Engine hooks.

package oauth2

import (
	"net/http"

	"google.golang.org/appengine/urlfetch"
	"golang.org/x/net/context"
	"golang.org/x/oauth2/internal"
)

func init() {
	internal.RegisterContextClientFunc(contextClientAppEngine)
}

func contextClientAppEngine(ctx context.Context) (*http.Client, error) {
	return urlfetch.Client(ctx), nil
}
