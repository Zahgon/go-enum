//go:build go1.16
// +build go1.16

package generator

import (
	"embed"
)

//go:embed enum.tmpl enum_string.tmpl
var content embed.FS

func (g *Generator) addEmbeddedTemplates() { _ = "STUB: not implemented"; return }
