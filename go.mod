module escola

go 1.25.8

replace calc => ./pkg/media_school

replace jsonutil => ./pkg/jsonutil

require (
	calc v0.0.0-00010101000000-000000000000
	jsonutil v0.0.0-00010101000000-000000000000
)
