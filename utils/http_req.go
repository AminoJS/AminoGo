package utils

import (
	"errors"
	"fmt"
	"github.com/AminoJS/AminoGo/stores"
	"github.com/imroc/req"
	"time"
)

func Get(url string) (*req.Resp, error) {
	if url == "" {
		return &req.Resp{}, errors.New("URL cannot be emtpy")
	}
	SID := stores.Get("SID")
	if SID == nil || SID == "" {
		return &req.Resp{}, errors.New("missing SID in state, try using aminogo.Login() first")
	}
	header := req.Header{
		"NDCAUTH": fmt.Sprintf("sid=%s", SID),
	}

	req.SetTimeout(30 * time.Second)
	res, err := req.Get(url, header)
	return res, err
}
