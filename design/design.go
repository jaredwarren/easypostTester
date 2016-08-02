package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

var _ = API("cellar", func() {
	Title("The virtual wine cellar")
	Description("A simple goa service")
	Version("2.0")
	Scheme("http")
	Host("localhost:8080")
	BasePath("/v2")
	Consumes("application/json")
	Produces("application/json")

	License(func() {
		Name("Apache License Version 2.0")
		URL("http://www.apache.org/licenses/LICENSE-2.0")
	})

})
