// Code generated by SQLBoiler 4.16.1 (https://github.com/volatiletech/sqlboiler). DO NOT EDIT.
// This file is meant to be re-generated in place and/or deleted at any time.

package models

import "testing"

// This test suite runs each operation test in parallel.
// Example, if your database has 3 tables, the suite will run:
// table1, table2 and table3 Delete in parallel
// table1, table2 and table3 Insert in parallel, and so forth.
// It does NOT run each operation group in parallel.
// Separating the tests thusly grants avoidance of Postgres deadlocks.
func TestParent(t *testing.T) {
	t.Run("Diaries", testDiaries)
	t.Run("DiaryColors", testDiaryColors)
	t.Run("DiaryKeywords", testDiaryKeywords)
	t.Run("Friends", testFriends)
	t.Run("FriendRequests", testFriendRequests)
	t.Run("MediaFiles", testMediaFiles)
	t.Run("MediaFileBlockedHistories", testMediaFileBlockedHistories)
	t.Run("MediaFileRecheckRequests", testMediaFileRecheckRequests)
	t.Run("MediaTypes", testMediaTypes)
	t.Run("Members", testMembers)
	t.Run("MemberSocialTypes", testMemberSocialTypes)
	t.Run("MemberVerifies", testMemberVerifies)
	t.Run("Notices", testNotices)
	t.Run("Notifications", testNotifications)
	t.Run("NotificationTypes", testNotificationTypes)
	t.Run("Records", testRecords)
	t.Run("RecordKeywords", testRecordKeywords)
	t.Run("RecordTests", testRecordTests)
	t.Run("Replies", testReplies)
}

func TestDelete(t *testing.T) {
	t.Run("Diaries", testDiariesDelete)
	t.Run("DiaryColors", testDiaryColorsDelete)
	t.Run("DiaryKeywords", testDiaryKeywordsDelete)
	t.Run("Friends", testFriendsDelete)
	t.Run("FriendRequests", testFriendRequestsDelete)
	t.Run("MediaFiles", testMediaFilesDelete)
	t.Run("MediaFileBlockedHistories", testMediaFileBlockedHistoriesDelete)
	t.Run("MediaFileRecheckRequests", testMediaFileRecheckRequestsDelete)
	t.Run("MediaTypes", testMediaTypesDelete)
	t.Run("Members", testMembersDelete)
	t.Run("MemberSocialTypes", testMemberSocialTypesDelete)
	t.Run("MemberVerifies", testMemberVerifiesDelete)
	t.Run("Notices", testNoticesDelete)
	t.Run("Notifications", testNotificationsDelete)
	t.Run("NotificationTypes", testNotificationTypesDelete)
	t.Run("Records", testRecordsDelete)
	t.Run("RecordKeywords", testRecordKeywordsDelete)
	t.Run("RecordTests", testRecordTestsDelete)
	t.Run("Replies", testRepliesDelete)
}

func TestQueryDeleteAll(t *testing.T) {
	t.Run("Diaries", testDiariesQueryDeleteAll)
	t.Run("DiaryColors", testDiaryColorsQueryDeleteAll)
	t.Run("DiaryKeywords", testDiaryKeywordsQueryDeleteAll)
	t.Run("Friends", testFriendsQueryDeleteAll)
	t.Run("FriendRequests", testFriendRequestsQueryDeleteAll)
	t.Run("MediaFiles", testMediaFilesQueryDeleteAll)
	t.Run("MediaFileBlockedHistories", testMediaFileBlockedHistoriesQueryDeleteAll)
	t.Run("MediaFileRecheckRequests", testMediaFileRecheckRequestsQueryDeleteAll)
	t.Run("MediaTypes", testMediaTypesQueryDeleteAll)
	t.Run("Members", testMembersQueryDeleteAll)
	t.Run("MemberSocialTypes", testMemberSocialTypesQueryDeleteAll)
	t.Run("MemberVerifies", testMemberVerifiesQueryDeleteAll)
	t.Run("Notices", testNoticesQueryDeleteAll)
	t.Run("Notifications", testNotificationsQueryDeleteAll)
	t.Run("NotificationTypes", testNotificationTypesQueryDeleteAll)
	t.Run("Records", testRecordsQueryDeleteAll)
	t.Run("RecordKeywords", testRecordKeywordsQueryDeleteAll)
	t.Run("RecordTests", testRecordTestsQueryDeleteAll)
	t.Run("Replies", testRepliesQueryDeleteAll)
}

