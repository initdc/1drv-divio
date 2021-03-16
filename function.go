package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func getEnv(envName string, envDef string) string {
	var env = os.Getenv(envName)
	if env == "" {
		env = envDef
		log.Printf("Defaulting %s to %s", envName, envDef)
	}
	return env
}

func check(object string, value interface{}, err error) {
	if err != nil {
		log.Fatal(object, ": ", err)
	}
	log.Println(object, ": ", value)
}

func catch(object string, err error) bool {
	if err != nil {
		log.Println(object, " got a error: ", err)
		return false
	}
	log.Println(object, " works well")
	return true
}

func writeOut(w http.ResponseWriter, resp *http.Response) {
	body, err := ioutil.ReadAll(resp.Body)
	check("body", string(body), err)

	//rHeader := resp.Header
	rCode := resp.StatusCode
	log.Println("rCode: ", rCode)

	w.WriteHeader(rCode)
	w.Write(body)

	log.Println("writeOut works well")
}

func servicing(service string, r *http.Request, port string) {
	log.Printf("servicing %s in http://localhost:%s%s\n", service, port, r.URL.RequestURI())
}

func serveComplete(service string) {
	log.Printf("serve %s complete!\n", service)
}
