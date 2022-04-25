#
# STAGE 1: SERVER
#
FROM golang:1.17-bullseye as server
WORKDIR /src

#
# install dependencies
RUN set -eux; apt-get update; \
    apt-get install -y --no-install-recommends git cmake make libx11-dev libxrandr-dev libxtst-dev \
    libgstreamer1.0-dev libgstreamer-plugins-base1.0-dev gstreamer1.0-plugins-base gstreamer1.0-plugins-good gstreamer1.0-plugins-bad gstreamer1.0-plugins-ugly; \
    #
    # install libclipboard
    set -eux; \
    cd /tmp; \
    git clone https://github.com/jtanx/libclipboard; \
    cd libclipboard; \
    cmake .; \
    make -j4; \
    make install; \
    rm -rf /tmp/libclipboard; \
    #
    # clean up
    apt-get clean -y; \
    rm -rf /var/lib/apt/lists/* /var/cache/apt/*

COPY server/ .
RUN go get -v -t -d . && go build -o bin/app cmd/app/main.go

FROM node:14-bullseye-slim as client
WORKDIR /src

COPY client/package*.json ./
RUN npm install

COPY client/ .
RUN npm run build

FROM debian:bullseye-slim

ENV DEBIAN_FRONTEND=noninteractive

ARG USERNAME=user
ARG USER_UID=1000
ARG USER_GID=$USER_UID

#
# install dependencies
RUN set -eux; apt-get update; \
    apt-get install -y --no-install-recommends wget ca-certificates supervisor; \
    apt-get install -y --no-install-recommends pulseaudio dbus-x11 xserver-xorg-video-dummy; \
    apt-get install -y --no-install-recommends libcairo2 libxcb1 libxrandr2 libxv1 libopus0 libvpx6; \
    #
    # gst
    apt-get install -y --no-install-recommends libgstreamer1.0-dev libgstreamer-plugins-base1.0-dev \
                    gstreamer1.0-plugins-base gstreamer1.0-plugins-good gstreamer1.0-plugins-bad gstreamer1.0-plugins-ugly gstreamer1.0-pulseaudio; \
    #
    # fonts
    apt-get install -y --no-install-recommends fonts-takao-mincho; \
    #
    # create a non-root user
    groupadd --gid $USER_GID $USERNAME; \
    useradd --uid $USER_UID --gid $USERNAME --shell /bin/bash --create-home $USERNAME; \
    adduser $USERNAME audio; \
    adduser $USERNAME video; \
    adduser $USERNAME pulse; \
    #
    mkdir -p /home/$USERNAME/.config/pulse/; \
    echo "default-server=unix:/tmp/pulseaudio.socket" > /home/$USERNAME/.config/pulse/client.conf; \
    #
    mkdir /tmp/.X11-unix; \
    chmod 1777 /tmp/.X11-unix; \
    chown $USERNAME /tmp/.X11-unix/; \
    #
    mkdir -p /etc/app /var/www /var/log/app; \
    chmod 1777 /var/log/app; \
    chown $USERNAME /var/log/app/; \
    chown -R $USERNAME:$USERNAME /home/$USERNAME; \
    #
    apt-get clean -y; \
    rm -rf /var/lib/apt/lists/* /var/cache/apt/*

#
# copy config files
COPY .docker/base/dbus /usr/bin/dbus
COPY .docker/base/default.pa /etc/pulse/default.pa
COPY .docker/base/supervisord.conf /etc/app/supervisord.conf
COPY .docker/base/xorg.conf /etc/app/xorg.conf

#
ENV USER=$USERNAME
ENV DISPLAY=:99.0
ENV APP_PASSWORD=password
ENV APP_PASSWORD_ADMIN=admin
ENV APP_BIND=:8090

#
COPY --from=server /src/bin/app /usr/bin/app
COPY --from=client /src/dist/ /var/www

HEALTHCHECK --interval=10s --timeout=5s --retries=8 \
    CMD wget -O - http://localhost:${APP_BIND#*:}/health || exit 1

#
CMD ["/usr/bin/supervisord", "-c", "/etc/app/supervisord.conf"]
