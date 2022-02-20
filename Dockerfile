FROM golang:alpine

# Define workdir
WORKDIR /app/src
# Install required applications
RUN apk update

RUN go get github.com/cosmtrek/air
COPY .air.toml /root/
# Copy the Source Code
COPY ./src ./
# Install the dependencies
RUN go get
# Run the program
CMD [ "air", "-c", "/root/.air.toml" ]