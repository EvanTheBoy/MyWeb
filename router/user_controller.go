package router

import "MyWeb/framework"

func UserLoginController(c *framework.Context) error {
	c.SetOkStatus().Json("ok, UserLoginController")
	return nil
}
