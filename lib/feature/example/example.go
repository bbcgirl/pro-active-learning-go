
package example_feature

import (
	"net/url"
	"strings"
	"sync"
	"unicode"

	"github.com/ikawaha/kagome/tokenizer"
	"github.com/jdkato/prose/tag"
	"github.com/jdkato/prose/tokenize"
	"github.com/syou6162/go-active-learning/lib/feature"
)

var excludingWordList = []string{
	`:`, `;`,
	`,`, `.`,
	`"`, `''`,
	`+`, `-`, `*`, `/`, `|`, `++`, `--`,
	`[`, `]`,
	`{`, `}`,
	`(`, `)`,
	`<`, `>`,
	`「`, `」`,