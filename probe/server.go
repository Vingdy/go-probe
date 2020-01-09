package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"probe/watcher"
	"time"
)

var ip = []Version{
	{IP:"111.230.186.23"},
	{IP:"47.96.233.11"},
}

type Version struct {
	HostName string `json:"hostname"`
	IP string `json:"ip"`
	version string `json:"version"`
}

type ServerVersion struct {
	 Version string `json:"version"`
	 VersionNum int `json:"version_num"`
}

type ServerInfo struct {
	ServerRegion string `json:"region"`
	Number int `json:"region_num"`
	VersionInfo []ServerVersion `json:"version_info"`
}

type FB struct{
	feedback []ServerVersion `json:"serverinfo"`
}

var ips = []Version{
	{HostName:"fk_adn_tktracking",IP:"18.184.151.114"},
	{HostName:"fk_adn_tktracking",IP:"18.184.188.80"},
	{HostName:"fk_adn_tktracking",IP:"18.184.213.26"},
	{HostName:"fk_adn_tktracking",IP:"18.184.4.108"},
	{HostName:"fk_adn_tktracking",IP:"18.184.4.76"},
	{HostName:"fk_adn_tktracking",IP:"18.195.146.141"},
	{HostName:"fk_adn_tktracking",IP:"18.196.235.75"},
	{HostName:"fk_adn_tktracking",IP:"18.197.151.122"},
	{HostName:"fk_adn_tktracking",IP:"18.197.35.18"},
	{HostName:"fk_adn_tktracking",IP:"3.120.191.64"},
	{HostName:"fk_adn_tktracking",IP:"3.122.231.89"},
	{HostName:"fk_adn_tktracking",IP:"3.123.24.245"},
	{HostName:"fk_adn_tktracking",IP:"3.125.4.23"},
	{HostName:"fk_adn_tktracking",IP:"35.156.81.86"},
	{HostName:"fk_adn_tktracking",IP:"35.159.31.9"},
	{HostName:"fk_adn_tktracking",IP:"52.58.166.130"},
	{HostName:"fk_adn_tktracking",IP:"54.93.191.30"},
	{HostName:"fk_adn_tktrackingStaticPostback",IP:"35.157.205.186"},
	{HostName:"fk_adn_tktrackingStaticPostback",IP:"35.157.218.38"},
	{HostName:"sg_adn_tktracking",IP:"13.228.72.88"},
	{HostName:"sg_adn_tktracking",IP:"13.229.107.79"},
	{HostName:"sg_adn_tktracking",IP:"13.229.212.123"},
	{HostName:"sg_adn_tktracking",IP:"13.229.218.203"},
	{HostName:"sg_adn_tktracking",IP:"13.229.230.183"},
	{HostName:"sg_adn_tktracking",IP:"13.229.94.42"},
	{HostName:"sg_adn_tktracking",IP:"13.250.107.82"},
	{HostName:"sg_adn_tktracking",IP:"13.250.17.13"},
	{HostName:"sg_adn_tktracking",IP:"13.250.21.154"},
	{HostName:"sg_adn_tktracking",IP:"13.250.29.16"},
	{HostName:"sg_adn_tktracking",IP:"18.140.65.241"},
	{HostName:"sg_adn_tktracking",IP:"18.141.13.144"},
	{HostName:"sg_adn_tktracking",IP:"18.141.138.172"},
	{HostName:"sg_adn_tktracking",IP:"3.0.103.40"},
	{HostName:"sg_adn_tktracking",IP:"3.0.182.229"},
	{HostName:"sg_adn_tktracking",IP:"3.0.51.121"},
	{HostName:"sg_adn_tktracking",IP:"3.1.5.146"},
	{HostName:"sg_adn_tktracking",IP:"52.77.20.174"},
	{HostName:"sg_adn_tktracking",IP:"52.77.232.135"},
	{HostName:"sg_adn_tktracking",IP:"52.77.255.120"},
	{HostName:"sg_adn_tktracking",IP:"54.169.43.120"},
	{HostName:"sg_adn_tktracking",IP:"54.169.74.207"},
	{HostName:"sg_adn_tktracking",IP:"54.254.250.44"},
	{HostName:"sg_adn_tktracking",IP:"54.255.219.53"},
	{HostName:"sg_adn_tktracking",IP:"54.255.222.126"},
	{HostName:"sg_adn_tktrackingStaticPostback",IP:"13.228.153.248"},
	{HostName:"sg_adn_tktrackingStaticPostback",IP:"52.76.52.15"},
	{HostName:"vg_adn_tktracking",IP:"100.24.50.62"},
	{HostName:"vg_adn_tktracking",IP:"18.206.190.64"},
	{HostName:"vg_adn_tktracking",IP:"18.209.94.44"},
	{HostName:"vg_adn_tktracking",IP:"18.215.156.212"},
	{HostName:"vg_adn_tktracking",IP:"184.72.94.227"},
	{HostName:"vg_adn_tktracking",IP:"3.85.23.196"},
	{HostName:"vg_adn_tktracking",IP:"3.87.245.37"},
	{HostName:"vg_adn_tktracking",IP:"3.91.35.179"},
	{HostName:"vg_adn_tktracking",IP:"34.192.196.101"},
	{HostName:"vg_adn_tktracking",IP:"34.204.44.192"},
	{HostName:"vg_adn_tktracking",IP:"34.204.98.129"},
	{HostName:"vg_adn_tktracking",IP:"34.228.43.64"},
	{HostName:"vg_adn_tktracking",IP:"34.229.146.189"},
	{HostName:"vg_adn_tktracking",IP:"34.229.176.242"},
	{HostName:"vg_adn_tktracking",IP:"34.230.25.184"},
	{HostName:"vg_adn_tktracking",IP:"34.230.67.102"},
	{HostName:"vg_adn_tktracking",IP:"35.175.197.155"},
	{HostName:"vg_adn_tktracking",IP:"52.206.242.134"},
	{HostName:"vg_adn_tktracking",IP:"52.90.198.72"},
	{HostName:"vg_adn_tktracking",IP:"52.91.109.144"},
	{HostName:"vg_adn_tktracking",IP:"54.145.154.124"},
	{HostName:"vg_adn_tktracking",IP:"54.145.24.226"},
	{HostName:"vg_adn_tktracking",IP:"54.147.22.63"},
	{HostName:"vg_adn_tktracking",IP:"54.152.152.111"},
	{HostName:"vg_adn_tktracking",IP:"54.152.192.110"},
	{HostName:"vg_adn_tktracking",IP:"54.160.203.223"},
	{HostName:"vg_adn_tktracking",IP:"54.161.40.138"},
	{HostName:"vg_adn_tktracking",IP:"54.162.152.199"},
	{HostName:"vg_adn_tktracking",IP:"54.166.168.199"},
	{HostName:"vg_adn_tktracking",IP:"54.172.51.171"},
	{HostName:"vg_adn_tktracking",IP:"54.175.72.156"},
	{HostName:"vg_adn_tktracking",IP:"54.196.15.160"},
	{HostName:"vg_adn_tktracking",IP:"54.196.209.157"},
	{HostName:"vg_adn_tktracking",IP:"54.221.130.246"},
	{HostName:"vg_adn_tktracking",IP:"54.235.55.225"},
	{HostName:"vg_adn_tktracking",IP:"54.81.187.245"},
	{HostName:"vg_adn_tktracking",IP:"54.83.120.171"},
	{HostName:"vg_adn_tktracking",IP:"54.84.7.122"},
	{HostName:"vg_adn_tktracking",IP:"54.86.112.54"},
	{HostName:"vg_adn_tktracking",IP:"54.89.242.129"},
	{HostName:"vg_adn_tktracking",IP:"54.90.190.60"},
	{HostName:"vg_adn_tktracking",IP:"54.90.214.46"},
	{HostName:"vg_adn_tktrackingForRta",IP:"54.209.54.80"},
	{HostName:"vg_adn_tktrackingInstallService",IP:"18.215.114.111"},
	{HostName:"vg_adn_tktrackingInstallService",IP:"18.215.97.1"},
	{HostName:"vg_adn_tktrackingInstallService",IP:"34.200.44.88"},
	{HostName:"vg_adn_tktrackingStaticPostback",IP:"34.199.85.129"},
	{HostName:"vg_adn_tktrackingStaticPostback",IP:"34.224.140.231"},
}

