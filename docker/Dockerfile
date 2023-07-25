# Download base image ubuntu 22.04
FROM ubuntu:22.04

# LABEL about the custom image
LABEL maintainer="akrck02@gmail.com"
LABEL version="0.1"
LABEL description="This is a custom Docker image for ServerDashboard application"

# Environment variables
ENV GITHUB_REFRESH_INTERVAL 60

ENV GITHUB_BACKUP_USER ""
ENV GITHUB_BACKUP_TOKEN "-"
ENV GITHUB_BACKUP_PATH "/app/backup"

ENV GITHUB_FETCH_USERNAMES ""
ENV GITHUB_FETCH_TOKEN "-"
ENV GITHUB_FETCH_JSON_PATH "./backup/Resources/"

# Disable Prompt During Packages Installation
ARG DEBIAN_FRONTEND=noninteractive

# Update Ubuntu Software repository
RUN apt update -y --allow-unauthenticated

# Install git
RUN apt install -y git

# Install curl
RUN apt install -y curl
    
# Install go (1.20.3) 
RUN curl -OL https://golang.org/dl/go1.20.3.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go1.20.3.linux-amd64.tar.gz

# Add go to the PATH environment variable
ENV PATH=$PATH:/usr/local/go/bin

# Add goPath to the PATH environment variable
RUN mkdir /cache
ENV GOPATH=/cache

# Make app directiory
RUN mkdir /app

# Volume configuration
VOLUME ["/app", "/app"]
WORKDIR /app

# Copy start.sh script and define default command for the container
COPY start.sh /start.sh
CMD ["/bin/bash","/start.sh"]