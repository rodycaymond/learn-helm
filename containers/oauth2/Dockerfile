FROM docker.io/bitnami/oauth2-proxy:7.6.0

ENV HOST=keycloak.zeus.socom.dev
ENV APP_NAME=generator-app
ENV OAUTH2_PROXY_CLIENT_ID=${APP_NAME}
ENV OAUTH2_PROXY_UPSTREAMS=http://${APP_NAME}.zeus.socom.dev
ENV OAUTH2_PROXY_COOKIE_NAME=__Secure-oauth_proxy_${APP_NAME}
ENV OAUTH2_PROXY_CLIENT_SECRET=52f81DodPlYafGR23xJenW6UXizwOv0H
ENV OAUTH2_PROXY_COOKIE_SECRET="68W9GCfJxthaPQFZ3itOJ48HBNFzDTt0iZtfP1L8YjI="
ENV OAUTH2_PROXY_OIDC_ISSUER_URL=https://${HOST}/realms/zeus-apps


COPY oauth2-proxy.cfg /oauth2-proxy.cfg
COPY rootCA.pem /etc/ssl/certs/rootCA.pem

CMD [ "--config", "/oauth2-proxy.cfg" ]