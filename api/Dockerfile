FROM node:12

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm config set registry http://registry.npmjs.org/
RUN npm ci --only=production

COPY . .

EXPOSE 3000
CMD [ "node", "index.js" ]