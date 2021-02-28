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
           go mod vendor
           echo "${PIPELINE_URL}"
           echo "${PRIVATE_TOKEN}"


'''
      }
    }

    stage('Unit Test') {
      steps {
        sh '''
        echo "${APP_PATH}"
        ls -la
        cd ${GO_PATH}/${APP_NAME}
        ls -la

'''
      }
    }

  }
  environment {
    APP_NAME = 'minimum-pod-scheduler'
    PIPELINE_URL = 'credentials(\'PIPELINE_URL\')'
    PRIVATE_TOKEN = 'credentials(\'PRIVATE_TOKEN\')'
  }
}