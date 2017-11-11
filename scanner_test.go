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
				t.Errorf("Scan(cases[%d]) got %v want %v", i, tst, c.tests[i])
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
			{"cases[1]", 3, []string{"-- a comment", ""}, []string{"SELECT * FROM tbl;"}, true},
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
			{"cases[2]", 3, []string{"-- top of file", "        "},
				[]string{"SELECT * FROM tbl", "WHERE x = 12", "ORDER BY y"}, true},
			{"cases[2]", 8, []string{"", "-- another comment"},
				[]string{"INSERT INTO tbl VALUES", "	(1, 2, 3),", "	(4, 5, 6)"}, false},
			{"cases[2]", 12, []string{""}, []string{"DELETE FROM tbl", "WHERE c2 = 2"}, false},
		},
	},
}
