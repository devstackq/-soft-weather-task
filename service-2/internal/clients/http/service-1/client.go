package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-resty/resty/v2"
	"soft-weater/internal/models"
	"time"
)

type Service interface {
	GetAccount(ctx context.Context, userID string) (models.Account, error)
	IncreaseDebt(ctx context.Context, account models.Account) error
}

type Client struct {
	cl *resty.Client
}

const (
	baseUrl = "http://localhost"
	port    = ":9090"
)

func NewService() Service {
	httpCl := resty.New().
		SetTimeout(10 * time.Second).
		SetBaseURL(baseUrl + port)

	return &Client{
		cl: httpCl,
	}
}

func (c *Client) GetAccount(ctx context.Context, userID string) (res models.Account, err error) {

	path := fmt.Sprint(`/v1/account/`, userID)

	resp, err := c.cl.R().
		SetResult(&res).
		Get(path)

	if err != nil {
		return res, err
	}
	if resp.IsError() {
		return models.Account{}, errors.New(resp.Status() + resp.String())
	}

	return res, nil
}

func (c *Client) IncreaseDebt(ctx context.Context, account models.Account) error {
	path := `/v1/account/debt/increase`

	resp, err := c.cl.R().
		SetBody(account).
		Put(path)
	if err != nil {
		return err
	}
	if resp.IsError() {
		return errors.New(resp.Status() + resp.String())
	}

	return nil
}
