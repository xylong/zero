package gormc

// Config mysql配置
type Config struct {
	Separation   bool     `json:",default=false"`
	Master       string   `json:",optional"`
	Sources      []string `json:",optional"`
	Replicas     []string `json:",optional"`
	DSN          string   `json:",optional"`
	Debug        bool     `json:",default=false"`
	MaxIdleConns int      `json:",default=10"`
	MaxOpenConns int      `json:",default=100"`
}
