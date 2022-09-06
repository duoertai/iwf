/*
 * Workflow APIs
 *
 * This APIs for iwf SDKs to operate workflows
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package serverapi

type EncodedObject struct {

	Encoding string `json:"encoding,omitempty"`

	Data string `json:"data,omitempty"`
}

// AssertEncodedObjectRequired checks if the required fields are not zero-ed
func AssertEncodedObjectRequired(obj EncodedObject) error {
	return nil
}

// AssertRecurseEncodedObjectRequired recursively checks if required fields are not zero-ed in a nested slice.
// Accepts only nested slice of EncodedObject (e.g. [][]EncodedObject), otherwise ErrTypeAssertionError is thrown.
func AssertRecurseEncodedObjectRequired(objSlice interface{}) error {
	return AssertRecurseInterfaceRequired(objSlice, func(obj interface{}) error {
		aEncodedObject, ok := obj.(EncodedObject)
		if !ok {
			return ErrTypeAssertionError
		}
		return AssertEncodedObjectRequired(aEncodedObject)
	})
}
