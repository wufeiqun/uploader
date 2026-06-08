package main

import (
	"fmt"
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)

func main() {

	port := 8080
	filePath := "/data/uploader"
	// enableAuth := false
	// username := "admin"
	// password := "admin"

	router := gin.Default()
	// 文件下载
	router.StaticFS("/", http.Dir(filePath))
	
	// 单文件上传
	router.POST("/upload", func(c *gin.Context) {
		// 获取上传的文件
		file, err := c.FormFile("file")
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "上传文件失败"})
			return
		}

		// 设置保存路径
		dst := filepath.Join(filePath, file.Filename)

		// 保存文件到服务器
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "保存文件失败"})
			return
		}

		// 返回成功信息
		c.JSON(http.StatusOK, gin.H{"message": "上传成功", "filename": file.Filename})
	})

	// 多文件上传
	router.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		files := form.File["files"] // 获取多个文件

		var uploadedFiles []string

		for _, file := range files {
			dst := filepath.Join(filePath, file.Filename)
			if err := c.SaveUploadedFile(file, dst); err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "保存文件失败"})
				return
			}
			uploadedFiles = append(uploadedFiles, file.Filename)
		}

		c.JSON(http.StatusOK, gin.H{"message": "上传成功", "files": uploadedFiles})
	})
	router.Run(":" + fmt.Sprint(port))
	fmt.Println("Server started on port", port)
}
