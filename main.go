// It creates a new Echo instance, adds some middleware, creates a new WhyPFS node, creates a new GatewayHandler, and then
// adds a route to the Echo instance
package main

import (
	"context"
	"errors"
	"fmt"
	"github.com/gabriel-vasile/mimetype"
	cid2 "github.com/ipfs/go-cid"
	mdagipld "github.com/ipfs/go-ipld-format"
	"github.com/ipfs/go-merkledag"
	"github.com/ipfs/go-unixfs"
	uio "github.com/ipfs/go-unixfs/io"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"html/template"
	"io"
	"net/http"
	_ "net/http"
	"net/url"
	"os"
	"os/signal"
	gopath "path"
	"strings"
	"syscall"
	"time"
	"whypfs-gateway/gateway"
	"whypfs-gateway/metrics"

	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/multiformats/go-multiaddr"

	whypfs "github.com/application-research/whypfs-core"
	logging "github.com/ipfs/go-log/v2"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/xerrors"
)

type HttpError struct {
	Code    int    `json:"code,omitempty"`
	Reason  string `json:"reason"`
	Details string `json:"details"`
}

func (he HttpError) Error() string {
	if he.Details == "" {
		return he.Reason
	}
	return he.Reason + ": " + he.Details
}

type HttpErrorResponse struct {
	Error HttpError `json:"error"`
}

type StagingBucket struct {
	// ID is the unique identifier for the bucket.
	ID         string `json:"id"`
	Name       string `json:"name"`
	UUID       string `json:"uuid"`
	Status     string `json:"status"` // open, in-progress, completed (closed).
	Cid        string `json:"cid"`    // car file of the consolidated content
	created_at time.Time
	updated_at time.Time
}

type Content struct {
	ID            uint   `gorm:"primaryKey"`
	Name          string `json:"name"`
	Size          int64  `json:"size"`
	Cid           string `json:"cid"`
	StagingBucket string `json:"staging_bucket"` // where this content will be associated
	created_at    time.Time
	updated_at    time.Time
}

type ContentDeals struct {
	ID          uint   `gorm:"primary_key"`
	ContentID   uint   `gorm:"content_id"`
	DealID      uint   `gorm:"deal_id"`
	Status      string `gorm:"status"` // active, inactive.
	Replication int    `gorm:"replication"`
	created_at  time.Time
	updated_at  time.Time
}

var (
	node *whypfs.Node
	gw   *gateway.GatewayHandler
	// OsSignal signal used to shutdown
	OsSignal chan os.Signal
	log      = logging.Logger("gateway")
)

var defaultTestBootstrapPeers []multiaddr.Multiaddr

// Creating a list of multiaddresses that are used to bootstrap the network.
func BootstrapEstuaryPeers() []peer.AddrInfo {

	for _, s := range []string{
		"/ip4/145.40.90.135/tcp/6746/p2p/12D3KooWNTiHg8eQsTRx8XV7TiJbq3379EgwG6Mo3V3MdwAfThsx",
		"/ip4/139.178.68.217/tcp/6744/p2p/12D3KooWCVXs8P7iq6ao4XhfAmKWrEeuKFWCJgqe9jGDMTqHYBjw",
		"/ip4/147.75.49.71/tcp/6745/p2p/12D3KooWGBWx9gyUFTVQcKMTenQMSyE2ad9m7c9fpjS4NMjoDien",
		"/ip4/147.75.86.255/tcp/6745/p2p/12D3KooWFrnuj5o3tx4fGD2ZVJRyDqTdzGnU3XYXmBbWbc8Hs8Nd",
		"/ip4/3.134.223.177/tcp/6745/p2p/12D3KooWN8vAoGd6eurUSidcpLYguQiGZwt4eVgDvbgaS7kiGTup",
		"/ip4/35.74.45.12/udp/6746/quic/p2p/12D3KooWLV128pddyvoG6NBvoZw7sSrgpMTPtjnpu3mSmENqhtL7",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmNnooDu7bfjPFoTZYxMNLWUQJyrVwtbZg5gBMjTezGAJN",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmQCU2EcMqAqQPR2i9bChDtGNJchTbq5TbXJJ16u19uLTa",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmbLHAnMoJPWSCR5Zhtx6BHJX9KiKNN6tpvbUcqanj75Nb",
		"/dnsaddr/bootstrap.libp2p.io/p2p/QmcZf59bWwK5XFi76CZX8cbJ4BhTzzA3gU1ZjYZcYW3dwt",
	} {
		ma, err := multiaddr.NewMultiaddr(s)
		if err != nil {
			panic(err)
		}
		defaultTestBootstrapPeers = append(defaultTestBootstrapPeers, ma)
	}

	peers, _ := peer.AddrInfosFromP2pAddrs(defaultTestBootstrapPeers...)
	return peers
}

func main() {
	OsSignal = make(chan os.Signal, 1)

	// add a command to run API node

	// add a cli command to run commands

	// initialize sqlite database using gorm

	// auto create staging buckets
	// gateways
	GatewayRoutersConfig()

	LoopForever()
}

