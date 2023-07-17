package plaid

import (
	"context"
	config "finance/config"
	"fmt"
	"net/http"
	"strings"
	"time"

	plaidAPI "github.com/plaid/plaid-go/v3/plaid"
)

type PlaidService struct {
	client *plaidAPI.APIClient
}

var environments = map[string]plaidAPI.Environment{
	"sandbox":     plaidAPI.Sandbox,
	"development": plaidAPI.Development,
	"production":  plaidAPI.Production,
}

func Init(conf config.Configuration) *PlaidService {
	plaidConfig := plaidAPI.NewConfiguration()
	plaidConfig.AddDefaultHeader("PLAID-CLIENT-ID", conf.ClientID)
	plaidConfig.AddDefaultHeader("PLAID-SECRET", conf.Secret)
	plaidConfig.UseEnvironment(environments[conf.Environment])
	client := plaidAPI.NewAPIClient(plaidConfig)
	return &PlaidService{
		client: client,
	}
}

func (s *PlaidService) CreateLinkToken(conf config.Configuration, req *http.Request) string {
	countries := conf.Countries
	countriesArr := strings.Split(countries, ",")
	countryCodes := []plaidAPI.CountryCode{}

	products := conf.Products
	productArr := strings.Split(products, ",")
	productCodes := []plaidAPI.Products{}

	ctx := context.Background()

	fmt.Println(countries)
	fmt.Println(countriesArr)

	for _, countryCodeStr := range countriesArr {
		countryCodes = append(countryCodes, plaidAPI.CountryCode(countryCodeStr))
	}

	user := plaidAPI.LinkTokenCreateRequestUser{
		ClientUserId: time.Now().String(),
	}

	for _, productStr := range productArr {
		productCodes = append(productCodes, plaidAPI.Products(productStr))
	}

	request := plaidAPI.NewLinkTokenCreateRequest("Finances", "en", countryCodes, user)
	request.SetProducts(productCodes)
	request.SetRedirectUri("http://localhost:3000/")

	linkTokenCreateResp, _, err := s.client.PlaidApi.LinkTokenCreate(ctx).LinkTokenCreateRequest(*request).Execute()

	fmt.Println(request.GetRedirectUri())

	if err != nil {
		fmt.Println(err)
		return ""
	}

	return linkTokenCreateResp.GetLinkToken()
}
