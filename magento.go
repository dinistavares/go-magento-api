package magento

import (
  "bytes"
  "crypto/hmac"
  "crypto/sha256"
  "encoding/base64"
  "encoding/json"
  "errors"
  "fmt"
  "io"
  "io/ioutil"
  "net/http"
  "net/url"
  "strconv"
  "strings"
  "time"

  "github.com/google/go-querystring/query"
  "github.com/google/uuid"
)

const (
  defaultRestEndpointVersion   = "V1"
  defaultHeaderName            = "Authorization"
  acceptedContentType          = "application/json"
  userAgent                    = "go-magento-api/1.1"
  clientRequestRetryAttempts   = 2
  clientRequestRetryHoldMillis = 1000
)

var (
  authHeaderParams = []string{
    "oauth_consumer_key",
    "oauth_nonce",
    "oauth_signature",
    "oauth_signature_method",
    "oauth_timestamp",
    "oauth_token",
    "oauth_version",
    "oauth_verifier",
  }

  errorDoAllAttemptsExhausted = errors.New("all request attempts were exhausted")
  errorDoAttemptNilRequest = errors.New("request could not be constructed")
)

type ClientConfig struct {
  HttpClient          *http.Client
  RestEndpointURL     string
  RestEndpointVersion string
}

type auth struct {
  HeaderName          string
  ConsumerKey         string
  ConsumerSecret      string
  AccessToken         string
  AccessTokenSecret   string
}

type Client struct {
  config   *ClientConfig
  client   *http.Client
  auth     *auth
  baseURL  *url.URL
  tokenURL *url.URL

  Tokens      *TokensService
  Customers   *CustomersService
  Orders      *OrdersService
}

type service struct {
  client *Client
}

type errorResponse struct {
  Response *http.Response

  Code    string    `json:"code"`
  Message string    `json:"message"`
  Data    ErrorData `json:"data"`
}

type ErrorData struct {
  Status int `json:"status"`
}

func (response *errorResponse) Error() string {
  return fmt.Sprintf("%v %v: %d %v",
    response.Response.Request.Method, response.Response.Request.URL,
    response.Response.StatusCode, response.Message)
}

func New(shopURL string) (*Client, error) {
  if shopURL == "" {
    return nil, errors.New("store url is required")
  }

  shopURL = strings.TrimSuffix(shopURL, "/")

  config := ClientConfig{
    HttpClient: http.DefaultClient,
  }

  config.HttpClient = http.DefaultClient
  config.RestEndpointURL = shopURL
  config.RestEndpointVersion = defaultRestEndpointVersion

  // Create client
  baseURL, _ := url.Parse(config.RestEndpointURL + "/rest/default/" + defaultRestEndpointVersion)
  tokenURL, err := url.Parse(config.RestEndpointURL + "/oauth/token/")

  if err != nil {
    return nil, err
  }

  client := &Client{config: &config, client: config.HttpClient, auth: &auth{}, baseURL: baseURL, tokenURL: tokenURL}

  // Map services
  client.Tokens = &TokensService{client: client}
  client.Customers = &CustomersService{client: client}
  client.Orders = &OrdersService{client: client}

  return client, nil
}

// Authenticate saves authenitcation parameters for user
func (client *Client) Authenticate(consumerKey, consumerSecret, accessToken, AccessTokenSecret string) {
  client.auth = &auth{
    ConsumerKey: consumerKey,
    ConsumerSecret: consumerSecret,
    AccessToken: accessToken,
    AccessTokenSecret: AccessTokenSecret,
    HeaderName: defaultHeaderName,
  }
}

// NewRequest creates an API request
func (client *Client) NewRequest(method, urlStr string, opts interface{}, body interface{}) (*http.Request, error) {
  params := make(map[string]string)
  // Append Query Params to URL
  if opts != nil {
    queryParams, err := query.Values(opts)

    if err != nil {
      return nil, err
    }

    for key, value := range queryParams {
      if len(value) > 0 {
        params[key] = value[0]
      }
    }

    rawQuery := queryParams.Encode()

    if rawQuery != "" {
      urlStr += "?" + rawQuery
    }
  }

  rel, err := url.Parse(client.config.RestEndpointVersion + urlStr)
  if err != nil {
    return nil, err
  }

  _url := client.baseURL.ResolveReference(rel)

  var buf io.ReadWriter
  if body != nil {
    buf = new(bytes.Buffer)

    err := json.NewEncoder(buf).Encode(body)
    if err != nil {
      return nil, err
    }
  }

  req, err := http.NewRequest(method, _url.String(), buf)
  if err != nil {
    return nil, err
  }

  authenticationHeader := client.acquireAuthenticationHeader(method, _url.String(), params)

  req.Header.Add(client.auth.HeaderName, authenticationHeader)
  req.Header.Add("Accept", acceptedContentType)
  req.Header.Add("Content-type", acceptedContentType)
  req.Header.Add("User-Agent", userAgent)

  return req, nil
}


