package helper

import (
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"ucode_go_query_service/models"

	"google.golang.org/protobuf/types/known/structpb"
)

const (
	key        string = "key"
	value      string = "value"
	BASEURL           = "base_url"
	BODYTYPE          = "body_type"
	PATH              = "path"
	PARAMS            = "params"
	HEADERS           = "headers"
	COOKIES           = "cookies"
	ACTIONTYPE        = "action_type"
	BODY              = "body"
)

type ParamsI interface {
	Add(key string, value string)
	Set(key string, value string)
	Get(key string) string
}

func ParseParams(param ParamsI, body interface{}) {
	m := body.([]interface{})
	for i := range m {
		mi := m[i].(map[string]interface{})

		if mi[key].(string) != "" && mi[value].(string) != "" {
			param.Add(mi[key].(string), mi[value].(string))
		}
	}
}

func ParseBody(body map[string]interface{}) (io.Reader, error) {
	bodyType, ok := body[BODYTYPE].(string)
	if !ok {
		return strings.NewReader(""), nil
	}
	switch bodyType {
	case "RAW":
		bodyBody, ok := body[BODY]
		if !ok {
			return strings.NewReader(""), errors.New("cant parse query body")
		}
		switch bodyBody.(type) {
		case string:
			return strings.NewReader(bodyBody.(string)), nil
		case structpb.Struct:
			res := bodyBody.(structpb.Struct)
			bodyJson, err := res.MarshalJSON()
			if err != nil {
				return strings.NewReader(""), errors.New("cant parse query body to json")
			}
			return strings.NewReader(string(bodyJson)), nil
		}
	}
	return strings.NewReader(""), nil
}

func ParseQueryBody(body map[string]interface{}) (*http.Request, error) {
	var (
		actionType = "GET"
		baseUrl    string
		path       string
		//strPb      = &structpb.Struct{}
	)

	bodyBaseUrl, ok := body[BASEURL]
	if !ok {
		err := errors.New("cant parse query body")
		return nil, err
		//s.logger.Error("Error while running query in service", l.Error(err))
		//return nil, status.Error(codes.Internal, err.Error())
	}
	baseUrl = bodyBaseUrl.(string)

	u, err := url.Parse(baseUrl)
	if err != nil {
		err := errors.New("cant parse query body")
		return nil, err
	}

	bodyPath, ok := body[PATH]
	if ok {
		path = bodyPath.(string)
	}

	u.Path += path
	fmt.Println("initial::::", u.String())

	bodyQuery, ok := body[PARAMS]
	if ok {
		q := u.Query()
		ParseParams(&q, bodyQuery)
		u.RawQuery = q.Encode()
	}

	payload, err := ParseBody(body)
	if err != nil {
		return nil, err
	}

	bodyActionType, ok := body[ACTIONTYPE]
	if ok {
		actionType = strings.ToUpper(bodyActionType.(string))
	}

	req, err := http.NewRequest(actionType, u.String(), payload)
	if err != nil {
		return nil, err
	}

	bodyHeader, ok := body[HEADERS]
	if ok {
		ParseParams(&req.Header, bodyHeader)
	}

	cookies, ok := body[COOKIES]
	if ok {
		c := cookies.([]interface{})
		for i := range c {
			ci := c[i].(map[string]interface{})

			if ci[key].(string) != "" && ci[value].(string) != "" {
				req.AddCookie(&http.Cookie{Name: ci[key].(string), Value: ci[value].(string)})
			}
		}
	}

	return req, nil
}

func ReplaceVariable(data []byte, variables []models.VariableSchema) ([]byte, error) {
	var (
		dataString    string
		replacingData []string
	)
	dataString = string(data)
	for _, v := range variables {
		if v.Key != "" {
			if v.Value != "" {
				replacingData = append(replacingData, []string{"{{" + v.Key + "}}", v.Value}...)
			} else if v.Default != "" {
				replacingData = append(replacingData, []string{"{{" + v.Key + "}}", v.Default}...)
			}
		}
	}

	replacer := strings.NewReplacer(replacingData...)
	dataString = replacer.Replace(dataString)

	return []byte(dataString), nil
}
