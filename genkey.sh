openssl genrsa -out key/key.rsa 1024
openssl rsa -in key/key.rsa -pubout > key/key.pub