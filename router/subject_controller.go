package router

import "MyWeb/framework"

func SubjectAddController(c *framework.Context) error {
	err := c.Json(200, "ok, SubjectAddController")
	if err != nil {
		return err
	}
	return nil
}

func SubjectListController(c *framework.Context) error {
	err := c.Json(200, "ok, SubjectListController")
	if err != nil {
		return err
	}
	return nil
}

func SubjectDelController(c *framework.Context) error {
	err := c.Json(200, "ok, SubjectDelController")
	if err != nil {
		return err
	}
	return nil
}

func SubjectUpdateController(c *framework.Context) error {
	err := c.Json(200, "ok, SubjectUpdateController")
	if err != nil {
		return err
	}
	return nil
}

func SubjectGetController(c *framework.Context) error {
	err := c.Json(200, "ok, SubjectGetController")
	if err != nil {
		return err
	}
	return nil
}

func SubjectNameController(c *framework.Context) error {
	err := c.Json(200, "ok, SubjectNameController")
	if err != nil {
		return err
	}
	return nil
}
