TIMESTAMP             	:= $(shell /bin/date "+%F %T")
NAME					:= regioncnclient
VERSION					:= 1.0.0

fmt:
	go fmt $(CURDIR)/...

protoc:
	protoc -I=$(CURDIR)/protobuf/ --go_out=$(CURDIR)/protobuf/ $(CURDIR)/protobuf/regioncn.proto

github: fmt protoc
	git add .
	git commit -m "$(TIMESTAMP)"
	git push

.PHONY: fmt protoc github
