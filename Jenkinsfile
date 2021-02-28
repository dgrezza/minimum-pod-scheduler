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
           export GO111MODULE=on
           ls -la
           go mod download
           go mod vendor
           '''
      }
    }

    stage('Unit Test') {
      steps {
        sh '''
        ls -la
        '''
      }
    }

  }
  environment {
    APP_NAME = 'minimum-pod-scheduler'
    GO_PATH = '/go/src/${APP_NAME}'
  }
}
