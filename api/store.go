package api

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"light-estuary-node/core"
	"time"
)

func ConfigStoreRouter(e *echo.Group, node *core.LightNode) {
	// upload
	e.POST("/upload", func(c echo.Context) error {
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}
		src, err := file.Open()
		if err != nil {
			return err
		}

		addNode, err := node.Node.AddPinFile(c.Request().Context(), src, nil)

		// get availabel staging buckets.
		// save the file to the database.
		content := core.Content{
			Name:          file.Filename,
			Size:          file.Size,
			Cid:           addNode.Cid().String(),
			StagingBucket: "",
			Created_at:    time.Time{},
			Updated_at:    time.Time{},
		}

		fmt.Println(content) //	save content

		if err != nil {
			return err
		}
		c.Response().Write([]byte(addNode.Cid().String()))
		return nil
	})
}
