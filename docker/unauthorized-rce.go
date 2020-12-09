package docker

import (
	"fmt"
	"net/http"
	"time"
)

func DockerUnauthorized(ip string, ports []string, timeOut time.Duration) ([]string, bool) {
	addr := make([]string, 0)
	if len(ports) == 0 {
		ports = []string{
			"2375",
		}
	}

	if timeOut <= 0 {
		timeOut = time.Second * 1
	}
	for _, port := range ports {

		host_addr := fmt.Sprintf("http://%s:%s/", ip, port)
		http_cli := http.Client{
			Timeout: timeOut,
		}
		resp, err := http_cli.Get(host_addr)
		_ = resp
		if err != nil {
			continue
		} else {
			addr = append(addr, host_addr)
		}

	}

	if len(addr) > 0 {
		return addr, true
	} else {
		return addr, false
	}
}

func DockerUnauthorized_RCE() {
	//transport := new(http.Transport)
	//sockets.ConfigureTransport(transport, "tcp", host_addr)
	//httpclient := &http.Client{
	//	Transport: transport,
	//	Timeout: time.Second * 3,
	//}
	//cli, err := client.NewClient("tcp://"+ host_addr, version, httpclient, nil)
	//if err != nil {
	//	continue
	//}
	//_, err = cli.ContainerList(context.Background(), types.ContainerListOptions{})
	//if err != nil {
	//	fmt.Println(err)
	//	continue
	//} else {
	//	addr = append(addr, host_addr)
	//}
}
