package top_accessed_example

import (
	"bufio"
	"fmt"
	"strconv"

	"os"

	"github.com/syou6162/go-active-learning/lib/service"
	"github.com/urfave/cli"
)

func parseLine(line string) (int, error) {
	exampleId, err := strconv.ParseInt(line, 10, 0)
	if err != nil {
		return 0, fmt.Errorf("Invalid line: %s", line)
	}
	return int(exampleId), nil
}

func readTopAccessedExampleIds(filename string) ([]int, error) {
	fp, err := os.Open(filename)
	defer fp.Close()
	if err != nil {
		return nil, err
	}

	exampleIds := ma