func TestSliceDeleteAll(t *testing.T) {
	t.Run("Diaries", testDiariesSliceDeleteAll)
	t.Run("DiaryColors", testDiaryColorsSliceDeleteAll)
	t.Run("DiaryKeywords", testDiaryKeywordsSliceDeleteAll)
	t.Run("Friends", testFriendsSliceDeleteAll)
	t.Run("FriendRequests", testFriendRequestsSliceDeleteAll)
	t.Run("MediaFiles", testMediaFilesSliceDeleteAll)
	t.Run("MediaFileBlockedHistories", testMediaFileBlockedHistoriesSliceDeleteAll)
	t.Run("MediaFileRecheckRequests", testMediaFileRecheckRequestsSliceDeleteAll)
	t.Run("MediaTypes", testMediaTypesSliceDeleteAll)
	t.Run("Members", testMembersSliceDeleteAll)
	t.Run("MemberSocialTypes", testMemberSocialTypesSliceDeleteAll)
	t.Run("MemberVerifies", testMemberVerifiesSliceDeleteAll)
	t.Run("Notices", testNoticesSliceDeleteAll)
	t.Run("Notifications", testNotificationsSliceDeleteAll)
	t.Run("NotificationTypes", testNotificationTypesSliceDeleteAll)
	t.Run("Records", testRecordsSliceDeleteAll)
	t.Run("RecordKeywords", testRecordKeywordsSliceDeleteAll)
	t.Run("RecordTests", testRecordTestsSliceDeleteAll)
	t.Run("Replies", testRepliesSliceDeleteAll)
}

func TestExists(t *testing.T) {
	t.Run("Diaries", testDiariesExists)
	t.Run("DiaryColors", testDiaryColorsExists)
	t.Run("DiaryKeywords", testDiaryKeywordsExists)
	t.Run("Friends", testFriendsExists)
	t.Run("FriendRequests", testFriendRequestsExists)
	t.Run("MediaFiles", testMediaFilesExists)
	t.Run("MediaFileBlockedHistories", testMediaFileBlockedHistoriesExists)
	t.Run("MediaFileRecheckRequests", testMediaFileRecheckRequestsExists)
	t.Run("MediaTypes", testMediaTypesExists)
	t.Run("Members", testMembersExists)
	t.Run("MemberSocialTypes", testMemberSocialTypesExists)
	t.Run("MemberVerifies", testMemberVerifiesExists)
	t.Run("Notices", testNoticesExists)
	t.Run("Notifications", testNotificationsExists)
	t.Run("NotificationTypes", testNotificationTypesExists)
	t.Run("Records", testRecordsExists)
	t.Run("RecordKeywords", testRecordKeywordsExists)
	t.Run("RecordTests", testRecordTestsExists)
	t.Run("Replies", testRepliesExists)
}

func TestFind(t *testing.T) {
	t.Run("Diaries", testDiariesFind)
	t.Run("DiaryColors", testDiaryColorsFind)
	t.Run("DiaryKeywords", testDiaryKeywordsFind)
	t.Run("Friends", testFriendsFind)
	t.Run("FriendRequests", testFriendRequestsFind)
	t.Run("MediaFiles", testMediaFilesFind)
	t.Run("MediaFileBlockedHistories", testMediaFileBlockedHistoriesFind)
	t.Run("MediaFileRecheckRequests", testMediaFileRecheckRequestsFind)
	t.Run("MediaTypes", testMediaTypesFind)
	t.Run("Members", testMembersFind)
	t.Run("MemberSocialTypes", testMemberSocialTypesFind)
	t.Run("MemberVerifies", testMemberVerifiesFind)
	t.Run("Notices", testNoticesFind)
	t.Run("Notifications", testNotificationsFind)
	t.Run("NotificationTypes", testNotificationTypesFind)
	t.Run("Records", testRecordsFind)
	t.Run("RecordKeywords", testRecordKeywordsFind)
	t.Run("RecordTests", testRecordTestsFind)
	t.Run("Replies", testRepliesFind)
}

