package tools

import (
	"store-chat/model/mysqls"
)

// 数据库的user表调试账号
var UserMap = userTest()
var UserMapById = userByIdTest()
var StoreMap = storeTest()

func userTest() map[string]mysqls.UserApi {
	userMaps := make(map[string]mysqls.UserApi)
	userMaps["2gDGQwDxsrX0UG8yRbophdHxHqD"] = mysqls.UserApi{
		UserID: int64(1788408218839183360),
		Token:  "2gDGQwDxsrX0UG8yRbophdHxHqD",
		Name:   "蜻蜓队长(管理员)",
		Fund:   int64(10000000),
	}
	userMaps["2gDGQugkyFF4MI10hK7WfT3W3Pe"] = mysqls.UserApi{
		UserID: int64(1788408218897903616),
		Token:  "2gDGQugkyFF4MI10hK7WfT3W3Pe",
		Name:   "和平星(管理员)",
		Fund:   int64(10000000),
	}
	userMaps["2gDGQvEugR6Y5riFp2kVLdc7J0O"] = mysqls.UserApi{
		UserID: int64(1788408218960818176),
		Token:  "2gDGQvEugR6Y5riFp2kVLdc7J0O",
		Name:   "蜘蛛侦探",
		Fund:   int64(10000000),
	}
	userMaps["2gDGQwhqJQczjkCikEvg3StOKSR"] = mysqls.UserApi{
		UserID: int64(1788408219027927040),
		Token:  "2gDGQwhqJQczjkCikEvg3StOKSR",
		Name:   "蝎子莱莱",
		Fund:   int64(10000000),
	}
	userMaps["2gDGQvpg5xTE3Qn0SIzbyDXpdma"] = mysqls.UserApi{
		UserID: int64(1788408219090841600),
		Token:  "2gDGQvpg5xTE3Qn0SIzbyDXpdma",
		Name:   "卡布达",
		Fund:   int64(10000000),
	}
	userMaps["2kJyfAuHY2JwTFMMyPwGFi61YkQ"] = mysqls.UserApi{
		UserID: int64(1821087524543295488),
		Token:  "2kJyfAuHY2JwTFMMyPwGFi61YkQ",
		Name:   "金龟次郎",
		Fund:   int64(10000000),
	}
	userMaps["2kJyf7WxQ8N1US5nyJjbStxLtf8"] = mysqls.UserApi{
		UserID: int64(1821087524656541696),
		Token:  "2kJyf7WxQ8N1US5nyJjbStxLtf8",
		Name:   "丸子龙",
		Fund:   int64(10000000),
	}
	userMaps["2kJyf7EbM6FQWCyrkM36JDYTShv"] = mysqls.UserApi{
		UserID: int64(1821087524857868288),
		Token:  "2kJyf7EbM6FQWCyrkM36JDYTShv",
		Name:   "蟑螂恶霸",
		Fund:   int64(10000000),
	}
	userMaps["2kJyf703zSyhdKFBa2b9l2bjxWX"] = mysqls.UserApi{
		UserID: int64(1821087524929171456),
		Token:  "2kJyf703zSyhdKFBa2b9l2bjxWX",
		Name:   "鲨鱼辣椒",
		Fund:   int64(10000000),
	}
	return userMaps
}

func userByIdTest() map[int64]mysqls.UserApi {
	userMaps := make(map[int64]mysqls.UserApi)
	userMaps[1788408218839183360] = mysqls.UserApi{
		UserID: int64(1788408218839183360),
		Token:  "2gDGQwDxsrX0UG8yRbophdHxHqD",
		Name:   "蜻蜓队长(管理员)",
		Fund:   int64(10000000),
	}
	userMaps[1788408218897903616] = mysqls.UserApi{
		UserID: int64(1788408218897903616),
		Token:  "2gDGQugkyFF4MI10hK7WfT3W3Pe",
		Name:   "愿望星(管理员)",
		Fund:   int64(10000000),
	}
	userMaps[1788408218960818176] = mysqls.UserApi{
		UserID: int64(1788408218960818176),
		Token:  "2gDGQvEugR6Y5riFp2kVLdc7J0O",
		Name:   "蜘蛛侦探",
		Fund:   int64(10000000),
	}
	userMaps[1788408219027927040] = mysqls.UserApi{
		UserID: int64(1788408219027927040),
		Token:  "2gDGQwhqJQczjkCikEvg3StOKSR",
		Name:   "蝎子莱莱",
		Fund:   int64(10000000),
	}
	userMaps[1788408219090841600] = mysqls.UserApi{
		UserID: int64(1788408219090841600),
		Token:  "2gDGQvpg5xTE3Qn0SIzbyDXpdma",
		Name:   "卡布达",
		Fund:   int64(10000000),
	}
	userMaps[1821087524543295488] = mysqls.UserApi{
		UserID: int64(1821087524543295488),
		Token:  "2kJyfAuHY2JwTFMMyPwGFi61YkQ",
		Name:   "金龟次郎",
		Fund:   int64(10000000),
	}
	userMaps[1821087524656541696] = mysqls.UserApi{
		UserID: int64(1821087524656541696),
		Token:  "2kJyf7WxQ8N1US5nyJjbStxLtf8",
		Name:   "丸子龙",
		Fund:   int64(10000000),
	}
	userMaps[1821087524857868288] = mysqls.UserApi{
		UserID: int64(1821087524857868288),
		Token:  "2kJyf7EbM6FQWCyrkM36JDYTShv",
		Name:   "蟑螂恶霸",
		Fund:   int64(10000000),
	}
	userMaps[1821087524929171456] = mysqls.UserApi{
		UserID: int64(1821087524929171456),
		Token:  "2kJyf703zSyhdKFBa2b9l2bjxWX",
		Name:   "鲨鱼辣椒",
		Fund:   int64(10000000),
	}
	return userMaps
}

func storeTest() map[int64]mysqls.StoresApi {
	storeMaps := make(map[int64]mysqls.StoresApi)
	storeMaps[int64(1810940924055547904)] = mysqls.StoresApi{
		StoreID: int64(1810940924055547904),
		Name:    "国服",
	}
	storeMaps[int64(1810941036622278656)] = mysqls.StoresApi{
		StoreID: int64(1810941036622278656),
		Name:    "欧服",
	}
	storeMaps[int64(1810941555327660032)] = mysqls.StoresApi{
		StoreID: int64(1810941555327660032),
		Name:    "美服",
	}
	return storeMaps
}
