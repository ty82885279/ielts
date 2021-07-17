package controller

import (
	"context"
	"fmt"
	"ielts/models"
	"io"
	"os"
	"strconv"
	"time"

	"github.com/qiniu/api.v7/storage"

	"github.com/qiniu/api.v7/auth/qbox"

	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {

	var user models.User
	user.Account = c.PostForm("account")
	user.Psw = c.PostForm("psw")

	fmt.Println("绑定---", user)

	fmt.Println("创建用户")
	code, user1, err2 := models.RegisterUser(&user)
	switch code {
	case 3001:
		{
			c.JSON(http.StatusOK, gin.H{
				"status": "3001",
				"msg":    "用户已存在",
			})
		}
	case 3002:
		{
			c.JSON(http.StatusOK, gin.H{
				"status": "3002",
				"msg":    "注册失败",
				"data":   err2.Error(),
			})
		}
	case 3000:
		{
			c.JSON(http.StatusOK, gin.H{
				"status": "3000",
				"msg":    "注册成功",
				"data":   user1,
			})
		}
	}
}

func UpdateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBind(&user)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	uid, _ := strconv.Atoi(c.PostForm("id"))
	user.ID = uint(uid)
	fmt.Println("user -", user)
	fmt.Println("id----", user.ID)
	updateUser, err := models.UpdateUser(&user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "3001",
			"data":   err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "3000",
		"data":   updateUser,
	})
}

//添加成绩
func CreatRecord(c *gin.Context) {

	var record models.Record
	uid, _ := strconv.Atoi(c.PostForm("id"))
	record.UserID = uid
	err := c.ShouldBind(&record)
	if err != nil {
		fmt.Println("数据绑定失败", err.Error())
		return
	}
	err = models.CreatRecord(&record)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": "3001",
			"msg":    "添加成绩失败",
			"data":   err.Error(),
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "3000",
			"msg":    "添加成绩成功",
			"data":   record,
		})
	}
}

func GetRecords(c *gin.Context) {

	id := c.Query("userid")
	intid, _ := strconv.Atoi(id)
	records, err := models.GetRecords(intid)
	if err != nil {
		fmt.Println("获取成绩列表失败")
		c.JSON(http.StatusOK, gin.H{
			"status": "3001",
			"msg":    "获取成绩列表失败",
			"data":   err.Error(),
		})
		return
	} else {
		c.JSON(http.StatusOK, gin.H{
			"status": "3000",
			"msg":    "获取成绩列表成功",
			"data":   records,
		})
	}
}

func DeleteARecord(c *gin.Context) {

	uid := c.PostForm("uid")
	rid := c.PostForm("rid")
	fmt.Println(uid + "123")

	println("ok1---ok2", uid, rid)
	fmt.Println("来了")
	intuid, _ := strconv.Atoi(uid)
	intrid, _ := strconv.Atoi(rid)
	err := models.DeleteARecord(intuid, intrid)
	if err != nil {
		fmt.Println("删除失败")
		c.JSON(http.StatusOK, gin.H{
			"status": "3001",
			"msg":    "删除失败",
			"data":   err.Error(),
		})
		return
	}
	fmt.Println("删除成功")
	c.JSON(http.StatusOK, gin.H{
		"status": "3000",
		"msg":    "删除成功",
		"data":   uid + " ," + rid,
	})

}

func UploadAvatar(c *gin.Context) {

	fmt.Println("开始上传头像")

	address := "http://qb2mcx3ln.bkt.clouddn.com"
	file, header, err := c.Request.FormFile("file")

	if err != nil {
		fmt.Println("222----")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	//fmt.Println("333----")
	//log.Println("123----", header.Filename)

	fileName := strconv.FormatInt(time.Now().Unix(), 10) + "_" + header.Filename
	fmt.Println("name----", fileName)
	out, err := os.Create("static/avatars/" + fileName)
	defer out.Close()
	//
	if err != nil {

		fmt.Println("错误", err.Error())
		return
	}
	_, err = io.Copy(out, file)
	if err == nil {

		//fmt.Println("上传成功---")
	}
	//
	//fmt.Println("路径：", fileName)

	accessKey := "p5xuh5QpdvEh_ZLuvvtqRVvjs8261sXfifdXevKK"
	secretKey := "F1uH_MhB0OiHmGSI2s0M7sIkB06x8fkN3zcKyyVh"

	bucket := "ty82885279"
	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}
	putPolicy.Expires = 864000 //有效期
	mac := qbox.NewMac(accessKey, secretKey)
	fmt.Println("凭证：", mac)

	upToken := putPolicy.UploadToken(mac)

	fmt.Println("token", upToken)

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuadong
	// 是否使用https域名
	cfg.UseHTTPS = false
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false
	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	localFile := "static/avatars/" + fileName
	qiniuKey := fileName
	err = formUploader.PutFile(context.Background(), &ret, upToken, qiniuKey, localFile, nil)
	if err != nil {
		fmt.Println("上传文件失败,原因:", err)
		c.JSON(http.StatusOK, gin.H{
			"status": "3001",
			"msg":    "avatar upload fail",
			"data":   err.Error(),
		})
		return
	}

	fmt.Println("上传成功,key为:", ret.Key)
	address = address + ret.Key

	c.JSON(http.StatusOK, gin.H{
		"status": "3000",
		"msg":    "avatar upload success",
		"data":   address,
	})

	//删除本地头像
	//err = os.Remove(localFile)
	//if err != nil {
	//	fmt.Println("删除失败--", err.Error())
	//} else {
	//	fmt.Println("删除成功")
	//}
}

func Login(c *gin.Context) {

	account := c.Query("account")
	psw := c.Query("psw")

	user := models.GEtAUser(account, psw)
	if user == nil {

		fmt.Println("账户或者密码错误")
		c.JSON(http.StatusOK, gin.H{
			"status": "3001",
			"msg":    "账户或者密码错误",
			"data":   "nil",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"status": "3000",
		"msg":    "登录成功",
		"data":   user,
	})
}

//
func Index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"msd": "来了老弟？",
	})
}
