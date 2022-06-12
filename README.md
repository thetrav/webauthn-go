

# windows docker CLI
docker run -it --rm -w /usr/src/app -v ${pwd}:/usr/src/app -p 8090:8090 golang:1.18 /bin/bash

