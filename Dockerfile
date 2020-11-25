FROM node:14-alpine

# Use tini to handle signal forwarding.
RUN apk add --no-cache tini
ENTRYPOINT ["/sbin/tini", "--"]
CMD ["node", "index.js"]
EXPOSE 5000
USER node

# Create the app and build dir first, otherwise WORKDIR will create them as root.
RUN mkdir /home/node/local-discovery
RUN mkdir /tmp/build

# Build the app, move the build result to the app dir.
WORKDIR /tmp/build
COPY --chown=node *.json ./
COPY --chown=node src src
RUN npm i
RUN npm run build
RUN mv build/* /home/node/local-discovery
RUN rm -rf /tmp/build

# Install the production dependencies.
WORKDIR /home/node/local-discovery
ENV NODE_ENV production
COPY package*.json ./
RUN npm ci
