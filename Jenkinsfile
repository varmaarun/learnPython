pipeline {
    agent any

    stages {
        stage('Build') {
            steps {
                bat '''
            	cd C:\\Users\\arunv\\Desktop\\MyBriefcase\\LEARNING\\LearnPython\\
            	python build.py 
            	'''
            }
        }
        stage('Test') {
            steps {
                echo 'Testing..'
		bat '''
            	cd C:\\Users\\arunv\\Desktop\\MyBriefcase\\LEARNING\\LearnPython\\
            	python test.py 
            	'''
            }
        }
        stage('Deploy') {
            steps {
                echo 'Deploying....'
            }
        }
    }
}