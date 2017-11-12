package sqltest_test

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"sqltest"
)

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
		s: `-- a comment
SELECT * FROM tbl;
`,
		tests: []sqltest.Test{
			{
				Filename:   "cases[0]",
				LineNumber: 1,
				Test: `-- a comment
SELECT * FROM tbl;`,
				IsQuery: true,
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
				Filename:   "cases[1]",
				LineNumber: 1,
				Test: `-- top of file
SELECT * FROM tbl
WHERE x = 12
ORDER BY y`,
				IsQuery: true,
			},
			{
				Filename:   "cases[1]",
				LineNumber: 6,
				Test: `-- another comment
INSERT INTO tbl VALUES
	(1, 2, 3),
	(4, 5, 6)`,
				IsQuery: false,
			},
			{
				Filename:   "cases[1]",
				LineNumber: 11,
				Test: `DELETE FROM tbl
WHERE c2 = 2`,
				IsQuery: false,
			},
		},
	},
}
