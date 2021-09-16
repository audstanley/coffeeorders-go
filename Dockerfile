FROM frolvlad/alpine-glibc
WORKDIR /root
COPY ./bin/coffeeorders-amd64-linux /usr/local/bin
CMD [ "coffeeorders-amd64-linux"]