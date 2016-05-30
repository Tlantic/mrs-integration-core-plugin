package configuration

import (
	"github.com/Tlantic/mrs-integration-utils/configuration"
	"fmt"
	"strings"
)

var Data *PConfig

type PConfig  struct{
	Couchbase Couchbase
	Rabbit Rabbit
	Elastic Elastic
	Endpoints Endpoints
	Server Server
	Mongodb Mongodb
	Flags Flags
	EventSourcing EventSourcing
}

type Server struct {
	Port string
}

type Couchbase struct{
	Url string
	Bucket string
	Password string

}

type Rabbit struct{
	Url string
}

type Elastic struct{
	Url string
}

type Endpoints struct{
	FullSingleProduct CustomEndPoint
}

type CustomEndPoint struct {
	Endpoint string
	Verb string
}

type Mongodb struct {
	Url string
	Database string
}

type Flags struct {
	Workers int
	QueueSize int
	UploadBaseBasePath string
	UploadBucketName string
	S3Region string
	RollBarToken string
	RemoteLogs bool
}

type EventSourcing struct{
	Exchange string
	Prefix string
}

func InitConfig(path string, prefix string, url string, contentType string){
	cf := configuration.GetConfiguration(path, prefix, url, contentType)
	cf.Load()

	cc := PConfig{}

	err := cf.Keys.Unmarshal(&cc)
	if err != nil {
		fmt.Println(err)
	}

	Data = &cc
}


func GetConfig() *PConfig{
	return Data
}


const (

	QUEUE_PRODUCT_CHANGE_DESCRIPTION = "product-change-description"
	EVENT_PREFIX_NAME = "product-integration"
	EVENT_QUEUE_NAME = "queue1"
	EVENT_QUEUE_EXCH = "testing.example"
)


var QueueNames = []string{QUEUE_PRODUCT_CHANGE_DESCRIPTION}


func GetProductServiceEndpoint(organization string) string{
	return strings.Replace(GetConfig().Endpoints.FullSingleProduct.Endpoint, "[org]", organization, 1)
}