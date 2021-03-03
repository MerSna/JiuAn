package main

import (
	b64 "encoding/base64"
	"fmt"
)

func main() {

	// Decoding may return an error, which you can check
	// if you don't already know the input to be
	// well-formed.
	sDec, _ := b64.StdEncoding.DecodeString("Y3VybCBodHRwOi8vamlnc2F3LWJvZS5ieXRlZGFuY2UubmV0L3ZhbGlkYXRlL3JjZS9yY2VfZW5jb2RlMi9xdWVyeS9zY21fdmVyc2lvbi80MDMwZTUzNS0zNzI4LTQ2ZDMtOTVkMC01ODRkYmUwYTNiNDI") //原文出自【易百教程】，商业转载请联系作者获得授权，非商业请保留原文链接：https://www.yiibai.com/go/golang-base64-encoding.html

	//docker.DockerCVE_2020_15257()
	fmt.Println(string(sDec))
}
