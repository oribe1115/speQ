package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var scorePerUser = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "speQ_score_per_user",
		Help: "Current scores of users",
	}, []string{"contestant"})

func NewScore(contestant string, score float64) {
	scorePerUser.WithLabelValues(contestant).Set(score)
}
