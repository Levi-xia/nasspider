FROM ubuntu:18.04

# 基础变量
ARG WORK_DIR=/home/work
ARG BIN_DIR=${WORK_DIR}
ARG LOG_DIR=${WORK_DIR}/logs
ARG VAR_DIR=${WORK_DIR}/var
ARG CONF_DIR=${WORK_DIR}/config
ARG TEMPLATES_DIR=${WORK_DIR}/templates
ARG STATIC_DIR=${WORK_DIR}/static

ARG SUPERVISOR_DIR=${WORK_DIR}/supervisor
ARG SUPERVISOR_LOG=${SUPERVISOR_DIR}/logs
ARG SUPERVISOR_RUN_DIR=${SUPERVISOR_DIR}/run
ARG SUPERVISOR_CONF_DIR=${SUPERVISOR_DIR}/conf
ARG SUPERVISOR_CONF_D_DIR=${SUPERVISOR_CONF_DIR}/conf.d

ARG SUPERVISOR_RUN_CHILD_LOG_DIR=${SUPERVISOR_RUN_DIR}/supervisor

# 创建文件夹
RUN mkdir -p ${WORK_DIR} ${BIN_DIR} ${LOG_DIR} ${VAR_DIR} \
    ${CONF_DIR} ${TEMPLATES_DIR} ${STATIC_DIR} ${SUPERVISOR_DIR} ${SUPERVISOR_LOG} ${SUPERVISOR_RUN_DIR} \
    ${SUPERVISOR_CONF_DIR} ${SUPERVISOR_CONF_D_DIR} ${SUPERVISOR_RUN_CHILD_LOG_DIR}

# 修改 apt 源为国内镜像
RUN sed -i -e 's@//ports.ubuntu.com/\? @//ports.ubuntu.com/ubuntu-ports @g' \
    -e 's@//ports.ubuntu.com@//mirrors.ustc.edu.cn@g' \
    /etc/apt/sources.list

# 安装必备依赖
RUN apt-get update && apt-get install -y \
    vim \
    curl \
    supervisor \
    wget \
    gnupg2 \
    chromium-browser \
    tzdata \
    && apt-get clean \
    && rm -rf /var/lib/apt/lists/* \
    && ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && echo "Asia/Shanghai" > /etc/timezone \
    && dpkg-reconfigure -f noninteractive tzdata

# 创建用户组和用户
RUN groupadd -r work && \
    useradd -r -g work -d ${WORK_DIR} -s /bin/bash -m work

# 复制二进制构件
COPY ./bin/nas-spider ${BIN_DIR}

# 复制配置文件
COPY ./config/config.yaml ${CONF_DIR}

# 复制模版文件
COPY ./templates ${TEMPLATES_DIR}

# 复制静态文件
COPY ./static ${STATIC_DIR}

# 创建环境文件
RUN mkdir -p ${BIN_DIR}/.deploy \
    && echo "prod" > ${BIN_DIR}/.deploy/service.cluster.txt

# 配置 Supervisor 来管理 Go 二进制程序
RUN echo "[unix_http_server]\nfile=${SUPERVISOR_RUN_DIR}/supervisor.sock\nchmod=0770\n\n[supervisord]\nnodaemon=true\nlogfile=${SUPERVISOR_LOG}/supervisord.log\npidfile=${SUPERVISOR_RUN_DIR}/supervisord.pid\nchildlogdir=${SUPERVISOR_RUN_CHILD_LOG_DIR}\n\n[rpcinterface:supervisor]\nsupervisor.rpcinterface_factory = supervisor.rpcinterface:make_main_rpcinterface\n\n[supervisorctl]\nserverurl=unix://${SUPERVISOR_RUN_DIR}/supervisor.sock\n\n[include]\nfiles = ${SUPERVISOR_CONF_D_DIR}/*.conf" > ${SUPERVISOR_CONF_DIR}/supervisord.conf && \
    echo "[program:nas-spider]\ncommand=${BIN_DIR}/nas-spider\nuser=work\nautostart=true\nautorestart=true\nstderr_logfile=${SUPERVISOR_LOG}/nas-spider.err.log\nstdout_logfile=${SUPERVISOR_LOG}/nas-spider.out.log" >> ${SUPERVISOR_CONF_D_DIR}/nas-spider.conf

# 设置文件权限
RUN chown -R work:work ${WORK_DIR}

# 切换用户
USER work

# 设置工作目录
WORKDIR ${BIN_DIR}

# 启动服务
CMD ["/usr/bin/supervisord", "-c", "/home/work/supervisor/conf/supervisord.conf"]