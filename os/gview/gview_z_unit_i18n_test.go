package gview_test

import (
	"context"
	"testing"

	"github.com/xrcn/cg/debug/gdebug"
	"github.com/xrcn/cg/frame/g"
	"github.com/xrcn/cg/i18n/gi18n"
	"github.com/xrcn/cg/os/gctx"
	"github.com/xrcn/cg/os/gfile"
	"github.com/xrcn/cg/os/gview"
	"github.com/xrcn/cg/test/gtest"
)

func Test_I18n(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		content := `{{.name}} says "{#hello}{#world}!"`
		expect1 := `john says "你好世界!"`
		expect2 := `john says "こんにちは世界!"`
		expect3 := `john says "{#hello}{#world}!"`

		g.I18n().SetPath(gtest.DataPath("i18n"))

		g.I18n().SetLanguage("zh-CN")
		result1, err := g.View().ParseContent(context.TODO(), content, g.Map{
			"name": "john",
		})
		t.AssertNil(err)
		t.Assert(result1, expect1)

		g.I18n().SetLanguage("ja")
		result2, err := g.View().ParseContent(context.TODO(), content, g.Map{
			"name": "john",
		})
		t.AssertNil(err)
		t.Assert(result2, expect2)

		g.I18n().SetLanguage("none")
		result3, err := g.View().ParseContent(context.TODO(), content, g.Map{
			"name": "john",
		})
		t.AssertNil(err)
		t.Assert(result3, expect3)
	})
	gtest.C(t, func(t *gtest.T) {
		content := `{{.name}} says "{#hello}{#world}!"`
		expect1 := `john says "你好世界!"`
		expect2 := `john says "こんにちは世界!"`
		expect3 := `john says "{#hello}{#world}!"`

		g.I18n().SetPath(gdebug.CallerDirectory() + gfile.Separator + "testdata" + gfile.Separator + "i18n")

		result1, err := g.View().ParseContent(context.TODO(), content, g.Map{
			"name":         "john",
			"I18nLanguage": "zh-CN",
		})
		t.AssertNil(err)
		t.Assert(result1, expect1)

		result2, err := g.View().ParseContent(context.TODO(), content, g.Map{
			"name":         "john",
			"I18nLanguage": "ja",
		})
		t.AssertNil(err)
		t.Assert(result2, expect2)

		result3, err := g.View().ParseContent(context.TODO(), content, g.Map{
			"name":         "john",
			"I18nLanguage": "none",
		})
		t.AssertNil(err)
		t.Assert(result3, expect3)
	})
	// gi18n manager is nil
	gtest.C(t, func(t *gtest.T) {
		content := `{{.name}} says "{#hello}{#world}!"`
		expect1 := `john says "{#hello}{#world}!"`

		g.I18n().SetPath(gdebug.CallerDirectory() + gfile.Separator + "testdata" + gfile.Separator + "i18n")

		view := gview.New()
		view.SetI18n(nil)
		result1, err := view.ParseContent(context.TODO(), content, g.Map{
			"name":         "john",
			"I18nLanguage": "zh-CN",
		})
		t.AssertNil(err)
		t.Assert(result1, expect1)
	})
	// SetLanguage in context
	gtest.C(t, func(t *gtest.T) {
		content := `{{.name}} says "{#hello}{#world}!"`
		expect1 := `john says "你好世界!"`
		ctx := gctx.New()
		g.I18n().SetPath(gdebug.CallerDirectory() + gfile.Separator + "testdata" + gfile.Separator + "i18n")
		ctx = gi18n.WithLanguage(ctx, "zh-CN")
		t.Log(gi18n.LanguageFromCtx(ctx))

		view := gview.New()

		result1, err := view.ParseContent(ctx, content, g.Map{
			"name": "john",
		})
		t.AssertNil(err)
		t.Assert(result1, expect1)
	})

}
