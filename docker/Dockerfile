# Copyright 2020 The golang.design Initiative authors.
# All rights reserved. Use of this source code is governed
# by a GNU GPL-3.0 license that can be found in the LICENSE file.

FROM chromedp/headless-shell
RUN apt update && apt install dumb-init
# RUN apt-get update -y \
#     && apt-get install -y fonts-noto \
#     && apt-get install -y fonts-noto-cjk \
#     && apt-get clean && rm -rf /var/lib/apt/lists/* /tmp/* /var/tmp/
ENTRYPOINT ["dumb-init", "--"]
WORKDIR /app
COPY . .
EXPOSE 80
CMD ["/app/code2img"]
