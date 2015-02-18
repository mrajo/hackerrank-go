package hrtest

import (
	"testing"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"io/ioutil"
	"path/filepath"
	"regexp"
)

type Case struct {
	input, output string
}

func GetPackage() string {
	return strings.Replace(filepath.Base(os.Args[0]), ".test", "", 1);
}

func FileExists(filename string) bool {
	if _, err := os.Stat(filename); err != nil {
		if os.IsNotExist(err) {
			return false;
		}
	}
	return true;
}

func NormalizeWhitespace(text string) string {
	re := regexp.MustCompile(`\r?\n`);
	return strings.TrimSpace(re.ReplaceAllString(text, "\n"));
}

func RunTests(t *testing.T) {
	var cases []Case;

	files, _ := filepath.Glob("./tests/*_in.txt");

	for _, file := range files {
		infile, _ := ioutil.ReadFile(file);
		outfile, _ := ioutil.ReadFile(strings.Replace(file, "_in", "_out", 1));
		cases = append(cases, Case{ NormalizeWhitespace(string(infile)), NormalizeWhitespace(string(outfile)) });
	}

	for n, c := range cases {
		var cmd *exec.Cmd

		if FileExists("./main.go") {
			cmd = exec.Command("go", "run", "./main.go");
		} else {
			cmd = exec.Command("go", "run", fmt.Sprintf("./%s.go", GetPackage()));
		}
		cmd.Stdin = strings.NewReader(c.input);
		out, err := cmd.Output();

		if err != nil {
			t.Errorf("Error running command: %s", err);
		}

		if strings.TrimSpace(string(out)) != c.output {
			t.Errorf("[Case %d] expected: %s, actual: %s", n + 1, strings.Replace(c.output, "\n", " ", -1), strings.Replace(string(out), "\n", " ", -1));
		}

		if (!t.Failed()) {
			t.Logf("[Case %d] passed", n + 1);
		}
	}
}