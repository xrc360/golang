// Package gi18n implements internationalization and localization.
package gi18n

import (
	"context"

	"github.com/xrc360/golang/os/gctx"
)

const (
	ctxLanguage gctx.StrKey = "I18nLanguage"
)

// WithLanguage append language setting to the context and returns a new context.
func WithLanguage(ctx context.Context, language string) context.Context {
	if ctx == nil {
		ctx = context.TODO()
	}
	return context.WithValue(ctx, ctxLanguage, language)
}

// LanguageFromCtx retrieves and returns language name from context.
// It returns an empty string if it is not set previously.
func LanguageFromCtx(ctx context.Context) string {
	if ctx == nil {
		return ""
	}
	v := ctx.Value(ctxLanguage)
	if v != nil {
		return v.(string)
	}
	return ""
}
