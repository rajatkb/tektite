package transport

const (
	HandlerIDControllerRegisterL0Table = iota + 10
	HandlerIDControllerApplyChanges
	HandlerIDControllerQueryTablesInRange
	HandlerIDControllerRegisterTableListener
	HandlerIDControllerGetOffsets
	HandlerIDControllerPollForJob
	HandlerIDControllerGetTopicInfo
	HandlerIDControllerCreateTopic
	HandlerIDControllerDeleteTopic
	HandlerIDControllerGetGroupCoordinatorInfo
	HandlerIDControllerGenerateSequence
	HandlerIDMetaLocalCacheTopicAdded
	HandlerIDMetaLocalCacheTopicDeleted
	HandlerIDFetchCacheGetTableBytes
	HandlerIDFetcherTableRegisteredNotification
	HandlerIDTablePusherDirectWrite
	HandlerIDTablePusherDirectProduce
)
