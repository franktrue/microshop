module github.com/franktrue/microshop/services/demo-cli

go 1.16

replace github.com/franktrue/microshop/services/demo-service => E:/WorkSpace/Project/MicroService/MicroShop/services/demo-service

require (
	github.com/franktrue/microshop/services/demo-service v0.0.0-20210812112115-20014105ea02
	github.com/micro/go-micro/v2 v2.9.1
)
