package grifts

import (
	"github.com/gobuffalo/buffalo"
	"github.com/t0m/code_resume/actions"
)

func init() {
	buffalo.Grifts(actions.App())
}
