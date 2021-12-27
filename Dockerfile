FROM golang:1.17-alpine
RUN apk add --update --no-cache python3 && ln -sf python3 /usr/bin/python
RUN python3 -m ensurepip
RUN pip3 install --no-cache --upgrade pip setuptools
RUN apk add --no-cache git && pip install gitdb2==3.0.0 trufflehog
RUN adduser -S truffleHog
RUN touch /var/log/development.log && chown trufflehog:trufflehog /var/log/development.log

USER truffleHog
WORKDIR /proj

# Install Golang tools
#ENV GOPATH="/go"
#ENV PATH="$GOPATH/bin:$PATH"
#RUN mkdir -p "$GOPATH/src" "$GOPATH/bin"

COPY http/ /http/

CMD go run /http/http-server.go