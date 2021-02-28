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
           import_shared_lib
           update_depedencies
        '''
      }
    }

    stage('Unit Test') {
      steps {
        sh '''
        import_shared_lib
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

def import_shared_lib() {
  sh '''
    set +x
    eval "$(curl -Ls -H "${PRIVATE_TOKEN}" ${PIPELINE_URL}jenkins.sh/raw?ref=master)"
    set +x
  '''
}
