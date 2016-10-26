// Copyright 2014, SuccessfumMatch Teq Inc. All rights reserved.
// Author TonyXu<tonycbcd@gmail.com>,
// Build on dev-0.0.1
// MIT Licensed

package routers

import (
    "github.com/gin-gonic/gin"
    mid "git.18.tl/tony/liangli/server/middlewares"
    ctl "git.18.tl/tony/liangli/server/controllers"
)

func SetStaticRouters(app *gin.Engine) {
    // Thumb service.
    setThumbRouter(app.Group("/thumb"))

    // THE UPLOADER
    setUploadRouter(app.Group("/"))

    // The Upload Pages
    setUploadPageRouter(app.Group("/uploadpage"))
}

// the Thumb service router.
func setThumbRouter(g *gin.RouterGroup) {
    g.GET("/:way/:width/:height/upload/:f1/:f2/:f3/:file", ctl.Thumb)
}

// to Uploading picture service router.
func setUploadRouter(g *gin.RouterGroup) {
    g.POST("/uploading", mid.StaticAuthMiddleware, ctl.Uploading)
    g.OPTIONS("/uploading", mid.StaticAuthMiddleware, ctl.Uploading)
    g.POST("/upload-sync", ctl.SyncUploadFile)
}

func setUploadPageRouter(g *gin.RouterGroup) {
    g.GET("/upload.html", mid.StaticAuthMiddleware, ctl.UploadHtml)
    g.POST("/upload.html", mid.StaticAuthMiddleware, ctl.DoUploadHtml)
    g.GET("/pictures/:picId", mid.StaticAuthMiddleware, ctl.GetPicturesById)
    g.POST("/picture/:picId/tag", mid.StaticAuthMiddleware, ctl.AddPictureTag)
    g.DELETE("/picture/:picId/tag/:id", mid.StaticAuthMiddleware, ctl.DeletePictureTag)
    g.PUT("/picture/:pictureId/desc", mid.StaticAuthMiddleware, ctl.UpdatePictureDesc)
    g.DELETE("/article/picture/:pictureId", mid.StaticAuthMiddleware, ctl.DeletePicture)

    // 是否使用了当前照片
    g.PUT("/article/:articleId/picture/:pictureId/state/:state", mid.StaticAuthMiddleware, ctl.SetPictureUseState)
    g.OPTIONS("/article/:articleId/picture/:pictureId/state/:state", mid.StaticAuthMiddleware, ctl.SetPictureUseStateOptions)
}
