FROM node:24 AS build

# Create app directory
WORKDIR /usr/src/app
COPY package*.json ./

# NodeJS Install
RUN npm ci

# Bundle app source and build typescript
FROM build AS typescript
COPY . .
RUN npm run build

# Remove development dependencies
FROM typescript AS pre-final
RUN npm prune --omit=dev

# Create a separate image with the build layers removed
FROM node:24 AS final
COPY --from=pre-final /usr/src/app/node_modules/ /usr/src/app/node_modules/
COPY --from=pre-final /usr/src/app/start.js /usr/src/app/

# Configure application startup
WORKDIR /usr/src/app
USER node
EXPOSE 8088
CMD [ "node", "start.js" ]
