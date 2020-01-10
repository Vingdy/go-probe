package watcher

import (
	"github.com/hashicorp/consul/api"
	"probe/model"
	"probe/utils"
	"strings"
	"time"
)

var (
	region_filter = []string{
		"vg",
		"sg",
		"fk",
	}
	ad_type_filter = []string{
		"adn",
	}
	server_filter = []string{
		"tktracking",
		"tktrackingStaticPostback",
		"tktrackingForRta",
		"tktrackingInstallService",
	}
)

//var DefaultWatcher = NewWatcher()



type Watcher struct {
}

//var ServerIP = ServerPraviteIP{}
var ServerIP = model.ServerPraviteIP{IpInfo:[]model.IPInfo{
		{
			"fk",
			3,
			[]string{"54.175.192.163","34.233.128.129","34.234.66.169",},
		},
},}

func NewWatcher() *Watcher {
	watcher := &Watcher{}

	go watcher.DoDiscover("localhost:8500", "consul")

	return watcher
}

func (w *Watcher) DoDiscover(consul_addr string, found_service string) {
	during := time.Duration(10)
	tick := time.NewTicker(time.Second * during)
	for {
		select {
		case <-tick.C:
			w.DiscoverServices(consul_addr, true, found_service)
		}
	}
}

func FindAdnTracking(node_name string) (string, string, bool) {
	s := strings.Split(node_name, "_")
	for _, region := range region_filter {
		if region == s[0] {
			return "", "", false
		}
	}
	for _, ad_type := range ad_type_filter {
		if ad_type == s[1] {
			return "", "", false
		}
	}
	adn_type := strings.Split(s[2], ":")

	for _, server := range server_filter {
		if server == adn_type[0] {
			return "", "", false
		}
	}
	return s[0], adn_type[1], true
}

func IsNewRegion(region string) bool {
	for _, ipinfo := range ServerIP.IpInfo {
		if region == ipinfo.Region {
			return false
		}
	}
	return true
}

func (w *Watcher) DiscoverServices(addr string, healthyOnly bool, service_name string) {
	consulConf := api.DefaultConfig()
	consulConf.Address = addr
	client, err := api.NewClient(consulConf)
	utils.CheckErr(err)

	services, _, err := client.Health().Service("adn_tracking", "", true, &api.QueryOptions{})
	utils.CheckErr(err)

	ServerIP = model.ServerPraviteIP{}
	for _, service := range services {

		region, real_ip, find := FindAdnTracking(service.Node.Node)
		if !find {
			continue
		}
		if IsNewRegion(region) {
			ServerIP.IpInfo = append(ServerIP.IpInfo, model.IPInfo{
				Region: region,
				Number: 0,
				IP:     nil,
			})
		}
		for ip := 0; ip < len(ServerIP.IpInfo); ip++ {
			if ServerIP.IpInfo[ip].Region == region {
				ServerIP.IpInfo[ip].IP = append(ServerIP.IpInfo[ip].IP, real_ip)
				ServerIP.IpInfo[ip].Number++
			}
		}
	}
}


