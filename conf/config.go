package conf

import (
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net"
	"net/http"
	"test_server/app"
	"test_server/db"
	"time"
)

//var DBcfg = &db.DBconf{}

//externalIP is TEMP function to get external ip from current machine
func externalIP() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		addrs, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range addrs {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network?")
}

//Config function gets http.Server pointer and write configuration from config file.
//TODO: create config.cfg file reader and parser. All current params now statically written in function
func Config(srv *http.Server) {
	fmt.Println("Yu are obosralsya!")

	ip, err := externalIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)

	mux := mux.NewRouter()
	//mux.HandleFunc("/lol", funccc)
	s := mux.PathPrefix("/app").Subrouter()

	app.Handler(s)

	srv.Addr = ip + ":7990"
	srv.Handler = mux
	srv.WriteTimeout = 15 * time.Second
	srv.ReadTimeout = 15 * time.Second

}

//CreateDBConfiguration func is read config for database and fill all config fields
//TODO: create filereader too!
func CreateDBConfiguration() {

	db.DBcfg.Host = "192.168.10.205"
	db.DBcfg.Port = 5432
	db.DBcfg.User = "adminpsg"
	db.DBcfg.Pass = "PadminSGG"
	db.DBcfg.DBname = "server_gredit_db"

	fmt.Println(db.DBcfg)
}
