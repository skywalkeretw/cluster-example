# Stage 1: Build an Angular Docker Image
FROM node:16 as builder

WORKDIR /app

COPY package*.json /app/

RUN npm install

COPY . /app

ARG configuration=production
RUN npm run build -- --outputPath=./dist/out --configuration $configuration

# Stage 2, use the compiled app, ready for production with Nginx

FROM nginx

COPY --from=builder /app/dist/out/ /usr/share/nginx/html
COPY /nginx.conf /etc/nginx/conf.d/default.conf