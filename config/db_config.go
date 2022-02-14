package config

/*
	Oracle ve Sqlite3 bağlantıları için gerekli
	bilgiler tutulur
*/

type OracleConfig struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	Port        int    `json:"port"`
	ServiceName string `json:"servicename"`
}

type SqliteConfig struct {
	DbPath string `json:"path"`
}
