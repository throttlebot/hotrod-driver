node {
    def scmVars = checkout scm
    def image
    def imageName = "willwangkelda/hotrod-driver:release-${scmVars.GIT_COMMIT}"
    stage('Build') {
        image = docker.build(imageName)
    }
    stage('Test') {
//        image.withRun('-p 8082:8082') { c ->
//          sh 'python route_unit_test.py localhost'
//        }
    }
    stage('Push') {
        image.push()
    }
}
