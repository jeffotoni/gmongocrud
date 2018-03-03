package template

import (
	"html/template"

	"github.com/jeffotoni/mercuriuscrud/lib/context"
)

// FuncMaps to view
func FuncMaps() []template.FuncMap {
	return []template.FuncMap{
		map[string]interface{}{
			"Tr": context.I18n,
		}}
}
