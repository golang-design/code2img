# Copyright 2020 The golang.design Initiative authors.
# All rights reserved. Use of this source code is governed
# by a GNU GPL-3.0 license that can be found in the LICENSE file.

all:
	GOOS=linux go build
	docker build -t code2img .
up: down
	docker-compose -f docker-compose.yml up -d
down:
	docker-compose -f docker-compose.yml down