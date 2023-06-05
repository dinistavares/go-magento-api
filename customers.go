package magento

import (
  "net/http"
)

// Customer service
type CustomersService service

type CustomerItems struct {
  Customers *[]Customer `json:"items,omitempty"`
}

type Customer struct {
  ID                     int                 `json:"id,omitempty"`
  GroupID                int                 `json:"group_id,omitempty"`
  DefaultBilling         string              `json:"default_billing,omitempty"`
  DefaultShipping        string              `json:"default_shipping,omitempty"`
  Confirmation           string              `json:"confirmation,omitempty"`
  CreatedAt              string              `json:"created_at,omitempty"`
  UpdatedAt              string              `json:"updated_at,omitempty"`
  CreatedIn              string              `json:"created_in,omitempty"`
  Dob                    string              `json:"dob,omitempty"`
  Email                  string              `json:"email,omitempty"`
  Firstname              string              `json:"firstname,omitempty"`
  Lastname               string              `json:"lastname,omitempty"`
  Middlename             string              `json:"middlename,omitempty"`
  Prefix                 string              `json:"prefix,omitempty"`
  Suffix                 string              `json:"suffix,omitempty"`
  Gender                 int                 `json:"gender,omitempty"`
  StoreID                int                 `json:"store_id,omitempty"`
  Taxvat                 string              `json:"taxvat,omitempty"`
  WebsiteID              int                 `json:"website_id,omitempty"`
  Addresses              *[]Addresses         `json:"addresses,omitempty"`
  DisableAutoGroupChange int                 `json:"disable_auto_group_change,omitempty"`
  CustomAttributes       *[]CustomAttributes  `json:"custom_attributes,omitempty"`
}

type Region struct {
  RegionCode          string              `json:"region_code,omitempty"`
  Region              string              `json:"region,omitempty"`
  RegionID            int                 `json:"region_id,omitempty"`
}

type CustomAttributes struct {
  AttributeCode string `json:"attribute_code,omitempty"`
  Value         string `json:"value,omitempty"`
}

type Addresses struct {
  ID                  int                 `json:"id,omitempty"`
  CustomerID          int                 `json:"customer_id,omitempty"`
  Region              *Region              `json:"region,omitempty"`
  RegionID            int                 `json:"region_id,omitempty"`
  CountryID           string              `json:"country_id,omitempty"`
  Street              *[]string            `json:"street,omitempty"`
  Company             string              `json:"company,omitempty"`
  Telephone           string              `json:"telephone,omitempty"`
  Fax                 string              `json:"fax,omitempty"`
  Postcode            string              `json:"postcode,omitempty"`
  City                string              `json:"city,omitempty"`
  Firstname           string              `json:"firstname,omitempty"`
  Lastname            string              `json:"lastname,omitempty"`
  Middlename          string              `json:"middlename,omitempty"`
  Prefix              string              `json:"prefix,omitempty"`
  Suffix              string              `json:"suffix,omitempty"`
  VatID               string              `json:"vat_id,omitempty"`
  DefaultShipping     bool                `json:"default_shipping,omitempty"`
  DefaultBilling      bool                `json:"default_billing,omitempty"`
  CustomAttributes    []CustomAttributes  `json:"custom_attributes,omitempty"`
}

type SearchParams struct {
  CurrentPage     string    `url:"searchCriteria[currentPage],omitempty"`
  ConditionType   string    `url:"searchCriteria[filterGroups][0][filters][0][conditionType],omitempty"`
  FilterField     string    `url:"searchCriteria[filterGroups][0][filters][0][field],omitempty"`
  FilterValue     string    `url:"searchCriteria[filterGroups][0][filters][0][value],omitempty"`
  PageSize        string    `url:"searchCriteria[pageSize],omitempty"`
  SortDirection   string    `url:"searchCriteria[sortOrders][0][direction],omitempty"`
  SortField       string    `url:"searchCriteria[sortOrders][0][field],omitempty"`
}

// Get a customer by ID. Reference: https://adobe-commerce.redoc.ly/2.4.6-admin/tag/customerscustomerId#operation/GetV1CustomersCustomerId
func (service *CustomersService) Get(customerID string) (*Customer, *http.Response, error) {
  _url := "/customers/" + customerID
  req, _ := service.client.NewRequest("GET", _url, nil, nil)

  customer := new(Customer)
  response, err := service.client.Do(req, customer)

  if err != nil {
    return nil, response, err
  }

  return customer, response, nil
}

// Search for a customer. Reference: https://adobe-commerce.redoc.ly/2.4.6-admin/tag/customerssearch#operation/GetV1CustomersSearch
func (service *CustomersService) Search(opts *SearchParams) (*[]Customer, *http.Response, error) {
  _url := "/customers/search"
  req, _ := service.client.NewRequest("GET", _url, opts, nil)

  customerItems := new(CustomerItems)
  response, err := service.client.Do(req, customerItems)

  if err != nil {
    return nil, response, err
  }

  return customerItems.Customers, response, nil
}