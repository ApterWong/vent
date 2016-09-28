package geography


import (
	"github.com/kataras/iris"
	"github.com/golang/protobuf/proto"
	pb "github.com/chenshaobo/vent/business_service/proto"
	"github.com/chenshaobo/vent/business_service/rpclient"
	"github.com/chenshaobo/vent/business_service/utils"
	"golang.org/x/net/context"
)

func SetupGeoApi(){
	userParty := iris.Party("/api/v1/geo")
	userParty.Put("/upload",GeoUpload)
}



func GeoUpload(c *iris.Context){
	body := c.PostBody()

	c2s := &pb.GeoUploadC2S{}
	s2c := &pb.CommonS2C{}


	err := proto.Unmarshal(body, c2s)
	if err != nil {
		s2c.ErrCode = utils.ErrParams
		utils.SetBody(c,s2c)
		return
	}

	conn := rpclient.Get("registerService")
	if conn == nil {
		s2c.ErrCode = utils.ErrServer
		utils.SetBody(c,s2c)
		return
	}
	utils.Info("register ")
	rc := pb.NewGeoManagerClient(conn)
	s2c, err = rc.UserGeoUpload(context.Background(), c2s)
	utils.Info("register rpc ok")
	utils.PrintErr(err)
	utils.SetBody(c,s2c)
}