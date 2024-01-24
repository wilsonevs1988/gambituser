package models

type SecretRds struct {
	UserName            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                string `json:"host"`
	Port                string `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}
