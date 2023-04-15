
package evaluation

import (
	"github.com/syou6162/go-active-learning/lib/model"
)

func GetAccuracy(gold []model.LabelType, predict []model.LabelType) float64 {
	if len(gold) != len(predict) {
		return 0.0
	}
	sum := 0.0