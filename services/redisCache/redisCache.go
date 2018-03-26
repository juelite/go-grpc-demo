package redisCache

import (
	"frrpc/services"
	"frrpc/const"
	pb "frrpc/protoFile"
)

var client Client

func init()  {
	var host , pass string
	baseServ := services.BaseService{}
	host = baseServ.GetVal("redishost")
	pass = baseServ.GetVal("redispass")
	client.Addr = host
	client.Password = pass
}

/**
 * 写入redis缓存
 * @param name string 键名
 * @param exp int64 有效期，0为不过期，单位秒
 * @param val string 值
 * @return err error 写入结果
 */
func RedisCache(name string , exp int64 , val string) (*pb.RedisCacheReply , error) {
	base := services.BaseService{}
	base.LogInfo("redis_cache_request" , "name:"+name+" val:"+val)
	var err error
	if exp > 0 {
		err = client.Setex(name , exp , val)
	} else {
		err = client.Set(name , val)
	}

	ret := pb.RedisCacheReply{
		Code:_const.STATUS_SUCCESS,
		Message:"success",
		Data:nil,
	}
	if err != nil {
		ret.Code = _const.REDIS_WRITE_ERR
		ret.Message = "缓存写入失败"
	}
	return &ret , nil
}