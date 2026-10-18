FROM golang:1.26
# Had to change from 1.25 version of go to support air!

WORKDIR /app

COPY . .

# AIR helps automatically recompile GO when *.go files are modified
RUN go install github.com/air-verse/air@latest
ENV PATH="/go/bin:${PATH}"

EXPOSE 8080

#CMD ["go", "run", "."]
CMD ["air"]
