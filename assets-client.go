package client

import (
	"encoding/json"
	"errors"
)

type assetsClient struct {
	Client *nativeAbstractClient
}

func newAssetsClient(nac *nativeAbstractClient) (assetsClient, error) {
	return assetsClient{nac}, nil
}

type CreateOptions struct {
	Filepath string
	Data     string
	Keywords []string
}

func (ac *assetsClient) Create(options CreateOptions) ([]byte, error) {
    if options.Filepath == "" && options.Data == "" {
		return nil, errors.New("Please provide publish options in order to publish")
	}


	opt := PublishRequestOptions{"provision", options.Data, options.Filepath, options.Keywords, ""}

	// Make a query request
	resp, err := ac.Client.publishRequest(opt)
	if err != nil {
		return nil, err
	}

	createResponse := make(map[string]interface{})

	// Transform response to json struct
	if err := json.Unmarshal(resp, &createResponse); err != nil {
		return nil, errors.New("Could not unmarshal query request response")
	}

	// Get the handler id
	resultOpt := GetResultOptions{createResponse["handler_id"].(string), "publish"}

	// Get the actual result
	respJson, err := ac.Client.getResult(resultOpt)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}

type UpdateOptions struct {
	Filepath string
	Data     string
	Keywords []string
}

func (ac *assetsClient) Update(ual string, options UpdateOptions) ([]byte, error) {
	if options.Filepath == "" && options.Data == "" {
		return nil, errors.New("Please provide update options in order to update")
	}

	opt := PublishRequestOptions{"update", options.Data, options.Filepath, options.Keywords, ual}

	// Make the request
	resp, err := ac.Client.publishRequest(opt)
	if err != nil {
		return nil, err
	}

	createResponse := make(map[string]interface{})

	// Transform response to json struct
	if err := json.Unmarshal(resp, &createResponse); err != nil {
		return nil, errors.New("Could not unmarshal query request response")
	}

	// Get the handler id
	resultOpt := GetResultOptions{createResponse["handler_id"].(string), "publish"}

	// Get the actual result
	respJson, err := ac.Client.getResult(resultOpt)
	if err != nil {
		return nil, err
	}

	return respJson, nil
}

type GetOptions struct {
	UAL        string
	CommitHash string
}

func (ac *assetsClient) Get(ual string, options GetOptions) ([]byte, error) {
	var resp []byte
	var err error

	if options.CommitHash != "" {
		opt := ResolveRequestOptions{[]string{options.CommitHash}}
		resp, err = ac.Client.Resolve(opt)
		if err != nil {
			return nil, err
		}
	} else {
		opt := ResolveRequestOptions{[]string{options.UAL}}
		resp, err = ac.Client.Resolve(opt)
		if err != nil {
			return nil, err
		}
	}

	return resp, nil
}

// The following are functions that were not used (or left in a TODO state)
// in the original implementation but are kept for future reference.

func (ac *assetsClient) GetStateCommitHashes(ual string) ([]byte, error) {
	opt := ResolveRequestOptions{[]string{ual}}
	resp, err := ac.Client.Resolve(opt)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func (ac *assetsClient) transfer() {}
func (ac *assetsClient) approve()  {}
