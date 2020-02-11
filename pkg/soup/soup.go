package soup

import (
	"context"
	"task3/pkg/api"
)

//GRPCServer ...
type GRPCServer struct{}

//Square ...
func (server *GRPCServer) BuildMenu(ctx context.Context, req *api.SoupRequest) (*api.SoupResponse, error) {
	ret := SoupMaker(int(req.GetDays()))
	return &api.SoupResponse{Answer: ret}, nil
}

func SoupMaker(days int) string {
	menu := [5]string{"щи", "борщ", "щавелевый суп", "овсяный суп", "суп по-чабански"}
	var result string
	for i := 0; i < days; i++ {
		result += menu[i%5] + "\n"
	}
	result = result[:len(result)-1]
	return result
}
