//

package ghttp

import (
	"github.com/xrc360/golang/os/gcfg"
	"github.com/xrc360/golang/os/gview"
	"github.com/xrc360/golang/util/gconv"
	"github.com/xrc360/golang/util/gmode"
	"github.com/xrc360/golang/util/gutil"
)

// WriteTpl parses and responses given template file.
// The parameter `params` specifies the template variables for parsing.
func (r *Response) WriteTpl(tpl string, params ...gview.Params) error {
	r.Header().Set("Content-Type", contentTypeHtml)
	b, err := r.ParseTpl(tpl, params...)
	if err != nil {
		if !gmode.IsProduct() {
			r.Write("Template Parsing Error: " + err.Error())
		}
		return err
	}
	r.Write(b)
	return nil
}

// WriteTplDefault parses and responses the default template file.
// The parameter `params` specifies the template variables for parsing.
func (r *Response) WriteTplDefault(params ...gview.Params) error {
	r.Header().Set("Content-Type", contentTypeHtml)
	b, err := r.ParseTplDefault(params...)
	if err != nil {
		if !gmode.IsProduct() {
			r.Write("Template Parsing Error: " + err.Error())
		}
		return err
	}
	r.Write(b)
	return nil
}

// WriteTplContent parses and responses the template content.
// The parameter `params` specifies the template variables for parsing.
func (r *Response) WriteTplContent(content string, params ...gview.Params) error {
	r.Header().Set("Content-Type", contentTypeHtml)
	b, err := r.ParseTplContent(content, params...)
	if err != nil {
		if !gmode.IsProduct() {
			r.Write("Template Parsing Error: " + err.Error())
		}
		return err
	}
	r.Write(b)
	return nil
}

// ParseTpl parses given template file `tpl` with given template variables `params`
// and returns the parsed template content.
func (r *Response) ParseTpl(tpl string, params ...gview.Params) (string, error) {
	return r.Request.GetView().Parse(r.Request.Context(), tpl, r.buildInVars(params...))
}

// ParseTplDefault parses the default template file with params.
func (r *Response) ParseTplDefault(params ...gview.Params) (string, error) {
	return r.Request.GetView().ParseDefault(r.Request.Context(), r.buildInVars(params...))
}

// ParseTplContent parses given template file `file` with given template parameters `params`
// and returns the parsed template content.
func (r *Response) ParseTplContent(content string, params ...gview.Params) (string, error) {
	return r.Request.GetView().ParseContent(r.Request.Context(), content, r.buildInVars(params...))
}

// buildInVars merges build-in variables into `params` and returns the new template variables.
// TODO performance improving.
func (r *Response) buildInVars(params ...map[string]interface{}) map[string]interface{} {
	m := gutil.MapMergeCopy(r.Request.viewParams)
	if len(params) > 0 {
		gutil.MapMerge(m, params[0])
	}
	// Retrieve custom template variables from request object.
	sessionMap := gconv.MapDeep(r.Request.Session.MustData())
	gutil.MapMerge(m, map[string]interface{}{
		"Form":    r.Request.GetFormMap(),
		"Query":   r.Request.GetQueryMap(),
		"Request": r.Request.GetMap(),
		"Cookie":  r.Request.Cookie.Map(),
		"Session": sessionMap,
	})
	// Note that it should assign no Config variable to a template
	// if there's no configuration file.
	if v, _ := gcfg.Instance().Data(r.Request.Context()); len(v) > 0 {
		m["Config"] = v
	}
	return m
}
