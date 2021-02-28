pipeline {
  agent {
    docker {
      image 'dgrlabs/base-runner:latest'
    }

  }
  stages {
    stage('Test') {
      steps {
        import_shared_lib()
        sh '''
           update_depedencies
           test_coverage
        '''
      }
    }

    stage('Build') {
      steps {
        import_shared_lib()
        sh "echo build"
      }
    }
    
    stage('Deploy') {
      steps {
        import_shared_lib()
        sh "echo deploy"
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
