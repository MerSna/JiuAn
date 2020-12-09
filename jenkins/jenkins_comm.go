package jenkins

var (
	default_ports = []string{
		"8080",
		//"8081",
		//"8088",
		//"8888",
		//"8880",
	}
	PREAMLE = []byte("<===[JENKINS REMOTING CAPACITY]===>rO0ABXNyABpodWRzb24ucmVtb3RpbmcuQ2FwYWJpbGl0eQAAAAAAAAABAgABSgAEbWFza3hwAAAAAAAAAH4=")
	PROTO   = []byte("\x00\x00\x00\x00")
)
