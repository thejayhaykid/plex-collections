package models

// RequestPayload is an interface that implements Validate which validates a model
// and returns if the payload was decoded and validated ok and returns a message if there was
// a validation error
type RequestPayload interface {
	Validate() (bool, string)
}
