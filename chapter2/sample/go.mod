module gia/main

go 1.18

replace matchers => ./matchers/

replace search => ./search/

require (
	matchers v0.0.0-00010101000000-000000000000
	search v0.0.0-00010101000000-000000000000
)
