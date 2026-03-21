module escola

go 1.25.8

replace calc => ./internal/media_school

replace jsonutil => ./internal/jsonutil

require (
	calc v0.0.0-00010101000000-000000000000
	jsonutil v0.0.0-00010101000000-000000000000
)
