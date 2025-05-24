FROM node:22.14.0-alpine

USER node
WORKDIR /app

COPY --chown=node package.json yarn.lock ./
RUN yarn

COPY --chown=node . .

# CMD yarn build
