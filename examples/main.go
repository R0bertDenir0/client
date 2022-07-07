package main

import (
	"fmt"

	d "dkg-client-go"

	"github.com/ivpusic/golog"
)

const OT_NODE_HOSTNAME string = "167.99.210.73"
const OT_NODE_PORT int = 8900

func main() {
	opt := d.DkgClientOptions{Endpoint: OT_NODE_HOSTNAME, Port: OT_NODE_PORT, UseSSL: false, LogLevel: golog.ERROR, MaxNumberOfRetries: 15}

	// Initialize connection to your DKG Node
	dkg, err := d.NewDkgClient(opt)
	if err != nil {
		panic(err)
	}

	// Get info about endpoint that you connected to ---WORKING---
	//out, err := dkg.Client.NodeInfo()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println("Node Info:")
	//fmt.Println(string(out))

	// Provisioning an asset ---WORKING---
	// createOpt := d.CreateOptions{Filepath: "./kg-example.json", Data: "", Keywords: []string{"Product", "Executive Objects", "ACME"}}
	// createOut, err := dkg.Assets.Create(createOpt)
	// if err != nil {
	//     panic(err)
	// }
	// fmt.Println(string(createOut))

	// Updating the previously provisioned asset ---WORKING---
	// Set this to the UAL returned when provisioning the asset in the function above
	// This value, the UAL, is used to identify certain asset and update it
	// The function above should have returned a JSON, searrch for "UALs" entry inside of the JSON, thats your UAL
	// ual := ""
	// updateOpt := d.UpdateOptions{Filepath: "./kg-example.json", Data: "", Keywords: []string{"Product", "Executive Objects", "ACME"}}
	// updateOut, err := dkg.Assets.Update(ual, updateOpt)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(updateOut))

	// Publishing a dataset ---WORKING---
	// publishOpt := d.PublishOptions{Filepath: "./kg-example.json", Data: "", Keywords: []string{"Product", "Executive Objects", "ACME"}}
	// publishOut, err := dkg.Client.Publish(publishOpt)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(string(publishOut))

	// Resolving assertion
	//resolveOpt := d.ResolveRequestOptions{IDS: []string{
	//	"066787bc7269c062fe73b0ebb004c258e07151777e6dfba027fea046df5caf7c",
	//	"2286826799d0a32a6f0eec7813fcb627910be45fca21f6378cb26ca95097c939"},
	//}
	//resolveOut, err := dkg.Client.Resolve(resolveOpt)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(resolveOut)

	// Search assertions
	//searchAssertionsOpt := d.SearchRequestOptions{
	//	Query:            "Product",
	//	ResultType:       "assertions",
	//	Prefix:           true,
	//	Limit:            20,
	//	Issuers:          []string{"Issuer 1", "Issuer 2"}, //Unused by the API, but was in the js-code
	//	SchemaTypes:      "Schema Type 1",                  //Unused by the API, but was in the js-code
	//	NumbersOfResults: 5,
	//	Timeout:          25,
	//}
	//searchAssertionsOut, err := dkg.Client.Search(searchAssertionsOpt)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(searchAssertionsOut)

	// Search entities
	//searchEntitiesOpt := d.SearchRequestOptions{
	//	Query:            "Product",
	//	ResultType:       "entities",
	//	Prefix:           true,
	//	Limit:            20,
	//	Issuers:          []string{"Issuer 1", "Issuer 2"}, //Unused by the API, but was in the js-code
	//	SchemaTypes:      "Schema Type 1",                  //Unused by the API, but was in the js-code
	//	NumbersOfResults: 5,
	//	Timeout:          25,
	//}
	//searchEntitiesOut, err := dkg.Client.Search(searchEntitiesOpt)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(searchEntitiesOut)

	// Execute sparql query on dkg ---WORKING---
// 	q := `PREFIX schema: <http://schema.org/>
// 	CONSTRUCT { ?s ?p ?o }
// 	WHERE {
// 		GRAPH ?g {
// 		?s ?p ?o .
// 		?s schema:hasVisibility ?v
// 	}
// }`
// 	queryOpt := d.QueryOptions{Query: q, Type: "construct"}
//
// 	queryOut, err := dkg.Client.Query(queryOpt)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(queryOut))

	// Validate some triples that we can get querying
	//validateOpt := d.ValidateOptions{Nquads: []string{
	//	"<did:dkg:87c4edd8695ab8a493015361b5a564c82f90f4c5e6c5e5cc9adccf4e11a63ad7> <http://schema.org/hasType> \"person\" .",
	//	"<did:dkg:25304bfd61ddcf490dfe852b883c01918768c114a84dcda0ac4aff179ff9ba65> <http://schema.org/hasType> \"person\" ."},
	//}
	//validateOut, err := dkg.Client.Validate(validateOpt)
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(validateOut)
}
