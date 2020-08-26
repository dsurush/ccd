package settings

// Settings app settings
type Settings struct {
	AppParams   Params      `json:"app"`
	CcsDbParams CcsDbParams `json:"ccsDb"`
}

// Params contains params of server meta data
type Params struct {
	ServerName string `json:"serverName"`
	PortRun    int    `json:"portRun"`
	LogFile    string `json:"logFile"`
}

// ccdDbParams conteins params of postgresql services server
type CcsDbParams struct {
	Server   string `json:"server"`
	Port     int    `json:"port"`
	User     string `json:"user"`
	Password string `json:"password"`
	Database string `json:"database"`
}
