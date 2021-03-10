# Quote Of The Day

This repository contains logic for a server that returns quotes of the day. It
is meant to run on port 17 and conform to RFC 865.
https://tools.ietf.org/html/rfc865

The repo also has a buggy branch that can be used to mimic an exploitable
vulnerability. Switch to the `vulnerability` branch to see more.

### Deployment

```
docker build -t logstoneducation/qotd:0.0.1 .
docker push logstoneducation/qotd:0.0.1
```
