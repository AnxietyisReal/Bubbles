FROM mongo:7.0.1-jammy

VOLUME [ "../../database" ]
EXPOSE 8800
CMD ["mongod", "--quiet", "--bind_ip_all", "--port", "8800"]