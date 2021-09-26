FROM node:16

WORKDIR /workspace

RUN apt update && apt upgrade -y

RUN npm install -g @vue/cli

COPY package*.json ./

CMD npm install \
    && vue ui --host 0.0.0.0