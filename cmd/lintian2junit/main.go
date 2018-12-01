package main

import (
	"bufio"
	"encoding/xml"
	"fmt"
	"os"
	"strings"

	jtypes "github.com/rodrigodiez/go-junit/types"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)
	testcases := make([]*jtypes.Testcase, 0)
	for scanner.Scan() {
		s := scanner.Text()
		ss := strings.SplitN(s, ": ", 3)
		tc := jtypes.Testcase{
			Id: ss[1],
			Name: ss[1],
			Failures: []*jtypes.Failure{{
				Message: ss[2],
			}},
		}
		testcases = append(testcases, &tc)
	}

	testsuites :=  &jtypes.Testsuites{
		Id: "Lintian2JUnit",
		Testsuites: []*jtypes.Testsuite {{
			Id: "lintian",
			Tests: len(testcases),
			Failures: len(testcases),
			Testcases: testcases,
		}},
	}

	output, _ := xml.MarshalIndent(testsuites, "  ", "    ")
	fmt.Println(string(output[:]))

	if scanner.Err() != nil {
		// handle error.
		fmt.Println(scanner.Err())
		os.Exit(1)
	}
}