// LoopForever on signal processing
func LoopForever() {
	fmt.Printf("Entering infinite loop\n")

	signal.Notify(OsSignal, syscall.SIGINT, syscall.SIGTERM, syscall.SIGUSR1)
	_ = <-OsSignal

	fmt.Printf("Exiting infinite loop received OsSignal\n")
}

func ErrorHandler(err error, c echo.Context) {
	var httpRespErr *HttpError
	if xerrors.As(err, &httpRespErr) {
		log.Errorf("handler error: %s", err)
		if err := c.JSON(httpRespErr.Code, HttpErrorResponse{Error: *httpRespErr}); err != nil {
			log.Errorf("handler error: %s", err)
			return
		}
		return
	}

	var echoErr *echo.HTTPError
	if xerrors.As(err, &echoErr) {
		if err := c.JSON(echoErr.Code, HttpErrorResponse{
			Error: HttpError{
				Code:    echoErr.Code,
				Reason:  http.StatusText(echoErr.Code),
				Details: echoErr.Message.(string),
			},
		}); err != nil {
			log.Errorf("handler error: %s", err)
			return
		}
		return
	}

	log.Errorf("handler error: %s", err)
	if err := c.JSON(http.StatusInternalServerError, HttpErrorResponse{
		Error: HttpError{
			Code:    http.StatusInternalServerError,
			Reason:  http.StatusText(http.StatusInternalServerError),
			Details: err.Error(),
		},
	}); err != nil {
		log.Errorf("handler error: %s", err)
		return
	}
}

func GatewayRoutersConfig() {
	// Echo instance
	e := echo.New()
	e.File("/", "templates/index.html")

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.RemoveTrailingSlash())
	e.HTTPErrorHandler = ErrorHandler

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	whypfsPeer, err := whypfs.NewNode(whypfs.NewNodeParams{
		Ctx:       ctx,
		Datastore: whypfs.NewInMemoryDatastore(),
	})

	whypfsPeer.BootstrapPeers(BootstrapEstuaryPeers())

	node = whypfsPeer

	gw = gateway.NewGatewayHandler(node)
	if err != nil {
		panic(err)
	}

	//	gateway
	e.GET("/gw/ipfs/:path", GatewayResolverCheckHandlerDirectPath)
	e.GET("/gw/:path", GatewayResolverCheckHandlerDirectPath)
	e.GET("/ipfs/:path", GatewayResolverCheckHandlerDirectPath)
	e.GET("/gw/dir/:path", GatewayDirResolverCheckHandler)
	e.GET("/gw/file/:path", GatewayFileResolverCheckHandler)

	// metrics
	phandle := promhttp.Handler()
	e.GET("/debug/metrics/prometheus", func(e echo.Context) error {
		phandle.ServeHTTP(e.Response().Writer, e.Request())

		return nil
	})

	e.GET("/debug/metrics", func(e echo.Context) error {
		return e.JSON(http.StatusOK, "Ok")
		//return nil
	})

	e.GET("/debug/metrics", func(e echo.Context) error {
		metrics.Exporter().ServeHTTP(e.Response().Writer, e.Request())
		return nil
	})
	e.GET("/debug/stack", func(e echo.Context) error {
		err := metrics.WriteAllGoroutineStacks(e.Response().Writer)
		if err != nil {
			log.Error(err)
		}
		return err
	})

	e.GET("/debug/pprof/:prof", metrics.ServeProfile) // Upload for testing

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

		addNode, err := node.AddPinFile(c.Request().Context(), src, nil)

		// get availabel staging buckets.
		// save the file to the database.
		content := Content{
			Name:          file.Filename,
			Size:          file.Size,
			Cid:           addNode.Cid().String(),
			StagingBucket: "",
			created_at:    time.Time{},
			updated_at:    time.Time{},
		}

		fmt.Println(content) //	save content

		if err != nil {
			return err
		}
		c.Response().Write([]byte(addNode.Cid().String()))
		return nil
	})

	// Start server
	e.Logger.Fatal(e.Start("0.0.0.0:1313"))
}

func GatewayDirResolverCheckHandler(c echo.Context) error {
	p := c.Param("path")
	req := c.Request().Clone(c.Request().Context())
	req.URL.Path = p

	fmt.Printf("Request path: " + p)
	cid, err := cid2.Decode(p)

	if err != nil {
		return err
	}
	//	 check if file or dir.

	rscDir, err := node.GetDirectoryWithCid(c.Request().Context(), cid)
	if err != nil {
		return err
	}

	rscDir.GetNode()

	c.Response().Write([]byte("nice dir"))
	return nil
}

