package main

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type Idp struct {
	MetadataB64                     string `json:"metadataB64,omitempty"`
	PartnerName                     string `json:"partnerName,omitempty"`
	NameIDFormat                    string `json:"nameIDFormat,omitempty"`
	SsoProfile                      string `json:"ssoProfile,omitempty"`
	ProviderID                      string `json:"providerID,omitempty"`
	AssertionConsumerURL            string `json:"assertionConsumerURL,omitempty"`
	LogoutRequestURL                string `json:"logoutRequestURL,omitempty"`
	LogoutResponseURL               string `json:"logoutResponseURL,omitempty"`
	AdminManualCreation             string `json:"adminManualCreation,omitempty"`
	DisplaySigningCertDN            string `json:"displaySigningCertDN,omitempty"`
	DisplaySigningCertIssuerDN      string `json:"displaySigningCertIssuerDN,omitempty"`
	DisplaySigningCertStart         string `json:"displaySigningCertStart,omitempty"`
	DisplaySigningCertExpiration    string `json:"displaySigningCertExpiration,omitempty"`
	DisplayEncryptionCertDN         string `json:"displayEncryptionCertDN,omitempty"`
	DisplayEncryptionCertIssuerDN   string `json:"displayEncryptionCertIssuerDN,omitempty"`
	DisplayEncryptionCertStart      string `json:"displayEncryptionCertStart,omitempty"`
	DisplayEncryptionCertExpiration string `json:"displayEncryptionCertExpiration,omitempty"`
}

type IDPPartner struct {
	MetadataB64     string `json:"metadataB64,omitempty"`
	MetadataURL     string `json:"metadataURL,omitempty"`
	PartnerType     string `json:"partnerType,omitempty"`
	TenantName      string `json:"tenantName,omitempty"`
	TenantURL       string `json:"tenantURL,omitempty"`
	PartnerName     string `json:"partnerName,omitempty"`
	NameIDFormat    string `json:"nameIDFormat,omitempty"`
	SsoProfile      string `json:"ssoProfile,omitempty"`
	AttributeLDAP   string `json:"attributeLDAP,omitempty"`
	AttributeSAML   string `json:"attributeSAML,omitempty"`
	FaWelcomePage   string `json:"faWelcomePage,omitempty"`
	GenerateNewKeys string `json:"generateNewKeys,omitempty"`
	ValidityNewKeys string `json:"validityNewKeys,omitempty"`
	Preverify       bool   `json:"preverify,omitempty"`
	ProviderID      string `json:"providerID,omitempty"`
	SsoURL          string `json:"ssoURL,omitempty"`
}

func readFromFile(filename string) []string {
	// Read the URL from a file
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}

	lines := strings.Split(string(data), "\n")

	//url := strings.TrimSpace(string(urlBytes))

	// Print the response data
	fmt.Println(lines)
	return lines
}

func getDataFromURL(url, username, password string) io.ReadCloser {
	// Define a JSON string
	jsonString := `{"MetadataB64": "value", "PartnerName": "IDP"}`

	// Create an io.ReadCloser that reads from the JSON string
	reader := ioutil.NopCloser(strings.NewReader(jsonString))
	defer reader.Close()
	return reader

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}

	// Set the request headers
	req.Header.Set("Content-Type", "application/json")

	// Set the basic authentication credentials
	req.SetBasicAuth(username, password)

	// Send the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	return resp.Body
}

func getData() {
	filename := os.Args[0]
	username := os.Args[1]
	password := os.Args[2]
	filename = "test"
	data := readFromFile(filename)
	url := data[0]

	for _, v := range data[1:] {
		fmt.Println("v")
		fmt.Println(v)
		splitted := strings.Split(v, ",")
		out := splitted[0]
		in := splitted[1]
		IDPPartner := new(IDPPartner)
		reader := getDataFromURL(url+"/"+out, username, password)

		// Decode the JSON response data into a MyData struct
		var dataParsed Idp
		if err := json.NewDecoder(reader).Decode(&dataParsed); err != nil {
			panic(err)
		}
		IDPPartner.MetadataB64 = cleanMetadata(dataParsed.MetadataB64)
		IDPPartner.PartnerName = in

		saveToFile(in, IDPPartner)

	}
}

func cleanMetadata(MetadataB64 string) string {
	MetadataB64 = strings.ReplaceAll(MetadataB64, "\n", "")
	MetadataB64 = strings.ReplaceAll(MetadataB64, "\t", "")
	MetadataB64 = strings.ReplaceAll(MetadataB64, "\\\"", "\"")
	MetadataB64 = base64.StdEncoding.EncodeToString([]byte(MetadataB64))
	return MetadataB64
}

func saveToFile(filename string, data interface{}) {
	// Open the file for writing
	file, err := os.Create(filename + ".json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Encode the data as JSON and write it to the file
	if err := json.NewEncoder(file).Encode(data); err != nil {
		panic(err)
	}
}
