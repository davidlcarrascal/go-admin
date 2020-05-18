package constant

import (
	"github.com/davidlcarrascal/go-admin/modules/constant"
	"github.com/davidlcarrascal/go-admin/modules/language"
	"github.com/davidlcarrascal/go-admin/template/icon"
)

const (
	// PjaxHeader is default pjax http header key.
	PjaxHeader = constant.PjaxHeader

	// PjaxUrlHeader is default pjax url http header key.
	PjaxUrlHeader = constant.PjaxUrlHeader

	EditPKKey   = "__goadmin_edit_pk"
	DetailPKKey = "__goadmin_detail_pk"
	PrefixKey   = "__prefix"
)

var (
	DefaultErrorMsg = icon.Icon(icon.Warning, 2) + language.GetFromHtml("error") + `!`
)
