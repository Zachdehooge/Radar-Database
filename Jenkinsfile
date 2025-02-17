pipeline {
    agent any
    tools { go '1.23.1'
            dockerTool "Docker" }

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
                    withCredentials([usernamePassword(credentialsId: 'DOCKERID', passwordVariable: 'DOCKERID_PSW', usernameVariable: 'DOCKERID_USR')]) 
                    {
                    // Setting up Docker
                    sh 'docker version'      
                    
                    sh "docker login -u ${env.DOCKERID_USR} -p ${env.DOCKERID_PSW}"
                    
                    // Build Docker image
                    sh "docker build -t ${env.DOCKERID_USR}/radar-database:latest ."

                    // Push Docker image
                    sh "docker push ${env.DOCKERID_USR}/radar-database:${tag}" 
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
