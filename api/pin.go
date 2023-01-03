package api

import (
	cid2 "github.com/ipfs/go-cid"
	"github.com/labstack/echo/v4"
	"light-estuary-node/core"
	"strings"
	"time"
)

func ConfigPinRouter(e *echo.Group, node *core.LightNode) {

	content := e.Group("/pin")
	content.POST("/add", func(c echo.Context) error {
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
			Name:       file.Filename,
			Size:       file.Size,
			Cid:        addNode.Cid().String(),
			Created_at: time.Time{},
			Updated_at: time.Time{},
		}

		node.DB.Create(&content)

		//node.Filclient.MinerByPeerID(nil).QueryStorageAskUnchecked(nil

		if err != nil {
			return err
		}
		c.Response().Write([]byte(addNode.Cid().String()))
		return nil
	})

	content.POST("/cid", func(c echo.Context) error {
		cidFromForm := c.FormValue("cid")
		cidNode, err := cid2.Decode(cidFromForm)
		if err != nil {
			return err
		}

		//	 get the node
		addNode, err := node.Node.Get(c.Request().Context(), cidNode)

		// get availabel staging buckets.
		// save the file to the database.
		size, err := addNode.Size()

		content := core.Content{
			Name:       addNode.Cid().String(),
			Size:       int64(size),
			Cid:        addNode.Cid().String(),
			Created_at: time.Time{},
			Updated_at: time.Time{},
		}

		node.DB.Create(&content)
		return nil
	})

	content.POST("/cids", func(c echo.Context) error {
		cids := c.FormValue("cids")

		// process each cids
		cidsArray := strings.Split(cids, ",")
		for _, cid := range cidsArray {
			cidNode, err := cid2.Decode(cid)
			if err != nil {
				return err
			}

			//	 get the node and save on the database
			addNode, err := node.Node.Get(c.Request().Context(), cidNode)

			// get availabel staging buckets.
			// save the file to the database.
			size, err := addNode.Size()

			content := core.Content{
				Name:       addNode.Cid().String(),
				Size:       int64(size),
				Cid:        addNode.Cid().String(),
				Created_at: time.Time{},
				Updated_at: time.Time{},
			}

			node.DB.Create(&content)
		}
		return nil
	})

}
