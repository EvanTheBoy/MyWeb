package router

import "MyWeb/framework"

func UserLoginController(c *framework.Context) error {
	err := c.Json(200, "ok, UserLoginController")
	if err != nil {
		return err
	}
	return nil
}
