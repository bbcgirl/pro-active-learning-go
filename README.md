# pro-active-learning-go
[![CircleCI](https://circleci.com/gh/bbcgirl/pro-active-learning-go.svg?style=shield)](https://circleci.com/gh/bbcgirl/pro-active-learning-go)
[![Go Report Card](https://goreportcard.com/badge/github.com/bbcgirl/pro-active-learning-go)](https://goreportcard.com/report/github.com/bbcgirl/pro-active-learning-go)
[![Coverage Status](https://coveralls.io/repos/github/bbcgirl/pro-active-learning-go/badge.svg?branch=master)](https://coveralls.io/github/bbcgirl/pro-active-learning-go?branch=master)

pro-active-learning-go is a powerful command line annotation tool intended for binary classification tasks, built with Go. It employs a simple yet effective active learning algorithm to minimize annotation time.

# Install

```console
% go get github.com/bbcgirl/pro-active-learning-go
```

## Build from source

```console
% git clone https://github.com/bbcgirl/pro-active-learning-go.git
% cd pro-active-learning-go
% createdb pro-active-learning-go
% createdb pro-active-learning-go-test
% sql-migrate up -env=local
% sql-migrate up -env=test
% make build
```

# Usage
pro-active-learning-go comes with `annotate` (for suggesting new examples through active learning) and `diagnose` (for checking label conflicts in training data) modes. For detailed options, type `./pro-active-learning-go --help`.

## Annotation model
For detailed options, type `./pro-active-learning-go annotate --help`.

## Annotate new examples from command line interface
For detailed options, type `./pro-active-learning-go annotate cli --help`.

```console
% ./pro-active-learning-go annotate cli --open-url
Loading cache...
Label this example (Score: 0.600): http://srdk.rakuten.jp/ (それどこ)

p: Label this example as positive.
n: Label this example as negative.
s: Skip this example.
h: Show this help.
e: Exit.

Label this example (Score: 1.000): http://srdk.rakuten.jp/ (それどこ)
Labeled as negative
```

## Annotate new examples from slack
For detailed options, type `./pro-active-learning-go annotate cli --help`. To annotate new examples from slack, you need to create a slack bot, and obtain a token from [here](https://my.slack.com/services/new/bot). You can pass the token via environment variable (`SLACK_TOKEN`).

```console
% export SLACK_TOKEN=xoxb-SLACK-TOKEN
% ./pro-active-learning-go annotate slack --filter-status-code-ok --channel CHANNEL_ID
```

## Diagnosis model
For detailed options, type `./pro-active-learning-go diagnose --help`.

### Diagnose training data
This subcommand diagnoses label conflicts in training data. 'conflict' refers to a situation where an annotated label is '-1/1', but a predicted label by the model is '1/-1'. In such cases, for example, `http://www3.nhk.or.jp/news/` is a conflict case ('Label' is -1, but 'Score' is positive). You may need to gather such news articles to train a good classifier.

```console
% ./pro-active-learning-go diagnose label-conflict
Loading cache...
Index   Label   Score   URL     Title
0       -1      0.491   http://www3.nhk.or.jp/news/
1       1       0.491   http://blog.yuuk.io/
2       1       0.491   http://www.yasuhisay.info/
3       -1      -3.057  http://r.gnavi.co.jp/g-interview/       ぐるなび みんなのごはん
4       1       4.264   http://hakobe932.hatenablog.com/        hakobe-blog ♨
5       -1      -7.151  http://suumo.jp/town/   SUUMOタウン
6       -1      -26.321 https://www