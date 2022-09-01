package utils

import (
	"html/template"

	"github.com/haoran-mc/gcbbs/pkg/utils/time"
	"github.com/haoran-mc/gcbbs/pkg/utils/view"
)

// GetTemplateFuncMap 获取模板中使用的函数
func GetTemplateFuncMap() template.FuncMap {
	return template.FuncMap{
		"DiffForHumans":    time.DiffForHumans,
		"ToDateTimeString": time.ToDateTimeString,
		"Html":             view.Html,
		"RemindName":       view.RemindName,
	}
}