/*type LongPollHandler struct {
}

var c = make(chan int)

var count = 0

func (a LongPollHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	go func() {
		for {
			fmt.Println(time.Now(), " begin...")
			count++
			go func() {
				time.Sleep(5 * time.Second)
				c <- count
			}()
			fmt.Println(time.Now(), " wait...")
			time.Sleep(5 * time.Second)
		}
	}()
	for {
		select {
		case <-c:
			{
				fmt.Println(time.Now(), " end...")
				resp.Write([]byte(strconv.Itoa(<-c)))
			}
		}
	}
}*/

/*type ServiceInfo struct {
	ServiceID string
	IP        string
	Port      int
	Load      int
	Timestamp int //load updated ts
}
type ServiceList []ServiceInfo

type KVData struct {
	Load      int `json:"load"`
	Timestamp int `json:"ts"`
}

var (
	servics_map     = make(map[string]ServiceList)
	service_locker  = new(sync.Mutex)
	consul_client   *api.Client
	my_service_id   string
	my_service_name string
	my_kv_key       string
)

func CheckErr(err error) {
	if err != nil {
		log.Printf("error: %v", err)
		os.Exit(1)
	}
}
func StatusHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("check status.")
	fmt.Fprint(w, "status ok!")
}

func StartService(addr string) {
	http.HandleFunc("/status", StatusHandler)
	fmt.Println("start listen...")
	err := http.ListenAndServe(addr, nil)
	CheckErr(err)
}

func main() {
	var status_monitor_addr, service_name, service_ip, consul_addr, found_service string
	var service_port int
	flag.StringVar(&consul_addr, "consul_addr", "localhost:8500", "host:port of the service stuats monitor interface")
	flag.StringVar(&status_monitor_addr, "monitor_addr", "127.0.0.1:54321", "host:port of the service stuats monitor interface")
	flag.StringVar(&service_name, "service_name", "worker", "name of the service")
	flag.StringVar(&service_ip, "ip", "127.0.0.1", "service serve ip")
	flag.StringVar(&found_service, "found_service", "worker", "found the target service")
	flag.IntVar(&service_port, "port", 4300, "service serve port")
	flag.Parse()

	my_service_name = service_name

	DoRegistService(consul_addr, status_monitor_addr, service_name, service_ip, service_port)

	go DoDiscover(consul_addr, found_service)

	go StartService(status_monitor_addr)

	go WaitToUnRegistService()

	go DoUpdateKeyValue(consul_addr, service_name, service_ip, service_port)

	select {}
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

func DoDiscover(consul_addr string, found_service string) {
	t := time.NewTicker(time.Second * 5)
	for {
		select {
		case <-t.C:
			DiscoverServices(consul_addr, true, found_service)
		}
	}
}

func DiscoverServices(addr string, healthyOnly bool, service_name string) {
	consulConf := api.DefaultConfig()
	consulConf.Address = addr
	client, err := api.NewClient(consulConf)
	CheckErr(err)

	services, _, err := client.Catalog().Services(&api.QueryOptions{})
	CheckErr(err)

	fmt.Println("--do discover ---:", addr)

	var sers ServiceList
	for name := range services {
		servicesData, _, err := client.Health().Service(name, "", healthyOnly,
			&api.QueryOptions{})
		CheckErr(err)
		for _, entry := range servicesData {
			if service_name != entry.Service.Service {
				continue
			}
			for _, health := range entry.Checks {
				if health.ServiceName != service_name {
					continue
				}
				fmt.Println("  health nodeid:", health.Node, " service_name:", health.ServiceName, " service_id:", health.ServiceID, " status:", health.Status, " ip:", entry.Service.Address, " port:", entry.Service.Port)

				var node ServiceInfo
				node.IP = entry.Service.Address
				node.Port = entry.Service.Port
				node.ServiceID = health.ServiceID

				//get data from kv store
				s := GetKeyValue(service_name, node.IP, node.Port)
				if len(s) > 0 {
					var data KVData
					err = json.Unmarshal([]byte(s), &data)
					if err == nil {
						node.Load = data.Load
						node.Timestamp = data.Timestamp
					}
				}
				fmt.Println("service node updated ip:", node.IP, " port:", node.Port, " serviceid:", node.ServiceID, " load:", node.Load, " ts:", node.Timestamp)
				sers = append(sers, node)
			}
		}
	}

	service_locker.Lock()
	servics_map[service_name] = sers
	service_locker.Unlock()
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

func GetKeyValue(service_name string, ip string, port int) string {
	key := service_name + "/" + ip + ":" + strconv.Itoa(port)

	kv, _, err := consul_client.KV().Get(key, nil)
	if kv == nil {
		return ""
	}
	CheckErr(err)

	return string(kv.Value)
}*/

