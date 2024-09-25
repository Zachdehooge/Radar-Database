pipeline {
    agent any
    tools { go '1.23.1' }

// Need to have Docker and GO plugins installed on Jenkins
    stages {

        stage('Unit Tests') {
            steps {
                script {
                    // Setting up Go environment
                    //def goVersion = ">=1.23.0"
                    sh "go version"
                    // Running unit tests
                    sh '''
                    git clone https://github.com/Zachdehooge/Radar-Database.git
                    cd Radar-Database
                    pwd
                    echo "Running unit tests..."
                    go test main_test.go
                    echo "Unit Tests Complete"
                    '''
                }
            }
        }

        // TODO: DOCKER DAEMON NEEDS FIXING
         stage('Docker Deploy') {
            steps {
                script {
                    docker.image('maven:3.3.3-jdk-8').inside {
                    git '…your-sources…'
                    sh '''
                    echo "Hello From Pipeline Docker
                    '''
                    }
                }
            }
        }
    } 
    post {
        always {
            cleanWs()
        }
    }
}
