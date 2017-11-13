package sqltest_test

import (
	"fmt"
	"reflect"
	"testing"

	"sqltest"
)

type templateDialect struct{}

func (_ templateDialect) DriverName() string {
	return "template"
}

func (_ templateDialect) ColumnType(typ string, arg []int) string {
	if len(arg) > 0 {
		return fmt.Sprintf("%s(%d)", typ, arg[0])
	}
	return typ
}

func TestTemplateExecute(t *testing.T) {
	cases := []struct {
		tmpl, result string
		tctx, gctx   sqltest.TestContext
		fail         bool
	}{
		{
			tmpl:   "nothing changed",
			result: "nothing changed",
		},
		{
			tmpl:   "{{$.Dialect}}",
			result: "template",
		},
		{
			tmpl:   "{{Fail $.Test true}}",
			result: "",
			tctx:   sqltest.TestContext{Fail: true},
		},
		{
			tmpl:   "{{Fail $.Global true}}",
			result: "",
			gctx:   sqltest.TestContext{Fail: true},
		},
		{
			tmpl:   `{{$.Test.Statement}}{{Statement $.Test "SELECT"}}{{$.Test.Statement}}`,
			result: "SELECT",
			tctx:   sqltest.TestContext{Statement: "SELECT"},
		},
		{
			tmpl:   `{{eq $.Dialect "template" | Fail $.Test}}`,
			result: "",
			tctx:   sqltest.TestContext{Fail: true},
		},
		{
			tmpl:   "{{Fail $.Test}}",
			result: "",
			tctx:   sqltest.TestContext{Fail: true},
		},
		{
			tmpl:   "{{Sort .Test false}}",
			result: "",
			tctx:   sqltest.TestContext{NoSort: true},
		},
		{
			tmpl:   `{{$.ColumnType "BINARY"}}`,
			result: "BINARY",
		},
		{
			tmpl:   `{{$.ColumnType "VARBINARY" 12}}`,
			result: "VARBINARY(12)",
		},
	}

	for _, c := range cases {
		var tctx, gctx sqltest.TestContext
		r, err := sqltest.TemplateExecute(c.tmpl, &tctx, &gctx, templateDialect{})
		if c.fail {
			if err == nil {
				t.Errorf("TemplateExecute(%q) did not fail", c.tmpl)
			}
		} else {
			if err != nil {
				t.Errorf("TemplateExecute(%q) failed with %s", c.tmpl, err)
			} else {
				if r != c.result {
					t.Errorf("TemplateExecute(%q) got %s want %s", c.tmpl, r, c.result)
				}
				if !reflect.DeepEqual(tctx, c.tctx) {
					t.Errorf("TemplateExecute(%q): tctx got %v want %v", c.tmpl, tctx, c.tctx)
				}
				if !reflect.DeepEqual(gctx, c.gctx) {
					t.Errorf("TemplateExecute(%q): gctx got %v want %v", c.tmpl, gctx, c.gctx)
				}
			}
		}
	}
}
