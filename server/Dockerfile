
FROM node:alpine
WORKDIR /usr/src/app
COPY package.json package.json
RUN yarn install
COPY protos protos
COPY src src
CMD [ "node", "./src" ]
