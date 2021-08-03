
FROM node:15-alpine

RUN npm install -g http-server

WORKDIR /app

COPY package*.json ./

RUN npm install

COPY . .

ARG API_BASE_URL
ENV API_BASE_URL=${API_BASE_URL}

RUN  npm run build

EXPOSE 3000

CMD [ "http-server", "-p3000" , "dist" ]
