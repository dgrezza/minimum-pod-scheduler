pipeline {
  agent {
    docker {
      image 'golang:1-alpine3.13'
    }

  }
  stages {
    stage('Update Dependencies') {
      steps {
        sh '''
           APP_PATH=${GO_PATH}/${APP_NAME}
           export GO111MODULE=on
           ls -la
           mkdir ${APP_PATH}
           cp -rf * ${APP_PATH}
           cd ${APP_PATH}
           if [ ! -f "go.mod" ]; then
              go mod init
           fi
           go mod download
           go mod vendor'''
      }
    }

    stage('Unit Test') {
      steps {
        sh '''
        echo "${APP_PATH}"
        ls -la
        cd ${GO_PATH}/${APP_NAME}
        la -la

'''
      }
    }

  }
  environment {
    APP_NAME = 'minimum-pod-scheduler'
    GO_PATH = '/go/src'
  }
}