FROM golang:alpine AS builder
LABEL stage=builder
RUN apk add --no-cache gcc libc-dev && mkdir /app
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/kube-watch main.go


FROM alpine AS final
RUN mkdir /app
WORKDIR /app
COPY --from=builder /app .
RUN ls -l /
RUN chmod +x /app/kube-watch
CMD [ "/app/./kube-watch" ]