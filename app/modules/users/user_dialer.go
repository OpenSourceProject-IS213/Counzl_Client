package users

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"utilities/cmd"
	"utilities/converter"
)

const (
	CA          = "./database/.certificate/CA/ca.crt"
	client_crt  = "./database/.certificate/client.crt"
	client_key  = "./database/.certificate/client.key"
	server_ip   = "158.39.77.189"
	server_port = ":8081"
)

func CheckCerts_client(brukernavn string) string {
	log.SetFlags(log.Lshortfile)

	cert, err := tls.LoadX509KeyPair(client_crt, client_key)
	if err != nil {
		log.Fatal(err)
	}
	caCert, err := ioutil.ReadFile(CA)
	if err != nil {
		log.Fatal(err)
	}
	caCertPool := x509.NewCertPool()
	caCertPool.AppendCertsFromPEM(caCert)

	conf := &tls.Config{
		RootCAs:      caCertPool,
		Certificates: []tls.Certificate{cert},
	}

	id, err := dialMUX(brukernavn, conf)
	if err != nil {
		fmt.Println(cmd.ChangeColor("404", "red") + ": Can't establish a connection, the server may be offline, try to register the user locally :'(")
		os.Exit(1)
	}
	return id

}

func dialMUX(brukernavn string, conf *tls.Config) (string, error) {

	conn, err := tls.Dial("tcp", server_ip+server_port, conf)
	if err != nil {

		return "", err
	}
	defer conn.Close()

	n, err := conn.Write([]byte(brukernavn + "\n"))
	if err != nil {
		log.Println(n, err)
	}

	buf := make([]byte, 100)
	n, err = conn.Read(buf)
	if err != nil {
		log.Println(n, err)
	}

	println("Your user ID" + converter.Remove_0a(string(buf[:n])))
	return converter.Remove_0a(string(buf[:n])), nil
}
