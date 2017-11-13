package sqltest

import (
	"bytes"
	"text/template"
)

type TestContext struct {
	Statement string
	Fail      bool
	NoSort    bool
}

type templateContext struct {
	Test    *TestContext
	Global  *TestContext
	Dialect string
	dialect Dialect
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

var templateFuncs = template.FuncMap{
	"Fail":      failFunc,
	"Statement": statementFunc,
	"Sort":      sortFunc,
}

func TemplateExecute(tmpl string, tctx, gctx *TestContext, dialect Dialect) (string, error) {
	t := template.New("sqltest").Funcs(templateFuncs)
	t, err := t.Parse(tmpl)

	tmplCtx := templateContext{Test: tctx, Global: gctx, Dialect: dialect.Name, dialect: dialect}
	var test bytes.Buffer
	err = t.Execute(&test, &tmplCtx)
	if err != nil {
		return tmpl, err
	}
	return test.String(), nil
}
