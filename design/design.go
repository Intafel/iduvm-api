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
		Response(NotFound)
		Response(InternalServerError)
	})
	Action("stop", func() {
		Routing(GET("stop/:guestID"))
		Description("Stop and restore a guest")
		Params(func() {
			Param("guestID", String, "The Guest ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(InternalServerError)
	})
	Action("OpenFile", func() {
		Routing(GET("openfile/:guestID/:filePath"))
		Params(func() {
			Param("guestID", String, "The Guest ID")
			Param("filePath", String, "The path to the file to be opened")
		})
		Response(NoContent)
		Response(NotFound)
		Response(InternalServerError)
	})
	Action("OpenURL", func() {
		Routing(GET("openurl/:GuestID/:URL"))
		Params(func() {
			Param("GuestID", String, "The Guest ID")
			Param("URL", String, "The URL to be opened")
		})
		Response(NoContent)
		Response(NotFound)
		Response(InternalServerError)
	})
	Action("OpenBrowser", func() {
		Routing(GET("openbrowser/:GuestID"))
		Params(func() {
			Param("GuestID", String, "The Guest ID")
		})
		Response(NoContent)
		Response(NotFound)
		Response(InternalServerError)
	})
})
