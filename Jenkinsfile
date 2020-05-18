pipeline {
    agent any 
    stages {
        stage('Build') {
            steps {
                echo 'Build Starts!'
                bat "go" build \"${workspace}/."
                echo 'Build Ends'
            }
        }
    }
}
