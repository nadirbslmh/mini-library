package logging

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	logmodel "logging-service/pkg/model"
	"math/rand"
	"net/http"
	"pkg-service/discovery"
	"pkg-service/model"
)

type LogHandler struct {
	registry discovery.Registry
}

func New(registry discovery.Registry) *LogHandler {
	return &LogHandler{
		registry: registry,
	}
}

func (handler *LogHandler) Write(ctx context.Context, logInput logmodel.LogInput) (*model.Response[logmodel.Log], error) {
	addrs, err := handler.registry.ServiceAddresses(ctx, "logging")

	if err != nil {
		return nil, err
	}

	url := "http://" + addrs[rand.Intn(len(addrs))] + "/logs"

	log.Printf("calling log service. Request: POST logs" + url)

	jsonBody, err := json.Marshal(&logInput)

	if err != nil {
		return nil, err
	}

	bodyReader := bytes.NewReader(jsonBody)

	req, err := http.NewRequest(http.MethodPost, url, bodyReader)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New("not found")
	} else if resp.StatusCode/100 != 2 {
		return nil, fmt.Errorf("non-2xx response: %v", resp)
	}

	var v *model.Response[logmodel.Log]

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}