func main() {
	http.HandleFunc("/", getIps)
	err := http.ListenAndServe(":54321",nil)
	if err != nil {
		log.Fatal(err)
	}
	go watcher.NewWatcher()
}

var serverInfo = []ServerInfo{}

func HandleVersion(version string) {

}

func getIps(w http.ResponseWriter, r *http.Request) {
	serverInfo = serverInfo[:0]
	for _, v := range watcher.ServerIP.IpInfo {
		for _, server := range serverInfo {
			server.Number++
			continue
		}
		serverInfo=append(serverInfo,ServerInfo{ServerRegion:v.Region})
	}
	for _, v := range watcher.ServerIP.IpInfo {
		/*if serverInfo == v.Region {

		}*/
		for _, ip := range v.IP {
			//fmt.Println(ip)
			Send(ip+"/version")
		}
		//_, re := Send(v.IP+"/version")
		//_,re := Send(v.IP+"")

	}
	fmt.Fprintln(w, serverInfo)
}

func exec_shell(cmd string, ip string) (string, error) {
	//cmd := exec.Command("/bin/bash", "-c", s)
	command := exec.Command(cmd, "-I", ip)
	var out bytes.Buffer
	var outerr bytes.Buffer
	command.Stdout = &out
	command.Stderr = &outerr
	command.Run()
	return out.String(), errors.New(outerr.String())
}

