package models

type SecretRds struct {
	UserName            string `json:"username"`
	Password            string `json:"password"`
	Engine              string `json:"engine"`
	Host                int    `json:"host"`
	Port                int    `json:"port"`
	DbClusterIdentifier string `json:"dbClusterIdentifier"`
}
