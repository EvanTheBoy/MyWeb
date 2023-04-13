// Package gin Copyright 2021 jianfengye.  All rights reserved.
// Use of this source code is governed by a MIT style
// license that can be found in the LICENSE file.
package gin

import (
	"context"
)

func (c *Context) BaseContext() context.Context {
	return c.Request.Context()
}
