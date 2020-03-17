package firebaseclient

import (
	"context"
	"encoding/json"

	"print-apple/internal/config"
	"print-apple/pkg/errors"

	"cloud.google.com/go/firestore"
	"google.golang.org/api/option"
)

var (
	sharedClient = &firestore.Client{}
	credentials  = map[string]string{
		"type":                        "service_account",
		"project_id":                  "neogenesis-2a947",
		"private_key_id":              "3b7bfff9bdac3643be886bbde4a8eab65282724f",
		"private_key":                 "-----BEGIN PRIVATE KEY-----\nMIIEuwIBADANBgkqhkiG9w0BAQEFAASCBKUwggShAgEAAoIBAQCZy4V2kYnezQ4h\nvk+bvR36z0fBWMWqlULMGJF2dhFH7JHKufdJtIjB7n+TQbxJa/nUIIw/7WMVbQ0x\nuKcthxSs+GTYhzoWEJQCMdbD/OIxa0TYOTXfnD/pR2kdyrt/LW39BHc14QrMzrLZ\nRiRDp2V+/rPtgKYLKK1abvYINWJWRZIsGBCExIBMxQ/U08UVr2teqx5O2EQEXzNz\nYWj3ulx8O6xEls2uXzhllAQb9e+mc/grIVt3xDsyFTmPnHSLazcnsphi6DFUxlJI\nRNZYR7ZFf1TnEtq2u8rN/2YWBDFxWtQvO5TRFh/mervc9EkE7UjCI3ihxNIVSlj1\nAveEDhWJAgMBAAECgf89bUyhTwACPWSnlZu2WJLg2mmOkCZjO7qySh3Pctj0IUkZ\n760Vg0YcL2ZFKrxJlxbx4w4hem9Y33WlplsflBRuAchFs70/kt+LAN0G6rptgqq8\nsh0ScxYT1rHldFPvN/X2WRJEMt5dbzGAFsPlQbcIQjMbaYbobhbPSAJ8xNXCHujR\nxzL962Ho1STHPwNpndzy/XQw3s7rNub5CbuyCr8+NdQQzV0ypXDOtmwWALwUxyIu\nC1n32RiXFGT4Yli6+zYLrhF7FkOFaMzvKjmngMzYOYNFWmCgUcPVmHk5lMotNL/P\ny+6CpyP+pNdYRP3dv/2Gcx3WCKvyqZ9u2IwcIBUCgYEA0E/jRdfGQMPfW0gvEb+L\nepi9XsV338bacEsYYOCC6thBQd4a10aBRMjH0eoVhUh5HdlB7/oq7Q4scAcsFqUn\nef/16GZOohEyNJNHqLve63zh7n8qR2GsRErqMa69HrNSEe8/XzpZl2rWgXC9Mkg7\n8vBRPd576zliPEIJRilQPCUCgYEAvQCqatxECiJVnHylOOKlcOvZ9ioNJYzz/2cG\ny2sGxCyI8ijRQ2S2LA+VKo/zpF3yzehF/eyDPMT0t+HfnNSLPWaZ+hF+CGvUZ/Tr\nQNdMvWUIKbP6+0lNIztzkr2EYFYkfBpQuYO/SOVVxgfKlrXUHq0K1UZZyDwWQMNT\n38h4hJUCgYEAw/dzkh/URNc/hysYBLVSbKnF9KMGC4GRu3QZ4gEzh+SbN3DPhVex\nglj0ChkR18n/DsJ00mJhAZOE4HsO0drakV3nI5MjRDmzJlyrXCQpKRXZobkFuBM9\nsR1cxhJhncEKYw7UaiyFXfnHBAxgIC5uHzRO6Uok/3uDW7av7M4uyfUCgYATKSId\nuz7amCh9uNU3MyL6k66BGjpC+Es0NUmnDa6d7LXlduXgIzGkvd+tdPKKU0vuPAH8\ngCG942m7ypZU2+dRzjkF9QgF6oiaEWZYKHuLJ9bwA2MKXqAHVludIMFu0szYGALf\nC9A0n6tWbCvJo51hjsFuZbdsaUsIPcUfBr/REQKBgDfx2lrn7YtzqFwYeE9Xz7Tq\n7pukWmVniF9CsUxCba31ML6o08coCnc9no+rc2MtnRF8B2m1KOUnfIxhUm/YATxC\n8ipjX4qU6qPCMWRoh4HBHjEQgO/NG1IInyLu6FBssVHd85+HaaAoDXe54PDkRYzV\nq7qaAgbEf8QY88Liv+LY\n-----END PRIVATE KEY-----\n",
		"client_email":                "firebase-adminsdk-7zlk2@neogenesis-2a947.iam.gserviceaccount.com",
		"client_id":                   "110912029148578276583",
		"auth_uri":                    "https://accounts.google.com/o/oauth2/auth",
		"token_uri":                   "https://oauth2.googleapis.com/token",
		"auth_provider_x509_cert_url": "https://www.googleapis.com/oauth2/v1/certs",
		"client_x509_cert_url":        "https://www.googleapis.com/robot/v1/metadata/x509/firebase-adminsdk-7zlk2%40neogenesis-2a947.iam.gserviceaccount.com",
	}
)

// Client ...
type Client struct {
	Client *firestore.Client
}

// NewClient ...
func NewClient(cfg *config.Config) (*Client, error) {
	var c Client
	cb, err := json.Marshal(credentials)
	if err != nil {
		return &c, errors.Wrap(err, "[FIREBASE] Failed to marshal credentials!")
	}
	option := option.WithCredentialsJSON(cb)
	c.Client, err = firestore.NewClient(context.Background(), cfg.Firebase.ProjectID, option)
	if err != nil {
		return &c, errors.Wrap(err, "[FIREBASE] Failed to initiate firebase client!")
	}
	return &c, err
}
