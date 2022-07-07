package main

import (
	"fmt"

	d "dkg-client-go"
	"encoding/json"

	"github.com/ivpusic/golog"
)

const OT_NODE_HOSTNAME string = "167.99.210.73"
const OT_NODE_PORT int = 8900

func main() {
	opt := d.DkgClientOptions{Endpoint: OT_NODE_HOSTNAME, Port: OT_NODE_PORT, UseSSL: false, LogLevel: golog.DEBUG, MaxNumberOfRetries: 15}

	// Initialize connection to your DKG Node
	dkg, err := d.NewDkgClient(opt)
	if err != nil {
		panic(err)
	}

	////////////////////////
	// Function: NodeInfo //
	////////////////////////

	// Get info about endpoint that you connected to
	out, err := dkg.Client.NodeInfo()
	if err != nil {
	panic(err)
	}
	fmt.Println(string(out))

	//////////////////////
	// Function: Create //
	//////////////////////

	// Provisioning an asset
	createOpt := d.CreateOptions{Filepath: "./kg-example.json", Data: "", Keywords: []string{"Product", "Executive Objects", "ACME"}}
	createOut, err := dkg.Assets.Create(createOpt)
	if err != nil {
	    panic(err)
	}
	fmt.Println(string(createOut))

	//////////////////////
	// Function: Update //
	//////////////////////

	// Updating the previously provisioned asset
	// Set this to the UAL returned when provisioning the asset in the function above
	// This value, the UAL, is used to identify certain asset and update it
	// The function above should have returned a JSON, searrch for "UALs" entry inside of the JSON, thats your UAL
	ual := ""
	updateOpt := d.UpdateOptions{Filepath: "./kg-example.json", Data: "", Keywords: []string{"Product", "Executive Objects", "ACME"}}
	updateOut, err := dkg.Assets.Update(ual, updateOpt)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(updateOut))

	///////////////////////
	// Function: Publish //
	///////////////////////

	// Publishing a dataset
	publishOpt := d.PublishOptions{Filepath: "./kg-example.json", Data: "", Keywords: []string{"Product", "Executive Objects", "ACME"}}
	publishOut, err := dkg.Client.Publish(publishOpt)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(publishOut))


	///////////////////////
	// Function: Resolve //
	///////////////////////

	// Resolving assertion
	resolveOpt := d.ResolveRequestOptions{IDS: []string{
	"066787bc7269c062fe73b0ebb004c258e07151777e6dfba027fea046df5caf7c",
	"2286826799d0a32a6f0eec7813fcb627910be45fca21f6378cb26ca95097c939"},
	}
	resolveOut, err := dkg.Client.Resolve(resolveOpt)
	if err != nil {
	    panic(err)
	}
	fmt.Println(string(resolveOut))

	//////////////////////
	// Function: Search //
	//////////////////////

	// In this example, we search for assertions
	searchAssertionsOpt := d.SearchRequestOptions{
	Query:            "Product",
	ResultType:       "assertions",
	Prefix:           true,
	Limit:            20,
	Issuers:          []string{"Issuer 1", "Issuer 2"}, //Unused by the API, but was in the js-code
	SchemaTypes:      "Schema Type 1",                  //Unused by the API, but was in the js-code
	NumbersOfResults: 10,
	Timeout:          25,
	}
	searchAssertionsOut, err := dkg.Client.Search(searchAssertionsOpt)
	if err != nil {
	    panic(err)
	}
	fmt.Println(string(searchAssertionsOut))

	// In this example, we search for entities
	searchEntitiesOpt := d.SearchRequestOptions{
	Query:            "Product",
	ResultType:       "entities",
	Prefix:           true,
	Limit:            20,
	Issuers:          []string{"Issuer 1", "Issuer 2"}, //Unused by the API, but was in the js-code
	SchemaTypes:      "Schema Type 1",                  //Unused by the API, but was in the js-code
	NumbersOfResults: 5,
	Timeout:          25,
	}
	searchEntitiesOut, err := dkg.Client.Search(searchEntitiesOpt)
	if err != nil {
	    panic(err)
	}
	fmt.Println(string(searchEntitiesOut))

	/////////////////////
	// Function: Query //
	/////////////////////

	// Execute sparql query on dkg

	q := `PREFIX schema: <http://schema.org/>
	CONSTRUCT { ?s ?p ?o }
	WHERE {
		GRAPH ?g {
		?s ?p ?o .
		?s schema:hasVisibility ?v
	}
}`
	queryOpt := d.QueryOptions{Query: q, Type: "construct"}

	queryOut, err := dkg.Client.Query(queryOpt)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(queryOut))

	////////////////////////
	// Function: Validate //
	////////////////////////

	// Validate some triples that we can get querying

	// Search assertions
	searchAssertionsOpt := d.SearchRequestOptions{
	Query:            "Product",
	ResultType:       "assertions",
	Prefix:           true,
	Limit:            20,
	Issuers:          []string{"Issuer 1", "Issuer 2"}, //Unused by the API, but was in the js-code
	SchemaTypes:      "Schema Type 1",                  //Unused by the API, but was in the js-code
	NumbersOfResults: 10,
	Timeout:          25,
	}
	searchAssertionsOut, err := dkg.Client.Search(searchAssertionsOpt)
	if err != nil {
	    panic(err)
	}
	
	r := make(map[string]interface{})
	if err := json.Unmarshal(searchAssertionsOut, &r); err != nil {
	    panic(err)
	}


	// Why do we do this?
	// Well, to validate an assertion, we first need to get it's ID and some property to test
	// in order to actually test one that exists, we need to fetch on from the node
	// that's why we are running first a search function to get these values

	// This is what an assertion triple looks like, we need to validate this string but with an actual ID and property to test

	// "<did:dkg:25304bfd61ddcf490dfe852b883c01918768c114a84dcda0ac4aff179ff9ba65> <http://schema.org/hasType> \"person\" ."},
	id := r["itemListElement"].([]interface{})[0].(map[string]interface{})["result"].(map[string]interface{})["@id"].(string)
	typ := r["itemListElement"].([]interface{})[0].(map[string]interface{})["result"].(map[string]interface{})["metadata"].(map[string]interface{})["type"].(string)
	validateOpt := d.ValidateOptions{Nquads: []string{
	    fmt.Sprintf("<did:dkg:%s> <http://schema.org/hasType> \"%s\" .", id, typ),
	}}

	// Now that we have the triple, we validate it with the node
	validateOut, err := dkg.Client.Validate(validateOpt)
	if err != nil {
	    panic(err)
	}
	fmt.Println(validateOut)
}
