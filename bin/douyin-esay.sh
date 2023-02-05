#!/bin/bash
echo "-------------go build running----------------"

# 项目名称
PROJECT_NAME="douyin-esay"
# 部署目录
DEPLOY_DIR="../../"
# 获取项目不是路径的绝对路径

# shellcheck disable=SC2164
cd "${DEPLOY_DIR}"
DEPLOY_DIR=$(pwd)

# 代码目录
CODE_DIR="${DEPLOY_DIR}/${PROJECT_NAME}"
# 配置文件目录
APP_CONFIG_FILE="${CODE_DIR}/config/application.yaml"
# 部署配置文件目录
DEPLOY_APP_CONFIG_DIR="${DEPLOY_DIR}/config"
# 编译后的执行文件名称
BUILD_APP_NAME="douyin-esay"
# 日志目录
LOG_DIR="${DEPLOY_DIR}/log"
# 打包后的执行文件
APP_BUILD_FILE="${CODE_DIR}/${BUILD_APP_NAME}"

echo " - PROJECT_NAME is ${PROJECT_NAME} - "
echo " - DEPLOY_DIR is ${DEPLOY_DIR} - "
echo " - CODE_DIR is ${CODE_DIR} - "
echo " - APP_CONFIG_FILE is ${APP_CONFIG_FILE} - "
echo " - BUILD_APP_NAME is ${BUILD_APP_NAME} - "
echo " - LOG_DIR is ${LOG_DIR} - "
echo " - APP_BUILD_FILE is ${APP_BUILD_FILE} - "

# 检查代码是否有更新
echo " - git update - "
# shellcheck disable=SC2164
cd "${CODE_DIR}"
git checkout
git pull

# 编译打包
echo " - go build - "
go build -o $BUILD_APP_NAME

# 检查编译结果
echo " - check build result - "
if [ ! -e "${APP_BUILD_FILE}" ]; then
  echo  " - the build result not find ! user : douyin-easy.sh - "
  exit 1
fi

# 检查日志目录
if [ ! -d "${LOG_DIR}" ]; then
  ehco " - mkdir ${LOG_DIR} - "
  mkdir "$LOG_DIR"
fi

# 拷贝配置文件
if [ ! -d "$DEPLOY_APP_CONFIG_DIR" ]; then
  echo " - mkdir ${DEPLOY_APP_CONFIG_DIR}"
  mkdir "$DEPLOY_APP_CONFIG_DIR"
fi
cp -f "${APP_CONFIG_FILE}" "${DEPLOY_APP_CONFIG_DIR}/config"

# 检查拷贝是否成功
# shellcheck disable=SC2164
cd "$DEPLOY_DIR"
if [ -e "./${BUILD_APP_NAME}" ]; then
  echo " - ${BUILD_APP_NAME} is copy success! - ";
else
  echo " - ${BUILD_APP_NAME} is not exist ! - ";
  exit 1
fi

# 删除pid文件
if [ -e "./pid.txt" ]; then
  PID=$(cat "./pid.txt")
  ehco " - kill pid ${PID} -";
  kill -9 $PID
  sleep 2
  rm -f "./pid.txt"
  rm -f "./stop.sh"
fi

# 日志
LOG_NAME="${BUILD_APP_NAME}.log"
# 日志输出脚本
if [ ! -e "./log.sh" ]; then
  echo "tail -f -n 100 ${LOG_NAME}" > "./log.sh"
  chmod +x "./log.sh"
fi

# 检查进程及其创建pid文件
# shellcheck disable=SC2009
N_PID=$(ps aux | grep "${BUILD_APP_NAME}" | grep -v grep|awk 'print $2');
if [ -n "$N_PID" ]; then
  #pid 文件
  echo "$N_PID" > "./pid.txt";

  # 进程终止脚本
  echo "echo - kill pid ${N_PID}" > "./stop.sh"
  # shellcheck disable=SC2129
  echo "kill -9 ${N_PID}" >> "./stop.sh"
  echo "sleep 2" >> "./stop.sh"
  ehco "rm -rf ./pid.txt" >> "./stop.sh"
  echo "rm -rf ./stop.sh" >> "./stop.sh"

  chmod +x "./stop.sh"
else
  echo " - save pid.txt fail ,pid is null ! -";
  exit 1
fi

echo " run ${BUILD_APP_NAME} , start success ! -";
ehco "------------------ go build end -----------------------------"
tail -f -n 100 ${LOG_NAME}









