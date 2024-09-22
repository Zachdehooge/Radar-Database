pipeline {
    agent any

    stages {
        
        stage('Install Golang') {
            steps {
                script{
                sh '''
                    # Define Go version and download URL
                    GO_VERSION=1.20.7
                    GO_TAR_FILE="go${GO_VERSION}.linux-amd64.tar.gz"
                    GO_DOWNLOAD_URL="https://go.dev/dl/$GO_TAR_FILE"
                    
                    # Download the Go tarball
                    curl -O $GO_DOWNLOAD_URL
                    
                    # Extract the Go tarball to /usr/local
                    tar -C /usr/local -xzf $GO_TAR_FILE
                    
                    # Set Go binary path
                    echo "export PATH=\$PATH:/usr/local/go/bin" >> ~/.bashrc
                    source ~/.bashrc
                    
                    # Verify installation
                    go version
                '''
                }
            }
        }

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

        stage('Docker Deploy') {
            steps {
                script {
                    withCredentials([usernamePassword(credentialsId: 'DOCKERID', passwordVariable: 'DOCKERID_PSW', usernameVariable: 'DOCKERID_USR')]) 
                    {
                    // Setting up Docker
                    sh 'docker version || sudo apt-get install docker.io -y'

                    // Set up Docker image tag
                    def tag = "latest"
                    if (env.GIT_BRANCH == "main") {
                        if (env.GIT_TAG_NAME) {
                            tag = env.GIT_TAG_NAME
                        }
                    }
                    
                    sh "docker login -u ${env.DOCKERID_USR} -p ${env.DOCKERID_PSW}"
                    
                    // Build Docker image
                    sh "docker build -t ${env.DOCKERID_USR}/radar-database:${tag} ."

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