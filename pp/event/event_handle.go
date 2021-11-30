package event

// client pp event handler
import (
	"github.com/stratosnet/sds/framework/core"
	"github.com/stratosnet/sds/msg/header"
)

// RegisterEventHandle
func RegisterEventHandle() {
	core.Register(header.RspGetPPList, RspGetPPList)
	core.Register(header.RspGetSPList, RspGetSPList)
	core.Register(header.RspRegister, RspRegister)
	core.Register(header.ReqRegister, ReqRegister)
	core.Register(header.RspActivatePP, RspActivate)
	core.Register(header.RspActivatedPP, RspActivated)
	core.Register(header.RspUpdateStakePP, RspUpdateStake)
	core.Register(header.RspUpdatedStakePP, RspUpdatedStake)
	core.Register(header.RspDeactivatePP, RspDeactivate)
	core.Register(header.RspDeactivatedPP, RspDeactivated)
	core.Register(header.RspPrepay, RspPrepay)
	core.Register(header.RspPrepaid, RspPrepaid)
	core.Register(header.RspMining, RspMining)
	core.Register(header.RspFindMyFileList, RspFindMyFileList)
	core.Register(header.ReqFindMyFileList, ReqFindMyFileList)
	core.Register(header.ReqUploadFileSlice, ReqUploadFileSlice)
	core.Register(header.RspUploadFile, RspUploadFile)
	core.Register(header.RspUploadFileSlice, RspUploadFileSlice)
	core.Register(header.RspReportUploadSliceResult, RspReportUploadSliceResult)
	core.Register(header.ReqFileStorageInfo, ReqFileStorageInfo)
	core.Register(header.ReqDownloadSlice, ReqDownloadSlice)
	core.Register(header.RspDownloadSlice, RspDownloadSlice)
	core.Register(header.RspReportDownloadResult, RspReportDownloadResult)
	core.Register(header.RspRegisterNewPP, RspRegisterNewPP)
	core.Register(header.ReqTransferNotice, ReqTransferNotice)
	core.Register(header.RspValidateTransferCer, RspValidateTransferCer)
	core.Register(header.ReqTransferDownload, ReqTransferDownload)
	core.Register(header.RspTransferDownload, RspTransferDownload)
	core.Register(header.RspTransferDownloadResult, RspTransferDownloadResult)
	core.Register(header.RspReportTransferResult, RspReportTransferResult)
	core.Register(header.RspDownloadSliceWrong, RspDownloadSliceWrong)
	core.Register(header.RspFileStorageInfo, RspFileStorageInfo)
	core.Register(header.ReqClearDownloadTask, ReqClearDownloadTask)
	core.Register(header.ReqGetHDInfo, ReqGetHDInfo)
	core.Register(header.RspGetHDInfo, RspGetHDInfo)
	core.Register(header.ReqDeleteSlice, ReqDeleteSlice)
	core.Register(header.RspDeleteSlice, RspDeleteSlice)
	core.Register(header.ReqMakeDirectory, ReqMakeDirectory)
	core.Register(header.RspMakeDirectory, RspMakeDirectory)
	core.Register(header.ReqRemoveDirectory, ReqRemoveDirectory)
	core.Register(header.RspRemoveDirectory, RspRemoveDirectory)
	core.Register(header.ReqMoveFileDirectory, ReqMoveFileDirectory)
	core.Register(header.RspMoveFileDirectory, RspMoveFileDirectory)
	core.Register(header.ReqDownloadSlicePause, ReqDownloadSlicePause)
	core.Register(header.RspDownloadSlicePause, RspDownloadSlicePause)
	core.Register(header.ReqCreateAlbum, ReqCreateAlbum)
	core.Register(header.RspCreateAlbum, RspCreateAlbum)
	core.Register(header.ReqFindMyAlbum, ReqFindMyAlbum)
	core.Register(header.RspFindMyAlbum, RspFindMyAlbum)
	core.Register(header.ReqEditAlbum, ReqEditAlbum)
	core.Register(header.RspEditAlbum, RspEditAlbum)
	core.Register(header.ReqAlbumContent, ReqAlbumContent)
	core.Register(header.RspAlbumContent, RspAlbumContent)
	core.Register(header.ReqSearchAlbum, ReqSearchAlbum)
	core.Register(header.RspSearchAlbum, RspSearchAlbum)
	core.Register(header.ReqCollectionAlbum, ReqCollectionAlbum)
	core.Register(header.RspCollectionAlbum, RspCollectionAlbum)
	core.Register(header.ReqAbstractAlbum, ReqAbstractAlbum)
	core.Register(header.RspAbstractAlbum, RspAbstractAlbum)
	core.Register(header.ReqMyCollectionAlbum, ReqMyCollectionAlbum)
	core.Register(header.RspMyCollectionAlbum, RspMyCollectionAlbum)
	core.Register(header.ReqDeleteAlbum, ReqDeleteAlbum)
	core.Register(header.RspDeleteAlbum, RspDeleteAlbum)
	core.Register(header.ReqSaveFolder, ReqSaveFolder)
	core.Register(header.RspSaveFolder, RspSaveFolder)
	core.Register(header.UploadSpeedOfProgress, UploadSpeedOfProgress)
	core.Register(header.ReqGetCapacity, ReqGetCapacity)
	core.Register(header.RspGetCapacity, RspGetCapacity)

	core.Register(header.ReqShareLink, ReqShareLink)
	core.Register(header.RspShareLink, RspShareLink)
	core.Register(header.ReqShareFile, ReqShareFile)
	core.Register(header.RspShareFile, RspShareFile)
	core.Register(header.ReqDeleteShare, ReqDeleteShare)
	core.Register(header.RspDeleteShare, RspDeleteShare)
	core.Register(header.ReqGetShareFile, ReqGetShareFile)
	core.Register(header.RspGetShareFile, RspGetShareFile)

	core.Register(header.ReqSaveFile, ReqSaveFile)
	core.Register(header.RspSaveFile, RspSaveFile)

	core.Register(header.ReqHeart, SendHeartBeat)
	core.Register(header.ReqSpLatencyCheck, ReqHBLatencyCheckSpList)
	core.Register(header.RspHeart, RspHBLatencyCheckSpList)
	core.Register(header.ReqDeleteFile, ReqDeleteFile)
	core.Register(header.RspDeleteFile, RspDeleteFile)
	core.Register(header.ReqConfig, ReqGetMyConfig)
	core.Register(header.RspConfig, RspGetMyConfig)

	core.Register(header.ReqInvite, ReqInvite)
	core.Register(header.RspInvite, RspInvite)
	core.Register(header.ReqGetReward, ReqGetReward)
	core.Register(header.RspGetReward, RspGetReward)

	core.Register(header.ReqFindDirectoryTree, ReqFindDirectoryTree)
	core.Register(header.RspFindDirectoryTree, RspFindDirectoryTree)

	core.Register(header.ReqFileSort, ReqFileSort)
	core.Register(header.RspFileSort, RspFileSort)

	core.Register(header.ReqFindDirectory, ReqFindDirectory)
	core.Register(header.RspFindDirectory, RspFindDirectory)
}
