FROM alpine
ADD post-srv /post-srv
ENTRYPOINT [ "/post-srv" ]
