package repository

import (
	"bufio"
	"database/sql"
	"fmt"
	"io"
	"time"

	"github.com/lib/pq"
	"github.com/syou6162/go-active-learning/lib/feature"
	"github.com/syou6162/go-active-learning/lib/model"
	"github.com/syou6162/go-active-learning/lib/util/file"
)

var exampleNotFoundError = model.NotFoundError("example")

// データが存在しなければ追加
// データが存在する場合は、以下の場合にのみ更新する
// - ラベルが正例か負例に変更された
// - クロール対象のサイトが一時的に200以外のステータスで前回データが取得できなかった
func (r *repository) UpdateOrCreateExample(e *model.Example) error {
	now := time.Now()
	e.UpdatedAt = now
	_, err := r.db.NamedExec(`
INSERT INTO example
( url,  final_url,  title,  description,  og_description,  og_type,  og_image,  body,  sco