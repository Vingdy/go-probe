package watcher

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"log"
	"math/rand"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"time"
)

type ServerPraviteIP struct {
	IpInfo []IPInfo
}

type IPInfo struct{
	Region string
	Number int
	IP []string
}

//var ServerPraviteIP []string

var defaultWatcher = NewWatcher()

var (
	servics_map     = make(map[string]ServiceList)
	service_locker  = new(sync.Mutex)
	consul_client   *api.Client
	my_service_id   string
	my_service_name string
	my_kv_key       string
)

type KVData struct {
	IP        string `json:"ip"`
	Load      int    `json:"load"`
	Timestamp int    `json:"ts"`
}

type Watcher struct {
	register ServiceList
}

type ServiceInfo struct {
	ServiceID string
	IP        string
	Port      int
	Load      int
	Timestamp int //load updated ts
}

type ServiceList []ServiceInfo

func NewWatcher() *Watcher {
	watcher := &Watcher{

	}

	my_service_name = "worker"

	//DoRegistService("localhost:8500", "127.0.0.1:54321", "worker", "127.0.0.1", 4300)
	//go watcher.watch()
	go DoDiscover("localhost:8500", "consul")
	//go discover("localhost:8500", "consul")
	//go WaitToUnRegistService()

	//go DoUpdateKeyValue("localhost:8500", "worker", "127.0.0.1", 4300)

	return watcher
}

func discover(consul_addr string, found_service string) {
	during := time.Duration(5)
	tick := time.NewTicker(time.Second * during)

	for range tick.C {
		discoverServices("localhost:8500",true, "worker")
	}
}

func discoverServices(addr string, healthyOnly bool, service_name string) {

	consulConf := api.DefaultConfig()
	consulConf.Address = addr
	//client, err := api.NewClient(consulConf)
	//CheckErr(err)

	/*services, _, err := client.Health().Service(&client, "", true, &api.QueryOptions{
		WaitIndex: w.lastIndex, // 同步点，这个调用将一直阻塞，直到有新的更新
	})
	if err != nil {
		//logrus.Warn("error retrieving instances from Consul: %v", err)
		log.Fatal("")
	}
	//w.lastIndex = metainfo.LastIndex

	addrs := map[string]struct{}{}
	for _, service := range services {
		addrs[net.JoinHostPort(service.Service.Address, strconv.Itoa(service.Service.Port))] = struct{}{}
	}*/
}
/*func (w *Watcher) watch() {
	during := time.Duration(5)
	tick := time.NewTicker(time.Second * during)
	for range tick.C {
		DiscoverServices("localhost:8500",true, "worker")
	}
}*/

func DoDiscover(consul_addr string, found_service string) {
	during := time.Duration(5)
	tick := time.NewTicker(time.Second * during)
	for {
		select {
		case <-tick.C:
			DiscoverServices(consul_addr, true, found_service)
		}
	}
}

func FindAdnTracking(node_name string) (string, bool) {
	s := strings.Split(node_name,"_")
	if s[0] != "vg" && s[0] != "fk" && s[0] != "sg" {
		return "", false
	}
	if s[1] != "adn" {
		return "", false
	}
	if !strings.Contains(s[2],"tracking") {
		return "", false
	}
	return s[0], true
}

func IsNewRegion(region string) bool {
	for _, ipinfo := range ServerIP.IpInfo{
		if region == ipinfo.Region {
			return false
		}
	}
	return true
}

var ServerIP = ServerPraviteIP{}

func DiscoverServices(addr string, healthyOnly bool, service_name string) {
	consulConf := api.DefaultConfig()
	consulConf.Address = addr
	client, err := api.NewClient(consulConf)
	CheckErr(err)

	nodes, _, err := client.Catalog().Nodes(&api.QueryOptions{})
	CheckErr(err)

	ServerIP=ServerPraviteIP{}
	for _, name := range nodes {
		region,find := FindAdnTracking(name.Node)
		//fmt.Println(region,find)
		if !find {
			continue
		}
		if IsNewRegion(region) {
			ServerIP.IpInfo=append(ServerIP.IpInfo,IPInfo{
				Region: region,
				Number: 0,
				IP:     nil,
			})
		}
		for ip:=0;ip<len(ServerIP.IpInfo);ip++ {
			if ServerIP.IpInfo[ip].Region == region {
				ServerIP.IpInfo[ip].IP = append(ServerIP.IpInfo[ip].IP, name.Address)
				ServerIP.IpInfo[ip].Number++
			}
		}
		//fmt.Println(region,ServerIP)
	}
	fmt.Println(ServerIP)
	service_locker.Lock()
	service_locker.Unlock()
}

func CheckErr(err error) {
	if err != nil {
		log.Printf("error: %v", err)
		os.Exit(1)
	}
}

func GetKeyValue(service_name string, ip string, port int) string {
	key := service_name + "/" + ip + ":" + strconv.Itoa(port)

	kv, _, err := consul_client.KV().Get(key, nil)
	if kv == nil {
		return ""
	}
	CheckErr(err)

	return string(kv.Value)
}

func DoRegistService(consul_addr string, monitor_addr string, service_name string, ip string, port int) {
	my_service_id = service_name + "-" + ip
	var tags []string
	service := &api.AgentServiceRegistration{
		ID:      my_service_id,
		Name:    service_name,
		Port:    port,
		Address: ip,
		Tags:    tags,
		Check: &api.AgentServiceCheck{
			HTTP:     "http://" + monitor_addr + "/status",
			Interval: "5s",
			Timeout:  "1s",
		},
	}

	client, err := api.NewClient(api.DefaultConfig())
	if err != nil {
		log.Fatal(err)
	}
	consul_client = client
	if err := consul_client.Agent().ServiceRegister(service); err != nil {
		log.Fatal(err)
	}
	log.Printf("Registered service %q in consul with tags %q", service_name, strings.Join(tags, ","))
}

func WaitToUnRegistService() {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Kill)
	<-quit

	if consul_client == nil {
		return
	}
	if err := consul_client.Agent().ServiceDeregister(my_service_id); err != nil {
		log.Fatal(err)
	}
}


func DoUpdateKeyValue(consul_addr string, service_name string, ip string, port int) {
	t := time.NewTicker(time.Second * 10)
	for {
		select {
		case <-t.C:
			StoreKeyValue(consul_addr, service_name, ip, port)
		}
	}
}

func StoreKeyValue(consul_addr string, service_name string, ip string, port int) {

	my_kv_key = my_service_name + "/" + ip + ":" + strconv.Itoa(port)

	var data KVData
	data.IP = ip
	data.Load = rand.Intn(100)
	data.Timestamp = int(time.Now().Unix())
	bys, _ := json.Marshal(&data)

	kv := &api.KVPair{
		Key:   my_kv_key,
		Flags: 0,
		Value: bys,
	}

	_, err := consul_client.KV().Put(kv, nil)
	CheckErr(err)
	fmt.Println(" store data key:", kv.Key, " value:", string(bys))
}