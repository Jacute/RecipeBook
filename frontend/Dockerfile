FROM node:22.4.0


WORKDIR /app

COPY dist dist
COPY views views
COPY package* .
COPY static static

RUN npm install --production

CMD ["npm", "run", "start"]