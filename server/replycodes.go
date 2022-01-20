package server

const (
	ReplyServiceReadyNMinutes = 120
	ReplyDataConnectionReady  = 125
	ReplyFileStatusOk         = 150

	ReplyCommandOk                    = 200
	ReplySystemStatus                 = 211
	ReplyDirectoryStatus              = 212
	ReplyFileStatus                   = 213
	ReplyHelpMessage                  = 214
	ReplyNameSystem                   = 215
	ReplyReadyForNewUser              = 220
	ReplyClosingConnection            = 221
	ReplyDataConnectionOpenNoTransfer = 225
	ReplyClosingDataConnection        = 226
	ReplyPassiveMode                  = 227
	ReplyUserLoggedIn                 = 230
	ReplyRequestedFileOk              = 250
	ReplyPathNameCreated              = 257

	ReplyNeedPassword = 331
	ReplyNeedAccount  = 332

	ReplyServiceNotAvailable            = 421
	ReplyCannotOpenDataConnection       = 425
	ReplyConnectionClosed               = 426
	ReplyRequestedFileUnavailableOrBusy = 450
	ReplyRequestedLocalProcesingError   = 451
	ReplyRequestedInsufficientSpace     = 452

	ReplyCommandUnrecognized               = 500
	ReplyArgumentSyntaxError               = 501
	ReplyCommandNotImplemented             = 502
	ReplyBadCommandSequence                = 503
	ReplyCommandNotImplementedForParameter = 504
	ReplyNotLoggedIn                       = 530
	ReplyAccountNeeded                     = 532
	ReplyFileUnavailable                   = 550
	ReplyUnknownPageType                   = 551
	ReplyExceededStorageAllocation         = 552
	ReplyFileNameNotAllowed                = 553
)
