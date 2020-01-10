package model

type ServerVersion struct {
	Version    string `json:"version"`
	VersionNum int    `json:"version_num"`
}

type ServerVersionInfo struct {
	ServerRegion string          `json:"region"`
	Number       int             `json:"region_num"`
	VersionInfo  []ServerVersion `json:"version_info"`
}

type ServerPraviteIP struct {
	IpInfo []IPInfo
}

type IPInfo struct {
	Region string
	Number int
	IP     []string
}
