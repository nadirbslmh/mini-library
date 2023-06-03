package http

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"math/rand"
	bookgateway "minilib/library/internal/gateway/book/http"
	"minilib/library/pkg/constant"
	"minilib/pkg/discovery"
	"minilib/pkg/model"
	rentmodel "minilib/rent/pkg/model"
	"net/http"
	"strconv"
)

type Gateway struct {
	bookgateway *bookgateway.Gateway
	registry    discovery.Registry
}

func New(registry discovery.Registry, bookgateway *bookgateway.Gateway) *Gateway {
	return &Gateway{
		registry:    registry,
		bookgateway: bookgateway,
	}
}

func (g *Gateway) GetAll(ctx context.Context) (*model.Response[[]rentmodel.Rent], error) {
	addrs, err := g.registry.ServiceAddresses(ctx, "rent")

	if err != nil {
		return nil, err
	}

	userCtx := ctx.Value(constant.USER_ID_KEY)

	userID, ok := userCtx.(string)

	if !ok {
		return nil, errors.New("id is invalid")
	}

	url := "http://" + addrs[rand.Intn(len(addrs))] + "/rents/" + userID

	log.Printf("calling rent service. Request: GET " + url)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}
	req = req.WithContext(ctx)

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

	var v *model.Response[[]rentmodel.Rent]

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}

func (g *Gateway) Create(ctx context.Context, rentInput rentmodel.RentInput) (*model.Response[rentmodel.Rent], error) {
	responseData, err := g.bookgateway.GetByID(ctx, strconv.Itoa(rentInput.BookID))

	isFailed := err != nil || responseData == nil

	if isFailed {
		return nil, err
	}

	rentedBook := responseData.Data

	rentInput.BookTitle = rentedBook.Title

	addrs, err := g.registry.ServiceAddresses(ctx, "rent")

	if err != nil {
		return nil, err
	}

	url := "http://" + addrs[rand.Intn(len(addrs))] + "/rents"

	log.Printf("calling rent service. Request: POST " + url)

	jsonBody, err := json.Marshal(&rentInput)

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

	var v *model.Response[rentmodel.Rent]

	if err := json.NewDecoder(resp.Body).Decode(&v); err != nil {
		return nil, err
	}
	return v, nil
}
