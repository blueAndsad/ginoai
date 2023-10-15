package controller

import (
	"awesomeProject/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"os/exec"
)

func Hello(c *gin.Context) {
	id := c.Param("id")
	if id == "1" {
		fmt.Println("1")
	} else {
		fmt.Print("2")
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}
func Hello1(c *gin.Context) {
	var d models.Test
	err := c.ShouldBindJSON(&d)
	if err != nil {
		c.JSON(200, gin.H{
			"message": fmt.Sprintf("Bug: %v", err),
			"data":    gin.H{},
		})
	}
	if d.Name == "mxz" {
		cmd := exec.Command("sudo", "docker-compose", "-f", "docker-compose-basic-vpp-nrf.yaml", "up", "-d")
		cmd.Dir = "/home/dell/projects/docker_oai/oai-cn5g-fed/docker-compose"
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// 执行命令
		if err := cmd.Run(); err != nil {
			fmt.Println("yes")
		}
		//     cmd2 := exec.Command("sudo", "docker", "ps", "-a")
		//      cmd2.Dir = "/home/dell/projects/docker_oai/oai-cn5g-fed/docker-compose"
		//     cmd2.Stdout = os.Stdout
		//     cmd2.Stderr = os.Stderr

		// //  查看容器
		//     if err := cmd2.Run(); err != nil {
		//          fmt.Println("Error executing second command:", err)
		//     }
	} else {
		fmt.Println("其他名字: %v", d.Name)
	}

}
