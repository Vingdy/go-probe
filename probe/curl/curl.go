package curl

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"os/exec"
	"probe/dingding"
	"probe/model"
	"probe/utils"
	"probe/watcher"
	"strings"
)



var ServerInfo = []model.ServerVersionInfo{}
/*var ServerInfo = []ServerVersionInfo{
	ServerVersionInfo{
		ServerRegion:"fk",
		Number:7,
		VersionInfo:[]ServerVersion{
			ServerVersion{
				Version:"version|3.17.9",
				VersionNum:7,
			},
		},
	},
}*/

func HandleVersion(re_version string) {
	version := strings.Split(re_version, " ")
	for region := 0; region < len(ServerInfo); region++ {
		for server_version := 0; server_version < len(ServerInfo[region].VersionInfo); server_version++ {
			if version[0] == ServerInfo[region].VersionInfo[server_version].Version {
				ServerInfo[region].VersionInfo[server_version].VersionNum++
				return
			}
		}
		ServerInfo[region].VersionInfo = append(ServerInfo[region].VersionInfo, model.ServerVersion{
			Version:    version[0],
			VersionNum: 1,
		})
	}
}

func GetIps(w http.ResponseWriter, r *http.Request) {
	ServerInfo = ServerInfo[:0]
	for _, region := range watcher.ServerIP.IpInfo {
		ServerInfo = append(ServerInfo, model.ServerVersionInfo{
			ServerRegion: region.Region,
			Number:       region.Number,
		})
	}
	for _, region := range watcher.ServerIP.IpInfo {
		for _, ip := range region.IP {
			re, err := Send(ip + "/version")
			if err == nil {
				HandleVersion(re)
			}
		}
	}
	fb, err := json.Marshal(ServerInfo)
	utils.CheckErr(err)
	dingding.SendPost(ServerInfo)
	fmt.Fprintln(w, string(fb))
}

func Exec_shell(cmd string, ip string) (string, error) {
	command := exec.Command(cmd,ip)
	var out bytes.Buffer
	var outerr bytes.Buffer
	command.Stdout = &out
	command.Stderr = &outerr
	command.Run()
	return out.String(), errors.New(outerr.String())
}

func Send(ip string) (string, error) {
	s, err := Exec_shell("curl", ip)
	if err != nil {
		return s, nil
	}
	return ip, err
}