func TestBind(t *testing.T) {
	t.Run("Diaries", testDiariesBind)
	t.Run("DiaryColors", testDiaryColorsBind)
	t.Run("DiaryKeywords", testDiaryKeywordsBind)
	t.Run("Friends", testFriendsBind)
	t.Run("FriendRequests", testFriendRequestsBind)
	t.Run("MediaFiles", testMediaFilesBind)
	t.Run("MediaFileBlockedHistories", testMediaFileBlockedHistoriesBind)
	t.Run("MediaFileRecheckRequests", testMediaFileRecheckRequestsBind)
	t.Run("MediaTypes", testMediaTypesBind)
	t.Run("Members", testMembersBind)
	t.Run("MemberSocialTypes", testMemberSocialTypesBind)
	t.Run("MemberVerifies", testMemberVerifiesBind)
	t.Run("Notices", testNoticesBind)
	t.Run("Notifications", testNotificationsBind)
	t.Run("NotificationTypes", testNotificationTypesBind)
	t.Run("Records", testRecordsBind)
	t.Run("RecordKeywords", testRecordKeywordsBind)
	t.Run("RecordTests", testRecordTestsBind)
	t.Run("Replies", testRepliesBind)
}

func TestOne(t *testing.T) {
	t.Run("Diaries", testDiariesOne)
	t.Run("DiaryColors", testDiaryColorsOne)
	t.Run("DiaryKeywords", testDiaryKeywordsOne)
	t.Run("Friends", testFriendsOne)
	t.Run("FriendRequests", testFriendRequestsOne)
	t.Run("MediaFiles", testMediaFilesOne)
	t.Run("MediaFileBlockedHistories", testMediaFileBlockedHistoriesOne)
	t.Run("MediaFileRecheckRequests", testMediaFileRecheckRequestsOne)
	t.Run("MediaTypes", testMediaTypesOne)
	t.Run("Members", testMembersOne)
	t.Run("MemberSocialTypes", testMemberSocialTypesOne)
	t.Run("MemberVerifies", testMemberVerifiesOne)
	t.Run("Notices", testNoticesOne)
	t.Run("Notifications", testNotificationsOne)
	t.Run("NotificationTypes", testNotificationTypesOne)
	t.Run("Records", testRecordsOne)
	t.Run("RecordKeywords", testRecordKeywordsOne)
	t.Run("RecordTests", testRecordTestsOne)
	t.Run("Replies", testRepliesOne)
}

func TestAll(t *testing.T) {
	t.Run("Diaries", testDiariesAll)
	t.Run("DiaryColors", testDiaryColorsAll)
	t.Run("DiaryKeywords", testDiaryKeywordsAll)
	t.Run("Friends", testFriendsAll)
	t.Run("FriendRequests", testFriendRequestsAll)
	t.Run("MediaFiles", testMediaFilesAll)
	t.Run("MediaFileBlockedHistories", testMediaFileBlockedHistoriesAll)
	t.Run("MediaFileRecheckRequests", testMediaFileRecheckRequestsAll)
	t.Run("MediaTypes", testMediaTypesAll)
	t.Run("Members", testMembersAll)
	t.Run("MemberSocialTypes", testMemberSocialTypesAll)
	t.Run("MemberVerifies", testMemberVerifiesAll)
	t.Run("Notices", testNoticesAll)
	t.Run("Notifications", testNotificationsAll)
	t.Run("NotificationTypes", testNotificationTypesAll)
	t.Run("Records", testRecordsAll)
	t.Run("RecordKeywords", testRecordKeywordsAll)
	t.Run("RecordTests", testRecordTestsAll)
	t.Run("Replies", testRepliesAll)
}

