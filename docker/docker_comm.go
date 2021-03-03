package docker

import (
	"errors"
	"io/ioutil"
	"log"
	"regexp"
)

func GetDockerAbsPath() string {
	data, err := ioutil.ReadFile("/proc/self/mounts")
	if err != nil {
		log.Println(err)
	}
	//fmt.Println(string(data))

	// workdir=/var/lib/docker/overlay2/9383b939bf4ed66b3f01990664d533f97f1cf9c544cb3f3d2830fe97136eb76f/work
	pattern := regexp.MustCompile("workdir=([\\w\\d/]+)/work")
	params := pattern.FindStringSubmatch(string(data))
	if len(params) < 2 {
		log.Fatal("failed to find docker abs path in /proc/self/mounts")
	}

	return params[1]
}

func GetShimSockets() ([][]byte, error) {
	re, err := regexp.Compile("@/containerd-shim/.*\\.sock")
	if err != nil {
		return nil, err
	}
	data, err := ioutil.ReadFile("/proc/net/unix")
	matches := re.FindAll(data, -1)
	if matches == nil {
		return nil, errors.New("Cannot find vulnerable sockets ")
	}
	return matches, nil
}
