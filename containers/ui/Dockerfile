FROM node:18-alpine as ui_builder

WORKDIR /app

COPY ui/package.json ui/yarn.lock /app/
RUN yarn install --frozen-lockfile

COPY ui/ /app/
RUN yarn build

FROM nginx:alpine

COPY --from=ui_builder /app/dist /usr/share/nginx/html
COPY containers/utils/default.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD [ "nginx", "-g", "daemon off;"]

