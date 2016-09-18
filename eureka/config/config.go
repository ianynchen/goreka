package config

import (
	"encoding/json"
	"encoding/xml"
)

type LeaseInfo struct {
	EvictionDurationInSecs int `json:"evictionDurationInSecs,omitempty"`
}

type DataCenterName string

const (
	MyOwn  DataCenterName = "MyOwn"
	Amazon DataCenterName = "Amazon"
)

type StatusType string

const (
	UP           StatusType = "UP"
	DOWN         StatusType = "DOWN"
	STARTING     StatusType = "STARTING"
	OUTOFSERVICE StatusType = "OUT_OF_SERVICE"
	UNKNOWN      StatusType = "UNKNOWN"
)

type DataCenterInfo struct {
	Name     DataCenterName  `json:"name"`
	Metadata *AmazonMetaData `json:"metadata,omitempty"`
}

type AmazonMetaData struct {
	AmiLaunchIndex   string `json:"ami-launch-index"`
	LocalHostname    string `json:"local-hostname"`
	AvailabilityZone string `json:"availability-zone"`
	InstanceId       string `json:"instance-id"`
	PublicIPv4       string `json:"public-ipv4"`
	PublicHostname   string `json:"public-hostname"`
	AmiManifestPath  string `json:"ami-manifest-path"`
	LocalIPv4        string `json:"local-ipv4"`
	Hostname         string `json:"hostname"`
	AmiId            string `json:"ami-id"`
	InstanceType     string `json:"instance-type"`
}

type JsonPort struct {
	Number  int  `json:"$"`
	Enabled bool `json:"@enabled"`
}

type InstanceMetadata struct {
	Raw    []byte `xml:",innerxml" json:"-"`
	parsed map[string]interface{}
}

type Instance struct {
	XMLName          struct{} `xml:"instance" json:"-"`
	HostName         string   `xml:"hostName" json:"hostName"`
	App              string   `xml:"app" json:"app"`
	IPAddr           string   `xml:"ipAddr" json:"ipAddr"`
	VipAddress       string   `xml:"vipAddress" json:"vipAddress"`
	SecureVipAddress string   `xml:"secureVipAddress" json:"secureVipAddress"`

	Status           StatusType `xml:"status" json:"status"`
	Overriddenstatus StatusType `xml:"overriddenstatus" json:"overriddenstatus,omitempty"`

	Port           int       `xml:"port" json:"-"`
	PortJson       *JsonPort `json:"port,omitempty" xml:"-"`
	SecurePort     int       `xml:"securePort" json:"-"`
	SecurePortJson *JsonPort `json:"securePort" xml:"-"`

	HomePageUrl    string `xml:"homePageUrl" json:"homePageUrl"`
	StatusPageUrl  string `xml:"statusPageUrl" json:"statusPageUrl"`
	HealthCheckUrl string `xml:"healthCheckUrl" json:"healthCheckUrl"`

	CountryId      int64          `xml:"countryId" json:"countryId"`
	DataCenterInfo DataCenterInfo `xml:"dataCenterInfo" json:"dataCenterInfo"`

	LeaseInfo *LeaseInfo        `xml:"leaseInfo" json:"leaseInfo,omitempty"`
	Metadata  *InstanceMetadata `xml:"metadata" json:"metadata,omitempty"`
}

type Instances struct {
	Instances []Instance `xml:"instances" json:"instances"`
}

func NewInstance() *Instance {

	instance := new(Instance)
	instance.Port = 80
	instance.SecurePort = 443

	instance.PortJson = new(JsonPort)
	instance.PortJson.Number = 80
	instance.PortJson.Enabled = true

	instance.Status = UP

	instance.SecurePortJson = new(JsonPort)
	instance.SecurePortJson.Number = 443
	instance.SecurePortJson.Enabled = false

	instance.DataCenterInfo.Name = MyOwn

	return instance
}

func (instance *Instance) MarshalXml() ([]byte, bool) {

	if content, err := xml.Marshal(*instance); err == nil {

		return content, true
	} else {
		return nil, false
	}
}

func UnmarshalXmlInstance(content []byte) (*Instance, bool) {

	return nil, false
}

func (instance *Instance) MarshalJson() ([]byte, bool) {

	if content, err := json.Marshal(*instance); err == nil {

		return content, true
	} else {
		return nil, false
	}
}

func UnmarshalJsonInstance(content []byte) (*Instance, bool) {

	return nil, false
}

type EurekaService struct {
	RootEurekaServiceUrls        []string `json:"serviceUrls"`
	EurekaServicePath            string   `json:"servicePath"`
	ConnectionTimeoutInSecs      int      `json:"connectionTimeoutInSeconds"`
	GlobalEvictionDurationInSecs int      `json:"evictionDurationInSeconds"`
}

var EurekaServiceConfiguration EurekaService