func TestCount(t *testing.T) {
	t.Run("Diaries", testDiariesCount)
	t.Run("DiaryColors", testDiaryColorsCount)
	t.Run("DiaryKeywords", testDiaryKeywordsCount)
	t.Run("Friends", testFriendsCount)
	t.Run("FriendRequests", testFriendRequestsCount)
	t.Run("MediaFiles", testMediaFilesCount)
	t.Run("MediaFileBlockedHistories", testMediaFileBlockedHistoriesCount)
	t.Run("MediaFileRecheckRequests", testMediaFileRecheckRequestsCount)
	t.Run("MediaTypes", testMediaTypesCount)
	t.Run("Members", testMembersCount)
	t.Run("MemberSocialTypes", testMemberSocialTypesCount)
	t.Run("MemberVerifies", testMemberVerifiesCount)
	t.Run("Notices", testNoticesCount)
	t.Run("Notifications", testNotificationsCount)
	t.Run("NotificationTypes", testNotificationTypesCount)
	t.Run("Records", testRecordsCount)
	t.Run("RecordKeywords", testRecordKeywordsCount)
	t.Run("RecordTests", testRecordTestsCount)
	t.Run("Replies", testRepliesCount)
}

func TestHooks(t *testing.T) {
	t.Run("Diaries", testDiariesHooks)
	t.Run("DiaryColors", testDiaryColorsHooks)
	t.Run("DiaryKeywords", testDiaryKeywordsHooks)
	t.Run("Friends", testFriendsHooks)
	t.Run("FriendRequests", testFriendRequestsHooks)
	t.Run("MediaFiles", testMediaFilesHooks)
	t.Run("MediaFileBlockedHistories", testMediaFileBlockedHistoriesHooks)
	t.Run("MediaFileRecheckRequests", testMediaFileRecheckRequestsHooks)
	t.Run("MediaTypes", testMediaTypesHooks)
	t.Run("Members", testMembersHooks)
	t.Run("MemberSocialTypes", testMemberSocialTypesHooks)
	t.Run("MemberVerifies", testMemberVerifiesHooks)
	t.Run("Notices", testNoticesHooks)
	t.Run("Notifications", testNotificationsHooks)
	t.Run("NotificationTypes", testNotificationTypesHooks)
	t.Run("Records", testRecordsHooks)
	t.Run("RecordKeywords", testRecordKeywordsHooks)
	t.Run("RecordTests", testRecordTestsHooks)
	t.Run("Replies", testRepliesHooks)
}

func TestInsert(t *testing.T) {
	t.Run("Diaries", testDiariesInsert)
	t.Run("Diaries", testDiariesInsertWhitelist)
	t.Run("DiaryColors", testDiaryColorsInsert)
	t.Run("DiaryColors", testDiaryColorsInsertWhitelist)
	t.Run("DiaryKeywords", testDiaryKeywordsInsert)
	t.Run("DiaryKeywords", testDiaryKeywordsInsertWhitelist)
	t.Run("Friends", testFriendsInsert)
	t.Run("Friends", testFriendsInsertWhitelist)
	t.Run("FriendRequests", testFriendRequestsInsert)
	t.Run("FriendRequests", testFriendRequestsInsertWhitelist)
	t.Run("MediaFiles", testMediaFilesInsert)
	t.Run("MediaFiles", testMediaFilesInsertWhitelist)
	t.Run("MediaFileBlockedHistories", testMediaFileBlockedHistoriesInsert)
	t.Run("MediaFileBlockedHistories", testMediaFileBlockedHistoriesInsertWhitelist)
	t.Run("MediaFileRecheckRequests", testMediaFileRecheckRequestsInsert)
	t.Run("MediaFileRecheckRequests", testMediaFileRecheckRequestsInsertWhitelist)
	t.Run("MediaTypes", testMediaTypesInsert)
	t.Run("MediaTypes", testMediaTypesInsertWhitelist)
	t.Run("Members", testMembersInsert)
	t.Run("Members", testMembersInsertWhitelist)
	t.Run("MemberSocialTypes", testMemberSocialTypesInsert)
	t.Run("MemberSocialTypes", testMemberSocialTypesInsertWhitelist)
	t.Run("MemberVerifies", testMemberVerifiesInsert)
	t.Run("MemberVerifies", testMemberVerifiesInsertWhitelist)
	t.Run("Notices", testNoticesInsert)
	t.Run("Notices", testNoticesInsertWhitelist)
	t.Run("Notifications", testNotificationsInsert)
	t.Run("Notifications", testNotificationsInsertWhitelist)
	t.Run("NotificationTypes", testNotificationTypesInsert)
	t.Run("NotificationTypes", testNotificationTypesInsertWhitelist)
	t.Run("Records", testRecordsInsert)
	t.Run("Records", testRecordsInsertWhitelist)
	t.Run("RecordKeywords", testRecordKeywordsInsert)
	t.Run("RecordKeywords", testRecordKeywordsInsertWhitelist)
	t.Run("RecordTests", testRecordTestsInsert)
	t.Run("RecordTests", testRecordTestsInsertWhitelist)
	t.Run("Replies", testRepliesInsert)
	t.Run("Replies", testRepliesInsertWhitelist)
}

