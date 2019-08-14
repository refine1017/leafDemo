package db

func init() {
	skeleton.RegisterChanRPC("FindPlayerByUsername", rpcFindPlayerByUsername)
	skeleton.RegisterChanRPC("CreatePlayerData", rpcCreatePlayerData)
	skeleton.RegisterChanRPC("LoadPlayerDataByUsername", rpcLoadPlayerDataByUsername)
	skeleton.RegisterChanRPC("SavePlayerData", rpcSavePlayerData)
}
