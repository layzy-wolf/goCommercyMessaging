FROM node:20.11.1

WORKDIR /usr/src/app

COPY ./app .

RUN npm install -g vite

RUN npm install

EXPOSE 5173

CMD ["npm", "run", "dev"]