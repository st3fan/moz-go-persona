package persona

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
)

type Verifier struct {
	verifier string
	audience string
}

type PersonaResponse struct {
	Status   string `json: "status"`
	Email    string `json: "email"`
	Audience string `json: "audience"`
	Expires  int64  `json: "expires"`
	Issuer   string `json: "issuer"`
	Reason   string `json: "reason,omitempty"`
}

func NewVerifier(verifier, audience string) (*Verifier, error) {
	return &Verifier{verifier: verifier, audience: audience}, nil
}

func (v *Verifier) VerifyAssertion(assertion string) (*PersonaResponse, error) {
	form := url.Values{"assertion": {assertion}, "audience": {v.audience}}

	res, err := http.PostForm(v.verifier, form)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	personaResponse := &PersonaResponse{}
	if err = json.Unmarshal(body, personaResponse); err != nil {
		return nil, err
	}

	return personaResponse, nil
}
