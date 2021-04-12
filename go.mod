module client

go 1.16

replace local/pokemon => ./pokemon

require (
	google.golang.org/grpc v1.37.0
	local/pokemon v0.0.0-00010101000000-000000000000
)
