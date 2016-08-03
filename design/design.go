package design

import (
	. "github.com/goadesign/goa/design"
	. "github.com/goadesign/goa/design/apidsl"
)

// build with: goagen bootstrap -d github.com/jaredwarren/easypostTester/design
var _ = API("easypost", func() {
	Title("Easy Post tester")
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
