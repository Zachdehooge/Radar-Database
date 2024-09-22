pipeline {
    agent any

    stages {
        
        stage('Install Golang') {
            steps {
                script {
                    sh '''
                    curl -OL https://go.dev/dl/go1.21.1.linux-amd64.tar.gz 
                    tar -C $HOME -xzf go1.21.1.linux-amd64.tar.gz 
					echo "export PATH=\$PATH:$HOME/go/bin" >> ~/.profile 
                    source ~/.bashrc
                    go -version
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