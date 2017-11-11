package sqltest

import (
	"bufio"
	"fmt"
	"io"
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
	Comments   []string
	Stmts      []string
	IsQuery    bool
}

func NewScanner(r io.Reader) *Scanner {
	return &Scanner{scanner: bufio.NewScanner(r)}
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

func (s *Scanner) notStatement() bool {
	return strings.TrimSpace(s.line) == "" ||
		(len(s.line) >= 2 && s.line[0] == '-' && s.line[1] == '-')
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

func (s *Scanner) Scan() (*Test, error) {
	var tst Test

	if s.done {
		return nil, s.err()
	}

	if s.lineNumber == 0 && !s.scanLine() {
		return nil, s.err()
	}

	// Gather comments (lines starting with --)  and blank lines into s.Comments

	for {
		if s.notStatement() {
			tst.Comments = append(tst.Comments, s.line)
		} else {
			break
		}
		if !s.scanLine() {
			return nil, s.err()
		}
	}

	// Gather everything which is not a comment (line starting with --) or a blank line into
	// s.Stmts.

	tst.Filename = s.Filename
	tst.LineNumber = s.lineNumber

	for {
		if s.notStatement() {
			break
		}
		tst.Stmts = append(tst.Stmts, s.line)
		if !s.scanLine() {
			break
		}
	}

	if strings.ToUpper(strings.Fields(tst.Stmts[0])[0]) == "SELECT" {
		tst.IsQuery = true
	}

	return &tst, nil
}
