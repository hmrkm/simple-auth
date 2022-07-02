FROM node:16

WORKDIR /workspace

RUN apt update && apt upgrade -y

COPY package*.json ./

RUN npm install

ENTRYPOINT npm start
