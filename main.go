package main

import (
	"fmt"
	"log"
	"net/http"
)

//const URLTemplate string = "https://1drv.ms/%s/s!%s"

func main() {
	port := getEnv("PORT", "8080")

	//r := http.NewServeMux()
	//pattern, err := regexp.Compile(`\/\w+\/s!\w+`)
	//check(err)
	//par := pattern.String()
	//log.Println("par: ", par)
	//
	//r.HandleFunc(par, oneDirectHandler)
	http.HandleFunc("/", indexHandler)
	//http.Handle("/", r)

	log.Printf("Listening on port %s", port)
	log.Printf("Open http://localhost:%s in the browser", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", port), nil))
}

//
//func getEnv(envName string, envDef string) string {
//	var env = os.Getenv(envName)
//	if env == "" {
//		env = envDef
//		log.Printf("Defaulting %s to %s", envName, envDef)
//	}
//	return env
//}
//
//func check(object string, value interface{}, err error) {
//	if err != nil {
//		log.Fatal(object, ": ", err)
//	}
//	log.Println(object, ": ", value)
//}
//
//func catch(object string , err error) bool {
//	if err != nil {
//		log.Println(object," got a error: ",err)
//		return false
//	}
//	log.Println(object," works well")
//	return true
//}
//
//func writeOut(w http.ResponseWriter,resp *http.Response) {
//	body, err := ioutil.ReadAll(resp.Body)
//	check("body", string(body), err)
//
//	//rHeader := resp.Header
//	rCode := resp.StatusCode
//	log.Println("rCode: ",rCode)
//
//	w.WriteHeader(rCode)
//	w.Write(body)
//
//	log.Println("writeOut works well")
//}
//
//func servicing(service string, r *http.Request, port string) {
//	log.Printf("servicing %s in http://localhost:%s%s\n", service, port, r.URL.RequestURI())
//}
//
//func serveComplete(service string){
//	log.Printf("serve %s complete!\n",service)
//}

//func faviconHandler(w http.ResponseWriter, r *http.Request, port string) {
//	servicing("favicon", r, port)
//	http.ServeFile(w, r, "./images/ori.svg")
//
//}
//
//func fileHandler(w http.ResponseWriter, r *http.Request, port string) {
//	servicing("file server", r, port)
//	http.ServeFile(w, r, "."+r.URL.RequestURI())
//}
//
//func oneDirectHandler(w http.ResponseWriter, r *http.Request, port string) {
//	servicing("oneDirect", r, port)
//
//	//err := r.ParseForm()
//	//check(err)
//
//	url := "https://1drv.ms" + r.URL.RequestURI()
//	log.Println("url: ",url)
//
//	//vars := mux.Vars(r)
//	//url := fmt.Sprintf(URLTemplate, vars["action"], vars["token"])
//
//	//req1, err := http.NewRequest("HEAD", url,nil)
//	//check(err)
//
//	client := &http.Client{
//		CheckRedirect: func(req *http.Request, via []*http.Request) error {
//			return http.ErrUseLastResponse
//		},
//	}
//
//	resp1, err := client.Get(url)
//	defer resp1.Body.Close()
//	check("req1","ok",err)
//
//	transURL, err := resp1.Location()
//	if !catch("transURL",err) {
//		writeOut(w,resp1)
//		return
//	}
//
//
//	check("resp1", transURL, err)
//
//	downloadURL := strings.ReplaceAll(transURL.String(), "/redir?", "/download?")
//	log.Println("downloadURL: ", downloadURL)
//
//	//req2, err := http.NewRequest("HEAD", downloadURL,nil)
//	//check(err)
//	resp2, err := client.Get(downloadURL)
//	defer resp2.Body.Close()
//	check("req2","ok",err)
//
//	directURL, err := resp2.Location()
//	if !catch("directURL",err) {
//		writeOut(w,resp2)
//		return
//	}
//
//	URI := r.URL.RequestURI()
//	plain, err := regexp.MatchString(`\?txt`, URI)
//	check("txt", plain,err)
//	if plain {
//		_, err = fmt.Fprint(w, directURL)
//		serveComplete("oneDirect")
//		return
//	}
//
//	http.Redirect(w, r, directURL.String(), 302)
//	serveComplete("oneDirect")
//}

//func indexHandler(w http.ResponseWriter, r *http.Request) {
//	port := getEnv("PORT", "8080")
//	servicing("index", r, port)
//
//	if r.URL.RequestURI() == "/favicon.ico" {
//		faviconHandler(w, r, port)
//		return
//	}
//
//	rPath := r.URL.Path
//
//	pattern, err := regexp.Compile(`\/\w+\/s!\w+`)
//	check("pattern", pattern, err)
//
//	oneD := pattern.MatchString(rPath)
//	log.Println(oneD)
//
//	if oneD {
//		oneDirectHandler(w, r, port)
//		return
//	}
//	//link := "1drv.ms/u/s!Aiw77soXua44hb4CEu6eSveUl0xUoA?txt"
//	//par ,err :=regexp.Compile(`\/\w+\/s!\w+\??\w+`)
//	//ca := par.Find([]byte(link))
//	//
//	//log.Println(par.String())
//	//log.Println(string(ca))
//
//	_, err = fmt.Fprint(w, "Hello, here is index")
//	check("index", "ok", err)
//	serveComplete("index")
//}
