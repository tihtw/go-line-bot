package lbot

// A config for set LINE service config
type Config struct {
	ChannelID     string
	ChannelSecret string
	MID           string
	ServerHost    string
	Debug         bool
}

const (
	DefaultServerHost = "https://trialbot-api.line.me"
)

func (c *Config) SetDefaults() {
	c.ServerHost = DefaultServerHost
	c.Debug = false
}

func NewConfig() *Config {
	return &Config{}
}