func GatewayFileResolverCheckHandler(c echo.Context) error {
	p := c.Param("path")
	req := c.Request().Clone(c.Request().Context())
	req.URL.Path = p

	fmt.Printf("Request path: " + p)
	cid, err := cid2.Decode(p)

	if err != nil {
		return err
	}
	//	 check if file or dir.
	rsc, err := node.GetFile(c.Request().Context(), cid)

	if err != nil {
		return err
	}
	content, err := io.ReadAll(rsc)
	if err != nil {
		return err
	}
	c.Response().Write(content)
	return nil
}

// `GatewayResolverCheckHandlerDirectPath` is a function that takes a `echo.Context` and returns an `error`
func GatewayResolverCheckHandlerDirectPath(c echo.Context) error {
	ctx := c.Request().Context()
	p := c.Param("path")
	req := c.Request().Clone(c.Request().Context())
	req.URL.Path = p

	sp := strings.Split(p, "/")
	cid, err := cid2.Decode(sp[0])
	if err != nil {
		return err
	}
	nd, err := node.Get(c.Request().Context(), cid)
	if err != nil {
		return err
	}

	if err != nil {
		panic(err)
	}

	switch nd := nd.(type) {
	case *merkledag.ProtoNode:
		n, err := unixfs.FSNodeFromBytes(nd.Data())
		if err != nil {
			panic(err)
		}
		if n.IsDir() {
			return ServeDir(ctx, nd, c.Response().Writer, req)
		}
		if n.Type() == unixfs.TSymlink {
			return fmt.Errorf("symlinks not supported")
		}
	case *merkledag.RawNode:
	default:
		return errors.New("unknown node type")
	}

	dr, err := uio.NewDagReader(ctx, nd, node.DAGService)
	if err != nil {
		return err
	}

	err = SniffMimeType(c.Response().Writer, dr)
	if err != nil {
		return err
	}

	http.ServeContent(c.Response().Writer, req, cid.String(), time.Time{}, dr)
	return nil
}

type Context struct {
	CustomLinks []CustomLinks
}

type CustomLinks struct {
	Href     string
	HrefCid  string
	LinkName string
	Size     string
}

func ServeDir(ctx context.Context, n mdagipld.Node, w http.ResponseWriter, req *http.Request) error {

	dir, err := uio.NewDirectoryFromNode(node.DAGService, n)
	if err != nil {
		return err
	}

	nd, err := dir.Find(ctx, "index.html")
	switch {
	case err == nil:
		dr, err := uio.NewDagReader(ctx, nd, node.DAGService)
		if err != nil {
			return err
		}

		http.ServeContent(w, req, "index.html", time.Time{}, dr)
		return nil
	default:
		return err
	case xerrors.Is(err, os.ErrNotExist):

	}

	templates, err := template.ParseFiles("templates/dir.html")
	if err != nil {
		return err
	}

	links := make([]CustomLinks, 0)
	templates.Lookup("dir.html")

	requestURI, err := url.ParseRequestURI(req.RequestURI)

	if err := dir.ForEachLink(ctx, func(lnk *mdagipld.Link) error {
		href := gopath.Join(requestURI.Path, lnk.Name)
		hrefCid := lnk.Cid.String()

		links = append(links, CustomLinks{Href: href, HrefCid: hrefCid, LinkName: lnk.Name})
		return nil
	}); err != nil {
		return err
	}

	//fmt.Fprintf(w, "</ul></body></html>")
	Context := Context{CustomLinks: links}
	templates.Execute(w, Context)

	return nil
}

func SniffMimeType(w http.ResponseWriter, dr uio.DagReader) error {
	// see kubo https://github.com/ipfs/kubo/blob/df222053856d3967ff0b4d6bc513bdb66ceedd6f/core/corehttp/gateway_handler_unixfs_file.go
	// see http ServeContent https://cs.opensource.google/go/go/+/refs/tags/go1.19.2:src/net/http/fs.go;l=221;drc=1f068f0dc7bc997446a7aac44cfc70746ad918e0

	// Calculate deterministic value for Content-Type HTTP header
	// (we prefer to do it here, rather than using implicit sniffing in http.ServeContent)
	var ctype string /**/
	// uses https://github.com/gabriel-vasile/mimetype library to determine the content type.
	// Fixes https://github.com/ipfs/kubo/issues/7252
	mimeType, err := mimetype.DetectReader(dr)
	if err != nil {
		http.Error(w, fmt.Sprintf("cannot detect content-type: %s", err.Error()), http.StatusInternalServerError)
		return err
	}

	ctype = mimeType.String()
	_, err = dr.Seek(0, io.SeekStart)
	if err != nil {
		http.Error(w, "seeker can't seek", http.StatusInternalServerError)
		return err
	}
	// Strip the encoding from the HTML Content-Type header and let the
	// browser figure it out.
	//
	// Fixes https://github.com/ipfs/kubo/issues/2203
	if strings.HasPrefix(ctype, "text/html;") {
		ctype = "text/html"
	}
	// Setting explicit Content-Type to avoid mime-type sniffing on the client
	// (unifies behavior across gateways and web browsers)
	w.Header().Set("Content-Type", ctype)
	return nil
}