// NewAuthRequest creates an API request
func (client *Client) NewAuthRequest(method, urlStr, oauthVerifier string) (*http.Request, error) {
  rel, err := url.Parse("." + urlStr)
  if err != nil {
    return nil, err
  }

  _url := client.tokenURL.ResolveReference(rel)

  var buf io.ReadWriter

  req, err := http.NewRequest(method, _url.String(), buf)
  if err != nil {
    return nil, err
  }

  params := make(map[string]string)
  if oauthVerifier != "" {
    params["oauth_verifier"] = oauthVerifier
  }

  authenticationHeader := client.acquireAuthenticationHeader(method, _url.String(), params)

  if authenticationHeader == "" {
    return nil, errors.New("authentication header missing")
  }

  req.Header.Add(defaultHeaderName, authenticationHeader)
  req.Header.Add("Accept", acceptedContentType)
  req.Header.Add("Content-type", acceptedContentType)
  req.Header.Add("User-Agent", userAgent)

  return req, nil
}

// Do sends an API request
func (client *Client) Do(req *http.Request, v interface{}) (*http.Response, error) {
  var lastErr error

  attempts := 0

  for attempts < clientRequestRetryAttempts {
    // Hold before this attempt? (ie. not first attempt)
    if attempts > 0 {
      time.Sleep(clientRequestRetryHoldMillis * time.Millisecond)
    }

    // Dispatch request attempt
    attempts++
    resp, shouldRetry, err := client.doAttempt(req, v)

    // Return response straight away? (we are done)
    if !shouldRetry {
      return resp, err
    }

    // Should retry: store last error (we are not done)
    lastErr = err
  }

  // Set default error? (all attempts failed, but no error is set)
  if lastErr == nil {
    lastErr = errorDoAllAttemptsExhausted
  }

  // All attempts failed, return last attempt error
  return nil, lastErr
}

func (client *Client) doAttempt(req *http.Request, v interface{}) (*http.Response, bool, error) {
  if req == nil {
    return nil, false, errorDoAttemptNilRequest
  }

  resp, err := client.client.Do(req)

  if checkRequestRetry(resp, err) {
    return nil, true, err
  }

  defer resp.Body.Close()

  err = checkResponse(resp)
  if err != nil {
    return resp, false, err
  }

  if v != nil {
    if w, ok := v.(io.Writer); ok {
      io.Copy(w, resp.Body)
    } else {
      err = json.NewDecoder(resp.Body).Decode(v)
      if err == io.EOF {
        err = nil
      }
    }
  }

  return resp, false, err
}

// checkRequestRetry checks if should retry request
func checkRequestRetry(response *http.Response, err error) bool {
  // Low-level error, or response status is a server error? (HTTP 5xx)
  if err != nil || response.StatusCode >= 500 {
    return true
  }

  // No low-level error (should not retry)
  return false
}

// checkResponse checks response for errors
func checkResponse(response *http.Response) error {
  // No error in response? (HTTP 2xx)
  if code := response.StatusCode; 200 <= code && code <= 299 {
    return nil
  }

  // Map response error data (eg. HTTP 4xx)
  errorResponse := &errorResponse{Response: response}

  data, err := ioutil.ReadAll(response.Body)
  if err == nil && data != nil {

    if strings.Contains(response.Header.Get("Content-Type"), "text/html") {
      data = []byte(`{"message": "` + string(data) + `"}`)
    }

    json.Unmarshal(data, errorResponse)
  }

  return errorResponse
}

func (client *Client) acquireAuthenticationHeader(method, path string, params map[string]string) string {
  nonce := uuid.New()

  vals := url.Values{}
  vals.Add("oauth_nonce", nonce.String())
  vals.Add("oauth_consumer_key", client.auth.ConsumerKey)
  vals.Add("oauth_signature_method", "HMAC-SHA256")
  vals.Add("oauth_timestamp", strconv.Itoa(int(time.Now().Unix())))
  vals.Add("oauth_version", "1.0")

  if client.auth.AccessToken != "" {
    vals.Add("oauth_token", client.auth.AccessToken)
  }

  for key, value := range params {
    vals.Add(key, value)
  }

  // url.QueryEscape escapes spaces into "+" characters
  parameterString := strings.Replace(vals.Encode(), "+", "%20", -1)
  signatureBase := strings.ToUpper(method) + "&" + url.QueryEscape(strings.Split(path, "?")[0]) + "&" + url.QueryEscape(parameterString)
  signingKey := strings.Join([]string{
      url.QueryEscape(client.auth.ConsumerSecret),
      url.QueryEscape(client.auth.AccessTokenSecret),
    },
  "&")

  h := hmac.New(sha256.New, []byte(signingKey))
  h.Write([]byte(signatureBase))

  signature := base64.StdEncoding.EncodeToString(h.Sum(nil))

  vals.Add("oauth_signature", signature)

  authHeader := "OAuth "

  for _, key := range authHeaderParams {
    if vals.Get(key) != "" {
      authHeader += key + "=" + url.QueryEscape(vals.Get(key))
      authHeader += ","
    }
  }

  authHeader = strings.TrimSuffix(authHeader, ",")

  return authHeader
}

func (client *Client) parseStringQueryParamsAsJson(body string) ([]byte, error) {
  _urlParams, _ := url.ParseQuery(body)
  queryMap := make(map[string]interface{})

  for key, value := range _urlParams {
    queryMap[key] = value[0]
  }

  _data, err := json.Marshal(queryMap)

  if err != nil {
    return nil, err
  }

  return _data, nil
}