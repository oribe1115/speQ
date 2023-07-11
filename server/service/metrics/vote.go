package metrics

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

var voteCountsPerUser = promauto.NewGaugeVec(
	prometheus.GaugeOpts{
		Name: "speQ_vote_counts_per_user",
		Help: "How many voted for the users",
	}, []string{"contestant"})

func IncVoteCount(contestant string) {
	voteCountsPerUser.WithLabelValues(contestant).Inc()
}

func DecVoteCount(contestant string) {
	voteCountsPerUser.WithLabelValues(contestant).Dec()
}

func ClearVoteCount(newContestants []string) {
	voteCountsPerUser.Reset()

	for _, contestant := range newContestants {
		voteCountsPerUser.WithLabelValues(contestant).Set(0)
	}
}

func SetVoteCounts(votesPerContestant map[string]int) {
	voteCountsPerUser.Reset()

	for contestant, count := range votesPerContestant {
		voteCountsPerUser.WithLabelValues(contestant).Set(float64(count))
	}
}
