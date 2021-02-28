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
           eval "$(curl -Ls -H "${PRIVATE_TOKEN}" ${PIPELINE_URL}jenkins.sh/raw?ref=master)"
           update_depedencies
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
    PIPELINE_URL = credentials('PIPELINE_URL')
    PRIVATE_TOKEN = credentials('PRIVATE_TOKEN')
    GO_PATH = '/go/src'
  }
}
