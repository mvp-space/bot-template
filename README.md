# Bot Template

Common template for creating Bots.

## Help

```bash
user@devs:~$ make 
build                          build the bot-tg binary
build-docker                   build the bot-tg as a docker image
clean                          remove temporary files
fmt                            run "go fmt" on all Go packages
help                           help information about make commands
lint                           run golint on all Go package
run-live                       run the bot-tg with live reload support (requires fswatch)
run-restart                    restart the bot-tg
run                            run the bot-tg
test-cover                     run unit tests and show test coverage information
test                           run unit tests
version                        display the version of the bot-tg
```

## Benchmark

```
go test -run=. -bench=. -benchtime=5s -count 5 -benchmem -cpuprofile=cpu.out -memprofile=mem.out -trace=trace.out ./pkg/cache/
go tool pprof -http :8080 cpu.out
go tool pprof -http :8081 mem.out
go tool trace trace.out
```