package tools

import (
	"store-chat/model/mysqls"
)

// 数据库的user表调试账号
var UserMap = userTest()
var StoreMap = storeTest()

func userTest() map[string]mysqls.UserApi {
	userMaps := make(map[string]mysqls.UserApi)
	userMaps["2gDGQwDxsrX0UG8yRbophdHxHqD"] = mysqls.UserApi{
		UserID: int64(1788408218839183360),
		Token:  "2gDGQwDxsrX0UG8yRbophdHxHqD",
		Name:   "爷爷",
		Fund:   int64(10000000),
	}
	userMaps["2gDGQugkyFF4MI10hK7WfT3W3Pe"] = mysqls.UserApi{
		UserID: int64(1788408218897903616),
		Token:  "2gDGQugkyFF4MI10hK7WfT3W3Pe",
		Name:   "姥姥",
		Fund:   int64(10000000),
	}
	userMaps["2gDGQvEugR6Y5riFp2kVLdc7J0O"] = mysqls.UserApi{
		UserID: int64(1788408218960818176),
		Token:  "2gDGQvEugR6Y5riFp2kVLdc7J0O",
		Name:   "爸爸",
		Fund:   int64(10000000),
	}
	userMaps["2gDGQwhqJQczjkCikEvg3StOKSR"] = mysqls.UserApi{
		UserID: int64(1788408219027927040),
		Token:  "2gDGQwhqJQczjkCikEvg3StOKSR",
		Name:   "妈妈",
		Fund:   int64(10000000),
	}
	userMaps["2gDGQvpg5xTE3Qn0SIzbyDXpdma"] = mysqls.UserApi{
		UserID: int64(1788408219090841600),
		Token:  "2gDGQvpg5xTE3Qn0SIzbyDXpdma",
		Name:   "儿子",
		Fund:   int64(10000000),
	}
	return userMaps
}

func storeTest() map[int64]mysqls.StoresApi {
	storeMaps := make(map[int64]mysqls.StoresApi)
	storeMaps[int64(1810940924055547904)] = mysqls.StoresApi{
		StoreID: int64(1810940924055547904),
		Name:    "蟑螂恶霸",
	}
	storeMaps[int64(1810941036622278656)] = mysqls.StoresApi{
		StoreID: int64(1810941036622278656),
		Name:    "鲨鱼辣椒",
	}
	storeMaps[int64(1810941555327660032)] = mysqls.StoresApi{
		StoreID: int64(1810941555327660032),
		Name:    "蜘蛛侦探",
	}
	return storeMaps
}
