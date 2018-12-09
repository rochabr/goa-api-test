package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// TestPayload defines the data structure used in the create Test request body.
var TestPayload = Type("TestPayload", func() {
	Attribute("testId", func() {
		MinLength(1)
		Example("example")
	})
	Attribute("sor", func() {
		MinLength(1)
		Example("example")
	})
	Attribute("createdAt", DateTime)
})
