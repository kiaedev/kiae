package loki

import (
	"fmt"
	"strconv"
	"time"

	"github.com/go-resty/resty/v2"
)

type Client struct {
	cli *resty.Client

	host string
}

func NewLoki(host string) *Client {
	return &Client{
		cli:  resty.New(),
		host: host,
	}
}

func (l *Client) QueryRange(query string, limit int64, start, end time.Time, direction string) ([]Result, error) {
	params := map[string]string{
		"query":     query,
		"limit":     strconv.FormatInt(limit, 10),
		"start":     strconv.FormatInt(start.Unix(), 10),
		"end":       strconv.FormatInt(end.Unix(), 10),
		"direction": direction,
	}

	var resp Response
	_, err := l.cli.R().SetQueryParams(params).SetResult(&resp).Get(fmt.Sprintf("%s/loki/api/v1/query_range", l.host))
	if err != nil {
		return nil, err
	}

	return resp.Data.Result, err
}
