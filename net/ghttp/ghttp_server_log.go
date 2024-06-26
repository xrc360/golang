package ghttp

import (
	"fmt"

	"github.com/xrcn/cg/errors/gerror"
	"github.com/xrcn/cg/internal/instance"
	"github.com/xrcn/cg/os/glog"
	"github.com/xrcn/cg/text/gstr"
)

// handleAccessLog handles the access logging for server.
func (s *Server) handleAccessLog(r *Request) {
	if !s.IsAccessLogEnabled() {
		return
	}
	var (
		scheme            = "http"
		proto             = r.Header.Get("X-Forwarded-Proto")
		loggerInstanceKey = fmt.Sprintf(`Acccess Logger Of Server:%s`, s.instance)
	)

	if r.TLS != nil || gstr.Equal(proto, "https") {
		scheme = "https"
	}
	content := fmt.Sprintf(
		`%d "%s %s %s %s %s" %.3f, %s, "%s", "%s"`,
		r.Response.Status, r.Method, scheme, r.Host, r.URL.String(), r.Proto,
		float64(r.LeaveTime-r.EnterTime)/1000,
		r.GetClientIp(), r.Referer(), r.UserAgent(),
	)
	logger := instance.GetOrSetFuncLock(loggerInstanceKey, func() interface{} {
		l := s.Logger().Clone()
		l.SetFile(s.config.AccessLogPattern)
		l.SetStdoutPrint(s.config.LogStdout)
		l.SetLevelPrint(false)
		return l
	}).(*glog.Logger)
	logger.Print(r.Context(), content)
}

// handleErrorLog handles the error logging for server.
func (s *Server) handleErrorLog(err error, r *Request) {
	// It does nothing if error logging is custom disabled.
	if !s.IsErrorLogEnabled() {
		return
	}
	var (
		code              = gerror.Code(err)
		scheme            = "http"
		codeDetail        = code.Detail()
		proto             = r.Header.Get("X-Forwarded-Proto")
		loggerInstanceKey = fmt.Sprintf(`Error Logger Of Server:%s`, s.instance)
		codeDetailStr     string
	)
	if r.TLS != nil || gstr.Equal(proto, "https") {
		scheme = "https"
	}
	if codeDetail != nil {
		codeDetailStr = gstr.Replace(fmt.Sprintf(`%+v`, codeDetail), "\n", " ")
	}
	content := fmt.Sprintf(
		`%d "%s %s %s %s %s" %.3f, %s, "%s", "%s", %d, "%s", "%+v"`,
		r.Response.Status, r.Method, scheme, r.Host, r.URL.String(), r.Proto,
		float64(r.LeaveTime-r.EnterTime)/1000,
		r.GetClientIp(), r.Referer(), r.UserAgent(),
		code.Code(), code.Message(), codeDetailStr,
	)
	if s.config.ErrorStack {
		if stack := gerror.Stack(err); stack != "" {
			content += "\nStack:\n" + stack
		} else {
			content += ", " + err.Error()
		}
	} else {
		content += ", " + err.Error()
	}
	logger := instance.GetOrSetFuncLock(loggerInstanceKey, func() interface{} {
		l := s.Logger().Clone()
		l.SetStack(false)
		l.SetFile(s.config.ErrorLogPattern)
		l.SetStdoutPrint(s.config.LogStdout)
		l.SetLevelPrint(false)
		return l
	}).(*glog.Logger)
	logger.Error(r.Context(), content)
}
