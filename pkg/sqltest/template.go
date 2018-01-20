package sqltest

import (
	"bytes"
	"text/template"
)

type TemplateContext struct {
	Global TestContext
	Skip   bool
	funcs  template.FuncMap
}

type TestContext struct {
	Statement string
	Fail      bool
	NoSort    bool
}

func failFunc(ctx *TestContext, fail ...bool) string {
	if len(fail) > 0 {
		ctx.Fail = fail[0]
	} else {
		ctx.Fail = true
	}
	return ""
}

func statementFunc(ctx *TestContext, stmt string) string {
	ctx.Statement = stmt
	return ""
}

func sortFunc(ctx *TestContext, sort bool) string {
	ctx.NoSort = !sort
	return ""
}

func NewTemplateContext(dialect Dialect) *TemplateContext {
	tmplCtx := &TemplateContext{
		funcs: template.FuncMap{
			"Fail":      failFunc,
			"Statement": statementFunc,
			"Sort":      sortFunc,
			"Dialect": func() string {
				return dialect.DriverName()
			},
			"BINARY": func(arg ...int) string {
				if len(arg) == 0 {
					return dialect.ColumnType("BINARY")
				}
				return dialect.ColumnTypeArg("BINARY", arg[0])
			},
			"VARBINARY": func(arg ...int) string {
				if len(arg) == 0 {
					return dialect.ColumnType("VARBINARY")
				}
				return dialect.ColumnTypeArg("VARBINARY", arg[0])
			},
			"BLOB": func(arg ...int) string {
				if len(arg) == 0 {
					return dialect.ColumnType("BLOB")
				}
				return dialect.ColumnTypeArg("BLOB", arg[0])
			},
			"TEXT": func(arg ...int) string {
				if len(arg) == 0 {
					return dialect.ColumnType("TEXT")
				}
				return dialect.ColumnTypeArg("TEXT", arg[0])
			},
		},
	}

	tmplCtx.funcs["Skip"] = func() string {
		tmplCtx.Skip = true
		return ""
	}
	return tmplCtx
}

func (tmplCtx *TemplateContext) Execute(tmpl string) (string, TestContext, error) {
	t := template.New("sqltest").Funcs(tmplCtx.funcs)
	t, err := t.Parse(tmpl)

	tctx := tmplCtx.Global
	tmplData := struct {
		Test, Global *TestContext
	}{
		Test:   &tctx,
		Global: &tmplCtx.Global,
	}
	var test bytes.Buffer
	err = t.Execute(&test, &tmplData)
	if err != nil {
		return tmpl, TestContext{}, err
	}
	return test.String(), tctx, nil
}
