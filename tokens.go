package magento

import (
  "bytes"
  "encoding/json"
  "net/http"
)

// Tokens service
type TokensService service

type AccessTokens struct {
  OAuthToken       string `json:"oauth_token"`
  OAuthTokenSecret string `json:"oauth_token_secret"`
}

// Get a request token. Reference: https://developer.adobe.com/commerce/webapi/get-started/authentication/gs-authentication-oauth/#get-a-request-token
func (service *TokensService) GetRequestTokens() (*AccessTokens, *http.Response, error) {
  _url := "/request"

  data := &bytes.Buffer{}
  req, _ := service.client.NewAuthRequest("POST", _url, "")
  response, err := service.client.Do(req, data)

  if err != nil {
    return nil, response, err
  }

  accessTokens := new(AccessTokens)
  _data, err := service.client.parseStringQueryParamsAsJson(data.String())

  if err != nil {
    return nil, response, err
  }

  json.Unmarshal(_data, accessTokens)

  return accessTokens, response, nil
}

// Get Access Tokens. Reference: https://crisp-plugin-dinis.ngrok.io/config/success?website_id=d429990d-5cc0-42cd-a513-39e7724c179a&shop=dinis-crisp-tests.myshopify.com&shopify_init=true
func (service *TokensService) GetAccessTokens(oauthVerifier string) (*AccessTokens, *http.Response, error) {
  _url := "/access"

  data := &bytes.Buffer{}
  req, _ := service.client.NewAuthRequest("POST", _url, oauthVerifier)
  response, err := service.client.Do(req, data)

  if err != nil {
    return nil, response, err
  }

  accessTokens := new(AccessTokens)
  _data, err := service.client.parseStringQueryParamsAsJson(data.String())

  if err != nil {
    return nil, response, err
  }

  json.Unmarshal(_data, accessTokens)

  return accessTokens, response, nil
}