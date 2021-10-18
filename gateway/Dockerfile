FROM node:16-stretch-slim

WORKDIR /project

COPY . ./

RUN npm ci

CMD ["/usr/local/bin/npm", "run", "start"]
