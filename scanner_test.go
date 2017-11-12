package sqltest_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"sqltest"
)

func TestProperty(t *testing.T) {
	cases := []struct {
		line, key, val string
		ok             bool
	}{
		{line: "-- this is a comment"},
		{line: "-- this: comment", key: "this", val: "comment", ok: true},
		{line: "-- key:"},
		{line: "-- key:  "},
		{line: "-- k12ey: val"},
		{line: "-- Key: val", key: "key", val: "val", ok: true},
		{line: "-- KEY: VAL", key: "key", val: "VAL", ok: true},
		{line: "-- key:val", key: "key", val: "val", ok: true},
		{line: "-- key:    val", key: "key", val: "val", ok: true},
		{line: "-- key:val     ", key: "key", val: "val", ok: true},
		{line: "-- key: val val    ", key: "key", val: "val val", ok: true},
	}

	for _, c := range cases {
		key, val, ok := sqltest.Property(c.line)
		if ok {
			if !c.ok {
				t.Errorf("Property(%q) should not have found a property", c.line)
			} else if key != c.key || val != c.val {
				t.Errorf("Property(%q) got %q=%q want %q=%q", c.line, key, val, c.key, c.val)
			}
		} else if c.ok {
			t.Errorf("Property(%q) should have found %q=%q", c.line, c.key, c.val)
		}
	}
}

func TestScanner(t *testing.T) {
	for i, c := range cases {
		scanner := sqltest.NewScanner(strings.NewReader(c.s))
		scanner.Filename = fmt.Sprintf("cases[%d]", i)
		for i := 0; ; i++ {
			tst, err := scanner.Scan()
			if err != nil {
				t.Errorf("Scan(cases[%d]) failed with %s", i, err)
			}
			if tst == nil {
				if i < len(c.tests) {
					t.Errorf("Scan(cases[%d]) did not return enough tests", i)
				}
				break
			}
			if i == len(c.tests) {
				t.Errorf("Scan(cases[%d]) returned too many tests", i)
				break
			}
			if !reflect.DeepEqual(*tst, c.tests[i]) {
				t.Errorf("Scan(cases[%d]) got %v want %v", i, *tst, c.tests[i])
			}
		}
	}
}

var cases = []struct {
	s     string
	tests []sqltest.Test
}{
	{
		s: `
-- this is a comment

-- empty line above`,
		tests: []sqltest.Test{},
	},
	{
		s: `-- a comment

SELECT * FROM tbl;
`,
		tests: []sqltest.Test{
			{
				Filename:   "cases[1]",
				LineNumber: 3,
				Comments:   []string{"-- a comment", ""},
				Stmts:      []string{"SELECT * FROM tbl;"},
				IsQuery:    true,
			},
		},
	},
	{
		s: `-- top of file
        
SELECT * FROM tbl
WHERE x = 12
ORDER BY y

-- another comment
INSERT INTO tbl VALUES
	(1, 2, 3),
	(4, 5, 6)

DELETE FROM tbl
WHERE c2 = 2
`,
		tests: []sqltest.Test{
			{
				Filename:   "cases[2]",
				LineNumber: 3,
				Comments:   []string{"-- top of file", "        "},
				Stmts:      []string{"SELECT * FROM tbl", "WHERE x = 12", "ORDER BY y"},
				IsQuery:    true,
			},
			{
				Filename:   "cases[2]",
				LineNumber: 8,
				Comments:   []string{"", "-- another comment"},
				Stmts: []string{"INSERT INTO tbl VALUES", "	(1, 2, 3),", "	(4, 5, 6)"},
				IsQuery: false,
			},
			{
				Filename:   "cases[2]",
				LineNumber: 12,
				Comments:   []string{""},
				Stmts:      []string{"DELETE FROM tbl", "WHERE c2 = 2"},
				IsQuery:    false,
			},
		},
	},
	{
		s: `-- a comment
-- stmt: query
-- identical:  true 
-- fail:false

SELECT * FROM tbl;
`,
		tests: []sqltest.Test{
			{
				Filename:   "cases[3]",
				LineNumber: 6,
				Comments: []string{"-- a comment", "-- stmt: query", "-- identical:  true ",
					"-- fail:false", ""},
				Stmts:   []string{"SELECT * FROM tbl;"},
				IsQuery: true,
				Properties: map[string]string{
					"stmt":      "query",
					"identical": "true",
					"fail":      "false",
				},
			},
		},
	},
}
