FROM node:17-alpine

WORKDIR /opt/app

COPY ./next /opt/app
RUN npm install -g npm@latest && npm install create-next-app

CMD ["npm", "update"]