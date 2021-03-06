package sqltestdb

import (
	"bufio"
	"fmt"
	"io"
	"regexp"
	"strings"
)

type Scanner struct {
	Filename   string
	lineNumber int
	scanner    *bufio.Scanner
	line       string
	done       bool
}

type Test struct {
	Filename   string
	LineNumber int
	Test       string
	Statement  string
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{scanner: bufio.NewScanner(bufio.NewReader(r))}
}

func (s *Scanner) scanLine() bool {
	if s.done {
		return false
	}
	if !s.scanner.Scan() {
		s.done = true
		return false
	}
	s.line = s.scanner.Text()
	s.lineNumber += 1
	return true
}

func (s *Scanner) err() error {
	err := s.scanner.Err()
	if err != nil {
		if s.Filename != "" {
			return fmt.Errorf("%s:%d: %s", s.Filename, s.lineNumber, err)
		}
		return fmt.Errorf("%d: %s", s.lineNumber, err)
	}
	return nil
}

var (
	stmtRegexp    = regexp.MustCompile(`(?m)^[a-zA-Z]+`)
	twoStmtRegexp = regexp.MustCompile(`(?m)^[a-zA-Z]+ +[a-zA-Z]+`)
)

func emptyLine(s string) bool {
	for _, r := range s {
		if r != ' ' && r != '\t' && r != '\n' && r != '\r' {
			return false
		}
	}

	return true
}

func (s *Scanner) Scan() (*Test, error) {
	var tst Test

	if s.done {
		return nil, s.err()
	}

	if s.lineNumber == 0 && !s.scanLine() {
		return nil, s.err()
	}

	// Skip blank lines.

	for {
		if strings.Contains(s.line, "/*") {
			for {
				if strings.Contains(s.line, "*/") {
					break
				}
				if !s.scanLine() {
					return nil, s.err()
				}
			}
		} else if !emptyLine(s.line) {
			break
		}
		if !s.scanLine() {
			return nil, s.err()
		}
	}

	// Gather everything until the next blank line into tst.Test

	tst.Filename = s.Filename
	tst.LineNumber = s.lineNumber

	var sb strings.Builder
	for {
		if emptyLine(s.line) {
			break
		}
		if sb.Len() > 0 {
			sb.WriteByte('\n')
		}
		sb.WriteString(s.line)
		if !s.scanLine() {
			break
		}
	}
	tst.Test = sb.String()

	tst.Statement = strings.ToUpper(stmtRegexp.FindString(tst.Test))
	if tst.Statement == "ALTER" || tst.Statement == "CREATE" || tst.Statement == "DROP" {
		tst.Statement = strings.ToUpper(twoStmtRegexp.FindString(tst.Test))
	}

	return &tst, nil
}
