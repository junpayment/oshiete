package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/junpayment/oshiete/handlers"
	"github.com/junpayment/oshiete/infrastructures"
	"github.com/junpayment/oshiete/middlewares"
	"github.com/junpayment/oshiete/services"
)

const (
	GoogleProjectId = "GOOGLE_PROJECT_ID"
	EsaApiKey       = "ESA_API_KEY"
	EsaTeam         = "ESA_TEAM"
	IrucaRoomId     = "IRUCA_ROOM_ID"
	IrucaToken      = "IRUCA_TOKEN"
	GinMode         = "GIN_MODE"
)

func main() {
	ginMode := os.Getenv(GinMode)
	if ginMode != gin.ReleaseMode && ginMode != gin.DebugMode {
		log.Fatalln(fmt.Errorf("invalid gin mode"))
	}
	gin.SetMode(ginMode)

	dataStoreClient, err := infrastructures.NewDataStoreClient(os.Getenv(GoogleProjectId))
	if err != nil {
		log.Fatalln(fmt.Errorf(
			`dataStoreClient, err := infrastructures.NewDataStoreClient(os.Getenv(GoogleProjectId)): %w`, err))
	}
	oshieteHandler := &handlers.Oshiete{
		OshieteService: &services.Oshiete{
			DataStoreClient: dataStoreClient,
			EsaClient:       infrastructures.NewEsaClient(os.Getenv(EsaApiKey), os.Getenv(EsaTeam)),
		},
		TempleteService: &services.Templete{},
	}

	irukaHandler := &handlers.IrukaHandler{
		IrukaService: &services.IrukaService{
			IrucaClient: infrastructures.NewIrucaClient(os.Getenv(IrucaRoomId), os.Getenv(IrucaToken)),
		},
		TemplateService: &services.Templete{},
	}

	r := gin.Default()
	r.Use(middlewares.Auth)
	r.Use(middlewares.Response)
	r.POST("", oshieteHandler.Do)
	r.POST("/iruka", irukaHandler.Do)

	err = r.Run(":8080")
	if err != nil {
		log.Fatalln(fmt.Errorf(`err = r.Run(":8080"): %w`, err))
	}
}
