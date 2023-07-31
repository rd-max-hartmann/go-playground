# Use nxgo/cli as the base image to do the build
FROM node:18-alpine as builder

# Create app directory
WORKDIR /workspace

RUN yarn config set cache-folder ~/.yarn

RUN apk add --update --no-cache git tar curl vim zsh the_silver_searcher shadow

RUN sh -c "$(curl -fsSL https://raw.github.com/beeman/server-shell/master/tools/install.sh)"

RUN usermod -s /bin/zsh root

RUN apk add --no-cache git make musl-dev go

# Configure Go
ENV GOROOT /usr/lib/go
ENV GOPATH /go
ENV PATH /go/bin:$PATH

RUN mkdir -p ${GOPATH}/src ${GOPATH}/bin

RUN yarn global add \
    nx@16.5.5

# Copy package.json and the lock file
COPY package.json yarn.lock /workspace/

# Install app dependencies
RUN yarn install --ignore-engines

# Copy source files
COPY . .

# Build apps
RUN yarn build

# This is the stage where the final production image is built
FROM golang:1.20-alpine as final

# Copy over artifacts from builder image
COPY --from=builder /workspace/dist/apps/api /workspace/api

# Set environment variables
ENV PORT=3000
ENV HOST=0.0.0.0

# Expose default port
EXPOSE 3000

# Start server
CMD [ "/workspace/api" ]
