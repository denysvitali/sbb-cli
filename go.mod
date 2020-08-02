module github.com/denysvitali/sbb-cli

go 1.14

replace github.com/denysvitali/go-sbb-api v1.0.0 => ../go-sbb-api

require (
	github.com/alexflint/go-arg v1.3.0
	github.com/denysvitali/go-sbb-api v1.0.0
	github.com/docker/distribution v2.7.1+incompatible // indirect
	github.com/sirupsen/logrus v1.6.0
)
