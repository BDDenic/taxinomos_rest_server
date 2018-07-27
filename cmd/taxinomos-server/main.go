package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/DENICeG/taxinomos_rest_server/logging"

	"github.com/alecthomas/kingpin"
	"github.com/gin-gonic/gin"
)

var (
	builddate     string
	revision      string
	version       string
	lifetime      int
	listenaddress string
	debuglevel    int
	configfile    string
	wg            sync.WaitGroup
)

func main() {
	kingpin.UsageTemplate(kingpin.CompactUsageTemplate)
	kingpin.Flag("listenaddress", "Socket for the server to listen on.").Default("0.0.0.0:8080").Short('l').StringVar(&listenaddress)
	kingpin.Parse()

	//

	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime | log.Lmicroseconds)

	gin.SetMode(gin.ReleaseMode)
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(logging.GinLogger())

	log.Println("=====  MagicLinkServer  =====")
	log.Printf("Builddate: %s", builddate)
	log.Printf("Version  : %s", version)
	log.Printf("Revision : %s", revision)
	log.Println(" ---")
	log.Println("ServerConfig:")
	log.Printf("  listenaddress: %s", listenaddress)

	apiGroup := router.Group("/api/v1")
	{
		apiGroup.GET("/fetch", Fetch)
		//apiGroup.GET("/categories", GetCategories)
		//apiGroup.GET("/statuses", GetStatuses)
	}

	httpsrv := &http.Server{
		Addr:    listenaddress,
		Handler: router,
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Fatal(httpsrv.ListenAndServe())
	}()
	log.Println("HTTP server started.")

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	log.Println("Shutdown signals registered.")
	<-signalChan
	log.Println("Shutdown signal received, exiting.")
	httpsrv.Shutdown(context.Background())
	wg.Wait()
	log.Println("Server exiting")
}

func Fetch(c *gin.Context) {
	c.Redirect(301, "https://www.google.de")
}
