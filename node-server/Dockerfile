FROM node:alpine

WORKDIR /usr/src/app

COPY package*.json ./

COPY app.js ./

RUN npm install

EXPOSE 8080

CMD [ "node", "app.js" ]