func TestReload(t *testing.T) {
	t.Run("Diaries", testDiariesReload)
	t.Run("DiaryColors", testDiaryColorsReload)
	t.Run("DiaryKeywords", testDiaryKeywordsReload)
	t.Run("Friends", testFriendsReload)
	t.Run("FriendRequests", testFriendRequestsReload)
	t.Run("MediaFiles", testMediaFilesReload)
	t.Run("MediaFileBlockedHistories", testMediaFileBlockedHistoriesReload)
	t.Run("MediaFileRecheckRequests", testMediaFileRecheckRequestsReload)
	t.Run("MediaTypes", testMediaTypesReload)
	t.Run("Members", testMembersReload)
	t.Run("MemberSocialTypes", testMemberSocialTypesReload)
	t.Run("MemberVerifies", testMemberVerifiesReload)
	t.Run("Notices", testNoticesReload)
	t.Run("Notifications", testNotificationsReload)
	t.Run("NotificationTypes", testNotificationTypesReload)
	t.Run("Records", testRecordsReload)
	t.Run("RecordKeywords", testRecordKeywordsReload)
	t.Run("RecordTests", testRecordTestsReload)
	t.Run("Replies", testRepliesReload)
}

func TestReloadAll(t *testing.T) {
	t.Run("Diaries", testDiariesReloadAll)
	t.Run("DiaryColors", testDiaryColorsReloadAll)
	t.Run("DiaryKeywords", testDiaryKeywordsReloadAll)
	t.Run("Friends", testFriendsReloadAll)
	t.Run("FriendRequests", testFriendRequestsReloadAll)
	t.Run("MediaFiles", testMediaFilesReloadAll)
	t.Run("MediaFileBlockedHistories", testMediaFileBlockedHistoriesReloadAll)
	t.Run("MediaFileRecheckRequests", testMediaFileRecheckRequestsReloadAll)
	t.Run("MediaTypes", testMediaTypesReloadAll)
	t.Run("Members", testMembersReloadAll)
	t.Run("MemberSocialTypes", testMemberSocialTypesReloadAll)
	t.Run("MemberVerifies", testMemberVerifiesReloadAll)
	t.Run("Notices", testNoticesReloadAll)
	t.Run("Notifications", testNotificationsReloadAll)
	t.Run("NotificationTypes", testNotificationTypesReloadAll)
	t.Run("Records", testRecordsReloadAll)
	t.Run("RecordKeywords", testRecordKeywordsReloadAll)
	t.Run("RecordTests", testRecordTestsReloadAll)
	t.Run("Replies", testRepliesReloadAll)
}

