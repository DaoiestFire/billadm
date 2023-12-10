#!/bin/bash
# a shell script to install/uninstall billadm
# please execute *uninstall* before *install*

readonly CURRENT_PATH=$(
  cd $(dirname $0)
  pwd
)

readonly BILL_CLIENT_PATH=${CURRENT_PATH}/../cmd/billclient
readonly BILL_SERVER_PATH=${CURRENT_PATH}/../cmd/billserver
readonly INSTALL_PATH=/opt/bill
readonly BIN_PATH=${INSTALL_PATH}/bin
readonly DATA_PATH=${INSTALL_PATH}/data
readonly LOG_PATH=${INSTALL_PATH}/log

function log_info() {
  echo "[$(date '+%Y-%m-%d %H:%M:%S')]" "[INFO]" "$@"
}

function log_error() {
  echo "[$(date '+%Y-%m-%d %H:%M:%S')]" "[ERROR]" "$@"
}

function build() {
  path=$1
  file=$2
  cd $path || return 1
  rm $file
  rm -rf go.sum
  go mod tidy
  go build -ldflags '-s -w' -o $file
  if [ ! -f $file ]; then
    log_error "build $file failed"
    return 1
  fi
  log_info "build $file success"
  cp $file ${BIN_PATH}
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

  mkdir -p ${BIN_PATH} || exit 1
  mkdir -p ${DATA_PATH} || exit 1
  mkdir -p ${LOG_PATH} || exit 1

  build ${BILL_CLIENT_PATH} billctl || exit 1
  build ${BILL_SERVER_PATH} billserver || exit 1

  find ${INSTALL_PATH} -type d | xargs -i chmod 750 {}
  chmod 500 ${BIN_PATH}/*
  log_info "install billadm success"
}

function uninstall() {
  rm -rf ${BIN_PATH}
  rm -rf ${LOG_PATH}
  log_info "uninstall billadm success"
}

function clean() {
  rm -rf ${INSTALL_PATH}
  log_info "clean billadm success"
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
  clean)
    log_info "start to clean billadm ..."
    clean
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
