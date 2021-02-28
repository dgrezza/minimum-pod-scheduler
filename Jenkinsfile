pipeline {
  agent {
    docker {
      image 'dgrlabs/base-runner:latest'
    }

  }
  stages {
    stage('Update Dependencies') {
      steps {
        sh '''
           set +x
           eval "$(curl -Ls -H "${PRIVATE_TOKEN}" ${PIPELINE_URL}jenkins.sh/raw?ref=master)" > /dev/null
           update_depedencies
        '''
      }
    }

    stage('Unit Test') {
      steps {
        sh '''
        set +x
        eval "$(curl -Ls -H "${PRIVATE_TOKEN}" ${PIPELINE_URL}jenkins.sh/raw?ref=master)" > /dev/null
        
        cd ${GO_PATH}/${APP_NAME}
        set +x
        test_coverage
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
