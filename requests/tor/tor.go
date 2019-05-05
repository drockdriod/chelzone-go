package tor

import (
	"github.com/cretz/bine/tor"
	"context"
	"net/http"
	"time"
)

var torClient *http.Client

func Connect(ctx context.Context) (*http.Client, error){
	t, err := tor.Start(ctx, nil)
	if err != nil {
		return nil,err
	}

	// Wait at most a minute to start network and get
	dialCtx, dialCancel := context.WithTimeout(context.Background(), time.Minute)

	defer dialCancel()

	// Make connection
	dialer, err := t.Dialer(dialCtx, nil)
	if err != nil {
		return nil,err
	}
	torClient = &http.Client{Transport: &http.Transport{DialContext: dialer.DialContext}}
	
	return torClient, nil
}