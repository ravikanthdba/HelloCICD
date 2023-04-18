version: '3'

services:
  app:
    build: .
    ports:
      - "8080:8080"
    environment:
      PROMETHEUS_JOB_NAME: myapp

  prometheus:
    image: prom/prometheus:v2.30.3
    volumes:
      - ./prometheus.yml:/etc/prometheus/prometheus.yml
    ports:
      - "9090:9090"
