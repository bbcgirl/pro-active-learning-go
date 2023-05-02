
package fetcher

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
	"strings"
	"time"

	"net/url"
	"unicode/utf8"

	"github.com/PuerkitoBio/goquery"
	goose "github.com/syou6162/GoOse"
)
