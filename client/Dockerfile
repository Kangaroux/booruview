FROM node:22.14.0-alpine

WORKDIR /app

COPY package.json yarn.lock ./
RUN yarn

COPY . .

# CMD yarn build
