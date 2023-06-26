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
( url,  final_url,  title,  description,  og_description,  og_type,  og_image,  body,  score,  is_new,  status_code,  favicon,  label,  created_at,  updated_at)
VALUES
(:url, :final_url, :title, :description, :og_description, :og_type, :og_image, :body, :score, :is_new, :status_code, :favicon, :label, :created_at, :updated_at)
ON CONFLICT (url)
DO UPDATE SET
url = :url, final_url = :final_url, title = :title,
description = :description, og_description = :og_description, og_type = :og_type, og_image = :og_image,
body = :body, score = :score, is_new = :is_new, status_code = :status_code, favicon = :favicon,
label = :label, created_at = :created_at, updated_at = :updated_at
WHERE
((EXCLUDED.label != 0) AND (example.label != EXCLUDED.label)) OR
((example.status_code != 200) AND (EXCLUDED.status_code = 200))
;`, e)
	if err != nil {
		return err
	}
	tmp, err := r.FindExampleByUlr(e.Url)
	if err != nil {
		return err
	}
	e.Id = tmp.Id
	return nil
}

func (r *repository) UpdateScore(e *model.Example) error {
	if _, err := r.FindExampleByUlr(e.Url); err != nil {
		return err
	}
	if _, err := r.db.Exec(`UPDATE example SET score = $1, updated_at = $2 WHERE url = $3;`, e.Score, time.Now(), e.Url); err != nil {
		return err
	}
	return nil
}

func (r *repository) IncErrorCount(e *model.Example) error {
	errorCount, err := r.GetErrorCount(e)
	if err != nil {
		return err
	}
	if _, err := r.db.Exec(`UPDATE example SET error_count = $1, updated_at = $2 WHERE url = $3;`, errorCount+1, time.Now(), e.Url); err != nil {
		return err
	}
	return nil
}

func (r *repository) GetErrorCount(e *model.Example) (int, error) {
	example, err := r.FindExampleByUlr(e.Url)
	if err != nil {
		if err == exampleNotFoundError {
			re