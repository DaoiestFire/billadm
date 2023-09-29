#!/bin/bash
# a shell script to install/uninstall billadm
# please execute *uninstall* before *install*

readonly CURRENT_PATH=$(
  cd $(dirname $0)
  pwd
)

readonly BILLADM_PATH=${CURRENT_PATH}/../cmd
readonly INSTALL_PATH=/opt/billadm
readonly CONFIG_PATH=${INSTALL_PATH}/config
readonly BIN_PATH=${INSTALL_PATH}/bin
readonly 
readonly EXECUTABLE_NAME=billadm

function build() {
  cd "${BILLADM_PATH}" || return 1
  rm ${EXECUTABLE_NAME}
  rm -rf go.sum
  go mod tidy
  go build -ldflags '-s -w' -o ${EXECUTABLE_NAME}
  if [ ! -f ${EXECUTABLE_NAME} ]; then
    echo "build ${EXECUTABLE_NAME} failed"
    return 1
  fi
  echo "build ${EXECUTABLE_NAME} success"
  return 0
}

function install() {
  if [ -z "${GOROOT}" ]; then
    echo "please install go environment first"
    exit 1
  fi

  build

  if [ $? == 1 ]; then
    exit 1
  fi
  if [ ! -d ${INSTALL_PATH} ]; then
    mkdir ${INSTALL_PATH}
  fi
  cp ${EXECUTABLE_NAME} ${INSTALL_PATH}

}

function uninstall() {
  if [ -d ${CONFIG_PATH} ]; then
    rm -rf ${CONFIG_PATH}
  fi

  if [ -f ${INSTALL_PATH}/${EXECUTABLE_NAME} ]; then
    # shellcheck disable=SC2115
    rm -rf "${INSTALL_PATH}/${EXECUTABLE_NAME}"
  fi

  echo "uninstall billadm success"
}

function main() {
  operation=$1

  case ${operation} in
  install)
    echo "start to install billadm ..."
    echo "uninstall billadm first, user's data won't be removed"
    uninstall
    install
    ;;
  uninstall)
    echo "start to uninstall billadm ..."
    uninstall
    ;;
  *)
    echo "${operation} is not supported"
    echo "supported operation: install/uninstall"
    exit 1
    ;;
  esac
}

main $1
