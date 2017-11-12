package sqltest_test

import (
	"reflect"
	"testing"

	"sqltest"
)

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
			result: "test",
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
			tmpl:   `{{eq $.Dialect "test" | Fail $.Test}}`,
			result: "",
			tctx:   sqltest.TestContext{Fail: true},
		},
		{
			tmpl:   "{{Fail $.Test}}",
			result: "",
			tctx:   sqltest.TestContext{Fail: true},
		},
	}

	for _, c := range cases {
		var tctx, gctx sqltest.TestContext
		r, err := sqltest.TemplateExecute(c.tmpl, &tctx, &gctx, sqltest.Dialect{"test"})
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
