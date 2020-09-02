package main

import (
	"github.com/labstack/echo"
)

func CC() string {
	_ = echo.New()
	return "cc"
}