func fetch(url string, ch chan<- string) {
	start := time.Now()
	resp, err := http.Get(url)
	fmt.Println(resp)
	if err != nil {
		ch <- fmt.Sprint(err) // send to channel ch
		return
	}
	s,_ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close() // don't leak resources
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs  %s  %s", secs, string(s), url)
}

func Send(ip string) (string,string) {
	s, err := exec_shell( "curl", ip)
	if err != nil {
		return ip, s
	}
	return ip, err.Error()
}
/*
func main(){
	fmt.Println("test begin .")
	config := consulapi.DefaultConfig()
	//config.Address = "localhost"
	fmt.Println("defautl config : ", config)
	client, err := consulapi.NewClient(config)
	if err != nil {
		log.Fatal("consul client error : ", err)
	}
	//创建一个新服务。
	registration := new(consulapi.AgentServiceRegistration)
	registration.ID = "123"
	registration.Name = "user-tomcat"
	registration.Port = 8080
	registration.Tags = []string{"user-tomcat"}
	registration.Address = "127.0.0.1"

	//增加check。
	check := new(consulapi.AgentServiceCheck)
	check.HTTP = fmt.Sprintf("http://%s:%d%s", registration.Address, registration.Port, "/check")
	//设置超时 5s。
	check.Timeout = "5s"
	//设置间隔 5s。
	check.Interval = "5s"
	//注册check服务。
	registration.Check = check
	log.Println("get check.HTTP:",check)

	err = client.Agent().ServiceRegister(registration)

	if err != nil {
		log.Fatal("register server error : ", err)
	}
}*/

/*func main() {
	var (
		err    error
		params map[string]interface{}
		plan   *watch.Plan
		ch     chan int
	)
	ch = make(chan int, 1)

	params = make(map[string]interface{})
	params["type"] = "service"
	params["service"] = "test"
	params["passingonly"] = false
	params["tag"] = "SERVER"
	plan, err = watch.Parse(params)
	if err != nil {
		panic(err)
	}
	plan.Handler = func(index uint64, result interface{}) {
		if entries, ok := result.([]*consulApi.ServiceEntry); ok {
			fmt.Printf("serviceEntries:%v", entries)
			// your code
			ch <- 1
		}
	}
	go func() {
		// your consul agent addr
		if err = plan.Run("127.0.0.1:7888"); err != nil {
			panic(err)
		}
	}()
	go http.ListenAndServe(":8080", nil)
	go register()
	for {
		<-ch
		fmt.Printf("get change")
	}
}

func register() {
	var (
		err    error
		client *consulApi.Client
	)
	client, err = consulApi.NewClient(&consulApi.Config{Address: "127.0.0.1:7888"})
	if err != nil {
		panic(err)
	}
	err = client.Agent().ServiceRegister(&consulApi.AgentServiceRegistration{
		ID:   "",
		Name: "test",
		Tags: []string{"SERVER"},
		Port: 8080,
		Check: &consulApi.AgentServiceCheck{
			HTTP: "",
		},
	})
	if err != nil {
		panic(err)
	}
}*/