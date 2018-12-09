package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// Test is the Test resource media type.
var Test = MediaType("application/vnd.test+json", func() {
	Description("A test unit")
	Attributes(func() {
		Attribute("id", String, "Unique id/number")
		Attribute("sor", String, "System of record. Usually who created the Test.")
		Attribute("createdAt", DateTime, "Test creation date/time")

		Required("id", "sor")
		Required("createdAt")

		View("default", func() {
			Attribute("id")
			Attribute("sor")
			Attribute("createdAt")
		})
	})
})
