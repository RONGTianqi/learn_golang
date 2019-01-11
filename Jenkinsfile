pipeline {
    agent any
    stages{

    stage("build") {
            steps {
            echo "构建中..."
            sh 'docker build -t gindocker .'
            echo "构建完成."
                // archiveArtifacts artifacts: '**/target/*.jar', fingerprint: true // 收集构建产物
            }
        }

        stage("测试") {
            steps {
                echo "单元测试中..."

                echo "单元测试完成."
                // junit 'target/surefire-reports/*.xml' // 收集单元测试报告的调用过程
            }
        }

        stage("部署") {
            steps {
                echo "部署中..."
                sh 'docker run -d -p 8000:8000 gindocker'
                echo "部署完成"
            }
        }
    }

}
