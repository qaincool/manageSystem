package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
)

func UploadFileHandler(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusInternalServerError, "读取file失败: "+err.Error())
		return
	}
	c.SaveUploadedFile(file, viper.GetString("video_path")+file.Filename)
	c.String(http.StatusOK, "上传成功！")
}
