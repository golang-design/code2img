// Copyright 2021 The golang.design Initiative authors.
// All rights reserved. Use of this source code is governed
// by a GNU GPL-3.0 license that can be found in the LICENSE file.
//
// Written by Changkun Ou <changkun.de>

package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"golang.design/x/code2img"
)

func main() {
	router := gin.Default()
	router.Static("/api/v1/code2img/data/code", "./data/code")
	router.Static("/api/v1/code2img/data/images", "./data/images")
	router.POST("/api/v1/code2img", func(c *gin.Context) {
		b := struct {
			Code string `json:"code"`
		}{}
		if err := c.ShouldBindJSON(&b); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err))
			return
		}
		id := uuid.New().String()
		gofile := "./data/code/" + id + ".go"

		err := ioutil.WriteFile(gofile, []byte(b.Code), os.ModePerm)
		if err != nil {
			log.Printf("[%s]: write file error %v", gofile, err)
			c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err))
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
		defer cancel()

		buf, err := code2img.Render(ctx, code2img.LangGo, b.Code)
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("Error: %s", err))
			return
		}
		imgfile := "./data/images/" + id + ".png"
		if err := ioutil.WriteFile(imgfile, buf, os.ModePerm); err != nil {
			log.Printf("[%s]: write screenshot error %v", imgfile, err)
			return
		}
		c.String(http.StatusOK, "https://"+c.Request.Host+"/api/v1/code2img/data/images/"+id+".png")
	})

	s := &http.Server{Addr: ":80", Handler: router}
	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("shutting down...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := s.Shutdown(ctx); err != nil {
		log.Fatal("forced to shutdown: ", err)
	}
	log.Println("server exiting, good bye!")
}
