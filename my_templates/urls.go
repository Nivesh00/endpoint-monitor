package my_templates

import (
	"fmt"
	"strings"
)

// URLs array for single URLs
type Urls struct {
    Urls []Url `json:"urls"`
}

// URLs to be probed
type Url struct {
    Endpoint          string `json:"endpoint"`
    ResponseFormat    string `json:"responseFormat"`
    Contains        []string `json:"contains"`
    NotContains     []string `json:"notContains"`
}

// Convert URL data to string
func (url *Url) ToStr() string {
    
    // Convert array attrs to string
    in      := "[ '" + strings.Join(url.Contains, "', '") + "' ]"
    not_in  := "[ '" + strings.Join(url.NotContains, "', '") + "' ]"

    // Convert all attrs to string
    str_val := fmt.Sprintf("Endpoint: %s, RespFmt: %s, In: %s, NotIn: %s", 
                url.Endpoint,
                url.ResponseFormat,
                in,
                not_in)

    return str_val
}