package sqltestdb_test

import (
	"reflect"
	"testing"

	"github.com/leftmike/sqltest/sqltestdb"
)

type templateDialect struct {
	sqltestdb.DefaultDialect
}

func (_ templateDialect) DriverName() string {
	return "template"
}

func TestTemplateExecute(t *testing.T) {
	cases := []struct {
		tmpl, result string
		tctx, gctx   sqltestdb.TestContext
		fail         bool
	}{
		{
			tmpl:   "nothing changed",
			result: "nothing changed",
		},
		{
			tmpl:   "{{Dialect}}",
			result: "template",
		},
		{
			tmpl:   "{{Fail .Test true}}",
			result: "",
			tctx:   sqltestdb.TestContext{Fail: true},
		},
		{
			tmpl:   "{{Fail .Global true}}",
			result: "",
			gctx:   sqltestdb.TestContext{Fail: true},
		},
		{
			tmpl:   `{{.Test.Statement}}{{Statement .Test "SELECT"}}{{.Test.Statement}}`,
			result: "SELECT",
			tctx:   sqltestdb.TestContext{Statement: "SELECT"},
		},
		{
			tmpl:   `{{eq Dialect "template" | Fail .Test}}`,
			result: "",
			tctx:   sqltestdb.TestContext{Fail: true},
		},
		{
			tmpl:   "{{Fail .Test}}",
			result: "",
			tctx:   sqltestdb.TestContext{Fail: true},
		},
		{
			tmpl:   "{{Sort .Test false}}",
			result: "",
			tctx:   sqltestdb.TestContext{NoSort: true},
		},
		{
			tmpl:   "{{BINARY}}",
			result: "BINARY",
		},
		{
			tmpl:   "{{VARBINARY 12}}",
			result: "VARBINARY(12)",
		},
	}

	for _, c := range cases {
		tmplCtx := sqltestdb.NewTemplateContext(templateDialect{})
		r, tctx, err := tmplCtx.Execute(c.tmpl)
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
				if !reflect.DeepEqual(tmplCtx.Global, c.gctx) {
					t.Errorf("TemplateExecute(%q): gctx got %v want %v", c.tmpl, tmplCtx.Global,
						c.gctx)
				}
			}
		}
	}
}
