mysqlHost=43.198.175.111
mysqlPort=3306
mysqlUser=go
mysqlPassword=tz#97-DKhpM3DS9g
mysqlDb=OGame_dev

gamePkgName=./game

VERSION=$(shell git describe --always --match "v[0-9]*" HEAD)
BUILD_INFO_IMPORT_PATH=ph-gitlab.vipsroom.net/bingo_client/c_engine/pkg
BUILD_INFO=-ldflags "-X $(BUILD_INFO_IMPORT_PATH).buildAppVersion='$(VERSION)'"
.PHONY: init
init:
	go install gorm.io/gen/tools/gentool@latest

genGame:
	gentool -dsn "${mysqlUser}:tz#97-DKhpM3DS9g@(${mysqlHost}:${mysqlPort})/${mysqlDb}?charset=utf8mb4&parseTime=True&loc=Local" -modelPkgName "${gamePkgName}" -tables="$(tables)" -onlyModel

.PHONY: build-user
build-user:
	go build  -o ../cmd/user/user $(BUILD_INFO) ../cmd/user/main.go

#make genIm tables="tb_mailbox_content,tb_mailbox_template,tb_mailbox,tb_mailbox_read"
