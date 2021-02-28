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
           pwd
           ls -la
           echo $GO_PATH
           mkdir /go/src/${APP_NAME}
           cp . /go/src/${APP_NAME}
           cd /go/src/${APP_NAME}
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
