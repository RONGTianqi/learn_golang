// Run on an agent where we want to use Go<font></font>
// node {
//     // Install the desired Go version<font></font>
//     def root = tool name: 'go-1.11', type: 'go'
//     // Export environment variables pointing to the directory where Go was installed<font></font>
//     withEnv(["GOROOT=${root}", "PATH+GO=${root}/bin"]) {
//         sh 'go version'
    
//     }
//     stage ('build'){
//         echo "构建中..."
//         sh 'go build'
//         echo "构建完成."
//     }

//     stage ('test'){
//         echo "单元测试中..."
//         sh 'go run main.go'
//         echo "单元测试完成."
//     }

//     stage("部署") {
//         echo "部署中..."
//         echo "部署完成"
//     }
// }
pipeline {
    agent any
    
    stages  {
        
        stage("checkout") {
            steps {
                sh 'export GOPATH=$WORKSPACE'
                sh 'export PATH=$PATH:$GOPATH/bin'
                )
            }
        }

        stage("build") {
            steps {
                echo "构建中..."
                sh 'go version'
                // sh 'go install'
                // 请在这里放置您项目代码的单元测试调用过程，例如:
                // sh 'mvn package' // mvn 示例
                // sh 'make' // make 示例
                echo "构建完成."
                // archiveArtifacts artifacts: '**/target/*.jar', fingerprint: true // 收集构建产物
            }
        }

        stage("测试") {
            steps {
                echo "单元测试中..."
                // sh 'go run main.go'
                // 请在这里放置您项目代码的单元测试调用过程，例如:
                // sh 'mvn test' // mvn 示例
                // sh 'make test' // make 示例
                echo "单元测试完成."
                // junit 'target/surefire-reports/*.xml' // 收集单元测试报告的调用过程
            }
        }

        stage("部署") {
            steps {
                echo "部署中..."
                // 请在这里放置收集单元测试报告的调用过程，例如:
                // sh 'mvn tomcat7:deploy' // Maven tomcat7 插件示例：
                // sh './deploy.sh' // 自研部署脚本
                echo "部署完成"
            }
        }
    }
}