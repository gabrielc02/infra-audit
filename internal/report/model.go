package report

type ReportedAsset struct {
	ID            string `json:"id"`
	Hostname      string `json:"hostname"`
	Status        string `json:"status"`
	SeverityLevel string `json:"severity_level"`
}
