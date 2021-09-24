FROM node:16

WORKDIR /workspace

RUN apt update && apt upgrade -y

RUN apt install -y libgtk2.0-0 libgtk-3-0 libgbm-dev libnotify-dev libgconf-2-4 libnss3 libxss1 libasound2 libxtst6 xauth xvfb

RUN npm install -g @vue/cli

COPY package*.json ./

CMD npm install \
    && vue ui --host 0.0.0.0