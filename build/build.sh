#!/bin/bash
# a shell script to install/uninstall billadm
# please execute *uninstall* before *install*

readonly CURRENT_PATH=$(
  cd $(dirname $0)
  pwd
)

readonly BILLADM_PATH=${CURRENT_PATH}/../cmd
readonly INSTALL_PATH=/opt/billadm
readonly BIN_PATH=${INSTALL_PATH}/bin
readonly DATA_PATH=${INSTALL_PATH}/data
readonly LOG_PATH=${INSTALL_PATH}/log
readonly EXECUTABLE_NAME=billadm

function log_info() {
  echo "[$(date '+%Y-%m-%d %H:%M:%S')]" "[INFO]" "$@"
}

function log_error() {
  echo "[$(date '+%Y-%m-%d %H:%M:%S')]" "[ERROR]" "$@"
}

function build() {
  cd ${BILLADM_PATH} || return 1
  rm ${EXECUTABLE_NAME}
  rm -rf go.sum
  go mod tidy
  go build -ldflags '-s -w' -o ${EXECUTABLE_NAME}
  if [ ! -f ${EXECUTABLE_NAME} ]; then
    log_error "build ${EXECUTABLE_NAME} failed"
    return 1
  fi
  log_info "build ${EXECUTABLE_NAME} success"
  return 0
}

function install() {
  if [ -d ${INSTALL_PATH} ]; then
    log_error "please uninstall billadm first"
  fi

  if [ -z "${GOROOT}" ]; then
    log_error "please install go environment first"
    exit 1
  fi

  build || exit 1

  mkdir -p ${BIN_PATH} || exit 1
  mkdir -p ${DATA_PATH} || exit 1
  mkdir -p ${LOG_PATH} || exit 1

  cp ${EXECUTABLE_NAME} ${BIN_PATH}
  find ${INSTALL_PATH} -type d | xargs -i chmod 750 {}
  chmod 500 ${BIN_PATH}/${EXECUTABLE_NAME}
  log_info "install billadm success"
}

function uninstall() {
  rm -rf ${BIN_PATH}
  rm -rf ${LOG_PATH}
  log_info "uninstall billadm success"
}

function main() {
  operation=$1

  case ${operation} in
  install)
    log_info "start to install billadm ..."
    log_info "uninstall billadm first, user's data won't be removed"
    uninstall
    install
    ;;
  uninstall)
    log_info "start to uninstall billadm ..."
    uninstall
    ;;
  *)
    log_info "${operation} is not supported"
    log_info "supported operation: install/uninstall"
    exit 1
    ;;
  esac
}

main $1

exit 0
