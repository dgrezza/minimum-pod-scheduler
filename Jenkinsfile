pipeline {
  agent {
    docker {
      image 'dgrlabs/base-runner:latest'
    }

  }
  stages {
    stage('Update Dependencies') {
      steps {
        sh '''export GO111MODULE=on
export ROOT_LOCATION=$PWD
  
mkdir ${GO_PATH}/
cp -fr ${CI_PROJECT_DIR}/* ${GO_PATH}/
cd $GO_PATH
go mod download
go mod vendor
rm -rf $ROOT_LOCATION/vendor/
mkdir $ROOT_LOCATION/vendor/
cp -fR vendor/. $ROOT_LOCATION/vendor/.
cd $ROOT_LOCATION
'''
      }
    }

    stage('Unit Test') {
      steps {
        sh '''ls -la
'''
      }
    }

  }
  environment {
    APP_NAME = 'minimum-pod-scheduler'
    GO_PATH = '/go/src/${APP_NAME}'
  }
}