package main

import (
	"fmt"

	"github.com/ovh/go-ovh/ovh"
	"github.com/spf13/viper"
)

func main() {
	// Setup configuration with viper
	viper.SetDefault("API_ENDPOINT", "ovh-eu")
	viper.BindEnv("API_ENDPOINT")
	viper.BindEnv("APPLICATION_KEY")
	viper.BindEnv("APPLICATION_SECRET")
	viper.BindEnv("CONSUMER_KEY")
	viper.BindEnv("DNS_ZONE")

	apiEndpoint := viper.GetString("API_ENDPOINT")
	appKey := viper.GetString("APPLICATION_KEY")
	appSecret := viper.GetString("APPLICATION_SECRET")
	consumerKey := viper.GetString("CONSUMER_KEY")

	client, err := ovh.NewClient(
		apiEndpoint,
		appKey,
		appSecret,
		consumerKey,
	)
	if err != nil {
		panic(err)
	}

	dnsZone := viper.GetString("DNS_ZONE")
	query := fmt.Sprintf("/domain/zone/%v/refresh", dnsZone)

	if err = client.Post(query, nil, nil); err != nil {
		panic(err)
	}
}
