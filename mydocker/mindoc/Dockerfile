FROM wolcen/calibre
LABEL author="wolcen@msn.com"

# 安装mindoc
ADD zmindoc /mindoc/
WORKDIR /mindoc
RUN chmod +x /mindoc/start.sh

EXPOSE 8181
ENTRYPOINT ["/mindoc/start.sh"]
