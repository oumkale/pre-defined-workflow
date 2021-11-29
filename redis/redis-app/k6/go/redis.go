package main

import (
	"bytes"
	"flag"
	"os/exec"
)

//AppVars maintaining all parameters require during application installation/deletion
type AppVars struct {
	host     string
	port     string
	rps      int
	password string
	clients  string
	loop     string
	scope    string
}

//GetData derive the application filePath and timeout
//it derive the filePath based on application scenario(week vs resilient)
func main() {

	appVars, err := GetData()

	flag.Parse()

	//application namespace having weak and resilient filePath
	//loadtest namespace having loadtest filePath
	//sock-shop namespace having sock-shop filePath
	//podtato-head namespace having podtato-head filePath

}

func GetData() (*AppVars, error) {
	host := flag.String("h", "127.0.0.1", "Server hostname.")
	port := flag.Int("p", 6379, "Server port.")
	rps := flag.Int64("rps", 0, "Max rps. If 0 no limit is applied and the DB is stressed up to maximum.")
	password := flag.String("a", "", "Password for Redis Auth.")
	seed := flag.Int64("random-seed", 12345, "random seed to be used.")
	clients := flag.Uint64("c", 50, "number of clients.")
	keyspacelen := flag.Uint64("r", 1000000, "keyspace length. The benchmark will expand the string __key__ inside an argument with a number in the specified range from 0 to keyspacelen-1. The substitution changes every time a command is executed.")
	datasize := flag.Uint64("d", 3, "Data size of the expanded string __data__ value in bytes. The benchmark will expand the string __data__ inside an argument with a charset with length specified by this parameter. The substitution changes every time a command is executed.")
	numberRequests := flag.Uint64("n", 10000000, "Total number of requests")
	debug := flag.Int("debug", 0, "Client debug level.")
	multi := flag.Bool("multi", false, "Run each command in multi-exec.")
	waitReplicas := flag.Int("wait-replicas", 0, "If larger than 0 will wait for the specified number of replicas.")
	waitReplicasMs := flag.Int("wait-replicas-timeout-ms", 1000, "WAIT timeout when used together with -wait-replicas.")
	loop := flag.Bool("l", false, "Loop. Run the tests forever.")
	flag.Parse()

}

//CreateApp create the application
func CreateApp(path, ns, operation string) error {
	command := exec.Command("kubectl", operation, "-f", path, "-n", ns)
	var out, stderr bytes.Buffer
	command.Stdout = &out
	command.Stderr = &stderr
	if err := command.Run(); err != nil {
		return err
	}
	return nil
}
