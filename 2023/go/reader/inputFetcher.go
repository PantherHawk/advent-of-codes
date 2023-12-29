package reader

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net"
	"net/http"
	"path/filepath"
	"time"
)

const timeout = 30

func fetchInput(day string) ([]string, error) {
	// http call
	client := &http.Client{
		Transport: &http.Transport{
			Dial: (&net.Dialer{
				Timeout:   time.Second * timeout,
				KeepAlive: time.Second * timeout,
			}).Dial,
			TLSHandshakeTimeout:   time.Second * timeout,
			ResponseHeaderTimeout: time.Second * timeout,
			ExpectContinueTimeout: time.Second * timeout,
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
		Timeout: time.Second * timeout,
	}
	url := formatQuery(day)
	fmt.Printf("url: %s\n\n", url)
	req, err := http.NewRequestWithContext(context.TODO(), "GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Add("Content-Type", "application/json")
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() { _ = resp.Body.Close() }()
	if resp.StatusCode == http.StatusNotFound {
		return nil, errors.New(fmt.Sprintf("Advent of Code for day %s not found", day))
	}
	if resp.StatusCode == http.StatusInternalServerError {
		return nil, errors.New(fmt.Sprintf("Advent of Code for day %s not available", day))
	}
	if resp.StatusCode == http.StatusBadRequest {
		b, err := io.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(string(b))
	}
	if resp.StatusCode != http.StatusOK {
		var b string

		decoder := json.NewDecoder(resp.Body)
		err = decoder.Decode(&b)
		if err != nil {
			return nil, err
		}

		return nil, errors.New(b)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	response := make([]string, 0)
	err = json.Unmarshal(body, &response)
	if err != nil {
		return nil, err
	}
	fmt.Printf("response: %+v/n", response)
	// return input
	return response, nil
}

func formatQuery(day string) string {
	// look up the input directory
	filePath, _ := filepath.Abs("../")
	year := filepath.Base(filePath)
	// from the input directory look up to see which year we are in
	s := fmt.Sprintf("https://adventofcode.com/%s/day/%s/input", year, day)
	return s
}
