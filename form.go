package podio

import (
	"encoding/json"
	"fmt"
)

type Form struct {
	Id      int   `json:"form_id"`
	AppID   int64 `json:"app_id"`
	SpaceID int64 `json:"space_id"`

	Fields   []FormField `json:"fields"`
	FieldIDs []int64     `json:"field_ids"`

	Settings    FormSettings `json:"settings"`
	Attachments bool         `json:"attachments"`

	Status string `json:"status"` // active / inactive
	// domains // probably X-orgin stuff
}

type FormSettings struct {
	Captcha     bool             `json:"captcha"`
	SuccessPage *string          `json:"success_page"` // used for redirect or nil
	Theme       string           `json:"theme"`        // e.g. "superhero"
	CSS         string           `json:"css"`
	Text        FormTextSettings `json:"text"`
}

type FormTextSettings struct {
	Heading     string `json:"heading"`
	Description string `json:"description"`
	Submit      string `json:"submit"`
	Success     string `json:"success"`
}

type FormField struct {
	FieldID  int64           `json:"field_id"`
	Settings json.RawMessage `json:"settings"` // to discover what values this can hold
}

// https://developers.podio.com/doc/forms/get-forms-53771
func (client *Client) GetForms(appId int64) (forms []*Form, err error) {
	path := fmt.Sprintf("/form/app/%d", appId)
	err = client.Request("GET", path, nil, nil, &forms)
	return
}
