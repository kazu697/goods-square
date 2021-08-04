package renderer

import (
	"io"

	"github.com/unrolled/render"
)

func RenderJson(w io.Writer, status int, v interface{}) {
	r := render.New()
	r.JSON(w, status, v)
}
