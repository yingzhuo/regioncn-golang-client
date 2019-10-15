package regioncn_golang_client

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	"github.com/yingzhuo/regioncn-golang-client/protobuf"
	"io/ioutil"
	"net/http"
	"strings"
)

type Model struct {
	Id        int64   `json:"id"`
	Code      string  `json:"code"`
	Name      string  `json:"name"`
	ShortName string  `json:"shortName"`
	Lat       float64 `json:"lat"`
	Lng       float64 `json:"lng"`
}

func (m *Model) String() string {
	bytes, _ := json.Marshal(m)
	return string(bytes)
}

// ---------------------------------------------------------------------------------------------------------------------

type Config struct {
	Host         string
	Port         int
	ResponseType string
}

const (
	Json     = "json"
	Protobuf = "protobuf"
)

// ---------------------------------------------------------------------------------------------------------------------

type Client struct {
	host         string
	port         int
	responseType string
}

func (cli *Client) GetAllProvince() []*Model {
	url := fmt.Sprintf("http://%s:%d/province", cli.host, cli.port)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	if strings.EqualFold(Json, cli.responseType) {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var ret []*Model
		if err = json.Unmarshal(body, &ret); err == nil {
			return ret
		} else {
			panic(err)
		}
	}

	if strings.EqualFold(Protobuf, cli.responseType) {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var ret []*Model
		var models = protobuf.Models{}
		if err = proto.Unmarshal(body, &models); err == nil {
			for _, it := range models.List {
				m := &Model{
					Id:        it.Id,
					Code:      it.Code,
					Name:      it.Name,
					ShortName: it.ShortName,
					Lat:       it.Lat,
					Lng:       it.Lng,
				}
				ret = append(ret, m)
			}
			return ret
		} else {
			panic(err)
		}
	}
	return nil
}

func (cli *Client) GetCityByProvinceCode(code string) []*Model {
	url := fmt.Sprintf("http://%s:%d/city?provinceCode=%s", cli.host, cli.port, code)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	if strings.EqualFold(Json, cli.responseType) {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var ret []*Model
		if err = json.Unmarshal(body, &ret); err == nil {
			return ret
		} else {
			panic(err)
		}
	}

	if strings.EqualFold(Protobuf, cli.responseType) {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var ret []*Model
		var models = protobuf.Models{}
		if err = proto.Unmarshal(body, &models); err == nil {
			for _, it := range models.List {
				m := &Model{
					Id:        it.Id,
					Code:      it.Code,
					Name:      it.Name,
					ShortName: it.ShortName,
					Lat:       it.Lat,
					Lng:       it.Lng,
				}
				ret = append(ret, m)
			}
			return ret
		} else {
			panic(err)
		}
	}
	return nil
}

func (cli *Client) GetAreaByCityCode(code string) []*Model {
	url := fmt.Sprintf("http://%s:%d/area?cityCode=%s", cli.host, cli.port, code)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	if strings.EqualFold(Json, cli.responseType) {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var ret []*Model
		if err = json.Unmarshal(body, &ret); err == nil {
			return ret
		} else {
			panic(err)
		}
	}

	if strings.EqualFold(Protobuf, cli.responseType) {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var ret []*Model
		var models = protobuf.Models{}
		if err = proto.Unmarshal(body, &models); err == nil {
			for _, it := range models.List {
				m := &Model{
					Id:        it.Id,
					Code:      it.Code,
					Name:      it.Name,
					ShortName: it.ShortName,
					Lat:       it.Lat,
					Lng:       it.Lng,
				}
				ret = append(ret, m)
			}
			return ret
		} else {
			panic(err)
		}
	}
	return nil
}

func (cli *Client) GetStreetByAreaCode(code string) []*Model {
	url := fmt.Sprintf("http://%s:%d/street?areaCode=%s", cli.host, cli.port, code)

	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}

	if strings.EqualFold(Json, cli.responseType) {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}
		var ret []*Model
		if err = json.Unmarshal(body, &ret); err == nil {
			return ret
		} else {
			panic(err)
		}
	}

	if strings.EqualFold(Protobuf, cli.responseType) {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			panic(err)
		}

		var ret []*Model
		var models = protobuf.Models{}
		if err = proto.Unmarshal(body, &models); err == nil {
			for _, it := range models.List {
				m := &Model{
					Id:        it.Id,
					Code:      it.Code,
					Name:      it.Name,
					ShortName: it.ShortName,
					Lat:       it.Lat,
					Lng:       it.Lng,
				}
				ret = append(ret, m)
			}
			return ret
		} else {
			panic(err)
		}
	}
	return nil
}

// ---------------------------------------------------------------------------------------------------------------------

func NewClient(config *Config) *Client {
	return &Client{
		host:         config.Host,
		port:         config.Port,
		responseType: config.ResponseType,
	}
}

func NewJsonClient(host string, port int) *Client {
	return NewClient(&Config{
		Host:         host,
		Port:         port,
		ResponseType: Json,
	})
}

func NewProtobufClient(host string, port int) *Client {
	return NewClient(&Config{
		Host:         host,
		Port:         port,
		ResponseType: Protobuf,
	})
}
