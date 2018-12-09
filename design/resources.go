package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = Resource("test", func() {

	BasePath("/tests")
	DefaultMedia(Test)

	Action("list", func() {
		Routing(
			GET(""),
		)
		Description("Retrieve all tests.")
		Response(OK, func() {
			Media(CollectionOf(Test, func() {
				View("default")
			}))
		})
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("show", func() {
		Description("Get test by id")
		Routing(GET("/:testId"))

		Params(func() {
			Param("testId", String, "Test ID")
		})

		Response(OK)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

	Action("create", func() {
		Routing(
			POST(""),
		)
		Description("Create new test")
		Payload(TestPayload, func() {
			Required("testId", "sor")
		})
		Response(Created, Test)
		Response(BadRequest, ErrorMedia)
	})

	Action("update", func() {
		Routing(
			PATCH("/:testId"),
		)
		Params(func() {
			Param("testId", String)
		})
		Payload(TestPayload)
		Response(NoContent)
		Response(NotFound)
		Response(BadRequest, ErrorMedia)
	})

})

var _ = Resource("swagger", func() {
	Origin("*", func() {
		Methods("GET", "OPTIONS")
	})
	Files("/swagger.json", "swagger/swagger.json")
	Files("/swagger-ui/*filepath", "swagger-ui/")
})
