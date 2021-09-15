package smoove

import (
	"fmt"
	"github.com/Buccaneer69/wc-api-go/client"
	"github.com/Buccaneer69/wc-api-go/options"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	factory := client.Factory{}
	c := factory.NewClient(options.Basic{
		URL:    "http://danzigermarket.local/",
		Key:    "ck_20bea696699e76d68cb94f64c3347a57b8b1033d",
		Secret: "cs_67a5b0363d1f7456a9088942cfdd09350565d9e6",
		Options: options.Advanced{
			WPAPI:       true,
			WPAPIPrefix: "/wp-json/",
			Version:     "wc/v3",
		},
	})

	// Further using of client ... 
	// POST /wp-json/wc/v3/products
	// req, err := c.Post(endpoint, nil, body)

	if r, err := c.Get("products", nil); err != nil {
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