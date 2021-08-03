
FROM node:15-alpine

RUN yarn global add http-server

WORKDIR /app

#COPY package*.json ./
ADD . /app/
RUN yarn install

COPY . .

ARG API_BASE_URL
ENV API_BASE_URL=${API_BASE_URL}

RUN  yarn build

EXPOSE 3000

CMD [ "http-server", "-p3000" , "dist" ]
