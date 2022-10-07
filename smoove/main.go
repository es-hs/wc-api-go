package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gempages/wc-api-go/client"
	"github.com/gempages/wc-api-go/options"
)

type WooWebhookBatchUpdateBody struct {
	Create []WooWebhookParams `json:"create"`
	Delete []int              `json:"delete"`
}

type WooWebhookParams struct {
	Name        string `json:"name"`
	Topic       string `json:"topic"`
	DeliveryURL string `json:"delivery_url"`
}

var (
	WOOCOMMERCE_WEBHOOK_TOPICS = []string{"product.created", "product.updated", "product.deleted"}
	WOOCOMMERCE_APP_NAME       = "ES Headless Storefront"
)

func main() {
	factory := client.Factory{}
	c := factory.NewClient(options.Basic{
		URL:    "http://localhost:8080",
		Key:    "ck_016bd5cf27654b63e7c66750375c9f2eb211c865",
		Secret: "cs_a704b4c54849898727686d879c0b5a128dbbe6f1",
		Options: options.Advanced{
			WPAPI:       true,
			WPAPIPrefix: "/wp-json/",
			Version:     "wc/v3",
		},
	})

	// Further using of client ...
	// POST /wp-json/wc/v3/products
	// req, err := c.Post(endpoint, nil, body)

	topics := WOOCOMMERCE_WEBHOOK_TOPICS
	webhooks := []WooWebhookParams{}
	for _, topic := range topics {
		deliveryUrl := "https://norahaha.ap.ngrok.io/webhook"
		webhook := WooWebhookParams{
			Name:        fmt.Sprintf("%s - %s", WOOCOMMERCE_APP_NAME, topic),
			Topic:       topic,
			DeliveryURL: deliveryUrl,
		}
		webhooks = append(webhooks, webhook)
	}

	body := WooWebhookBatchUpdateBody{
		Create: webhooks,
		Delete: []int{},
	}

	if r, err := c.Post("webhooks/batch", nil, body); err != nil {
		log.Fatal(err)
	} else if r.StatusCode != http.StatusOK {
		log.Fatal("Unexpected StatusCode:", r)
	} else {
		defer r.Body.Close()
		if bodyBytes, err := ioutil.ReadAll(r.Body); err != nil {
			log.Fatal(err)
		} else {
			fmt.Println(string(bodyBytes))
		}
	}
}
