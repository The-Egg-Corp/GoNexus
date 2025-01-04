module github.com/the-egg-corp/gonexus

replace github.com/the-egg-corp/gonexus => ./

go 1.22.1

require github.com/go-resty/resty/v2 v2.16.2

require (
	github.com/joho/godotenv v1.5.1
	github.com/sanity-io/litter v1.5.5
	golang.org/x/net v0.33.0 // indirect
)
