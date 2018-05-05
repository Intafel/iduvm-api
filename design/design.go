package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("iduvm", func() {
	Title("IDUVM API")
	Description("Initial implementation of iduvm API")
	Version("0.1.0")
	Host("localhost:8080")
	Scheme("http")
})

var _ = Resource("guest", func() {
	Action("start", func() {
		Routing(GET("start/:guestID"))
		Description("Start a guest")
		Params(func() {
			Param("guestID", String, "The Guest ID")
		})
		Response(NoContent)
	})
	Action("stop", func() {
		Routing(GET("stop/:guestID"))
		Description("Stop and restore a guest")
		Params(func() {
			Param("guestID", String, "The Guest ID")
		})
		Response(NoContent)
	})
})
