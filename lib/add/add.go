package add

import (
	"fmt"
	"log"
	"time"

	"os"

	mkr "github.com/mackerelio/mackerel-client-go"
	"github.com/syou6162/go-active-learning/lib/classifier"
	"github.com/syou6162/go-active-learning/lib/hatena_bookmark"
	"github.com/syou6162/go-active-learning/lib/service"
	"github.com/syou