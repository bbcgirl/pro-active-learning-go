package related_example

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"os"

	"github.com/syou6162/go-active-learning/lib/model"
	"github.com/syou6162/go-active-learning/lib/service"
	"github.com/urfave/cli"
)

func parseLine(line string) (int, int, error) {
	tokens := strings.Split(line, "\t")
	if len(tokens) == 2 {
		exampleId, _ := strconv.ParseInt(tokens[0], 10, 0)
		relatedExampleId, _ := strconv.ParseInt(tokens[1], 10, 0)
		return int(exampleId), int(relatedExampleId), nil
	}
	return 0, 0, fmt.Errorf("Invalid line: %s", line)
}

func readRelatedExamples(filename string) ([]*model.RelatedExamples, error) {
	fp, err := os.Open(filename)
	defer fp.Close()
	if err != nil {
		return nil, err
	}

	exampleId2RelatedExampleIds := make(map[int][]int)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		exampleId, relatedExampleId, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		if _, ok := exampleId2RelatedExampleIds[exampleId]; ok {
			exampleId2RelatedExampleIds[exampleId] = append(exampleId2RelatedExampleIds[exampleId], relatedExampleId)
		} else {
			exampleId2RelatedExampleIds[exampleId] = []int{relatedExampleId}
		}
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	result := make([]*model.RelatedExamples, 0)
	for exampleId, relatedExampleIds := range exampleId2RelatedExampleId