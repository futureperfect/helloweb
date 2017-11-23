FROM       scratch
MAINTAINER Erik Hollembeak <erik.hollembeak@gmail.com>
ADD        helloweb helloweb
ENV        PORT 8080
EXPOSE     8080
ENTRYPOINT ["/helloweb"]
