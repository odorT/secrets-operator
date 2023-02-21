## CI agnostic Dockerfile

FROM node:lts AS deps
WORKDIR /home/node
COPY package.json /home/node/
RUN yarn --only=production
COPY . .
RUN yarn run build

FROM registry.access.redhat.com/ubi8/nginx-120
COPY --from=deps /home/node/build /usr/share/nginx/html
COPY ./nginx.AppConf /etc/nginx
USER nginx
EXPOSE 3000
CMD nginx -g 'daemon off;'
