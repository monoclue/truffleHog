# Trufflehog Setup
FROM python:3-alpine
RUN apk add --no-cache git && pip install gitdb2==3.0.0 trufflehog
RUN adduser -S truffleHog
USER truffleHog
WORKDIR /proj
ENTRYPOINT [ "/usr/local/bin/trufflehog","https://github.com/monoclue/truffleHog.git" ]
#ENTRYPOINT [ "tail", "-f", "/dev/null" ]
