package test

import (
	"fmt"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"io/ioutil"
	"log"
	"os"
	"path"
	"testing"
	"time"
)

// kong的插件上传
func TestBuildKongPlugPush(t *testing.T) {
	buildInfo := &BuildInfo{
		SshHost:               "10.0.40.30",
		SshUser:               "root",
		SshPassword:           "123.com",
		SshPort:               22,
		PluginsPath:           []string{"./service/token/valid_token.go"}, //该项目的插件路径
		RemoteFilePath:        "",                                         //服务器存放该项目的文件夹路径
		DockerAndKongConfPath: "",                                         // 服务器存放kong.conf和Docker的路径
	}
	DockerBuild(*buildInfo)
}

type BuildInfo struct {
	SshHost               string   //ssh ip地址
	SshUser               string   //ssh 账号
	SshPassword           string   //ssh 密码
	SshPort               int      //ssh port
	RemoteFilePath        string   //服务器文件夹路径
	PluginsPath           []string // 目录下插件路径
	DockerRunCommand      string   //DockerRun命令
	DockerAndKongConfPath string   // 服务器存放kong.conf和Docker的路径
}

func DockerBuild(paramInfo BuildInfo) {

	//dir, err := os.Getwd() // 获取当前路径
	//if err != nil {
	//	panic(err)
	//}
	//projectPath := dir[:strings.Index(dir, conf.ServiceName)] + conf.ServiceName //获取项目本地路径
	//通过ftp将可执行文件上传服务器
	client, err := CreateSftp(paramInfo.SshHost, paramInfo.SshUser, paramInfo.SshPassword, paramInfo.SshPort) //创建ftp连接
	if err != nil {
		panic(err)
	}
	defer client.Close()
	//if paramInfo.RemoteFilePath == "" {
	//	paramInfo.RemoteFilePath = "/mnt/go_kong/yk_kong/"
	//}
	//if paramInfo.DockerAndKongConfPath == "" {
	//	paramInfo.DockerAndKongConfPath = "/mnt/Do_Kong/"
	//}
	//err = client.MkdirAll(paramInfo.RemoteFilePath) // 服务器创建文件夹
	//if err != nil {
	//	panic(err)
	//}
	SrcFile(client, "/Users/fanyahui/learn_golang/main.go", "/mnt/go_kong/")
	////todo 上传有问题，上传的文件不能build,待修复
	//err = uploadDirectory(client, projectPath, paramInfo.RemoteFilePath) // 上传项目到服务器

	//var commondDocker strings.Builder
	//把 kong.conf 和Dockerfile文件移到指定目录
	//commondDocker.WriteString("mv " + paramInfo.RemoteFilePath + "kong.conf " + paramInfo.DockerAndKongConfPath + "\n")
	//commondDocker.WriteString("mv " + paramInfo.RemoteFilePath + "Dockerfile " + paramInfo.DockerAndKongConfPath + "\n")

	// 编译插件并且把文件移到指定目录
	//commondDocker.WriteString("cd " + paramInfo.RemoteFilePath + "\n")
	//for _, v := range paramInfo.PluginsPath {
	//	fmt.Println(v)
	//	commondDocker.WriteString("git clone ")
	//}
	////cd到指定目录
	////commondDocker.WriteString("cd " + paramInfo.DockerAndKongConfPath + "\n")
	////指定镜像编译并上传
	////commondDocker.WriteString("")
	//shell := commondDocker.String() //docker命令行参数拼接
	//cli := NewCon(paramInfo.SshHost, paramInfo.SshUser, paramInfo.SshPassword, paramInfo.SshPort)
	//output, err := cli.Run(shell) //执行docker命令
	//if err == nil {
	//	log.Println(output, "执行成功---build success--")
	//} else {
	//	log.Fatalf("%v\n%v\n%v", output, err, "执行失败!")
	//}
	//sess, err := SSHConnect(paramInfo.SshUser, paramInfo.SshPassword, paramInfo.SshHost, paramInfo.SshPort)
	//if err != nil {
	//	fmt.Println("789456", err)
	//}
	//err = sess.Run("ls")
	//if err != nil {
	//	fmt.Println("45613", err)
	//}
}

//创建sftp会话
func CreateSftp(sshHost, sshUser, sshPassword string, sshPort int) (*sftp.Client, error) {
	// 连接Linux服务器
	conn, err := pwdAuthConnect(sshHost, sshUser, sshPassword, sshPort)
	if err != nil {
		log.Fatal("连接Linux服务器失败", err)
		panic(err)
	}
	// 创建sftp会话
	client, err := sftp.NewClient(conn)
	if err != nil {
		return nil, err
	}
	return client, nil
}

//创建连接
func pwdAuthConnect(sshHost, sshUser, sshPassword string, sshPort int) (*ssh.Client, error) {
	config := ssh.ClientConfig{
		Timeout:         5 * time.Second,
		User:            sshUser,
		Auth:            []ssh.AuthMethod{ssh.Password((sshPassword))},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	addr := fmt.Sprintf("%s:%d", sshHost, sshPort)
	Client, err := ssh.Dial("tcp", addr, &config)
	if err != nil {
		log.Fatal("连接服务器失败", err)
		return nil, err
	}
	return Client, err
}

// 上传文件夹到服务器
func uploadDirectory(sftpClient *sftp.Client, localPath string, remotePath string) error {
	localFiles, err := ioutil.ReadDir(localPath)
	if err != nil {
		log.Fatal("read dir list fail ", err)
		return err
	}

	for _, backupDir := range localFiles {
		localFilePath := path.Join(localPath, backupDir.Name())
		remoteFilePath := path.Join(remotePath, backupDir.Name())
		if backupDir.IsDir() {
			err := sftpClient.Mkdir(remoteFilePath)
			if err != nil {
				return err
			}
			err = uploadDirectory(sftpClient, localFilePath, remoteFilePath)
			if err != nil {
				return err
			}
		} else {
			err := SrcFile(sftpClient, path.Join(localPath, backupDir.Name()), remotePath)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

//上传本地文件至服务器
func SrcFile(client *sftp.Client, localFilePath string, remoteFilePath string) error {
	//打开本地文件
	srcFile, err := os.Open(localFilePath)
	if err != nil {
		panic(err)
	}
	defer srcFile.Close()
	//绝对路径中获取文件名
	remoteFileName := path.Base(localFilePath)
	//打开服务器文件
	dstFile, err := client.Create(path.Join(remoteFilePath, remoteFileName))
	if err != nil {
		panic(err)
	}
	defer dstFile.Close()
	//将本地文件内容写入服务器文件
	buf := make([]byte, 1024)
	for {
		//将文件二进制内容写入buf字节切片中
		n, _ := srcFile.Read(buf)
		if n == 0 {
			break
		}
		dstFile.Write(buf)
	}
	return nil
}
