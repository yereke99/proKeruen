package models

type Settings struct {
	AppParams      Params           `json:"app"`
	PostgresParams PostgresSettings `json:"pg_params"`
}

type Params struct {
	ServerName string `json:"server_name"`
	PortRun    string `json:"port_run"`
	LogFile    string `json:"log_file"`
	ServerURL  string `json:"server_url"`
}

type PostgresSettings struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Host     string `json:"host"`
	Port     string `json:"port"`
	DataBase string `json:"database"`
}