func TestSelect(t *testing.T) {
	t.Run("Diaries", testDiariesSelect)
	t.Run("DiaryColors", testDiaryColorsSelect)
	t.Run("DiaryKeywords", testDiaryKeywordsSelect)
	t.Run("Friends", testFriendsSelect)
	t.Run("FriendRequests", testFriendRequestsSelect)
	t.Run("MediaFiles", testMediaFilesSelect)
	t.Run("MediaFileBlockedHistories", testMediaFileBlockedHistoriesSelect)
	t.Run("MediaFileRecheckRequests", testMediaFileRecheckRequestsSelect)
	t.Run("MediaTypes", testMediaTypesSelect)
	t.Run("Members", testMembersSelect)
	t.Run("MemberSocialTypes", testMemberSocialTypesSelect)
	t.Run("MemberVerifies", testMemberVerifiesSelect)
	t.Run("Notices", testNoticesSelect)
	t.Run("Notifications", testNotificationsSelect)
	t.Run("NotificationTypes", testNotificationTypesSelect)
	t.Run("Records", testRecordsSelect)
	t.Run("RecordKeywords", testRecordKeywordsSelect)
	t.Run("RecordTests", testRecordTestsSelect)
	t.Run("Replies", testRepliesSelect)
}

func TestUpdate(t *testing.T) {
	t.Run("Diaries", testDiariesUpdate)
	t.Run("DiaryColors", testDiaryColorsUpdate)
	t.Run("DiaryKeywords", testDiaryKeywordsUpdate)
	t.Run("Friends", testFriendsUpdate)
	t.Run("FriendRequests", testFriendRequestsUpdate)
	t.Run("MediaFiles", testMediaFilesUpdate)
	t.Run("MediaFileBlockedHistories", testMediaFileBlockedHistoriesUpdate)
	t.Run("MediaFileRecheckRequests", testMediaFileRecheckRequestsUpdate)
	t.Run("MediaTypes", testMediaTypesUpdate)
	t.Run("Members", testMembersUpdate)
	t.Run("MemberSocialTypes", testMemberSocialTypesUpdate)
	t.Run("MemberVerifies", testMemberVerifiesUpdate)
	t.Run("Notices", testNoticesUpdate)
	t.Run("Notifications", testNotificationsUpdate)
	t.Run("NotificationTypes", testNotificationTypesUpdate)
	t.Run("Records", testRecordsUpdate)
	t.Run("RecordKeywords", testRecordKeywordsUpdate)
	t.Run("RecordTests", testRecordTestsUpdate)
	t.Run("Replies", testRepliesUpdate)
}

func TestSliceUpdateAll(t *testing.T) {
	t.Run("Diaries", testDiariesSliceUpdateAll)
	t.Run("DiaryColors", testDiaryColorsSliceUpdateAll)
	t.Run("DiaryKeywords", testDiaryKeywordsSliceUpdateAll)
	t.Run("Friends", testFriendsSliceUpdateAll)
	t.Run("FriendRequests", testFriendRequestsSliceUpdateAll)
	t.Run("MediaFiles", testMediaFilesSliceUpdateAll)
	t.Run("MediaFileBlockedHistories", testMediaFileBlockedHistoriesSliceUpdateAll)
	t.Run("MediaFileRecheckRequests", testMediaFileRecheckRequestsSliceUpdateAll)
	t.Run("MediaTypes", testMediaTypesSliceUpdateAll)
	t.Run("Members", testMembersSliceUpdateAll)
	t.Run("MemberSocialTypes", testMemberSocialTypesSliceUpdateAll)
	t.Run("MemberVerifies", testMemberVerifiesSliceUpdateAll)
	t.Run("Notices", testNoticesSliceUpdateAll)
	t.Run("Notifications", testNotificationsSliceUpdateAll)
	t.Run("NotificationTypes", testNotificationTypesSliceUpdateAll)
	t.Run("Records", testRecordsSliceUpdateAll)
	t.Run("RecordKeywords", testRecordKeywordsSliceUpdateAll)
	t.Run("RecordTests", testRecordTestsSliceUpdateAll)
	t.Run("Replies", testRepliesSliceUpdateAll)
}
