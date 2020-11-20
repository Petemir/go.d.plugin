package powerdns

// https://doc.powerdns.com/authoritative/http-api/statistics.html#objects
type (
	statisticMetrics []statisticMetric
	statisticMetric  struct {
		Name  string
		Type  string
		Value interface{}
	}
)
