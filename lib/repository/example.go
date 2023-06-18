package repository

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"time"

	"github.com/lib/pq"
	"github.com/syou6162/go-active-learning