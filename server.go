package main

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os/exec"
	"time"
)

/*var ips = []string{
	"111.230.186.23",
	"47.96.233.11",
}*/

type Version struct {
	HostName string `json:"hostname"`
	IP string `json:"ip"`
	version string `json:"version"`
}

type FB struct{
	feedback []Version `json:"feedback"`
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


func main() {
	/*command := exec.Command("curl", "54.81.187.245/version")
	outinfo,stderr := bytes.Buffer{},bytes.Buffer{}
	command.Stdout = &outinfo
	command.Stderr = &stderr
	err := command.Start()
	if err != nil{
		fmt.Println(err.Error(),stderr.String())
	}
	if err = command.Wait();err!=nil{
		fmt.Println(err.Error(),stderr.String())
	}else{
		fmt.Println(command.ProcessState.Pid())
		fmt.Println(outinfo.String())
	}*/
	/*addr, err := net.LookupAddr("54.81.187.24")
	fmt.Println(addr, err)*/
	http.HandleFunc("/", getIps)
	err := http.ListenAndServe(":9999",nil)
	if err != nil {
		log.Fatal(err)
	}
}

func getIps(w http.ResponseWriter, r *http.Request) {
	fb := FB{}
	for _, v := range ips {
		_,re := Send(v.IP+"/version")
		fb.feedback = append(fb.feedback,Version{
			HostName:v.HostName,
			IP:v.IP,
			version:re,
		})
	}
	fmt.Fprintln(w, fb)
}

func exec_shell(cmd string, ip string) (string, error) {
	//cmd := exec.Command("/bin/bash", "-c", s)
	command := exec.Command(cmd, ip)
	//cmd := exec.Command(s)
	var out bytes.Buffer
	var outerr bytes.Buffer
	command.Stdout = &out
	command.Stderr = &outerr
	//可能会阻塞
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
	//fmt.Println(ip)
	//ch := make(chan string)
	//go fetch("http://"+ip, ch)
	//s, err := exec_shell("hostname",IP:"")
	//fmt.Println(s, err)
	s, err := exec_shell("curl",ip)
	//s := <-ch
	//fmt.Println("end")
	if err != nil {
		return ip, s
	}
	return ip, err.Error()
}

func Resp() {

}