package mtproto

import "fmt"

const (
	crc_boolFalse                              = 0xbc799737
	crc_boolTrue                               = 0x997275b5
	crc_error                                  = 0xc4b9f9bb
	crc_null                                   = 0x56730bcc
	crc_inputPeerEmpty                         = 0x7f3b18ea
	crc_inputPeerSelf                          = 0x7da07ec9
	crc_inputPeerContact                       = 0x1023dbe8
	crc_inputPeerForeign                       = 0x9b447325
	crc_inputPeerChat                          = 0x179be863
	crc_inputUserEmpty                         = 0xb98886cf
	crc_inputUserSelf                          = 0xf7c1b13f
	crc_inputUserContact                       = 0x86e94f65
	crc_inputUserForeign                       = 0x655e74ff
	crc_inputPhoneContact                      = 0xf392b7f4
	crc_inputFile                              = 0xf52ff27f
	crc_inputMediaEmpty                        = 0x9664f57f
	crc_inputMediaUploadedPhoto                = 0x2dc53a7d
	crc_inputMediaPhoto                        = 0x8f2ab2ec
	crc_inputMediaGeoPoint                     = 0xf9c44144
	crc_inputMediaContact                      = 0xa6e45987
	crc_inputMediaUploadedVideo                = 0x133ad6f6
	crc_inputMediaUploadedThumbVideo           = 0x9912dabf
	crc_inputMediaVideo                        = 0x7f023ae6
	crc_inputChatPhotoEmpty                    = 0x1ca48f57
	crc_inputChatUploadedPhoto                 = 0x94254732
	crc_inputChatPhoto                         = 0xb2e1bf08
	crc_inputGeoPointEmpty                     = 0xe4c123d6
	crc_inputGeoPoint                          = 0xf3b7acc9
	crc_inputPhotoEmpty                        = 0x1cd7bf0d
	crc_inputPhoto                             = 0xfb95c6c4
	crc_inputVideoEmpty                        = 0x5508ec75
	crc_inputVideo                             = 0xee579652
	crc_inputFileLocation                      = 0x14637196
	crc_inputVideoFileLocation                 = 0x3d0364ec
	crc_inputPhotoCropAuto                     = 0xade6b004
	crc_inputPhotoCrop                         = 0xd9915325
	crc_inputAppEvent                          = 0x770656a8
	crc_peerUser                               = 0x9db1bc6d
	crc_peerChat                               = 0xbad0e5bb
	crc_storage_fileUnknown                    = 0xaa963b05
	crc_storage_fileJpeg                       = 0x007efe0e
	crc_storage_fileGif                        = 0xcae1aadf
	crc_storage_filePng                        = 0x0a4f63c0
	crc_storage_filePdf                        = 0xae1e508d
	crc_storage_fileMp3                        = 0x528a0677
	crc_storage_fileMov                        = 0x4b09ebbc
	crc_storage_filePartial                    = 0x40bc6f52
	crc_storage_fileMp4                        = 0xb3cea0e4
	crc_storage_fileWebp                       = 0x1081464c
	crc_fileLocationUnavailable                = 0x7c596b46
	crc_fileLocation                           = 0x53d69076
	crc_userEmpty                              = 0x200250ba
	crc_userSelf                               = 0x7007b451
	crc_userContact                            = 0xcab35e18
	crc_userRequest                            = 0xd9ccc4ef
	crc_userForeign                            = 0x075cf7a8
	crc_userDeleted                            = 0xd6016d7a
	crc_userProfilePhotoEmpty                  = 0x4f11bae1
	crc_userProfilePhoto                       = 0xd559d8c8
	crc_userStatusEmpty                        = 0x09d05049
	crc_userStatusOnline                       = 0xedb93949
	crc_userStatusOffline                      = 0x008c703f
	crc_chatEmpty                              = 0x9ba2d800
	crc_chat                                   = 0x6e9c9bc7
	crc_chatForbidden                          = 0xfb0ccc41
	crc_chatFull                               = 0x630e61be
	crc_chatParticipant                        = 0xc8d7493e
	crc_chatParticipantsForbidden              = 0x0fd2bb8a
	crc_chatParticipants                       = 0x7841b415
	crc_chatPhotoEmpty                         = 0x37c1011c
	crc_chatPhoto                              = 0x6153276a
	crc_messageEmpty                           = 0x83e5de54
	crc_message                                = 0x567699b3
	crc_messageForwarded                       = 0xa367e716
	crc_messageService                         = 0x1d86f70e
	crc_messageMediaEmpty                      = 0x3ded6320
	crc_messageMediaPhoto                      = 0xc8c45a2a
	crc_messageMediaVideo                      = 0xa2d24290
	crc_messageMediaGeo                        = 0x56e0d474
	crc_messageMediaContact                    = 0x5e7d2f39
	crc_messageMediaUnsupported                = 0x29632a36
	crc_messageActionEmpty                     = 0xb6aef7b0
	crc_messageActionChatCreate                = 0xa6638b9a
	crc_messageActionChatEditTitle             = 0xb5a1ce5a
	crc_messageActionChatEditPhoto             = 0x7fcb13a8
	crc_messageActionChatDeletePhoto           = 0x95e3fbef
	crc_messageActionChatAddUser               = 0x5e3cfc4b
	crc_messageActionChatDeleteUser            = 0xb2ae9b0c
	crc_dialog                                 = 0xab3a99ac
	crc_photoEmpty                             = 0x2331b22d
	crc_photo                                  = 0x22b56751
	crc_photoSizeEmpty                         = 0x0e17e23c
	crc_photoSize                              = 0x77bfb61b
	crc_photoCachedSize                        = 0xe9a734fa
	crc_videoEmpty                             = 0xc10658a8
	crc_video                                  = 0x388fa391
	crc_geoPointEmpty                          = 0x1117dd5f
	crc_geoPoint                               = 0x2049d70c
	crc_auth_checkedPhone                      = 0xe300cc3b
	crc_auth_sentCode                          = 0xefed51d9
	crc_auth_authorization                     = 0xf6b673a4
	crc_auth_exportedAuthorization             = 0xdf969c2d
	crc_inputNotifyPeer                        = 0xb8bc5b0c
	crc_inputNotifyUsers                       = 0x193b4417
	crc_inputNotifyChats                       = 0x4a95e84e
	crc_inputNotifyAll                         = 0xa429b886
	crc_inputPeerNotifyEventsEmpty             = 0xf03064d8
	crc_inputPeerNotifyEventsAll               = 0xe86a2c74
	crc_inputPeerNotifySettings                = 0x46a2ce98
	crc_peerNotifyEventsEmpty                  = 0xadd53cb3
	crc_peerNotifyEventsAll                    = 0x6d1ded88
	crc_peerNotifySettingsEmpty                = 0x70a68512
	crc_peerNotifySettings                     = 0x8d5e11ee
	crc_wallPaper                              = 0xccb03657
	crc_userFull                               = 0x771095da
	crc_contact                                = 0xf911c994
	crc_importedContact                        = 0xd0028438
	crc_contactBlocked                         = 0x561bc879
	crc_contactSuggested                       = 0x3de191a1
	crc_contactStatus                          = 0xd3680c61
	crc_chatLocated                            = 0x3631cf4c
	crc_contacts_foreignLinkUnknown            = 0x133421f8
	crc_contacts_foreignLinkRequested          = 0xa7801f47
	crc_contacts_foreignLinkMutual             = 0x1bea8ce1
	crc_contacts_myLinkEmpty                   = 0xd22a1c60
	crc_contacts_myLinkRequested               = 0x6c69efee
	crc_contacts_myLinkContact                 = 0xc240ebd9
	crc_contacts_link                          = 0xeccea3f5
	crc_contacts_contactsNotModified           = 0xb74ba9d2
	crc_contacts_contacts                      = 0x6f8b8cb2
	crc_contacts_importedContacts              = 0xad524315
	crc_contacts_blocked                       = 0x1c138d15
	crc_contacts_blockedSlice                  = 0x900802a1
	crc_contacts_suggested                     = 0x5649dcc5
	crc_messages_dialogs                       = 0x15ba6c40
	crc_messages_dialogsSlice                  = 0x71e094f3
	crc_messages_messages                      = 0x8c718e87
	crc_messages_messagesSlice                 = 0x0b446ae3
	crc_messages_messageEmpty                  = 0x3f4e0648
	crc_messages_statedMessages                = 0x969478bb
	crc_messages_statedMessage                 = 0xd07ae726
	crc_messages_sentMessage                   = 0xd1f4d35c
	crc_messages_chats                         = 0x8150cbd8
	crc_messages_chatFull                      = 0xe5d7d19c
	crc_messages_affectedHistory               = 0xb7de36f2
	crc_inputMessagesFilterEmpty               = 0x57e2f66c
	crc_inputMessagesFilterPhotos              = 0x9609a51c
	crc_inputMessagesFilterVideo               = 0x9fc00e65
	crc_inputMessagesFilterPhotoVideo          = 0x56e9f0e4
	crc_inputMessagesFilterPhotoVideoDocuments = 0xd95e73bb
	crc_inputMessagesFilterDocument            = 0x9eddf188
	crc_inputMessagesFilterAudio               = 0xcfc87522
	crc_updateNewMessage                       = 0x013abdb3
	crc_updateMessageID                        = 0x4e90bfd6
	crc_updateReadMessages                     = 0xc6649e31
	crc_updateDeleteMessages                   = 0xa92bfe26
	crc_updateUserTyping                       = 0x5c486927
	crc_updateChatUserTyping                   = 0x9a65ea1f
	crc_updateChatParticipants                 = 0x07761198
	crc_updateUserStatus                       = 0x1bfbd823
	crc_updateUserName                         = 0xa7332b73
	crc_updateUserPhoto                        = 0x95313b0c
	crc_updateContactRegistered                = 0x2575bbb9
	crc_updateContactLink                      = 0x51a48a9a
	crc_updateNewAuthorization                 = 0x8f06529a
	crc_updates_state                          = 0xa56c2a3e
	crc_updates_differenceEmpty                = 0x5d75a138
	crc_updates_difference                     = 0x00f49ca0
	crc_updates_differenceSlice                = 0xa8fb1981
	crc_updatesTooLong                         = 0xe317af7e
	crc_updateShortMessage                     = 0xd3f45784
	crc_updateShortChatMessage                 = 0x2b2fbd4e
	crc_updateShort                            = 0x78d4dec1
	crc_updatesCombined                        = 0x725b04c3
	crc_updates                                = 0x74ae4240
	crc_photos_photos                          = 0x8dca6aa5
	crc_photos_photosSlice                     = 0x15051f54
	crc_photos_photo                           = 0x20212ca8
	crc_upload_file                            = 0x096a18d5
	crc_dcOption                               = 0x2ec2a43c
	crc_config                                 = 0x7dae33e0
	crc_nearestDc                              = 0x8e1a1775
	crc_help_appUpdate                         = 0x8987f311
	crc_help_noAppUpdate                       = 0xc45a6536
	crc_help_inviteText                        = 0x18cb9f78
	crc_messages_statedMessagesLinks           = 0x3e74f5c6
	crc_messages_statedMessageLink             = 0xa9af2881
	crc_messages_sentMessageLink               = 0xe9db4a3f
	crc_inputGeoChat                           = 0x74d456fa
	crc_inputNotifyGeoChatPeer                 = 0x4d8ddec8
	crc_geoChat                                = 0x75eaea5a
	crc_geoChatMessageEmpty                    = 0x60311a9b
	crc_geoChatMessage                         = 0x4505f8e1
	crc_geoChatMessageService                  = 0xd34fa24e
	crc_geochats_statedMessage                 = 0x17b1578b
	crc_geochats_located                       = 0x48feb267
	crc_geochats_messages                      = 0xd1526db1
	crc_geochats_messagesSlice                 = 0xbc5863e8
	crc_messageActionGeoChatCreate             = 0x6f038ebc
	crc_messageActionGeoChatCheckin            = 0x0c7d53de
	crc_updateNewGeoChatMessage                = 0x5a68e3f7
	crc_wallPaperSolid                         = 0x63117f24
	crc_updateNewEncryptedMessage              = 0x12bcbd9a
	crc_updateEncryptedChatTyping              = 0x1710f156
	crc_updateEncryption                       = 0xb4a2e88d
	crc_updateEncryptedMessagesRead            = 0x38fe25b7
	crc_encryptedChatEmpty                     = 0xab7ec0a0
	crc_encryptedChatWaiting                   = 0x3bf703dc
	crc_encryptedChatRequested                 = 0xc878527e
	crc_encryptedChat                          = 0xfa56ce36
	crc_encryptedChatDiscarded                 = 0x13d6dd27
	crc_inputEncryptedChat                     = 0xf141b5e1
	crc_encryptedFileEmpty                     = 0xc21f497e
	crc_encryptedFile                          = 0x4a70994c
	crc_inputEncryptedFileEmpty                = 0x1837c364
	crc_inputEncryptedFileUploaded             = 0x64bd0306
	crc_inputEncryptedFile                     = 0x5a17b5e5
	crc_inputEncryptedFileLocation             = 0xf5235d55
	crc_encryptedMessage                       = 0xed18c118
	crc_encryptedMessageService                = 0x23734b06
	crc_messages_dhConfigNotModified           = 0xc0e24635
	crc_messages_dhConfig                      = 0x2c221edd
	crc_messages_sentEncryptedMessage          = 0x560f8935
	crc_messages_sentEncryptedFile             = 0x9493ff32
	crc_inputFileBig                           = 0xfa4f0bb5
	crc_inputEncryptedFileBigUploaded          = 0x2dc173c8
	crc_updateChatParticipantAdd               = 0x3a0eeb22
	crc_updateChatParticipantDelete            = 0x6e5f8c22
	crc_updateDcOptions                        = 0x8e5e9873
	crc_inputMediaUploadedAudio                = 0x4e498cab
	crc_inputMediaAudio                        = 0x89938781
	crc_inputMediaUploadedDocument             = 0xffe76b78
	crc_inputMediaUploadedThumbDocument        = 0x41481486
	crc_inputMediaDocument                     = 0xd184e841
	crc_messageMediaDocument                   = 0x2fda2204
	crc_messageMediaAudio                      = 0xc6b68300
	crc_inputAudioEmpty                        = 0xd95adc84
	crc_inputAudio                             = 0x77d440ff
	crc_inputDocumentEmpty                     = 0x72f0eaae
	crc_inputDocument                          = 0x18798952
	crc_inputAudioFileLocation                 = 0x74dc404d
	crc_inputDocumentFileLocation              = 0x4e45abe9
	crc_audioEmpty                             = 0x586988d8
	crc_audio                                  = 0xc7ac6496
	crc_documentEmpty                          = 0x36f8c871
	crc_document                               = 0xf9a39f4f
	crc_help_support                           = 0x17c6b5f6
	crc_notifyPeer                             = 0x9fd40bd8
	crc_notifyUsers                            = 0xb4c83b4c
	crc_notifyChats                            = 0xc007cec3
	crc_notifyAll                              = 0x74d07c60
	crc_updateUserBlocked                      = 0x80ece81a
	crc_updateNotifySettings                   = 0xbec268ef
	crc_auth_sentAppCode                       = 0xe325edcf
	crc_sendMessageTypingAction                = 0x16bf744e
	crc_sendMessageCancelAction                = 0xfd5ec8f5
	crc_sendMessageRecordVideoAction           = 0xa187d66f
	crc_sendMessageUploadVideoAction           = 0x92042ff7
	crc_sendMessageRecordAudioAction           = 0xd52f73f7
	crc_sendMessageUploadAudioAction           = 0xe6ac8a6f
	crc_sendMessageUploadPhotoAction           = 0x990a3c1a
	crc_sendMessageUploadDocumentAction        = 0x8faee98e
	crc_sendMessageGeoLocationAction           = 0x176f8ba1
	crc_sendMessageChooseContactAction         = 0x628cbc6f
	crc_contactFound                           = 0xea879f95
	crc_contacts_found                         = 0x0566000e
	crc_updateServiceNotification              = 0x382dd3e4
	crc_userStatusRecently                     = 0xe26f42f1
	crc_userStatusLastWeek                     = 0x07bf09fc
	crc_userStatusLastMonth                    = 0x77ebc742
	crc_updatePrivacy                          = 0xee3b272a
	crc_inputPrivacyKeyStatusTimestamp         = 0x4f96cb18
	crc_privacyKeyStatusTimestamp              = 0xbc2eab30
	crc_inputPrivacyValueAllowContacts         = 0x0d09e07b
	crc_inputPrivacyValueAllowAll              = 0x184b35ce
	crc_inputPrivacyValueAllowUsers            = 0x131cc67f
	crc_inputPrivacyValueDisallowContacts      = 0x0ba52007
	crc_inputPrivacyValueDisallowAll           = 0xd66b66c9
	crc_inputPrivacyValueDisallowUsers         = 0x90110467
	crc_privacyValueAllowContacts              = 0xfffe1bac
	crc_privacyValueAllowAll                   = 0x65427b82
	crc_privacyValueAllowUsers                 = 0x4d5bbe0c
	crc_privacyValueDisallowContacts           = 0xf888fa1a
	crc_privacyValueDisallowAll                = 0x8b73e763
	crc_privacyValueDisallowUsers              = 0x0c7f49b7
	crc_account_privacyRules                   = 0x554abb6f
	crc_accountDaysTTL                         = 0xb8d0afdf
	crc_account_sentChangePhoneCode            = 0xa4f58c4c
	crc_updateUserPhone                        = 0x12b9417b
	crc_documentAttributeImageSize             = 0x6c37c15c
	crc_documentAttributeAnimated              = 0x11b58939
	crc_documentAttributeSticker               = 0xfb0a5727
	crc_documentAttributeVideo                 = 0x5910cccb
	crc_documentAttributeAudio                 = 0x051448e5
	crc_documentAttributeFilename              = 0x15590068
	crc_messages_stickersNotModified           = 0xf1749a22
	crc_messages_stickers                      = 0x8a8ecd32
	crc_stickerPack                            = 0x12b299d4
	crc_messages_allStickersNotModified        = 0xe86602c3
	crc_messages_allStickers                   = 0xdcef3102
	crc_disabledFeature                        = 0xae636f24
	crc_invokeAfterMsg                         = 0xcb9f372d
	crc_invokeAfterMsgs                        = 0x3dc4b4f0
	crc_auth_checkPhone                        = 0x6fe51dfb
	crc_auth_sendCode                          = 0x768d5f4d
	crc_auth_sendCall                          = 0x03c51564
	crc_auth_signUp                            = 0x1b067634
	crc_auth_signIn                            = 0xbcd51581
	crc_auth_logOut                            = 0x5717da40
	crc_auth_resetAuthorizations               = 0x9fab0d1a
	crc_auth_sendInvites                       = 0x771c1d97
	crc_auth_exportAuthorization               = 0xe5bfffcd
	crc_auth_importAuthorization               = 0xe3ef9613
	crc_auth_bindTempAuthKey                   = 0xcdd42a05
	crc_account_registerDevice                 = 0x446c712c
	crc_account_unregisterDevice               = 0x65c55b40
	crc_account_updateNotifySettings           = 0x84be5b93
	crc_account_getNotifySettings              = 0x12b3ad31
	crc_account_resetNotifySettings            = 0xdb7e1747
	crc_account_updateProfile                  = 0xf0888d68
	crc_account_updateStatus                   = 0x6628562c
	crc_account_getWallPapers                  = 0xc04cfac2
	crc_users_getUsers                         = 0x0d91a548
	crc_users_getFullUser                      = 0xca30a5b1
	crc_contacts_getStatuses                   = 0xc4a353ee
	crc_contacts_getContacts                   = 0x22c6aa08
	crc_contacts_importContacts                = 0xda30b32d
	crc_contacts_getSuggested                  = 0xcd773428
	crc_contacts_deleteContact                 = 0x8e953744
	crc_contacts_deleteContacts                = 0x59ab389e
	crc_contacts_block                         = 0x332b49fc
	crc_contacts_unblock                       = 0xe54100bd
	crc_contacts_getBlocked                    = 0xf57c350f
	crc_contacts_exportCard                    = 0x84e53737
	crc_contacts_importCard                    = 0x4fe196fe
	crc_messages_getMessages                   = 0x4222fa74
	crc_messages_getDialogs                    = 0xeccf1df6
	crc_messages_getHistory                    = 0x92a1df2f
	crc_messages_search                        = 0x07e9f2ab
	crc_messages_readHistory                   = 0xeed884c6
	crc_messages_deleteHistory                 = 0xf4f8fb61
	crc_messages_deleteMessages                = 0x14f2dd0a
	crc_messages_receivedMessages              = 0x28abcb68
	crc_messages_setTyping                     = 0xa3825e50
	crc_messages_sendMessage                   = 0x4cde0aab
	crc_messages_sendMedia                     = 0xa3c85d76
	crc_messages_forwardMessages               = 0x514cd10f
	crc_messages_getChats                      = 0x3c6aa187
	crc_messages_getFullChat                   = 0x3b831c66
	crc_messages_editChatTitle                 = 0xb4bc68b5
	crc_messages_editChatPhoto                 = 0xd881821d
	crc_messages_addChatUser                   = 0x2ee9ee9e
	crc_messages_deleteChatUser                = 0xc3c5cd23
	crc_messages_createChat                    = 0x419d9aee
	crc_updates_getState                       = 0xedd4882a
	crc_updates_getDifference                  = 0x0a041495
	crc_photos_updateProfilePhoto              = 0xeef579a0
	crc_photos_uploadProfilePhoto              = 0xd50f9c88
	crc_photos_deletePhotos                    = 0x87cf7f2f
	crc_upload_saveFilePart                    = 0xb304a621
	crc_upload_getFile                         = 0xe3a6cfb5
	crc_help_getConfig                         = 0xc4f9186b
	crc_help_getNearestDc                      = 0x1fb33026
	crc_help_getAppUpdate                      = 0xc812ac7e
	crc_help_saveAppLog                        = 0x6f02f748
	crc_help_getInviteText                     = 0xa4a95186
	crc_photos_getUserPhotos                   = 0xb7ee553c
	crc_messages_forwardMessage                = 0x03f3f4f2
	crc_messages_sendBroadcast                 = 0x41bb0972
	crc_geochats_getLocated                    = 0x7f192d8f
	crc_geochats_getRecents                    = 0xe1427e6f
	crc_geochats_checkin                       = 0x55b3e8fb
	crc_geochats_getFullChat                   = 0x6722dd6f
	crc_geochats_editChatTitle                 = 0x4c8e2273
	crc_geochats_editChatPhoto                 = 0x35d81a95
	crc_geochats_search                        = 0xcfcdc44d
	crc_geochats_getHistory                    = 0xb53f7a68
	crc_geochats_setTyping                     = 0x08b8a729
	crc_geochats_sendMessage                   = 0x061b0044
	crc_geochats_sendMedia                     = 0xb8f0deff
	crc_geochats_createGeoChat                 = 0x0e092e16
	crc_messages_getDhConfig                   = 0x26cf8950
	crc_messages_requestEncryption             = 0xf64daf43
	crc_messages_acceptEncryption              = 0x3dbc0415
	crc_messages_discardEncryption             = 0xedd923c5
	crc_messages_setEncryptedTyping            = 0x791451ed
	crc_messages_readEncryptedHistory          = 0x7f4b690a
	crc_messages_sendEncrypted                 = 0xa9776773
	crc_messages_sendEncryptedFile             = 0x9a901b66
	crc_messages_sendEncryptedService          = 0x32d439a4
	crc_messages_receivedQueue                 = 0x55a5bb66
	crc_upload_saveBigFilePart                 = 0xde7b673d
	crc_initConnection                         = 0x69796de9
	crc_help_getSupport                        = 0x9cdf08cd
	crc_auth_sendSms                           = 0x0da9f3e8
	crc_messages_readMessageContents           = 0x354b5bc2
	crc_account_checkUsername                  = 0x2714d86c
	crc_account_updateUsername                 = 0x3e0bdd7c
	crc_contacts_search                        = 0x11f812d8
	crc_account_getPrivacy                     = 0xdadbc950
	crc_account_setPrivacy                     = 0xc9f81ce8
	crc_account_deleteAccount                  = 0x418d4e0b
	crc_account_getAccountTTL                  = 0x08fc711d
	crc_account_setAccountTTL                  = 0x2442485e
	crc_invokeWithLayer                        = 0xda9b0d0d
	crc_contacts_resolveUsername               = 0x0bf0131c
	crc_account_sendChangePhoneCode            = 0xa407a8f4
	crc_account_changePhone                    = 0x70c32edb
	crc_messages_getStickers                   = 0xae22e045
	crc_messages_getAllStickers                = 0xaa3bc868
	crc_account_updateDeviceLocked             = 0x38df3532
)

type TL_boolFalse struct {
}

type TL_boolTrue struct {
}

type TL_error struct {
	code int32
	text string
}

type TL_null struct {
}

type TL_inputPeerEmpty struct {
}

type TL_inputPeerSelf struct {
}

type TL_inputPeerContact struct {
	user_id int32
}

type TL_inputPeerForeign struct {
	user_id     int32
	access_hash int64
}

type TL_inputPeerChat struct {
	chat_id int32
}

type TL_inputUserEmpty struct {
}

type TL_inputUserSelf struct {
}

type TL_inputUserContact struct {
	user_id int32
}

type TL_inputUserForeign struct {
	user_id     int32
	access_hash int64
}

type TL_inputPhoneContact struct {
	client_id  int64
	phone      string
	first_name string
	last_name  string
}

type TL_inputFile struct {
	id           int64
	parts        int32
	name         string
	md5_checksum string
}

type TL_inputMediaEmpty struct {
}

type TL_inputMediaUploadedPhoto struct {
	file TL // InputFile
}

type TL_inputMediaPhoto struct {
	id TL // InputPhoto
}

type TL_inputMediaGeoPoint struct {
	geo_point TL // InputGeoPoint
}

type TL_inputMediaContact struct {
	phone_number string
	first_name   string
	last_name    string
}

type TL_inputMediaUploadedVideo struct {
	file      TL // InputFile
	duration  int32
	w         int32
	h         int32
	mime_type string
}

type TL_inputMediaUploadedThumbVideo struct {
	file      TL // InputFile
	thumb     TL // InputFile
	duration  int32
	w         int32
	h         int32
	mime_type string
}

type TL_inputMediaVideo struct {
	id TL // InputVideo
}

type TL_inputChatPhotoEmpty struct {
}

type TL_inputChatUploadedPhoto struct {
	file TL // InputFile
	crop TL // InputPhotoCrop
}

type TL_inputChatPhoto struct {
	id   TL // InputPhoto
	crop TL // InputPhotoCrop
}

type TL_inputGeoPointEmpty struct {
}

type TL_inputGeoPoint struct {
	lat  float64
	long float64
}

type TL_inputPhotoEmpty struct {
}

type TL_inputPhoto struct {
	id          int64
	access_hash int64
}

type TL_inputVideoEmpty struct {
}

type TL_inputVideo struct {
	id          int64
	access_hash int64
}

type TL_inputFileLocation struct {
	volume_id int64
	local_id  int32
	secret    int64
}

type TL_inputVideoFileLocation struct {
	id          int64
	access_hash int64
}

type TL_inputPhotoCropAuto struct {
}

type TL_inputPhotoCrop struct {
	crop_left  float64
	crop_top   float64
	crop_width float64
}

type TL_inputAppEvent struct {
	time  float64
	_type string
	peer  int64
	data  string
}

type TL_peerUser struct {
	user_id int32
}

type TL_peerChat struct {
	chat_id int32
}

type TL_storage_fileUnknown struct {
}

type TL_storage_fileJpeg struct {
}

type TL_storage_fileGif struct {
}

type TL_storage_filePng struct {
}

type TL_storage_filePdf struct {
}

type TL_storage_fileMp3 struct {
}

type TL_storage_fileMov struct {
}

type TL_storage_filePartial struct {
}

type TL_storage_fileMp4 struct {
}

type TL_storage_fileWebp struct {
}

type TL_fileLocationUnavailable struct {
	volume_id int64
	local_id  int32
	secret    int64
}

type TL_fileLocation struct {
	dc_id     int32
	volume_id int64
	local_id  int32
	secret    int64
}

type TL_userEmpty struct {
	id int32
}

type TL_userSelf struct {
	id         int32
	first_name string
	last_name  string
	username   string
	phone      string
	photo      TL // UserProfilePhoto
	status     TL // UserStatus
	inactive   TL // Bool
}

type TL_userContact struct {
	id          int32
	first_name  string
	last_name   string
	username    string
	access_hash int64
	phone       string
	photo       TL // UserProfilePhoto
	status      TL // UserStatus
}

type TL_userRequest struct {
	id          int32
	first_name  string
	last_name   string
	username    string
	access_hash int64
	phone       string
	photo       TL // UserProfilePhoto
	status      TL // UserStatus
}

type TL_userForeign struct {
	id          int32
	first_name  string
	last_name   string
	username    string
	access_hash int64
	photo       TL // UserProfilePhoto
	status      TL // UserStatus
}

type TL_userDeleted struct {
	id         int32
	first_name string
	last_name  string
	username   string
}

type TL_userProfilePhotoEmpty struct {
}

type TL_userProfilePhoto struct {
	photo_id    int64
	photo_small TL // FileLocation
	photo_big   TL // FileLocation
}

type TL_userStatusEmpty struct {
}

type TL_userStatusOnline struct {
	expires int32
}

type TL_userStatusOffline struct {
	was_online int32
}

type TL_chatEmpty struct {
	id int32
}

type TL_chat struct {
	id                 int32
	title              string
	photo              TL // ChatPhoto
	participants_count int32
	date               int32
	left               TL // Bool
	version            int32
}

type TL_chatForbidden struct {
	id    int32
	title string
	date  int32
}

type TL_chatFull struct {
	id              int32
	participants    TL // ChatParticipants
	chat_photo      TL // Photo
	notify_settings TL // PeerNotifySettings
}

type TL_chatParticipant struct {
	user_id    int32
	inviter_id int32
	date       int32
}

type TL_chatParticipantsForbidden struct {
	chat_id int32
}

type TL_chatParticipants struct {
	chat_id      int32
	admin_id     int32
	participants []TL // ChatParticipant
	version      int32
}

type TL_chatPhotoEmpty struct {
}

type TL_chatPhoto struct {
	photo_small TL // FileLocation
	photo_big   TL // FileLocation
}

type TL_messageEmpty struct {
	id int32
}

type TL_message struct {
	flags   int32
	id      int32
	from_id int32
	to_id   TL // Peer
	date    int32
	message string
	media   TL // MessageMedia
}

type TL_messageForwarded struct {
	flags       int32
	id          int32
	fwd_from_id int32
	fwd_date    int32
	from_id     int32
	to_id       TL // Peer
	date        int32
	message     string
	media       TL // MessageMedia
}

type TL_messageService struct {
	flags   int32
	id      int32
	from_id int32
	to_id   TL // Peer
	date    int32
	action  TL // MessageAction
}

type TL_messageMediaEmpty struct {
}

type TL_messageMediaPhoto struct {
	photo TL // Photo
}

type TL_messageMediaVideo struct {
	video TL // Video
}

type TL_messageMediaGeo struct {
	geo TL // GeoPoint
}

type TL_messageMediaContact struct {
	phone_number string
	first_name   string
	last_name    string
	user_id      int32
}

type TL_messageMediaUnsupported struct {
	bytes []byte
}

type TL_messageActionEmpty struct {
}

type TL_messageActionChatCreate struct {
	title string
	users []int32
}

type TL_messageActionChatEditTitle struct {
	title string
}

type TL_messageActionChatEditPhoto struct {
	photo TL // Photo
}

type TL_messageActionChatDeletePhoto struct {
}

type TL_messageActionChatAddUser struct {
	user_id int32
}

type TL_messageActionChatDeleteUser struct {
	user_id int32
}

type TL_dialog struct {
	peer            TL // Peer
	top_message     int32
	unread_count    int32
	notify_settings TL // PeerNotifySettings
}

type TL_photoEmpty struct {
	id int64
}

type TL_photo struct {
	id          int64
	access_hash int64
	user_id     int32
	date        int32
	caption     string
	geo         TL   // GeoPoint
	sizes       []TL // PhotoSize
}

type TL_photoSizeEmpty struct {
	_type string
}

type TL_photoSize struct {
	_type    string
	location TL // FileLocation
	w        int32
	h        int32
	size     int32
}

type TL_photoCachedSize struct {
	_type    string
	location TL // FileLocation
	w        int32
	h        int32
	bytes    []byte
}

type TL_videoEmpty struct {
	id int64
}

type TL_video struct {
	id          int64
	access_hash int64
	user_id     int32
	date        int32
	caption     string
	duration    int32
	mime_type   string
	size        int32
	thumb       TL // PhotoSize
	dc_id       int32
	w           int32
	h           int32
}

type TL_geoPointEmpty struct {
}

type TL_geoPoint struct {
	long float64
	lat  float64
}

type TL_auth_checkedPhone struct {
	phone_registered TL // Bool
	phone_invited    TL // Bool
}

type TL_auth_sentCode struct {
	phone_registered  TL // Bool
	phone_code_hash   string
	send_call_timeout int32
	is_password       TL // Bool
}

type TL_auth_authorization struct {
	expires int32
	user    TL // User
}

type TL_auth_exportedAuthorization struct {
	id    int32
	bytes []byte
}

type TL_inputNotifyPeer struct {
	peer TL // InputPeer
}

type TL_inputNotifyUsers struct {
}

type TL_inputNotifyChats struct {
}

type TL_inputNotifyAll struct {
}

type TL_inputPeerNotifyEventsEmpty struct {
}

type TL_inputPeerNotifyEventsAll struct {
}

type TL_inputPeerNotifySettings struct {
	mute_until    int32
	sound         string
	show_previews TL // Bool
	events_mask   int32
}

type TL_peerNotifyEventsEmpty struct {
}

type TL_peerNotifyEventsAll struct {
}

type TL_peerNotifySettingsEmpty struct {
}

type TL_peerNotifySettings struct {
	mute_until    int32
	sound         string
	show_previews TL // Bool
	events_mask   int32
}

type TL_wallPaper struct {
	id    int32
	title string
	sizes []TL // PhotoSize
	color int32
}

type TL_userFull struct {
	user            TL // User
	link            TL // contacts_Link
	profile_photo   TL // Photo
	notify_settings TL // PeerNotifySettings
	blocked         TL // Bool
	real_first_name string
	real_last_name  string
}

type TL_contact struct {
	user_id int32
	mutual  TL // Bool
}

type TL_importedContact struct {
	user_id   int32
	client_id int64
}

type TL_contactBlocked struct {
	user_id int32
	date    int32
}

type TL_contactSuggested struct {
	user_id         int32
	mutual_contacts int32
}

type TL_contactStatus struct {
	user_id int32
	status  TL // UserStatus
}

type TL_chatLocated struct {
	chat_id  int32
	distance int32
}

type TL_contacts_foreignLinkUnknown struct {
}

type TL_contacts_foreignLinkRequested struct {
	has_phone TL // Bool
}

type TL_contacts_foreignLinkMutual struct {
}

type TL_contacts_myLinkEmpty struct {
}

type TL_contacts_myLinkRequested struct {
	contact TL // Bool
}

type TL_contacts_myLinkContact struct {
}

type TL_contacts_link struct {
	my_link      TL // contacts_MyLink
	foreign_link TL // contacts_ForeignLink
	user         TL // User
}

type TL_contacts_contactsNotModified struct {
}

type TL_contacts_contacts struct {
	contacts []TL // Contact
	users    []TL // User
}

type TL_contacts_importedContacts struct {
	imported       []TL // ImportedContact
	retry_contacts []int64
	users          []TL // User
}

type TL_contacts_blocked struct {
	blocked []TL // ContactBlocked
	users   []TL // User
}

type TL_contacts_blockedSlice struct {
	count   int32
	blocked []TL // ContactBlocked
	users   []TL // User
}

type TL_contacts_suggested struct {
	results []TL // ContactSuggested
	users   []TL // User
}

type TL_messages_dialogs struct {
	dialogs  []TL // Dialog
	messages []TL // Message
	chats    []TL // Chat
	users    []TL // User
}

type TL_messages_dialogsSlice struct {
	count    int32
	dialogs  []TL // Dialog
	messages []TL // Message
	chats    []TL // Chat
	users    []TL // User
}

type TL_messages_messages struct {
	messages []TL // Message
	chats    []TL // Chat
	users    []TL // User
}

type TL_messages_messagesSlice struct {
	count    int32
	messages []TL // Message
	chats    []TL // Chat
	users    []TL // User
}

type TL_messages_messageEmpty struct {
}

type TL_messages_statedMessages struct {
	messages []TL // Message
	chats    []TL // Chat
	users    []TL // User
	pts      int32
	seq      int32
}

type TL_messages_statedMessage struct {
	message TL   // Message
	chats   []TL // Chat
	users   []TL // User
	pts     int32
	seq     int32
}

type TL_messages_sentMessage struct {
	id   int32
	date int32
	pts  int32
	seq  int32
}

type TL_messages_chats struct {
	chats []TL // Chat
	users []TL // User
}

type TL_messages_chatFull struct {
	full_chat TL   // ChatFull
	chats     []TL // Chat
	users     []TL // User
}

type TL_messages_affectedHistory struct {
	pts    int32
	seq    int32
	offset int32
}

type TL_inputMessagesFilterEmpty struct {
}

type TL_inputMessagesFilterPhotos struct {
}

type TL_inputMessagesFilterVideo struct {
}

type TL_inputMessagesFilterPhotoVideo struct {
}

type TL_inputMessagesFilterPhotoVideoDocuments struct {
}

type TL_inputMessagesFilterDocument struct {
}

type TL_inputMessagesFilterAudio struct {
}

type TL_updateNewMessage struct {
	message TL // Message
	pts     int32
}

type TL_updateMessageID struct {
	id        int32
	random_id int64
}

type TL_updateReadMessages struct {
	messages []int32
	pts      int32
}

type TL_updateDeleteMessages struct {
	messages []int32
	pts      int32
}

type TL_updateUserTyping struct {
	user_id int32
	action  TL // SendMessageAction
}

type TL_updateChatUserTyping struct {
	chat_id int32
	user_id int32
	action  TL // SendMessageAction
}

type TL_updateChatParticipants struct {
	participants TL // ChatParticipants
}

type TL_updateUserStatus struct {
	user_id int32
	status  TL // UserStatus
}

type TL_updateUserName struct {
	user_id    int32
	first_name string
	last_name  string
	username   string
}

type TL_updateUserPhoto struct {
	user_id  int32
	date     int32
	photo    TL // UserProfilePhoto
	previous TL // Bool
}

type TL_updateContactRegistered struct {
	user_id int32
	date    int32
}

type TL_updateContactLink struct {
	user_id      int32
	my_link      TL // contacts_MyLink
	foreign_link TL // contacts_ForeignLink
}

type TL_updateNewAuthorization struct {
	auth_key_id int64
	date        int32
	device      string
	location    string
}

type TL_updates_state struct {
	pts          int32
	qts          int32
	date         int32
	seq          int32
	unread_count int32
}

type TL_updates_differenceEmpty struct {
	date int32
	seq  int32
}

type TL_updates_difference struct {
	new_messages           []TL // Message
	new_encrypted_messages []TL // EncryptedMessage
	other_updates          []TL // Update
	chats                  []TL // Chat
	users                  []TL // User
	state                  TL   // updates_State
}

type TL_updates_differenceSlice struct {
	new_messages           []TL // Message
	new_encrypted_messages []TL // EncryptedMessage
	other_updates          []TL // Update
	chats                  []TL // Chat
	users                  []TL // User
	intermediate_state     TL   // updates_State
}

type TL_updatesTooLong struct {
}

type TL_updateShortMessage struct {
	id      int32
	from_id int32
	message string
	pts     int32
	date    int32
	seq     int32
}

type TL_updateShortChatMessage struct {
	id      int32
	from_id int32
	chat_id int32
	message string
	pts     int32
	date    int32
	seq     int32
}

type TL_updateShort struct {
	update TL // Update
	date   int32
}

type TL_updatesCombined struct {
	updates   []TL // Update
	users     []TL // User
	chats     []TL // Chat
	date      int32
	seq_start int32
	seq       int32
}

type TL_updates struct {
	updates []TL // Update
	users   []TL // User
	chats   []TL // Chat
	date    int32
	seq     int32
}

type TL_photos_photos struct {
	photos []TL // Photo
	users  []TL // User
}

type TL_photos_photosSlice struct {
	count  int32
	photos []TL // Photo
	users  []TL // User
}

type TL_photos_photo struct {
	photo TL   // Photo
	users []TL // User
}

type TL_upload_file struct {
	_type TL // storage_FileType
	mtime int32
	bytes []byte
}

type TL_dcOption struct {
	id         int32
	hostname   string
	ip_address string
	port       int32
}

type TL_config struct {
	date               int32
	expires            int32
	test_mode          TL // Bool
	this_dc            int32
	dc_options         []TL // DcOption
	chat_big_size      int32
	chat_size_max      int32
	broadcast_size_max int32
	disabled_features  []TL // DisabledFeature
}

type TL_nearestDc struct {
	country    string
	this_dc    int32
	nearest_dc int32
}

type TL_help_appUpdate struct {
	id       int32
	critical TL // Bool
	url      string
	text     string
}

type TL_help_noAppUpdate struct {
}

type TL_help_inviteText struct {
	message string
}

type TL_messages_statedMessagesLinks struct {
	messages []TL // Message
	chats    []TL // Chat
	users    []TL // User
	links    []TL // contacts_Link
	pts      int32
	seq      int32
}

type TL_messages_statedMessageLink struct {
	message TL   // Message
	chats   []TL // Chat
	users   []TL // User
	links   []TL // contacts_Link
	pts     int32
	seq     int32
}

type TL_messages_sentMessageLink struct {
	id    int32
	date  int32
	pts   int32
	seq   int32
	links []TL // contacts_Link
}

type TL_inputGeoChat struct {
	chat_id     int32
	access_hash int64
}

type TL_inputNotifyGeoChatPeer struct {
	peer TL // InputGeoChat
}

type TL_geoChat struct {
	id                 int32
	access_hash        int64
	title              string
	address            string
	venue              string
	geo                TL // GeoPoint
	photo              TL // ChatPhoto
	participants_count int32
	date               int32
	checked_in         TL // Bool
	version            int32
}

type TL_geoChatMessageEmpty struct {
	chat_id int32
	id      int32
}

type TL_geoChatMessage struct {
	chat_id int32
	id      int32
	from_id int32
	date    int32
	message string
	media   TL // MessageMedia
}

type TL_geoChatMessageService struct {
	chat_id int32
	id      int32
	from_id int32
	date    int32
	action  TL // MessageAction
}

type TL_geochats_statedMessage struct {
	message TL   // GeoChatMessage
	chats   []TL // Chat
	users   []TL // User
	seq     int32
}

type TL_geochats_located struct {
	results  []TL // ChatLocated
	messages []TL // GeoChatMessage
	chats    []TL // Chat
	users    []TL // User
}

type TL_geochats_messages struct {
	messages []TL // GeoChatMessage
	chats    []TL // Chat
	users    []TL // User
}

type TL_geochats_messagesSlice struct {
	count    int32
	messages []TL // GeoChatMessage
	chats    []TL // Chat
	users    []TL // User
}

type TL_messageActionGeoChatCreate struct {
	title   string
	address string
}

type TL_messageActionGeoChatCheckin struct {
}

type TL_updateNewGeoChatMessage struct {
	message TL // GeoChatMessage
}

type TL_wallPaperSolid struct {
	id       int32
	title    string
	bg_color int32
	color    int32
}

type TL_updateNewEncryptedMessage struct {
	message TL // EncryptedMessage
	qts     int32
}

type TL_updateEncryptedChatTyping struct {
	chat_id int32
}

type TL_updateEncryption struct {
	chat TL // EncryptedChat
	date int32
}

type TL_updateEncryptedMessagesRead struct {
	chat_id  int32
	max_date int32
	date     int32
}

type TL_encryptedChatEmpty struct {
	id int32
}

type TL_encryptedChatWaiting struct {
	id             int32
	access_hash    int64
	date           int32
	admin_id       int32
	participant_id int32
}

type TL_encryptedChatRequested struct {
	id             int32
	access_hash    int64
	date           int32
	admin_id       int32
	participant_id int32
	g_a            []byte
}

type TL_encryptedChat struct {
	id              int32
	access_hash     int64
	date            int32
	admin_id        int32
	participant_id  int32
	g_a_or_b        []byte
	key_fingerprint int64
}

type TL_encryptedChatDiscarded struct {
	id int32
}

type TL_inputEncryptedChat struct {
	chat_id     int32
	access_hash int64
}

type TL_encryptedFileEmpty struct {
}

type TL_encryptedFile struct {
	id              int64
	access_hash     int64
	size            int32
	dc_id           int32
	key_fingerprint int32
}

type TL_inputEncryptedFileEmpty struct {
}

type TL_inputEncryptedFileUploaded struct {
	id              int64
	parts           int32
	md5_checksum    string
	key_fingerprint int32
}

type TL_inputEncryptedFile struct {
	id          int64
	access_hash int64
}

type TL_inputEncryptedFileLocation struct {
	id          int64
	access_hash int64
}

type TL_encryptedMessage struct {
	random_id int64
	chat_id   int32
	date      int32
	bytes     []byte
	file      TL // EncryptedFile
}

type TL_encryptedMessageService struct {
	random_id int64
	chat_id   int32
	date      int32
	bytes     []byte
}

type TL_messages_dhConfigNotModified struct {
	random []byte
}

type TL_messages_dhConfig struct {
	g       int32
	p       []byte
	version int32
	random  []byte
}

type TL_messages_sentEncryptedMessage struct {
	date int32
}

type TL_messages_sentEncryptedFile struct {
	date int32
	file TL // EncryptedFile
}

type TL_inputFileBig struct {
	id    int64
	parts int32
	name  string
}

type TL_inputEncryptedFileBigUploaded struct {
	id              int64
	parts           int32
	key_fingerprint int32
}

type TL_updateChatParticipantAdd struct {
	chat_id    int32
	user_id    int32
	inviter_id int32
	version    int32
}

type TL_updateChatParticipantDelete struct {
	chat_id int32
	user_id int32
	version int32
}

type TL_updateDcOptions struct {
	dc_options []TL // DcOption
}

type TL_inputMediaUploadedAudio struct {
	file      TL // InputFile
	duration  int32
	mime_type string
}

type TL_inputMediaAudio struct {
	id TL // InputAudio
}

type TL_inputMediaUploadedDocument struct {
	file       TL // InputFile
	mime_type  string
	attributes []TL // DocumentAttribute
}

type TL_inputMediaUploadedThumbDocument struct {
	file       TL // InputFile
	thumb      TL // InputFile
	mime_type  string
	attributes []TL // DocumentAttribute
}

type TL_inputMediaDocument struct {
	id TL // InputDocument
}

type TL_messageMediaDocument struct {
	document TL // Document
}

type TL_messageMediaAudio struct {
	audio TL // Audio
}

type TL_inputAudioEmpty struct {
}

type TL_inputAudio struct {
	id          int64
	access_hash int64
}

type TL_inputDocumentEmpty struct {
}

type TL_inputDocument struct {
	id          int64
	access_hash int64
}

type TL_inputAudioFileLocation struct {
	id          int64
	access_hash int64
}

type TL_inputDocumentFileLocation struct {
	id          int64
	access_hash int64
}

type TL_audioEmpty struct {
	id int64
}

type TL_audio struct {
	id          int64
	access_hash int64
	user_id     int32
	date        int32
	duration    int32
	mime_type   string
	size        int32
	dc_id       int32
}

type TL_documentEmpty struct {
	id int64
}

type TL_document struct {
	id          int64
	access_hash int64
	date        int32
	mime_type   string
	size        int32
	thumb       TL // PhotoSize
	dc_id       int32
	attributes  []TL // DocumentAttribute
}

type TL_help_support struct {
	phone_number string
	user         TL // User
}

type TL_notifyPeer struct {
	peer TL // Peer
}

type TL_notifyUsers struct {
}

type TL_notifyChats struct {
}

type TL_notifyAll struct {
}

type TL_updateUserBlocked struct {
	user_id int32
	blocked TL // Bool
}

type TL_updateNotifySettings struct {
	peer            TL // NotifyPeer
	notify_settings TL // PeerNotifySettings
}

type TL_auth_sentAppCode struct {
	phone_registered  TL // Bool
	phone_code_hash   string
	send_call_timeout int32
	is_password       TL // Bool
}

type TL_sendMessageTypingAction struct {
}

type TL_sendMessageCancelAction struct {
}

type TL_sendMessageRecordVideoAction struct {
}

type TL_sendMessageUploadVideoAction struct {
}

type TL_sendMessageRecordAudioAction struct {
}

type TL_sendMessageUploadAudioAction struct {
}

type TL_sendMessageUploadPhotoAction struct {
}

type TL_sendMessageUploadDocumentAction struct {
}

type TL_sendMessageGeoLocationAction struct {
}

type TL_sendMessageChooseContactAction struct {
}

type TL_contactFound struct {
	user_id int32
}

type TL_contacts_found struct {
	results []TL // ContactFound
	users   []TL // User
}

type TL_updateServiceNotification struct {
	_type   string
	message string
	media   TL // MessageMedia
	popup   TL // Bool
}

type TL_userStatusRecently struct {
}

type TL_userStatusLastWeek struct {
}

type TL_userStatusLastMonth struct {
}

type TL_updatePrivacy struct {
	key   TL   // PrivacyKey
	rules []TL // PrivacyRule
}

type TL_inputPrivacyKeyStatusTimestamp struct {
}

type TL_privacyKeyStatusTimestamp struct {
}

type TL_inputPrivacyValueAllowContacts struct {
}

type TL_inputPrivacyValueAllowAll struct {
}

type TL_inputPrivacyValueAllowUsers struct {
	users []TL // InputUser
}

type TL_inputPrivacyValueDisallowContacts struct {
}

type TL_inputPrivacyValueDisallowAll struct {
}

type TL_inputPrivacyValueDisallowUsers struct {
	users []TL // InputUser
}

type TL_privacyValueAllowContacts struct {
}

type TL_privacyValueAllowAll struct {
}

type TL_privacyValueAllowUsers struct {
	users []int32
}

type TL_privacyValueDisallowContacts struct {
}

type TL_privacyValueDisallowAll struct {
}

type TL_privacyValueDisallowUsers struct {
	users []int32
}

type TL_account_privacyRules struct {
	rules []TL // PrivacyRule
	users []TL // User
}

type TL_accountDaysTTL struct {
	days int32
}

type TL_account_sentChangePhoneCode struct {
	phone_code_hash   string
	send_call_timeout int32
}

type TL_updateUserPhone struct {
	user_id int32
	phone   string
}

type TL_documentAttributeImageSize struct {
	w int32
	h int32
}

type TL_documentAttributeAnimated struct {
}

type TL_documentAttributeSticker struct {
}

type TL_documentAttributeVideo struct {
	duration int32
	w        int32
	h        int32
}

type TL_documentAttributeAudio struct {
	duration int32
}

type TL_documentAttributeFilename struct {
	file_name string
}

type TL_messages_stickersNotModified struct {
}

type TL_messages_stickers struct {
	hash     string
	stickers []TL // Document
}

type TL_stickerPack struct {
	emoticon  string
	documents []int64
}

type TL_messages_allStickersNotModified struct {
}

type TL_messages_allStickers struct {
	hash      string
	packs     []TL // StickerPack
	documents []TL // Document
}

type TL_disabledFeature struct {
	feature     string
	description string
}

type TL_invokeAfterMsg struct {
	msg_id int64
	query  TL
}

type TL_invokeAfterMsgs struct {
	msg_ids []int64
	query   TL
}

type TL_auth_checkPhone struct {
	phone_number string
}

type TL_auth_sendCode struct {
	phone_number string
	sms_type     int32
	api_id       int32
	api_hash     string
	lang_code    string
}

type TL_auth_sendCall struct {
	phone_number    string
	phone_code_hash string
}

type TL_auth_signUp struct {
	phone_number    string
	phone_code_hash string
	phone_code      string
	first_name      string
	last_name       string
}

type TL_auth_signIn struct {
	phone_number    string
	phone_code_hash string
	phone_code      string
}

type TL_auth_logOut struct {
}

type TL_auth_resetAuthorizations struct {
}

type TL_auth_sendInvites struct {
	phone_numbers []string
	message       string
}

type TL_auth_exportAuthorization struct {
	dc_id int32
}

type TL_auth_importAuthorization struct {
	id    int32
	bytes []byte
}

type TL_auth_bindTempAuthKey struct {
	perm_auth_key_id  int64
	nonce             int64
	expires_at        int32
	encrypted_message []byte
}

type TL_account_registerDevice struct {
	token_type     int32
	token          string
	device_model   string
	system_version string
	app_version    string
	app_sandbox    TL // Bool
	lang_code      string
}

type TL_account_unregisterDevice struct {
	token_type int32
	token      string
}

type TL_account_updateNotifySettings struct {
	peer     TL // InputNotifyPeer
	settings TL // InputPeerNotifySettings
}

type TL_account_getNotifySettings struct {
	peer TL // InputNotifyPeer
}

type TL_account_resetNotifySettings struct {
}

type TL_account_updateProfile struct {
	first_name string
	last_name  string
}

type TL_account_updateStatus struct {
	offline TL // Bool
}

type TL_account_getWallPapers struct {
}

type TL_users_getUsers struct {
	id []TL // InputUser
}

type TL_users_getFullUser struct {
	id TL // InputUser
}

type TL_contacts_getStatuses struct {
}

type TL_contacts_getContacts struct {
	hash string
}

type TL_contacts_importContacts struct {
	contacts []TL // InputContact
	replace  TL   // Bool
}

type TL_contacts_getSuggested struct {
	limit int32
}

type TL_contacts_deleteContact struct {
	id TL // InputUser
}

type TL_contacts_deleteContacts struct {
	id []TL // InputUser
}

type TL_contacts_block struct {
	id TL // InputUser
}

type TL_contacts_unblock struct {
	id TL // InputUser
}

type TL_contacts_getBlocked struct {
	offset int32
	limit  int32
}

type TL_contacts_exportCard struct {
}

type TL_contacts_importCard struct {
	export_card []int32
}

type TL_messages_getMessages struct {
	id []int32
}

type TL_messages_getDialogs struct {
	offset int32
	max_id int32
	limit  int32
}

type TL_messages_getHistory struct {
	peer   TL // InputPeer
	offset int32
	max_id int32
	limit  int32
}

type TL_messages_search struct {
	peer     TL // InputPeer
	q        string
	filter   TL // MessagesFilter
	min_date int32
	max_date int32
	offset   int32
	max_id   int32
	limit    int32
}

type TL_messages_readHistory struct {
	peer          TL // InputPeer
	max_id        int32
	offset        int32
	read_contents TL // Bool
}

type TL_messages_deleteHistory struct {
	peer   TL // InputPeer
	offset int32
}

type TL_messages_deleteMessages struct {
	id []int32
}

type TL_messages_receivedMessages struct {
	max_id int32
}

type TL_messages_setTyping struct {
	peer   TL // InputPeer
	action TL // SendMessageAction
}

type TL_messages_sendMessage struct {
	peer      TL // InputPeer
	message   string
	random_id int64
}

type TL_messages_sendMedia struct {
	peer      TL // InputPeer
	media     TL // InputMedia
	random_id int64
}

type TL_messages_forwardMessages struct {
	peer TL // InputPeer
	id   []int32
}

type TL_messages_getChats struct {
	id []int32
}

type TL_messages_getFullChat struct {
	chat_id int32
}

type TL_messages_editChatTitle struct {
	chat_id int32
	title   string
}

type TL_messages_editChatPhoto struct {
	chat_id int32
	photo   TL // InputChatPhoto
}

type TL_messages_addChatUser struct {
	chat_id   int32
	user_id   TL // InputUser
	fwd_limit int32
}

type TL_messages_deleteChatUser struct {
	chat_id int32
	user_id TL // InputUser
}

type TL_messages_createChat struct {
	users []TL // InputUser
	title string
}

type TL_updates_getState struct {
}

type TL_updates_getDifference struct {
	pts  int32
	date int32
	qts  int32
}

type TL_photos_updateProfilePhoto struct {
	id   TL // InputPhoto
	crop TL // InputPhotoCrop
}

type TL_photos_uploadProfilePhoto struct {
	file      TL // InputFile
	caption   string
	geo_point TL // InputGeoPoint
	crop      TL // InputPhotoCrop
}

type TL_photos_deletePhotos struct {
	id []TL // InputPhoto
}

type TL_upload_saveFilePart struct {
	file_id   int64
	file_part int32
	bytes     []byte
}

type TL_upload_getFile struct {
	location TL // InputFileLocation
	offset   int32
	limit    int32
}

type TL_help_getConfig struct {
}

type TL_help_getNearestDc struct {
}

type TL_help_getAppUpdate struct {
	device_model   string
	system_version string
	app_version    string
	lang_code      string
}

type TL_help_saveAppLog struct {
	events []TL // InputAppEvent
}

type TL_help_getInviteText struct {
	lang_code string
}

type TL_photos_getUserPhotos struct {
	user_id TL // InputUser
	offset  int32
	max_id  int32
	limit   int32
}

type TL_messages_forwardMessage struct {
	peer      TL // InputPeer
	id        int32
	random_id int64
}

type TL_messages_sendBroadcast struct {
	contacts []TL // InputUser
	message  string
	media    TL // InputMedia
}

type TL_geochats_getLocated struct {
	geo_point TL // InputGeoPoint
	radius    int32
	limit     int32
}

type TL_geochats_getRecents struct {
	offset int32
	limit  int32
}

type TL_geochats_checkin struct {
	peer TL // InputGeoChat
}

type TL_geochats_getFullChat struct {
	peer TL // InputGeoChat
}

type TL_geochats_editChatTitle struct {
	peer    TL // InputGeoChat
	title   string
	address string
}

type TL_geochats_editChatPhoto struct {
	peer  TL // InputGeoChat
	photo TL // InputChatPhoto
}

type TL_geochats_search struct {
	peer     TL // InputGeoChat
	q        string
	filter   TL // MessagesFilter
	min_date int32
	max_date int32
	offset   int32
	max_id   int32
	limit    int32
}

type TL_geochats_getHistory struct {
	peer   TL // InputGeoChat
	offset int32
	max_id int32
	limit  int32
}

type TL_geochats_setTyping struct {
	peer   TL // InputGeoChat
	typing TL // Bool
}

type TL_geochats_sendMessage struct {
	peer      TL // InputGeoChat
	message   string
	random_id int64
}

type TL_geochats_sendMedia struct {
	peer      TL // InputGeoChat
	media     TL // InputMedia
	random_id int64
}

type TL_geochats_createGeoChat struct {
	title     string
	geo_point TL // InputGeoPoint
	address   string
	venue     string
}

type TL_messages_getDhConfig struct {
	version       int32
	random_length int32
}

type TL_messages_requestEncryption struct {
	user_id   TL // InputUser
	random_id int32
	g_a       []byte
}

type TL_messages_acceptEncryption struct {
	peer            TL // InputEncryptedChat
	g_b             []byte
	key_fingerprint int64
}

type TL_messages_discardEncryption struct {
	chat_id int32
}

type TL_messages_setEncryptedTyping struct {
	peer   TL // InputEncryptedChat
	typing TL // Bool
}

type TL_messages_readEncryptedHistory struct {
	peer     TL // InputEncryptedChat
	max_date int32
}

type TL_messages_sendEncrypted struct {
	peer      TL // InputEncryptedChat
	random_id int64
	data      []byte
}

type TL_messages_sendEncryptedFile struct {
	peer      TL // InputEncryptedChat
	random_id int64
	data      []byte
	file      TL // InputEncryptedFile
}

type TL_messages_sendEncryptedService struct {
	peer      TL // InputEncryptedChat
	random_id int64
	data      []byte
}

type TL_messages_receivedQueue struct {
	max_qts int32
}

type TL_upload_saveBigFilePart struct {
	file_id          int64
	file_part        int32
	file_total_parts int32
	bytes            []byte
}

type TL_initConnection struct {
	api_id         int32
	device_model   string
	system_version string
	app_version    string
	lang_code      string
	query          TL
}

type TL_help_getSupport struct {
}

type TL_auth_sendSms struct {
	phone_number    string
	phone_code_hash string
}

type TL_messages_readMessageContents struct {
	id []int32
}

type TL_account_checkUsername struct {
	username string
}

type TL_account_updateUsername struct {
	username string
}

type TL_contacts_search struct {
	q     string
	limit int32
}

type TL_account_getPrivacy struct {
	key TL // InputPrivacyKey
}

type TL_account_setPrivacy struct {
	key   TL   // InputPrivacyKey
	rules []TL // InputPrivacyRule
}

type TL_account_deleteAccount struct {
	reason string
}

type TL_account_getAccountTTL struct {
}

type TL_account_setAccountTTL struct {
	ttl TL // AccountDaysTTL
}

type TL_invokeWithLayer struct {
	layer int32
	query TL
}

type TL_contacts_resolveUsername struct {
	username string
}

type TL_account_sendChangePhoneCode struct {
	phone_number string
}

type TL_account_changePhone struct {
	phone_number    string
	phone_code_hash string
	phone_code      string
}

type TL_messages_getStickers struct {
	emoticon string
	hash     string
}

type TL_messages_getAllStickers struct {
	hash string
}

type TL_account_updateDeviceLocked struct {
	period int32
}

func (e TL_boolFalse) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_boolFalse)
	return x.buf
}

func (e TL_boolTrue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_boolTrue)
	return x.buf
}

func (e TL_error) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_error)
	x.Int(e.code)
	x.String(e.text)
	return x.buf
}

func (e TL_null) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_null)
	return x.buf
}

func (e TL_inputPeerEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerEmpty)
	return x.buf
}

func (e TL_inputPeerSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerSelf)
	return x.buf
}

func (e TL_inputPeerContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerContact)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_inputPeerForeign) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerForeign)
	x.Int(e.user_id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputPeerChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerChat)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_inputUserEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUserEmpty)
	return x.buf
}

func (e TL_inputUserSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUserSelf)
	return x.buf
}

func (e TL_inputUserContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUserContact)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_inputUserForeign) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputUserForeign)
	x.Int(e.user_id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputPhoneContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhoneContact)
	x.Long(e.client_id)
	x.String(e.phone)
	x.String(e.first_name)
	x.String(e.last_name)
	return x.buf
}

func (e TL_inputFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputFile)
	x.Long(e.id)
	x.Int(e.parts)
	x.String(e.name)
	x.String(e.md5_checksum)
	return x.buf
}

func (e TL_inputMediaEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaEmpty)
	return x.buf
}

func (e TL_inputMediaUploadedPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaUploadedPhoto)
	x.Bytes(e.file.encode())
	return x.buf
}

func (e TL_inputMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaPhoto)
	x.Bytes(e.id.encode())
	return x.buf
}

func (e TL_inputMediaGeoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaGeoPoint)
	x.Bytes(e.geo_point.encode())
	return x.buf
}

func (e TL_inputMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaContact)
	x.String(e.phone_number)
	x.String(e.first_name)
	x.String(e.last_name)
	return x.buf
}

func (e TL_inputMediaUploadedVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaUploadedVideo)
	x.Bytes(e.file.encode())
	x.Int(e.duration)
	x.Int(e.w)
	x.Int(e.h)
	x.String(e.mime_type)
	return x.buf
}

func (e TL_inputMediaUploadedThumbVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaUploadedThumbVideo)
	x.Bytes(e.file.encode())
	x.Bytes(e.thumb.encode())
	x.Int(e.duration)
	x.Int(e.w)
	x.Int(e.h)
	x.String(e.mime_type)
	return x.buf
}

func (e TL_inputMediaVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaVideo)
	x.Bytes(e.id.encode())
	return x.buf
}

func (e TL_inputChatPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChatPhotoEmpty)
	return x.buf
}

func (e TL_inputChatUploadedPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChatUploadedPhoto)
	x.Bytes(e.file.encode())
	x.Bytes(e.crop.encode())
	return x.buf
}

func (e TL_inputChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputChatPhoto)
	x.Bytes(e.id.encode())
	x.Bytes(e.crop.encode())
	return x.buf
}

func (e TL_inputGeoPointEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGeoPointEmpty)
	return x.buf
}

func (e TL_inputGeoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGeoPoint)
	x.Double(e.lat)
	x.Double(e.long)
	return x.buf
}

func (e TL_inputPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhotoEmpty)
	return x.buf
}

func (e TL_inputPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhoto)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputVideoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputVideoEmpty)
	return x.buf
}

func (e TL_inputVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputVideo)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputFileLocation)
	x.Long(e.volume_id)
	x.Int(e.local_id)
	x.Long(e.secret)
	return x.buf
}

func (e TL_inputVideoFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputVideoFileLocation)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputPhotoCropAuto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhotoCropAuto)
	return x.buf
}

func (e TL_inputPhotoCrop) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPhotoCrop)
	x.Double(e.crop_left)
	x.Double(e.crop_top)
	x.Double(e.crop_width)
	return x.buf
}

func (e TL_inputAppEvent) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputAppEvent)
	x.Double(e.time)
	x.String(e._type)
	x.Long(e.peer)
	x.String(e.data)
	return x.buf
}

func (e TL_peerUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerUser)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_peerChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerChat)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_storage_fileUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileUnknown)
	return x.buf
}

func (e TL_storage_fileJpeg) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileJpeg)
	return x.buf
}

func (e TL_storage_fileGif) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileGif)
	return x.buf
}

func (e TL_storage_filePng) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_filePng)
	return x.buf
}

func (e TL_storage_filePdf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_filePdf)
	return x.buf
}

func (e TL_storage_fileMp3) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileMp3)
	return x.buf
}

func (e TL_storage_fileMov) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileMov)
	return x.buf
}

func (e TL_storage_filePartial) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_filePartial)
	return x.buf
}

func (e TL_storage_fileMp4) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileMp4)
	return x.buf
}

func (e TL_storage_fileWebp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_storage_fileWebp)
	return x.buf
}

func (e TL_fileLocationUnavailable) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_fileLocationUnavailable)
	x.Long(e.volume_id)
	x.Int(e.local_id)
	x.Long(e.secret)
	return x.buf
}

func (e TL_fileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_fileLocation)
	x.Int(e.dc_id)
	x.Long(e.volume_id)
	x.Int(e.local_id)
	x.Long(e.secret)
	return x.buf
}

func (e TL_userEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userEmpty)
	x.Int(e.id)
	return x.buf
}

func (e TL_userSelf) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userSelf)
	x.Int(e.id)
	x.String(e.first_name)
	x.String(e.last_name)
	x.String(e.username)
	x.String(e.phone)
	x.Bytes(e.photo.encode())
	x.Bytes(e.status.encode())
	x.Bytes(e.inactive.encode())
	return x.buf
}

func (e TL_userContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userContact)
	x.Int(e.id)
	x.String(e.first_name)
	x.String(e.last_name)
	x.String(e.username)
	x.Long(e.access_hash)
	x.String(e.phone)
	x.Bytes(e.photo.encode())
	x.Bytes(e.status.encode())
	return x.buf
}

func (e TL_userRequest) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userRequest)
	x.Int(e.id)
	x.String(e.first_name)
	x.String(e.last_name)
	x.String(e.username)
	x.Long(e.access_hash)
	x.String(e.phone)
	x.Bytes(e.photo.encode())
	x.Bytes(e.status.encode())
	return x.buf
}

func (e TL_userForeign) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userForeign)
	x.Int(e.id)
	x.String(e.first_name)
	x.String(e.last_name)
	x.String(e.username)
	x.Long(e.access_hash)
	x.Bytes(e.photo.encode())
	x.Bytes(e.status.encode())
	return x.buf
}

func (e TL_userDeleted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userDeleted)
	x.Int(e.id)
	x.String(e.first_name)
	x.String(e.last_name)
	x.String(e.username)
	return x.buf
}

func (e TL_userProfilePhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userProfilePhotoEmpty)
	return x.buf
}

func (e TL_userProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userProfilePhoto)
	x.Long(e.photo_id)
	x.Bytes(e.photo_small.encode())
	x.Bytes(e.photo_big.encode())
	return x.buf
}

func (e TL_userStatusEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusEmpty)
	return x.buf
}

func (e TL_userStatusOnline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusOnline)
	x.Int(e.expires)
	return x.buf
}

func (e TL_userStatusOffline) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusOffline)
	x.Int(e.was_online)
	return x.buf
}

func (e TL_chatEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatEmpty)
	x.Int(e.id)
	return x.buf
}

func (e TL_chat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chat)
	x.Int(e.id)
	x.String(e.title)
	x.Bytes(e.photo.encode())
	x.Int(e.participants_count)
	x.Int(e.date)
	x.Bytes(e.left.encode())
	x.Int(e.version)
	return x.buf
}

func (e TL_chatForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatForbidden)
	x.Int(e.id)
	x.String(e.title)
	x.Int(e.date)
	return x.buf
}

func (e TL_chatFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatFull)
	x.Int(e.id)
	x.Bytes(e.participants.encode())
	x.Bytes(e.chat_photo.encode())
	x.Bytes(e.notify_settings.encode())
	return x.buf
}

func (e TL_chatParticipant) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipant)
	x.Int(e.user_id)
	x.Int(e.inviter_id)
	x.Int(e.date)
	return x.buf
}

func (e TL_chatParticipantsForbidden) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipantsForbidden)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_chatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatParticipants)
	x.Int(e.chat_id)
	x.Int(e.admin_id)
	x.Vector(e.participants)
	x.Int(e.version)
	return x.buf
}

func (e TL_chatPhotoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatPhotoEmpty)
	return x.buf
}

func (e TL_chatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatPhoto)
	x.Bytes(e.photo_small.encode())
	x.Bytes(e.photo_big.encode())
	return x.buf
}

func (e TL_messageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageEmpty)
	x.Int(e.id)
	return x.buf
}

func (e TL_message) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_message)
	x.Int(e.flags)
	x.Int(e.id)
	x.Int(e.from_id)
	x.Bytes(e.to_id.encode())
	x.Int(e.date)
	x.String(e.message)
	x.Bytes(e.media.encode())
	return x.buf
}

func (e TL_messageForwarded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageForwarded)
	x.Int(e.flags)
	x.Int(e.id)
	x.Int(e.fwd_from_id)
	x.Int(e.fwd_date)
	x.Int(e.from_id)
	x.Bytes(e.to_id.encode())
	x.Int(e.date)
	x.String(e.message)
	x.Bytes(e.media.encode())
	return x.buf
}

func (e TL_messageService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageService)
	x.Int(e.flags)
	x.Int(e.id)
	x.Int(e.from_id)
	x.Bytes(e.to_id.encode())
	x.Int(e.date)
	x.Bytes(e.action.encode())
	return x.buf
}

func (e TL_messageMediaEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaEmpty)
	return x.buf
}

func (e TL_messageMediaPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaPhoto)
	x.Bytes(e.photo.encode())
	return x.buf
}

func (e TL_messageMediaVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaVideo)
	x.Bytes(e.video.encode())
	return x.buf
}

func (e TL_messageMediaGeo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaGeo)
	x.Bytes(e.geo.encode())
	return x.buf
}

func (e TL_messageMediaContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaContact)
	x.String(e.phone_number)
	x.String(e.first_name)
	x.String(e.last_name)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_messageMediaUnsupported) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaUnsupported)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_messageActionEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionEmpty)
	return x.buf
}

func (e TL_messageActionChatCreate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatCreate)
	x.String(e.title)
	x.VectorInt(e.users)
	return x.buf
}

func (e TL_messageActionChatEditTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatEditTitle)
	x.String(e.title)
	return x.buf
}

func (e TL_messageActionChatEditPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatEditPhoto)
	x.Bytes(e.photo.encode())
	return x.buf
}

func (e TL_messageActionChatDeletePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatDeletePhoto)
	return x.buf
}

func (e TL_messageActionChatAddUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatAddUser)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_messageActionChatDeleteUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionChatDeleteUser)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_dialog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_dialog)
	x.Bytes(e.peer.encode())
	x.Int(e.top_message)
	x.Int(e.unread_count)
	x.Bytes(e.notify_settings.encode())
	return x.buf
}

func (e TL_photoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoEmpty)
	x.Long(e.id)
	return x.buf
}

func (e TL_photo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photo)
	x.Long(e.id)
	x.Long(e.access_hash)
	x.Int(e.user_id)
	x.Int(e.date)
	x.String(e.caption)
	x.Bytes(e.geo.encode())
	x.Vector(e.sizes)
	return x.buf
}

func (e TL_photoSizeEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoSizeEmpty)
	x.String(e._type)
	return x.buf
}

func (e TL_photoSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoSize)
	x.String(e._type)
	x.Bytes(e.location.encode())
	x.Int(e.w)
	x.Int(e.h)
	x.Int(e.size)
	return x.buf
}

func (e TL_photoCachedSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photoCachedSize)
	x.String(e._type)
	x.Bytes(e.location.encode())
	x.Int(e.w)
	x.Int(e.h)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_videoEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_videoEmpty)
	x.Long(e.id)
	return x.buf
}

func (e TL_video) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_video)
	x.Long(e.id)
	x.Long(e.access_hash)
	x.Int(e.user_id)
	x.Int(e.date)
	x.String(e.caption)
	x.Int(e.duration)
	x.String(e.mime_type)
	x.Int(e.size)
	x.Bytes(e.thumb.encode())
	x.Int(e.dc_id)
	x.Int(e.w)
	x.Int(e.h)
	return x.buf
}

func (e TL_geoPointEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geoPointEmpty)
	return x.buf
}

func (e TL_geoPoint) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geoPoint)
	x.Double(e.long)
	x.Double(e.lat)
	return x.buf
}

func (e TL_auth_checkedPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_checkedPhone)
	x.Bytes(e.phone_registered.encode())
	x.Bytes(e.phone_invited.encode())
	return x.buf
}

func (e TL_auth_sentCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentCode)
	x.Bytes(e.phone_registered.encode())
	x.String(e.phone_code_hash)
	x.Int(e.send_call_timeout)
	x.Bytes(e.is_password.encode())
	return x.buf
}

func (e TL_auth_authorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_authorization)
	x.Int(e.expires)
	x.Bytes(e.user.encode())
	return x.buf
}

func (e TL_auth_exportedAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_exportedAuthorization)
	x.Int(e.id)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_inputNotifyPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyPeer)
	x.Bytes(e.peer.encode())
	return x.buf
}

func (e TL_inputNotifyUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyUsers)
	return x.buf
}

func (e TL_inputNotifyChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyChats)
	return x.buf
}

func (e TL_inputNotifyAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyAll)
	return x.buf
}

func (e TL_inputPeerNotifyEventsEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerNotifyEventsEmpty)
	return x.buf
}

func (e TL_inputPeerNotifyEventsAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerNotifyEventsAll)
	return x.buf
}

func (e TL_inputPeerNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPeerNotifySettings)
	x.Int(e.mute_until)
	x.String(e.sound)
	x.Bytes(e.show_previews.encode())
	x.Int(e.events_mask)
	return x.buf
}

func (e TL_peerNotifyEventsEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifyEventsEmpty)
	return x.buf
}

func (e TL_peerNotifyEventsAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifyEventsAll)
	return x.buf
}

func (e TL_peerNotifySettingsEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifySettingsEmpty)
	return x.buf
}

func (e TL_peerNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_peerNotifySettings)
	x.Int(e.mute_until)
	x.String(e.sound)
	x.Bytes(e.show_previews.encode())
	x.Int(e.events_mask)
	return x.buf
}

func (e TL_wallPaper) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_wallPaper)
	x.Int(e.id)
	x.String(e.title)
	x.Vector(e.sizes)
	x.Int(e.color)
	return x.buf
}

func (e TL_userFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userFull)
	x.Bytes(e.user.encode())
	x.Bytes(e.link.encode())
	x.Bytes(e.profile_photo.encode())
	x.Bytes(e.notify_settings.encode())
	x.Bytes(e.blocked.encode())
	x.String(e.real_first_name)
	x.String(e.real_last_name)
	return x.buf
}

func (e TL_contact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contact)
	x.Int(e.user_id)
	x.Bytes(e.mutual.encode())
	return x.buf
}

func (e TL_importedContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_importedContact)
	x.Int(e.user_id)
	x.Long(e.client_id)
	return x.buf
}

func (e TL_contactBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactBlocked)
	x.Int(e.user_id)
	x.Int(e.date)
	return x.buf
}

func (e TL_contactSuggested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactSuggested)
	x.Int(e.user_id)
	x.Int(e.mutual_contacts)
	return x.buf
}

func (e TL_contactStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactStatus)
	x.Int(e.user_id)
	x.Bytes(e.status.encode())
	return x.buf
}

func (e TL_chatLocated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_chatLocated)
	x.Int(e.chat_id)
	x.Int(e.distance)
	return x.buf
}

func (e TL_contacts_foreignLinkUnknown) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_foreignLinkUnknown)
	return x.buf
}

func (e TL_contacts_foreignLinkRequested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_foreignLinkRequested)
	x.Bytes(e.has_phone.encode())
	return x.buf
}

func (e TL_contacts_foreignLinkMutual) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_foreignLinkMutual)
	return x.buf
}

func (e TL_contacts_myLinkEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_myLinkEmpty)
	return x.buf
}

func (e TL_contacts_myLinkRequested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_myLinkRequested)
	x.Bytes(e.contact.encode())
	return x.buf
}

func (e TL_contacts_myLinkContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_myLinkContact)
	return x.buf
}

func (e TL_contacts_link) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_link)
	x.Bytes(e.my_link.encode())
	x.Bytes(e.foreign_link.encode())
	x.Bytes(e.user.encode())
	return x.buf
}

func (e TL_contacts_contactsNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_contactsNotModified)
	return x.buf
}

func (e TL_contacts_contacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_contacts)
	x.Vector(e.contacts)
	x.Vector(e.users)
	return x.buf
}

func (e TL_contacts_importedContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_importedContacts)
	x.Vector(e.imported)
	x.VectorLong(e.retry_contacts)
	x.Vector(e.users)
	return x.buf
}

func (e TL_contacts_blocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_blocked)
	x.Vector(e.blocked)
	x.Vector(e.users)
	return x.buf
}

func (e TL_contacts_blockedSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_blockedSlice)
	x.Int(e.count)
	x.Vector(e.blocked)
	x.Vector(e.users)
	return x.buf
}

func (e TL_contacts_suggested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_suggested)
	x.Vector(e.results)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messages_dialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dialogs)
	x.Vector(e.dialogs)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messages_dialogsSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dialogsSlice)
	x.Int(e.count)
	x.Vector(e.dialogs)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messages_messages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_messages)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messages_messagesSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_messagesSlice)
	x.Int(e.count)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messages_messageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_messageEmpty)
	return x.buf
}

func (e TL_messages_statedMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_statedMessages)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	x.Int(e.pts)
	x.Int(e.seq)
	return x.buf
}

func (e TL_messages_statedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_statedMessage)
	x.Bytes(e.message.encode())
	x.Vector(e.chats)
	x.Vector(e.users)
	x.Int(e.pts)
	x.Int(e.seq)
	return x.buf
}

func (e TL_messages_sentMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sentMessage)
	x.Int(e.id)
	x.Int(e.date)
	x.Int(e.pts)
	x.Int(e.seq)
	return x.buf
}

func (e TL_messages_chats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_chats)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messages_chatFull) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_chatFull)
	x.Bytes(e.full_chat.encode())
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messages_affectedHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_affectedHistory)
	x.Int(e.pts)
	x.Int(e.seq)
	x.Int(e.offset)
	return x.buf
}

func (e TL_inputMessagesFilterEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterEmpty)
	return x.buf
}

func (e TL_inputMessagesFilterPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhotos)
	return x.buf
}

func (e TL_inputMessagesFilterVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterVideo)
	return x.buf
}

func (e TL_inputMessagesFilterPhotoVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhotoVideo)
	return x.buf
}

func (e TL_inputMessagesFilterPhotoVideoDocuments) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterPhotoVideoDocuments)
	return x.buf
}

func (e TL_inputMessagesFilterDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterDocument)
	return x.buf
}

func (e TL_inputMessagesFilterAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMessagesFilterAudio)
	return x.buf
}

func (e TL_updateNewMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewMessage)
	x.Bytes(e.message.encode())
	x.Int(e.pts)
	return x.buf
}

func (e TL_updateMessageID) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateMessageID)
	x.Int(e.id)
	x.Long(e.random_id)
	return x.buf
}

func (e TL_updateReadMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateReadMessages)
	x.VectorInt(e.messages)
	x.Int(e.pts)
	return x.buf
}

func (e TL_updateDeleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDeleteMessages)
	x.VectorInt(e.messages)
	x.Int(e.pts)
	return x.buf
}

func (e TL_updateUserTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserTyping)
	x.Int(e.user_id)
	x.Bytes(e.action.encode())
	return x.buf
}

func (e TL_updateChatUserTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatUserTyping)
	x.Int(e.chat_id)
	x.Int(e.user_id)
	x.Bytes(e.action.encode())
	return x.buf
}

func (e TL_updateChatParticipants) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipants)
	x.Bytes(e.participants.encode())
	return x.buf
}

func (e TL_updateUserStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserStatus)
	x.Int(e.user_id)
	x.Bytes(e.status.encode())
	return x.buf
}

func (e TL_updateUserName) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserName)
	x.Int(e.user_id)
	x.String(e.first_name)
	x.String(e.last_name)
	x.String(e.username)
	return x.buf
}

func (e TL_updateUserPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserPhoto)
	x.Int(e.user_id)
	x.Int(e.date)
	x.Bytes(e.photo.encode())
	x.Bytes(e.previous.encode())
	return x.buf
}

func (e TL_updateContactRegistered) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateContactRegistered)
	x.Int(e.user_id)
	x.Int(e.date)
	return x.buf
}

func (e TL_updateContactLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateContactLink)
	x.Int(e.user_id)
	x.Bytes(e.my_link.encode())
	x.Bytes(e.foreign_link.encode())
	return x.buf
}

func (e TL_updateNewAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewAuthorization)
	x.Long(e.auth_key_id)
	x.Int(e.date)
	x.String(e.device)
	x.String(e.location)
	return x.buf
}

func (e TL_updates_state) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_state)
	x.Int(e.pts)
	x.Int(e.qts)
	x.Int(e.date)
	x.Int(e.seq)
	x.Int(e.unread_count)
	return x.buf
}

func (e TL_updates_differenceEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_differenceEmpty)
	x.Int(e.date)
	x.Int(e.seq)
	return x.buf
}

func (e TL_updates_difference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_difference)
	x.Vector(e.new_messages)
	x.Vector(e.new_encrypted_messages)
	x.Vector(e.other_updates)
	x.Vector(e.chats)
	x.Vector(e.users)
	x.Bytes(e.state.encode())
	return x.buf
}

func (e TL_updates_differenceSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_differenceSlice)
	x.Vector(e.new_messages)
	x.Vector(e.new_encrypted_messages)
	x.Vector(e.other_updates)
	x.Vector(e.chats)
	x.Vector(e.users)
	x.Bytes(e.intermediate_state.encode())
	return x.buf
}

func (e TL_updatesTooLong) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesTooLong)
	return x.buf
}

func (e TL_updateShortMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShortMessage)
	x.Int(e.id)
	x.Int(e.from_id)
	x.String(e.message)
	x.Int(e.pts)
	x.Int(e.date)
	x.Int(e.seq)
	return x.buf
}

func (e TL_updateShortChatMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShortChatMessage)
	x.Int(e.id)
	x.Int(e.from_id)
	x.Int(e.chat_id)
	x.String(e.message)
	x.Int(e.pts)
	x.Int(e.date)
	x.Int(e.seq)
	return x.buf
}

func (e TL_updateShort) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateShort)
	x.Bytes(e.update.encode())
	x.Int(e.date)
	return x.buf
}

func (e TL_updatesCombined) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatesCombined)
	x.Vector(e.updates)
	x.Vector(e.users)
	x.Vector(e.chats)
	x.Int(e.date)
	x.Int(e.seq_start)
	x.Int(e.seq)
	return x.buf
}

func (e TL_updates) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates)
	x.Vector(e.updates)
	x.Vector(e.users)
	x.Vector(e.chats)
	x.Int(e.date)
	x.Int(e.seq)
	return x.buf
}

func (e TL_photos_photos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_photos)
	x.Vector(e.photos)
	x.Vector(e.users)
	return x.buf
}

func (e TL_photos_photosSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_photosSlice)
	x.Int(e.count)
	x.Vector(e.photos)
	x.Vector(e.users)
	return x.buf
}

func (e TL_photos_photo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_photo)
	x.Bytes(e.photo.encode())
	x.Vector(e.users)
	return x.buf
}

func (e TL_upload_file) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_file)
	x.Bytes(e._type.encode())
	x.Int(e.mtime)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_dcOption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_dcOption)
	x.Int(e.id)
	x.String(e.hostname)
	x.String(e.ip_address)
	x.Int(e.port)
	return x.buf
}

func (e TL_config) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_config)
	x.Int(e.date)
	x.Int(e.expires)
	x.Bytes(e.test_mode.encode())
	x.Int(e.this_dc)
	x.Vector(e.dc_options)
	x.Int(e.chat_big_size)
	x.Int(e.chat_size_max)
	x.Int(e.broadcast_size_max)
	x.Vector(e.disabled_features)
	return x.buf
}

func (e TL_nearestDc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_nearestDc)
	x.String(e.country)
	x.Int(e.this_dc)
	x.Int(e.nearest_dc)
	return x.buf
}

func (e TL_help_appUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_appUpdate)
	x.Int(e.id)
	x.Bytes(e.critical.encode())
	x.String(e.url)
	x.String(e.text)
	return x.buf
}

func (e TL_help_noAppUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_noAppUpdate)
	return x.buf
}

func (e TL_help_inviteText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_inviteText)
	x.String(e.message)
	return x.buf
}

func (e TL_messages_statedMessagesLinks) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_statedMessagesLinks)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	x.Vector(e.links)
	x.Int(e.pts)
	x.Int(e.seq)
	return x.buf
}

func (e TL_messages_statedMessageLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_statedMessageLink)
	x.Bytes(e.message.encode())
	x.Vector(e.chats)
	x.Vector(e.users)
	x.Vector(e.links)
	x.Int(e.pts)
	x.Int(e.seq)
	return x.buf
}

func (e TL_messages_sentMessageLink) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sentMessageLink)
	x.Int(e.id)
	x.Int(e.date)
	x.Int(e.pts)
	x.Int(e.seq)
	x.Vector(e.links)
	return x.buf
}

func (e TL_inputGeoChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputGeoChat)
	x.Int(e.chat_id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputNotifyGeoChatPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputNotifyGeoChatPeer)
	x.Bytes(e.peer.encode())
	return x.buf
}

func (e TL_geoChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geoChat)
	x.Int(e.id)
	x.Long(e.access_hash)
	x.String(e.title)
	x.String(e.address)
	x.String(e.venue)
	x.Bytes(e.geo.encode())
	x.Bytes(e.photo.encode())
	x.Int(e.participants_count)
	x.Int(e.date)
	x.Bytes(e.checked_in.encode())
	x.Int(e.version)
	return x.buf
}

func (e TL_geoChatMessageEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geoChatMessageEmpty)
	x.Int(e.chat_id)
	x.Int(e.id)
	return x.buf
}

func (e TL_geoChatMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geoChatMessage)
	x.Int(e.chat_id)
	x.Int(e.id)
	x.Int(e.from_id)
	x.Int(e.date)
	x.String(e.message)
	x.Bytes(e.media.encode())
	return x.buf
}

func (e TL_geoChatMessageService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geoChatMessageService)
	x.Int(e.chat_id)
	x.Int(e.id)
	x.Int(e.from_id)
	x.Int(e.date)
	x.Bytes(e.action.encode())
	return x.buf
}

func (e TL_geochats_statedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geochats_statedMessage)
	x.Bytes(e.message.encode())
	x.Vector(e.chats)
	x.Vector(e.users)
	x.Int(e.seq)
	return x.buf
}

func (e TL_geochats_located) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geochats_located)
	x.Vector(e.results)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_geochats_messages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geochats_messages)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_geochats_messagesSlice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geochats_messagesSlice)
	x.Int(e.count)
	x.Vector(e.messages)
	x.Vector(e.chats)
	x.Vector(e.users)
	return x.buf
}

func (e TL_messageActionGeoChatCreate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionGeoChatCreate)
	x.String(e.title)
	x.String(e.address)
	return x.buf
}

func (e TL_messageActionGeoChatCheckin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageActionGeoChatCheckin)
	return x.buf
}

func (e TL_updateNewGeoChatMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewGeoChatMessage)
	x.Bytes(e.message.encode())
	return x.buf
}

func (e TL_wallPaperSolid) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_wallPaperSolid)
	x.Int(e.id)
	x.String(e.title)
	x.Int(e.bg_color)
	x.Int(e.color)
	return x.buf
}

func (e TL_updateNewEncryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNewEncryptedMessage)
	x.Bytes(e.message.encode())
	x.Int(e.qts)
	return x.buf
}

func (e TL_updateEncryptedChatTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEncryptedChatTyping)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_updateEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEncryption)
	x.Bytes(e.chat.encode())
	x.Int(e.date)
	return x.buf
}

func (e TL_updateEncryptedMessagesRead) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateEncryptedMessagesRead)
	x.Int(e.chat_id)
	x.Int(e.max_date)
	x.Int(e.date)
	return x.buf
}

func (e TL_encryptedChatEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatEmpty)
	x.Int(e.id)
	return x.buf
}

func (e TL_encryptedChatWaiting) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatWaiting)
	x.Int(e.id)
	x.Long(e.access_hash)
	x.Int(e.date)
	x.Int(e.admin_id)
	x.Int(e.participant_id)
	return x.buf
}

func (e TL_encryptedChatRequested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatRequested)
	x.Int(e.id)
	x.Long(e.access_hash)
	x.Int(e.date)
	x.Int(e.admin_id)
	x.Int(e.participant_id)
	x.StringBytes(e.g_a)
	return x.buf
}

func (e TL_encryptedChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChat)
	x.Int(e.id)
	x.Long(e.access_hash)
	x.Int(e.date)
	x.Int(e.admin_id)
	x.Int(e.participant_id)
	x.StringBytes(e.g_a_or_b)
	x.Long(e.key_fingerprint)
	return x.buf
}

func (e TL_encryptedChatDiscarded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedChatDiscarded)
	x.Int(e.id)
	return x.buf
}

func (e TL_inputEncryptedChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedChat)
	x.Int(e.chat_id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_encryptedFileEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedFileEmpty)
	return x.buf
}

func (e TL_encryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedFile)
	x.Long(e.id)
	x.Long(e.access_hash)
	x.Int(e.size)
	x.Int(e.dc_id)
	x.Int(e.key_fingerprint)
	return x.buf
}

func (e TL_inputEncryptedFileEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileEmpty)
	return x.buf
}

func (e TL_inputEncryptedFileUploaded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileUploaded)
	x.Long(e.id)
	x.Int(e.parts)
	x.String(e.md5_checksum)
	x.Int(e.key_fingerprint)
	return x.buf
}

func (e TL_inputEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFile)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputEncryptedFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileLocation)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_encryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedMessage)
	x.Long(e.random_id)
	x.Int(e.chat_id)
	x.Int(e.date)
	x.StringBytes(e.bytes)
	x.Bytes(e.file.encode())
	return x.buf
}

func (e TL_encryptedMessageService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_encryptedMessageService)
	x.Long(e.random_id)
	x.Int(e.chat_id)
	x.Int(e.date)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_messages_dhConfigNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dhConfigNotModified)
	x.StringBytes(e.random)
	return x.buf
}

func (e TL_messages_dhConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_dhConfig)
	x.Int(e.g)
	x.StringBytes(e.p)
	x.Int(e.version)
	x.StringBytes(e.random)
	return x.buf
}

func (e TL_messages_sentEncryptedMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sentEncryptedMessage)
	x.Int(e.date)
	return x.buf
}

func (e TL_messages_sentEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sentEncryptedFile)
	x.Int(e.date)
	x.Bytes(e.file.encode())
	return x.buf
}

func (e TL_inputFileBig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputFileBig)
	x.Long(e.id)
	x.Int(e.parts)
	x.String(e.name)
	return x.buf
}

func (e TL_inputEncryptedFileBigUploaded) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputEncryptedFileBigUploaded)
	x.Long(e.id)
	x.Int(e.parts)
	x.Int(e.key_fingerprint)
	return x.buf
}

func (e TL_updateChatParticipantAdd) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipantAdd)
	x.Int(e.chat_id)
	x.Int(e.user_id)
	x.Int(e.inviter_id)
	x.Int(e.version)
	return x.buf
}

func (e TL_updateChatParticipantDelete) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateChatParticipantDelete)
	x.Int(e.chat_id)
	x.Int(e.user_id)
	x.Int(e.version)
	return x.buf
}

func (e TL_updateDcOptions) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateDcOptions)
	x.Vector(e.dc_options)
	return x.buf
}

func (e TL_inputMediaUploadedAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaUploadedAudio)
	x.Bytes(e.file.encode())
	x.Int(e.duration)
	x.String(e.mime_type)
	return x.buf
}

func (e TL_inputMediaAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaAudio)
	x.Bytes(e.id.encode())
	return x.buf
}

func (e TL_inputMediaUploadedDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaUploadedDocument)
	x.Bytes(e.file.encode())
	x.String(e.mime_type)
	x.Vector(e.attributes)
	return x.buf
}

func (e TL_inputMediaUploadedThumbDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaUploadedThumbDocument)
	x.Bytes(e.file.encode())
	x.Bytes(e.thumb.encode())
	x.String(e.mime_type)
	x.Vector(e.attributes)
	return x.buf
}

func (e TL_inputMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputMediaDocument)
	x.Bytes(e.id.encode())
	return x.buf
}

func (e TL_messageMediaDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaDocument)
	x.Bytes(e.document.encode())
	return x.buf
}

func (e TL_messageMediaAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messageMediaAudio)
	x.Bytes(e.audio.encode())
	return x.buf
}

func (e TL_inputAudioEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputAudioEmpty)
	return x.buf
}

func (e TL_inputAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputAudio)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputDocumentEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDocumentEmpty)
	return x.buf
}

func (e TL_inputDocument) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDocument)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputAudioFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputAudioFileLocation)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_inputDocumentFileLocation) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputDocumentFileLocation)
	x.Long(e.id)
	x.Long(e.access_hash)
	return x.buf
}

func (e TL_audioEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_audioEmpty)
	x.Long(e.id)
	return x.buf
}

func (e TL_audio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_audio)
	x.Long(e.id)
	x.Long(e.access_hash)
	x.Int(e.user_id)
	x.Int(e.date)
	x.Int(e.duration)
	x.String(e.mime_type)
	x.Int(e.size)
	x.Int(e.dc_id)
	return x.buf
}

func (e TL_documentEmpty) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentEmpty)
	x.Long(e.id)
	return x.buf
}

func (e TL_document) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_document)
	x.Long(e.id)
	x.Long(e.access_hash)
	x.Int(e.date)
	x.String(e.mime_type)
	x.Int(e.size)
	x.Bytes(e.thumb.encode())
	x.Int(e.dc_id)
	x.Vector(e.attributes)
	return x.buf
}

func (e TL_help_support) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_support)
	x.String(e.phone_number)
	x.Bytes(e.user.encode())
	return x.buf
}

func (e TL_notifyPeer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyPeer)
	x.Bytes(e.peer.encode())
	return x.buf
}

func (e TL_notifyUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyUsers)
	return x.buf
}

func (e TL_notifyChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyChats)
	return x.buf
}

func (e TL_notifyAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_notifyAll)
	return x.buf
}

func (e TL_updateUserBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserBlocked)
	x.Int(e.user_id)
	x.Bytes(e.blocked.encode())
	return x.buf
}

func (e TL_updateNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateNotifySettings)
	x.Bytes(e.peer.encode())
	x.Bytes(e.notify_settings.encode())
	return x.buf
}

func (e TL_auth_sentAppCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sentAppCode)
	x.Bytes(e.phone_registered.encode())
	x.String(e.phone_code_hash)
	x.Int(e.send_call_timeout)
	x.Bytes(e.is_password.encode())
	return x.buf
}

func (e TL_sendMessageTypingAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageTypingAction)
	return x.buf
}

func (e TL_sendMessageCancelAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageCancelAction)
	return x.buf
}

func (e TL_sendMessageRecordVideoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageRecordVideoAction)
	return x.buf
}

func (e TL_sendMessageUploadVideoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadVideoAction)
	return x.buf
}

func (e TL_sendMessageRecordAudioAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageRecordAudioAction)
	return x.buf
}

func (e TL_sendMessageUploadAudioAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadAudioAction)
	return x.buf
}

func (e TL_sendMessageUploadPhotoAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadPhotoAction)
	return x.buf
}

func (e TL_sendMessageUploadDocumentAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageUploadDocumentAction)
	return x.buf
}

func (e TL_sendMessageGeoLocationAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageGeoLocationAction)
	return x.buf
}

func (e TL_sendMessageChooseContactAction) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_sendMessageChooseContactAction)
	return x.buf
}

func (e TL_contactFound) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contactFound)
	x.Int(e.user_id)
	return x.buf
}

func (e TL_contacts_found) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_found)
	x.Vector(e.results)
	x.Vector(e.users)
	return x.buf
}

func (e TL_updateServiceNotification) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateServiceNotification)
	x.String(e._type)
	x.String(e.message)
	x.Bytes(e.media.encode())
	x.Bytes(e.popup.encode())
	return x.buf
}

func (e TL_userStatusRecently) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusRecently)
	return x.buf
}

func (e TL_userStatusLastWeek) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusLastWeek)
	return x.buf
}

func (e TL_userStatusLastMonth) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_userStatusLastMonth)
	return x.buf
}

func (e TL_updatePrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updatePrivacy)
	x.Bytes(e.key.encode())
	x.Vector(e.rules)
	return x.buf
}

func (e TL_inputPrivacyKeyStatusTimestamp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyKeyStatusTimestamp)
	return x.buf
}

func (e TL_privacyKeyStatusTimestamp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyKeyStatusTimestamp)
	return x.buf
}

func (e TL_inputPrivacyValueAllowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueAllowContacts)
	return x.buf
}

func (e TL_inputPrivacyValueAllowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueAllowAll)
	return x.buf
}

func (e TL_inputPrivacyValueAllowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueAllowUsers)
	x.Vector(e.users)
	return x.buf
}

func (e TL_inputPrivacyValueDisallowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueDisallowContacts)
	return x.buf
}

func (e TL_inputPrivacyValueDisallowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueDisallowAll)
	return x.buf
}

func (e TL_inputPrivacyValueDisallowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_inputPrivacyValueDisallowUsers)
	x.Vector(e.users)
	return x.buf
}

func (e TL_privacyValueAllowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueAllowContacts)
	return x.buf
}

func (e TL_privacyValueAllowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueAllowAll)
	return x.buf
}

func (e TL_privacyValueAllowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueAllowUsers)
	x.VectorInt(e.users)
	return x.buf
}

func (e TL_privacyValueDisallowContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueDisallowContacts)
	return x.buf
}

func (e TL_privacyValueDisallowAll) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueDisallowAll)
	return x.buf
}

func (e TL_privacyValueDisallowUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_privacyValueDisallowUsers)
	x.VectorInt(e.users)
	return x.buf
}

func (e TL_account_privacyRules) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_privacyRules)
	x.Vector(e.rules)
	x.Vector(e.users)
	return x.buf
}

func (e TL_accountDaysTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_accountDaysTTL)
	x.Int(e.days)
	return x.buf
}

func (e TL_account_sentChangePhoneCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_sentChangePhoneCode)
	x.String(e.phone_code_hash)
	x.Int(e.send_call_timeout)
	return x.buf
}

func (e TL_updateUserPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updateUserPhone)
	x.Int(e.user_id)
	x.String(e.phone)
	return x.buf
}

func (e TL_documentAttributeImageSize) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeImageSize)
	x.Int(e.w)
	x.Int(e.h)
	return x.buf
}

func (e TL_documentAttributeAnimated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeAnimated)
	return x.buf
}

func (e TL_documentAttributeSticker) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeSticker)
	return x.buf
}

func (e TL_documentAttributeVideo) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeVideo)
	x.Int(e.duration)
	x.Int(e.w)
	x.Int(e.h)
	return x.buf
}

func (e TL_documentAttributeAudio) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeAudio)
	x.Int(e.duration)
	return x.buf
}

func (e TL_documentAttributeFilename) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_documentAttributeFilename)
	x.String(e.file_name)
	return x.buf
}

func (e TL_messages_stickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickersNotModified)
	return x.buf
}

func (e TL_messages_stickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_stickers)
	x.String(e.hash)
	x.Vector(e.stickers)
	return x.buf
}

func (e TL_stickerPack) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_stickerPack)
	x.String(e.emoticon)
	x.VectorLong(e.documents)
	return x.buf
}

func (e TL_messages_allStickersNotModified) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_allStickersNotModified)
	return x.buf
}

func (e TL_messages_allStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_allStickers)
	x.String(e.hash)
	x.Vector(e.packs)
	x.Vector(e.documents)
	return x.buf
}

func (e TL_disabledFeature) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_disabledFeature)
	x.String(e.feature)
	x.String(e.description)
	return x.buf
}

func (e TL_invokeAfterMsg) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeAfterMsg)
	x.Long(e.msg_id)
	x.Bytes(e.query.encode())
	return x.buf
}

func (e TL_invokeAfterMsgs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeAfterMsgs)
	x.VectorLong(e.msg_ids)
	x.Bytes(e.query.encode())
	return x.buf
}

func (e TL_auth_checkPhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_checkPhone)
	x.String(e.phone_number)
	return x.buf
}

func (e TL_auth_sendCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sendCode)
	x.String(e.phone_number)
	x.Int(e.sms_type)
	x.Int(e.api_id)
	x.String(e.api_hash)
	x.String(e.lang_code)
	return x.buf
}

func (e TL_auth_sendCall) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sendCall)
	x.String(e.phone_number)
	x.String(e.phone_code_hash)
	return x.buf
}

func (e TL_auth_signUp) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_signUp)
	x.String(e.phone_number)
	x.String(e.phone_code_hash)
	x.String(e.phone_code)
	x.String(e.first_name)
	x.String(e.last_name)
	return x.buf
}

func (e TL_auth_signIn) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_signIn)
	x.String(e.phone_number)
	x.String(e.phone_code_hash)
	x.String(e.phone_code)
	return x.buf
}

func (e TL_auth_logOut) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_logOut)
	return x.buf
}

func (e TL_auth_resetAuthorizations) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_resetAuthorizations)
	return x.buf
}

func (e TL_auth_sendInvites) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sendInvites)
	x.VectorString(e.phone_numbers)
	x.String(e.message)
	return x.buf
}

func (e TL_auth_exportAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_exportAuthorization)
	x.Int(e.dc_id)
	return x.buf
}

func (e TL_auth_importAuthorization) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_importAuthorization)
	x.Int(e.id)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_auth_bindTempAuthKey) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_bindTempAuthKey)
	x.Long(e.perm_auth_key_id)
	x.Long(e.nonce)
	x.Int(e.expires_at)
	x.StringBytes(e.encrypted_message)
	return x.buf
}

func (e TL_account_registerDevice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_registerDevice)
	x.Int(e.token_type)
	x.String(e.token)
	x.String(e.device_model)
	x.String(e.system_version)
	x.String(e.app_version)
	x.Bytes(e.app_sandbox.encode())
	x.String(e.lang_code)
	return x.buf
}

func (e TL_account_unregisterDevice) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_unregisterDevice)
	x.Int(e.token_type)
	x.String(e.token)
	return x.buf
}

func (e TL_account_updateNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateNotifySettings)
	x.Bytes(e.peer.encode())
	x.Bytes(e.settings.encode())
	return x.buf
}

func (e TL_account_getNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getNotifySettings)
	x.Bytes(e.peer.encode())
	return x.buf
}

func (e TL_account_resetNotifySettings) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_resetNotifySettings)
	return x.buf
}

func (e TL_account_updateProfile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateProfile)
	x.String(e.first_name)
	x.String(e.last_name)
	return x.buf
}

func (e TL_account_updateStatus) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateStatus)
	x.Bytes(e.offline.encode())
	return x.buf
}

func (e TL_account_getWallPapers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getWallPapers)
	return x.buf
}

func (e TL_users_getUsers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_users_getUsers)
	x.Vector(e.id)
	return x.buf
}

func (e TL_users_getFullUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_users_getFullUser)
	x.Bytes(e.id.encode())
	return x.buf
}

func (e TL_contacts_getStatuses) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getStatuses)
	return x.buf
}

func (e TL_contacts_getContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getContacts)
	x.String(e.hash)
	return x.buf
}

func (e TL_contacts_importContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_importContacts)
	x.Vector(e.contacts)
	x.Bytes(e.replace.encode())
	return x.buf
}

func (e TL_contacts_getSuggested) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getSuggested)
	x.Int(e.limit)
	return x.buf
}

func (e TL_contacts_deleteContact) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_deleteContact)
	x.Bytes(e.id.encode())
	return x.buf
}

func (e TL_contacts_deleteContacts) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_deleteContacts)
	x.Vector(e.id)
	return x.buf
}

func (e TL_contacts_block) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_block)
	x.Bytes(e.id.encode())
	return x.buf
}

func (e TL_contacts_unblock) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_unblock)
	x.Bytes(e.id.encode())
	return x.buf
}

func (e TL_contacts_getBlocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_getBlocked)
	x.Int(e.offset)
	x.Int(e.limit)
	return x.buf
}

func (e TL_contacts_exportCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_exportCard)
	return x.buf
}

func (e TL_contacts_importCard) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_importCard)
	x.VectorInt(e.export_card)
	return x.buf
}

func (e TL_messages_getMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getMessages)
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_messages_getDialogs) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getDialogs)
	x.Int(e.offset)
	x.Int(e.max_id)
	x.Int(e.limit)
	return x.buf
}

func (e TL_messages_getHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getHistory)
	x.Bytes(e.peer.encode())
	x.Int(e.offset)
	x.Int(e.max_id)
	x.Int(e.limit)
	return x.buf
}

func (e TL_messages_search) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_search)
	x.Bytes(e.peer.encode())
	x.String(e.q)
	x.Bytes(e.filter.encode())
	x.Int(e.min_date)
	x.Int(e.max_date)
	x.Int(e.offset)
	x.Int(e.max_id)
	x.Int(e.limit)
	return x.buf
}

func (e TL_messages_readHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readHistory)
	x.Bytes(e.peer.encode())
	x.Int(e.max_id)
	x.Int(e.offset)
	x.Bytes(e.read_contents.encode())
	return x.buf
}

func (e TL_messages_deleteHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_deleteHistory)
	x.Bytes(e.peer.encode())
	x.Int(e.offset)
	return x.buf
}

func (e TL_messages_deleteMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_deleteMessages)
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_messages_receivedMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_receivedMessages)
	x.Int(e.max_id)
	return x.buf
}

func (e TL_messages_setTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setTyping)
	x.Bytes(e.peer.encode())
	x.Bytes(e.action.encode())
	return x.buf
}

func (e TL_messages_sendMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendMessage)
	x.Bytes(e.peer.encode())
	x.String(e.message)
	x.Long(e.random_id)
	return x.buf
}

func (e TL_messages_sendMedia) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendMedia)
	x.Bytes(e.peer.encode())
	x.Bytes(e.media.encode())
	x.Long(e.random_id)
	return x.buf
}

func (e TL_messages_forwardMessages) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_forwardMessages)
	x.Bytes(e.peer.encode())
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_messages_getChats) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getChats)
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_messages_getFullChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getFullChat)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_messages_editChatTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editChatTitle)
	x.Int(e.chat_id)
	x.String(e.title)
	return x.buf
}

func (e TL_messages_editChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_editChatPhoto)
	x.Int(e.chat_id)
	x.Bytes(e.photo.encode())
	return x.buf
}

func (e TL_messages_addChatUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_addChatUser)
	x.Int(e.chat_id)
	x.Bytes(e.user_id.encode())
	x.Int(e.fwd_limit)
	return x.buf
}

func (e TL_messages_deleteChatUser) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_deleteChatUser)
	x.Int(e.chat_id)
	x.Bytes(e.user_id.encode())
	return x.buf
}

func (e TL_messages_createChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_createChat)
	x.Vector(e.users)
	x.String(e.title)
	return x.buf
}

func (e TL_updates_getState) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_getState)
	return x.buf
}

func (e TL_updates_getDifference) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_updates_getDifference)
	x.Int(e.pts)
	x.Int(e.date)
	x.Int(e.qts)
	return x.buf
}

func (e TL_photos_updateProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_updateProfilePhoto)
	x.Bytes(e.id.encode())
	x.Bytes(e.crop.encode())
	return x.buf
}

func (e TL_photos_uploadProfilePhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_uploadProfilePhoto)
	x.Bytes(e.file.encode())
	x.String(e.caption)
	x.Bytes(e.geo_point.encode())
	x.Bytes(e.crop.encode())
	return x.buf
}

func (e TL_photos_deletePhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_deletePhotos)
	x.Vector(e.id)
	return x.buf
}

func (e TL_upload_saveFilePart) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_saveFilePart)
	x.Long(e.file_id)
	x.Int(e.file_part)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_upload_getFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_getFile)
	x.Bytes(e.location.encode())
	x.Int(e.offset)
	x.Int(e.limit)
	return x.buf
}

func (e TL_help_getConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getConfig)
	return x.buf
}

func (e TL_help_getNearestDc) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getNearestDc)
	return x.buf
}

func (e TL_help_getAppUpdate) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getAppUpdate)
	x.String(e.device_model)
	x.String(e.system_version)
	x.String(e.app_version)
	x.String(e.lang_code)
	return x.buf
}

func (e TL_help_saveAppLog) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_saveAppLog)
	x.Vector(e.events)
	return x.buf
}

func (e TL_help_getInviteText) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getInviteText)
	x.String(e.lang_code)
	return x.buf
}

func (e TL_photos_getUserPhotos) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_photos_getUserPhotos)
	x.Bytes(e.user_id.encode())
	x.Int(e.offset)
	x.Int(e.max_id)
	x.Int(e.limit)
	return x.buf
}

func (e TL_messages_forwardMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_forwardMessage)
	x.Bytes(e.peer.encode())
	x.Int(e.id)
	x.Long(e.random_id)
	return x.buf
}

func (e TL_messages_sendBroadcast) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendBroadcast)
	x.Vector(e.contacts)
	x.String(e.message)
	x.Bytes(e.media.encode())
	return x.buf
}

func (e TL_geochats_getLocated) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geochats_getLocated)
	x.Bytes(e.geo_point.encode())
	x.Int(e.radius)
	x.Int(e.limit)
	return x.buf
}

func (e TL_geochats_getRecents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geochats_getRecents)
	x.Int(e.offset)
	x.Int(e.limit)
	return x.buf
}

func (e TL_geochats_checkin) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geochats_checkin)
	x.Bytes(e.peer.encode())
	return x.buf
}

func (e TL_geochats_getFullChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geochats_getFullChat)
	x.Bytes(e.peer.encode())
	return x.buf
}

func (e TL_geochats_editChatTitle) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geochats_editChatTitle)
	x.Bytes(e.peer.encode())
	x.String(e.title)
	x.String(e.address)
	return x.buf
}

func (e TL_geochats_editChatPhoto) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geochats_editChatPhoto)
	x.Bytes(e.peer.encode())
	x.Bytes(e.photo.encode())
	return x.buf
}

func (e TL_geochats_search) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geochats_search)
	x.Bytes(e.peer.encode())
	x.String(e.q)
	x.Bytes(e.filter.encode())
	x.Int(e.min_date)
	x.Int(e.max_date)
	x.Int(e.offset)
	x.Int(e.max_id)
	x.Int(e.limit)
	return x.buf
}

func (e TL_geochats_getHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geochats_getHistory)
	x.Bytes(e.peer.encode())
	x.Int(e.offset)
	x.Int(e.max_id)
	x.Int(e.limit)
	return x.buf
}

func (e TL_geochats_setTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geochats_setTyping)
	x.Bytes(e.peer.encode())
	x.Bytes(e.typing.encode())
	return x.buf
}

func (e TL_geochats_sendMessage) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geochats_sendMessage)
	x.Bytes(e.peer.encode())
	x.String(e.message)
	x.Long(e.random_id)
	return x.buf
}

func (e TL_geochats_sendMedia) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geochats_sendMedia)
	x.Bytes(e.peer.encode())
	x.Bytes(e.media.encode())
	x.Long(e.random_id)
	return x.buf
}

func (e TL_geochats_createGeoChat) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_geochats_createGeoChat)
	x.String(e.title)
	x.Bytes(e.geo_point.encode())
	x.String(e.address)
	x.String(e.venue)
	return x.buf
}

func (e TL_messages_getDhConfig) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getDhConfig)
	x.Int(e.version)
	x.Int(e.random_length)
	return x.buf
}

func (e TL_messages_requestEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_requestEncryption)
	x.Bytes(e.user_id.encode())
	x.Int(e.random_id)
	x.StringBytes(e.g_a)
	return x.buf
}

func (e TL_messages_acceptEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_acceptEncryption)
	x.Bytes(e.peer.encode())
	x.StringBytes(e.g_b)
	x.Long(e.key_fingerprint)
	return x.buf
}

func (e TL_messages_discardEncryption) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_discardEncryption)
	x.Int(e.chat_id)
	return x.buf
}

func (e TL_messages_setEncryptedTyping) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_setEncryptedTyping)
	x.Bytes(e.peer.encode())
	x.Bytes(e.typing.encode())
	return x.buf
}

func (e TL_messages_readEncryptedHistory) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readEncryptedHistory)
	x.Bytes(e.peer.encode())
	x.Int(e.max_date)
	return x.buf
}

func (e TL_messages_sendEncrypted) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendEncrypted)
	x.Bytes(e.peer.encode())
	x.Long(e.random_id)
	x.StringBytes(e.data)
	return x.buf
}

func (e TL_messages_sendEncryptedFile) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendEncryptedFile)
	x.Bytes(e.peer.encode())
	x.Long(e.random_id)
	x.StringBytes(e.data)
	x.Bytes(e.file.encode())
	return x.buf
}

func (e TL_messages_sendEncryptedService) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_sendEncryptedService)
	x.Bytes(e.peer.encode())
	x.Long(e.random_id)
	x.StringBytes(e.data)
	return x.buf
}

func (e TL_messages_receivedQueue) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_receivedQueue)
	x.Int(e.max_qts)
	return x.buf
}

func (e TL_upload_saveBigFilePart) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_upload_saveBigFilePart)
	x.Long(e.file_id)
	x.Int(e.file_part)
	x.Int(e.file_total_parts)
	x.StringBytes(e.bytes)
	return x.buf
}

func (e TL_initConnection) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_initConnection)
	x.Int(e.api_id)
	x.String(e.device_model)
	x.String(e.system_version)
	x.String(e.app_version)
	x.String(e.lang_code)
	x.Bytes(e.query.encode())
	return x.buf
}

func (e TL_help_getSupport) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_help_getSupport)
	return x.buf
}

func (e TL_auth_sendSms) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_auth_sendSms)
	x.String(e.phone_number)
	x.String(e.phone_code_hash)
	return x.buf
}

func (e TL_messages_readMessageContents) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_readMessageContents)
	x.VectorInt(e.id)
	return x.buf
}

func (e TL_account_checkUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_checkUsername)
	x.String(e.username)
	return x.buf
}

func (e TL_account_updateUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateUsername)
	x.String(e.username)
	return x.buf
}

func (e TL_contacts_search) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_search)
	x.String(e.q)
	x.Int(e.limit)
	return x.buf
}

func (e TL_account_getPrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getPrivacy)
	x.Bytes(e.key.encode())
	return x.buf
}

func (e TL_account_setPrivacy) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_setPrivacy)
	x.Bytes(e.key.encode())
	x.Vector(e.rules)
	return x.buf
}

func (e TL_account_deleteAccount) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_deleteAccount)
	x.String(e.reason)
	return x.buf
}

func (e TL_account_getAccountTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_getAccountTTL)
	return x.buf
}

func (e TL_account_setAccountTTL) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_setAccountTTL)
	x.Bytes(e.ttl.encode())
	return x.buf
}

func (e TL_invokeWithLayer) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_invokeWithLayer)
	x.Int(e.layer)
	x.Bytes(e.query.encode())
	return x.buf
}

func (e TL_contacts_resolveUsername) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_contacts_resolveUsername)
	x.String(e.username)
	return x.buf
}

func (e TL_account_sendChangePhoneCode) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_sendChangePhoneCode)
	x.String(e.phone_number)
	return x.buf
}

func (e TL_account_changePhone) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_changePhone)
	x.String(e.phone_number)
	x.String(e.phone_code_hash)
	x.String(e.phone_code)
	return x.buf
}

func (e TL_messages_getStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getStickers)
	x.String(e.emoticon)
	x.String(e.hash)
	return x.buf
}

func (e TL_messages_getAllStickers) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_messages_getAllStickers)
	x.String(e.hash)
	return x.buf
}

func (e TL_account_updateDeviceLocked) encode() []byte {
	x := NewEncodeBuf(512)
	x.UInt(crc_account_updateDeviceLocked)
	x.Int(e.period)
	return x.buf
}

func (m *DecodeBuf) ObjectGenerated(constructor uint32) (r TL) {
	switch constructor {
	case crc_boolFalse:
		r = TL_boolFalse{}

	case crc_boolTrue:
		r = TL_boolTrue{}

	case crc_error:
		r = TL_error{
			m.Int(),
			m.String(),
		}

	case crc_null:
		r = TL_null{}

	case crc_inputPeerEmpty:
		r = TL_inputPeerEmpty{}

	case crc_inputPeerSelf:
		r = TL_inputPeerSelf{}

	case crc_inputPeerContact:
		r = TL_inputPeerContact{
			m.Int(),
		}

	case crc_inputPeerForeign:
		r = TL_inputPeerForeign{
			m.Int(),
			m.Long(),
		}

	case crc_inputPeerChat:
		r = TL_inputPeerChat{
			m.Int(),
		}

	case crc_inputUserEmpty:
		r = TL_inputUserEmpty{}

	case crc_inputUserSelf:
		r = TL_inputUserSelf{}

	case crc_inputUserContact:
		r = TL_inputUserContact{
			m.Int(),
		}

	case crc_inputUserForeign:
		r = TL_inputUserForeign{
			m.Int(),
			m.Long(),
		}

	case crc_inputPhoneContact:
		r = TL_inputPhoneContact{
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_inputFile:
		r = TL_inputFile{
			m.Long(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case crc_inputMediaEmpty:
		r = TL_inputMediaEmpty{}

	case crc_inputMediaUploadedPhoto:
		r = TL_inputMediaUploadedPhoto{
			m.Object(),
		}

	case crc_inputMediaPhoto:
		r = TL_inputMediaPhoto{
			m.Object(),
		}

	case crc_inputMediaGeoPoint:
		r = TL_inputMediaGeoPoint{
			m.Object(),
		}

	case crc_inputMediaContact:
		r = TL_inputMediaContact{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_inputMediaUploadedVideo:
		r = TL_inputMediaUploadedVideo{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
		}

	case crc_inputMediaUploadedThumbVideo:
		r = TL_inputMediaUploadedThumbVideo{
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
		}

	case crc_inputMediaVideo:
		r = TL_inputMediaVideo{
			m.Object(),
		}

	case crc_inputChatPhotoEmpty:
		r = TL_inputChatPhotoEmpty{}

	case crc_inputChatUploadedPhoto:
		r = TL_inputChatUploadedPhoto{
			m.Object(),
			m.Object(),
		}

	case crc_inputChatPhoto:
		r = TL_inputChatPhoto{
			m.Object(),
			m.Object(),
		}

	case crc_inputGeoPointEmpty:
		r = TL_inputGeoPointEmpty{}

	case crc_inputGeoPoint:
		r = TL_inputGeoPoint{
			m.Double(),
			m.Double(),
		}

	case crc_inputPhotoEmpty:
		r = TL_inputPhotoEmpty{}

	case crc_inputPhoto:
		r = TL_inputPhoto{
			m.Long(),
			m.Long(),
		}

	case crc_inputVideoEmpty:
		r = TL_inputVideoEmpty{}

	case crc_inputVideo:
		r = TL_inputVideo{
			m.Long(),
			m.Long(),
		}

	case crc_inputFileLocation:
		r = TL_inputFileLocation{
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case crc_inputVideoFileLocation:
		r = TL_inputVideoFileLocation{
			m.Long(),
			m.Long(),
		}

	case crc_inputPhotoCropAuto:
		r = TL_inputPhotoCropAuto{}

	case crc_inputPhotoCrop:
		r = TL_inputPhotoCrop{
			m.Double(),
			m.Double(),
			m.Double(),
		}

	case crc_inputAppEvent:
		r = TL_inputAppEvent{
			m.Double(),
			m.String(),
			m.Long(),
			m.String(),
		}

	case crc_peerUser:
		r = TL_peerUser{
			m.Int(),
		}

	case crc_peerChat:
		r = TL_peerChat{
			m.Int(),
		}

	case crc_storage_fileUnknown:
		r = TL_storage_fileUnknown{}

	case crc_storage_fileJpeg:
		r = TL_storage_fileJpeg{}

	case crc_storage_fileGif:
		r = TL_storage_fileGif{}

	case crc_storage_filePng:
		r = TL_storage_filePng{}

	case crc_storage_filePdf:
		r = TL_storage_filePdf{}

	case crc_storage_fileMp3:
		r = TL_storage_fileMp3{}

	case crc_storage_fileMov:
		r = TL_storage_fileMov{}

	case crc_storage_filePartial:
		r = TL_storage_filePartial{}

	case crc_storage_fileMp4:
		r = TL_storage_fileMp4{}

	case crc_storage_fileWebp:
		r = TL_storage_fileWebp{}

	case crc_fileLocationUnavailable:
		r = TL_fileLocationUnavailable{
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case crc_fileLocation:
		r = TL_fileLocation{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Long(),
		}

	case crc_userEmpty:
		r = TL_userEmpty{
			m.Int(),
		}

	case crc_userSelf:
		r = TL_userSelf{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_userContact:
		r = TL_userContact{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.Long(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_userRequest:
		r = TL_userRequest{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.Long(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_userForeign:
		r = TL_userForeign{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.Long(),
			m.Object(),
			m.Object(),
		}

	case crc_userDeleted:
		r = TL_userDeleted{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_userProfilePhotoEmpty:
		r = TL_userProfilePhotoEmpty{}

	case crc_userProfilePhoto:
		r = TL_userProfilePhoto{
			m.Long(),
			m.Object(),
			m.Object(),
		}

	case crc_userStatusEmpty:
		r = TL_userStatusEmpty{}

	case crc_userStatusOnline:
		r = TL_userStatusOnline{
			m.Int(),
		}

	case crc_userStatusOffline:
		r = TL_userStatusOffline{
			m.Int(),
		}

	case crc_chatEmpty:
		r = TL_chatEmpty{
			m.Int(),
		}

	case crc_chat:
		r = TL_chat{
			m.Int(),
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_chatForbidden:
		r = TL_chatForbidden{
			m.Int(),
			m.String(),
			m.Int(),
		}

	case crc_chatFull:
		r = TL_chatFull{
			m.Int(),
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_chatParticipant:
		r = TL_chatParticipant{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_chatParticipantsForbidden:
		r = TL_chatParticipantsForbidden{
			m.Int(),
		}

	case crc_chatParticipants:
		r = TL_chatParticipants{
			m.Int(),
			m.Int(),
			m.Vector(),
			m.Int(),
		}

	case crc_chatPhotoEmpty:
		r = TL_chatPhotoEmpty{}

	case crc_chatPhoto:
		r = TL_chatPhoto{
			m.Object(),
			m.Object(),
		}

	case crc_messageEmpty:
		r = TL_messageEmpty{
			m.Int(),
		}

	case crc_message:
		r = TL_message{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.String(),
			m.Object(),
		}

	case crc_messageForwarded:
		r = TL_messageForwarded{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.String(),
			m.Object(),
		}

	case crc_messageService:
		r = TL_messageService{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Object(),
		}

	case crc_messageMediaEmpty:
		r = TL_messageMediaEmpty{}

	case crc_messageMediaPhoto:
		r = TL_messageMediaPhoto{
			m.Object(),
		}

	case crc_messageMediaVideo:
		r = TL_messageMediaVideo{
			m.Object(),
		}

	case crc_messageMediaGeo:
		r = TL_messageMediaGeo{
			m.Object(),
		}

	case crc_messageMediaContact:
		r = TL_messageMediaContact{
			m.String(),
			m.String(),
			m.String(),
			m.Int(),
		}

	case crc_messageMediaUnsupported:
		r = TL_messageMediaUnsupported{
			m.StringBytes(),
		}

	case crc_messageActionEmpty:
		r = TL_messageActionEmpty{}

	case crc_messageActionChatCreate:
		r = TL_messageActionChatCreate{
			m.String(),
			m.VectorInt(),
		}

	case crc_messageActionChatEditTitle:
		r = TL_messageActionChatEditTitle{
			m.String(),
		}

	case crc_messageActionChatEditPhoto:
		r = TL_messageActionChatEditPhoto{
			m.Object(),
		}

	case crc_messageActionChatDeletePhoto:
		r = TL_messageActionChatDeletePhoto{}

	case crc_messageActionChatAddUser:
		r = TL_messageActionChatAddUser{
			m.Int(),
		}

	case crc_messageActionChatDeleteUser:
		r = TL_messageActionChatDeleteUser{
			m.Int(),
		}

	case crc_dialog:
		r = TL_dialog{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_photoEmpty:
		r = TL_photoEmpty{
			m.Long(),
		}

	case crc_photo:
		r = TL_photo{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.String(),
			m.Object(),
			m.Vector(),
		}

	case crc_photoSizeEmpty:
		r = TL_photoSizeEmpty{
			m.String(),
		}

	case crc_photoSize:
		r = TL_photoSize{
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_photoCachedSize:
		r = TL_photoCachedSize{
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_videoEmpty:
		r = TL_videoEmpty{
			m.Long(),
		}

	case crc_video:
		r = TL_video{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_geoPointEmpty:
		r = TL_geoPointEmpty{}

	case crc_geoPoint:
		r = TL_geoPoint{
			m.Double(),
			m.Double(),
		}

	case crc_auth_checkedPhone:
		r = TL_auth_checkedPhone{
			m.Object(),
			m.Object(),
		}

	case crc_auth_sentCode:
		r = TL_auth_sentCode{
			m.Object(),
			m.String(),
			m.Int(),
			m.Object(),
		}

	case crc_auth_authorization:
		r = TL_auth_authorization{
			m.Int(),
			m.Object(),
		}

	case crc_auth_exportedAuthorization:
		r = TL_auth_exportedAuthorization{
			m.Int(),
			m.StringBytes(),
		}

	case crc_inputNotifyPeer:
		r = TL_inputNotifyPeer{
			m.Object(),
		}

	case crc_inputNotifyUsers:
		r = TL_inputNotifyUsers{}

	case crc_inputNotifyChats:
		r = TL_inputNotifyChats{}

	case crc_inputNotifyAll:
		r = TL_inputNotifyAll{}

	case crc_inputPeerNotifyEventsEmpty:
		r = TL_inputPeerNotifyEventsEmpty{}

	case crc_inputPeerNotifyEventsAll:
		r = TL_inputPeerNotifyEventsAll{}

	case crc_inputPeerNotifySettings:
		r = TL_inputPeerNotifySettings{
			m.Int(),
			m.String(),
			m.Object(),
			m.Int(),
		}

	case crc_peerNotifyEventsEmpty:
		r = TL_peerNotifyEventsEmpty{}

	case crc_peerNotifyEventsAll:
		r = TL_peerNotifyEventsAll{}

	case crc_peerNotifySettingsEmpty:
		r = TL_peerNotifySettingsEmpty{}

	case crc_peerNotifySettings:
		r = TL_peerNotifySettings{
			m.Int(),
			m.String(),
			m.Object(),
			m.Int(),
		}

	case crc_wallPaper:
		r = TL_wallPaper{
			m.Int(),
			m.String(),
			m.Vector(),
			m.Int(),
		}

	case crc_userFull:
		r = TL_userFull{
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.Object(),
			m.String(),
			m.String(),
		}

	case crc_contact:
		r = TL_contact{
			m.Int(),
			m.Object(),
		}

	case crc_importedContact:
		r = TL_importedContact{
			m.Int(),
			m.Long(),
		}

	case crc_contactBlocked:
		r = TL_contactBlocked{
			m.Int(),
			m.Int(),
		}

	case crc_contactSuggested:
		r = TL_contactSuggested{
			m.Int(),
			m.Int(),
		}

	case crc_contactStatus:
		r = TL_contactStatus{
			m.Int(),
			m.Object(),
		}

	case crc_chatLocated:
		r = TL_chatLocated{
			m.Int(),
			m.Int(),
		}

	case crc_contacts_foreignLinkUnknown:
		r = TL_contacts_foreignLinkUnknown{}

	case crc_contacts_foreignLinkRequested:
		r = TL_contacts_foreignLinkRequested{
			m.Object(),
		}

	case crc_contacts_foreignLinkMutual:
		r = TL_contacts_foreignLinkMutual{}

	case crc_contacts_myLinkEmpty:
		r = TL_contacts_myLinkEmpty{}

	case crc_contacts_myLinkRequested:
		r = TL_contacts_myLinkRequested{
			m.Object(),
		}

	case crc_contacts_myLinkContact:
		r = TL_contacts_myLinkContact{}

	case crc_contacts_link:
		r = TL_contacts_link{
			m.Object(),
			m.Object(),
			m.Object(),
		}

	case crc_contacts_contactsNotModified:
		r = TL_contacts_contactsNotModified{}

	case crc_contacts_contacts:
		r = TL_contacts_contacts{
			m.Vector(),
			m.Vector(),
		}

	case crc_contacts_importedContacts:
		r = TL_contacts_importedContacts{
			m.Vector(),
			m.VectorLong(),
			m.Vector(),
		}

	case crc_contacts_blocked:
		r = TL_contacts_blocked{
			m.Vector(),
			m.Vector(),
		}

	case crc_contacts_blockedSlice:
		r = TL_contacts_blockedSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crc_contacts_suggested:
		r = TL_contacts_suggested{
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_dialogs:
		r = TL_messages_dialogs{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_dialogsSlice:
		r = TL_messages_dialogsSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_messages:
		r = TL_messages_messages{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_messagesSlice:
		r = TL_messages_messagesSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_messageEmpty:
		r = TL_messages_messageEmpty{}

	case crc_messages_statedMessages:
		r = TL_messages_statedMessages{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_statedMessage:
		r = TL_messages_statedMessage{
			m.Object(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_sentMessage:
		r = TL_messages_sentMessage{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_chats:
		r = TL_messages_chats{
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_chatFull:
		r = TL_messages_chatFull{
			m.Object(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messages_affectedHistory:
		r = TL_messages_affectedHistory{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputMessagesFilterEmpty:
		r = TL_inputMessagesFilterEmpty{}

	case crc_inputMessagesFilterPhotos:
		r = TL_inputMessagesFilterPhotos{}

	case crc_inputMessagesFilterVideo:
		r = TL_inputMessagesFilterVideo{}

	case crc_inputMessagesFilterPhotoVideo:
		r = TL_inputMessagesFilterPhotoVideo{}

	case crc_inputMessagesFilterPhotoVideoDocuments:
		r = TL_inputMessagesFilterPhotoVideoDocuments{}

	case crc_inputMessagesFilterDocument:
		r = TL_inputMessagesFilterDocument{}

	case crc_inputMessagesFilterAudio:
		r = TL_inputMessagesFilterAudio{}

	case crc_updateNewMessage:
		r = TL_updateNewMessage{
			m.Object(),
			m.Int(),
		}

	case crc_updateMessageID:
		r = TL_updateMessageID{
			m.Int(),
			m.Long(),
		}

	case crc_updateReadMessages:
		r = TL_updateReadMessages{
			m.VectorInt(),
			m.Int(),
		}

	case crc_updateDeleteMessages:
		r = TL_updateDeleteMessages{
			m.VectorInt(),
			m.Int(),
		}

	case crc_updateUserTyping:
		r = TL_updateUserTyping{
			m.Int(),
			m.Object(),
		}

	case crc_updateChatUserTyping:
		r = TL_updateChatUserTyping{
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_updateChatParticipants:
		r = TL_updateChatParticipants{
			m.Object(),
		}

	case crc_updateUserStatus:
		r = TL_updateUserStatus{
			m.Int(),
			m.Object(),
		}

	case crc_updateUserName:
		r = TL_updateUserName{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_updateUserPhoto:
		r = TL_updateUserPhoto{
			m.Int(),
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case crc_updateContactRegistered:
		r = TL_updateContactRegistered{
			m.Int(),
			m.Int(),
		}

	case crc_updateContactLink:
		r = TL_updateContactLink{
			m.Int(),
			m.Object(),
			m.Object(),
		}

	case crc_updateNewAuthorization:
		r = TL_updateNewAuthorization{
			m.Long(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case crc_updates_state:
		r = TL_updates_state{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updates_differenceEmpty:
		r = TL_updates_differenceEmpty{
			m.Int(),
			m.Int(),
		}

	case crc_updates_difference:
		r = TL_updates_difference{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case crc_updates_differenceSlice:
		r = TL_updates_differenceSlice{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Object(),
		}

	case crc_updatesTooLong:
		r = TL_updatesTooLong{}

	case crc_updateShortMessage:
		r = TL_updateShortMessage{
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateShortChatMessage:
		r = TL_updateShortChatMessage{
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateShort:
		r = TL_updateShort{
			m.Object(),
			m.Int(),
		}

	case crc_updatesCombined:
		r = TL_updatesCombined{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updates:
		r = TL_updates{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
		}

	case crc_photos_photos:
		r = TL_photos_photos{
			m.Vector(),
			m.Vector(),
		}

	case crc_photos_photosSlice:
		r = TL_photos_photosSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
		}

	case crc_photos_photo:
		r = TL_photos_photo{
			m.Object(),
			m.Vector(),
		}

	case crc_upload_file:
		r = TL_upload_file{
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_dcOption:
		r = TL_dcOption{
			m.Int(),
			m.String(),
			m.String(),
			m.Int(),
		}

	case crc_config:
		r = TL_config{
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Vector(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Vector(),
		}

	case crc_nearestDc:
		r = TL_nearestDc{
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_help_appUpdate:
		r = TL_help_appUpdate{
			m.Int(),
			m.Object(),
			m.String(),
			m.String(),
		}

	case crc_help_noAppUpdate:
		r = TL_help_noAppUpdate{}

	case crc_help_inviteText:
		r = TL_help_inviteText{
			m.String(),
		}

	case crc_messages_statedMessagesLinks:
		r = TL_messages_statedMessagesLinks{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_statedMessageLink:
		r = TL_messages_statedMessageLink{
			m.Object(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_sentMessageLink:
		r = TL_messages_sentMessageLink{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Vector(),
		}

	case crc_inputGeoChat:
		r = TL_inputGeoChat{
			m.Int(),
			m.Long(),
		}

	case crc_inputNotifyGeoChatPeer:
		r = TL_inputNotifyGeoChatPeer{
			m.Object(),
		}

	case crc_geoChat:
		r = TL_geoChat{
			m.Int(),
			m.Long(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_geoChatMessageEmpty:
		r = TL_geoChatMessageEmpty{
			m.Int(),
			m.Int(),
		}

	case crc_geoChatMessage:
		r = TL_geoChatMessage{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
			m.Object(),
		}

	case crc_geoChatMessageService:
		r = TL_geoChatMessageService{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_geochats_statedMessage:
		r = TL_geochats_statedMessage{
			m.Object(),
			m.Vector(),
			m.Vector(),
			m.Int(),
		}

	case crc_geochats_located:
		r = TL_geochats_located{
			m.Vector(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_geochats_messages:
		r = TL_geochats_messages{
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_geochats_messagesSlice:
		r = TL_geochats_messagesSlice{
			m.Int(),
			m.Vector(),
			m.Vector(),
			m.Vector(),
		}

	case crc_messageActionGeoChatCreate:
		r = TL_messageActionGeoChatCreate{
			m.String(),
			m.String(),
		}

	case crc_messageActionGeoChatCheckin:
		r = TL_messageActionGeoChatCheckin{}

	case crc_updateNewGeoChatMessage:
		r = TL_updateNewGeoChatMessage{
			m.Object(),
		}

	case crc_wallPaperSolid:
		r = TL_wallPaperSolid{
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_updateNewEncryptedMessage:
		r = TL_updateNewEncryptedMessage{
			m.Object(),
			m.Int(),
		}

	case crc_updateEncryptedChatTyping:
		r = TL_updateEncryptedChatTyping{
			m.Int(),
		}

	case crc_updateEncryption:
		r = TL_updateEncryption{
			m.Object(),
			m.Int(),
		}

	case crc_updateEncryptedMessagesRead:
		r = TL_updateEncryptedMessagesRead{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_encryptedChatEmpty:
		r = TL_encryptedChatEmpty{
			m.Int(),
		}

	case crc_encryptedChatWaiting:
		r = TL_encryptedChatWaiting{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_encryptedChatRequested:
		r = TL_encryptedChatRequested{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_encryptedChat:
		r = TL_encryptedChat{
			m.Int(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Long(),
		}

	case crc_encryptedChatDiscarded:
		r = TL_encryptedChatDiscarded{
			m.Int(),
		}

	case crc_inputEncryptedChat:
		r = TL_inputEncryptedChat{
			m.Int(),
			m.Long(),
		}

	case crc_encryptedFileEmpty:
		r = TL_encryptedFileEmpty{}

	case crc_encryptedFile:
		r = TL_encryptedFile{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_inputEncryptedFileEmpty:
		r = TL_inputEncryptedFileEmpty{}

	case crc_inputEncryptedFileUploaded:
		r = TL_inputEncryptedFileUploaded{
			m.Long(),
			m.Int(),
			m.String(),
			m.Int(),
		}

	case crc_inputEncryptedFile:
		r = TL_inputEncryptedFile{
			m.Long(),
			m.Long(),
		}

	case crc_inputEncryptedFileLocation:
		r = TL_inputEncryptedFileLocation{
			m.Long(),
			m.Long(),
		}

	case crc_encryptedMessage:
		r = TL_encryptedMessage{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_encryptedMessageService:
		r = TL_encryptedMessageService{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messages_dhConfigNotModified:
		r = TL_messages_dhConfigNotModified{
			m.StringBytes(),
		}

	case crc_messages_dhConfig:
		r = TL_messages_dhConfig{
			m.Int(),
			m.StringBytes(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messages_sentEncryptedMessage:
		r = TL_messages_sentEncryptedMessage{
			m.Int(),
		}

	case crc_messages_sentEncryptedFile:
		r = TL_messages_sentEncryptedFile{
			m.Int(),
			m.Object(),
		}

	case crc_inputFileBig:
		r = TL_inputFileBig{
			m.Long(),
			m.Int(),
			m.String(),
		}

	case crc_inputEncryptedFileBigUploaded:
		r = TL_inputEncryptedFileBigUploaded{
			m.Long(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChatParticipantAdd:
		r = TL_updateChatParticipantAdd{
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateChatParticipantDelete:
		r = TL_updateChatParticipantDelete{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_updateDcOptions:
		r = TL_updateDcOptions{
			m.Vector(),
		}

	case crc_inputMediaUploadedAudio:
		r = TL_inputMediaUploadedAudio{
			m.Object(),
			m.Int(),
			m.String(),
		}

	case crc_inputMediaAudio:
		r = TL_inputMediaAudio{
			m.Object(),
		}

	case crc_inputMediaUploadedDocument:
		r = TL_inputMediaUploadedDocument{
			m.Object(),
			m.String(),
			m.Vector(),
		}

	case crc_inputMediaUploadedThumbDocument:
		r = TL_inputMediaUploadedThumbDocument{
			m.Object(),
			m.Object(),
			m.String(),
			m.Vector(),
		}

	case crc_inputMediaDocument:
		r = TL_inputMediaDocument{
			m.Object(),
		}

	case crc_messageMediaDocument:
		r = TL_messageMediaDocument{
			m.Object(),
		}

	case crc_messageMediaAudio:
		r = TL_messageMediaAudio{
			m.Object(),
		}

	case crc_inputAudioEmpty:
		r = TL_inputAudioEmpty{}

	case crc_inputAudio:
		r = TL_inputAudio{
			m.Long(),
			m.Long(),
		}

	case crc_inputDocumentEmpty:
		r = TL_inputDocumentEmpty{}

	case crc_inputDocument:
		r = TL_inputDocument{
			m.Long(),
			m.Long(),
		}

	case crc_inputAudioFileLocation:
		r = TL_inputAudioFileLocation{
			m.Long(),
			m.Long(),
		}

	case crc_inputDocumentFileLocation:
		r = TL_inputDocumentFileLocation{
			m.Long(),
			m.Long(),
		}

	case crc_audioEmpty:
		r = TL_audioEmpty{
			m.Long(),
		}

	case crc_audio:
		r = TL_audio{
			m.Long(),
			m.Long(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Int(),
		}

	case crc_documentEmpty:
		r = TL_documentEmpty{
			m.Long(),
		}

	case crc_document:
		r = TL_document{
			m.Long(),
			m.Long(),
			m.Int(),
			m.String(),
			m.Int(),
			m.Object(),
			m.Int(),
			m.Vector(),
		}

	case crc_help_support:
		r = TL_help_support{
			m.String(),
			m.Object(),
		}

	case crc_notifyPeer:
		r = TL_notifyPeer{
			m.Object(),
		}

	case crc_notifyUsers:
		r = TL_notifyUsers{}

	case crc_notifyChats:
		r = TL_notifyChats{}

	case crc_notifyAll:
		r = TL_notifyAll{}

	case crc_updateUserBlocked:
		r = TL_updateUserBlocked{
			m.Int(),
			m.Object(),
		}

	case crc_updateNotifySettings:
		r = TL_updateNotifySettings{
			m.Object(),
			m.Object(),
		}

	case crc_auth_sentAppCode:
		r = TL_auth_sentAppCode{
			m.Object(),
			m.String(),
			m.Int(),
			m.Object(),
		}

	case crc_sendMessageTypingAction:
		r = TL_sendMessageTypingAction{}

	case crc_sendMessageCancelAction:
		r = TL_sendMessageCancelAction{}

	case crc_sendMessageRecordVideoAction:
		r = TL_sendMessageRecordVideoAction{}

	case crc_sendMessageUploadVideoAction:
		r = TL_sendMessageUploadVideoAction{}

	case crc_sendMessageRecordAudioAction:
		r = TL_sendMessageRecordAudioAction{}

	case crc_sendMessageUploadAudioAction:
		r = TL_sendMessageUploadAudioAction{}

	case crc_sendMessageUploadPhotoAction:
		r = TL_sendMessageUploadPhotoAction{}

	case crc_sendMessageUploadDocumentAction:
		r = TL_sendMessageUploadDocumentAction{}

	case crc_sendMessageGeoLocationAction:
		r = TL_sendMessageGeoLocationAction{}

	case crc_sendMessageChooseContactAction:
		r = TL_sendMessageChooseContactAction{}

	case crc_contactFound:
		r = TL_contactFound{
			m.Int(),
		}

	case crc_contacts_found:
		r = TL_contacts_found{
			m.Vector(),
			m.Vector(),
		}

	case crc_updateServiceNotification:
		r = TL_updateServiceNotification{
			m.String(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_userStatusRecently:
		r = TL_userStatusRecently{}

	case crc_userStatusLastWeek:
		r = TL_userStatusLastWeek{}

	case crc_userStatusLastMonth:
		r = TL_userStatusLastMonth{}

	case crc_updatePrivacy:
		r = TL_updatePrivacy{
			m.Object(),
			m.Vector(),
		}

	case crc_inputPrivacyKeyStatusTimestamp:
		r = TL_inputPrivacyKeyStatusTimestamp{}

	case crc_privacyKeyStatusTimestamp:
		r = TL_privacyKeyStatusTimestamp{}

	case crc_inputPrivacyValueAllowContacts:
		r = TL_inputPrivacyValueAllowContacts{}

	case crc_inputPrivacyValueAllowAll:
		r = TL_inputPrivacyValueAllowAll{}

	case crc_inputPrivacyValueAllowUsers:
		r = TL_inputPrivacyValueAllowUsers{
			m.Vector(),
		}

	case crc_inputPrivacyValueDisallowContacts:
		r = TL_inputPrivacyValueDisallowContacts{}

	case crc_inputPrivacyValueDisallowAll:
		r = TL_inputPrivacyValueDisallowAll{}

	case crc_inputPrivacyValueDisallowUsers:
		r = TL_inputPrivacyValueDisallowUsers{
			m.Vector(),
		}

	case crc_privacyValueAllowContacts:
		r = TL_privacyValueAllowContacts{}

	case crc_privacyValueAllowAll:
		r = TL_privacyValueAllowAll{}

	case crc_privacyValueAllowUsers:
		r = TL_privacyValueAllowUsers{
			m.VectorInt(),
		}

	case crc_privacyValueDisallowContacts:
		r = TL_privacyValueDisallowContacts{}

	case crc_privacyValueDisallowAll:
		r = TL_privacyValueDisallowAll{}

	case crc_privacyValueDisallowUsers:
		r = TL_privacyValueDisallowUsers{
			m.VectorInt(),
		}

	case crc_account_privacyRules:
		r = TL_account_privacyRules{
			m.Vector(),
			m.Vector(),
		}

	case crc_accountDaysTTL:
		r = TL_accountDaysTTL{
			m.Int(),
		}

	case crc_account_sentChangePhoneCode:
		r = TL_account_sentChangePhoneCode{
			m.String(),
			m.Int(),
		}

	case crc_updateUserPhone:
		r = TL_updateUserPhone{
			m.Int(),
			m.String(),
		}

	case crc_documentAttributeImageSize:
		r = TL_documentAttributeImageSize{
			m.Int(),
			m.Int(),
		}

	case crc_documentAttributeAnimated:
		r = TL_documentAttributeAnimated{}

	case crc_documentAttributeSticker:
		r = TL_documentAttributeSticker{}

	case crc_documentAttributeVideo:
		r = TL_documentAttributeVideo{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_documentAttributeAudio:
		r = TL_documentAttributeAudio{
			m.Int(),
		}

	case crc_documentAttributeFilename:
		r = TL_documentAttributeFilename{
			m.String(),
		}

	case crc_messages_stickersNotModified:
		r = TL_messages_stickersNotModified{}

	case crc_messages_stickers:
		r = TL_messages_stickers{
			m.String(),
			m.Vector(),
		}

	case crc_stickerPack:
		r = TL_stickerPack{
			m.String(),
			m.VectorLong(),
		}

	case crc_messages_allStickersNotModified:
		r = TL_messages_allStickersNotModified{}

	case crc_messages_allStickers:
		r = TL_messages_allStickers{
			m.String(),
			m.Vector(),
			m.Vector(),
		}

	case crc_disabledFeature:
		r = TL_disabledFeature{
			m.String(),
			m.String(),
		}

	case crc_invokeAfterMsg:
		r = TL_invokeAfterMsg{
			m.Long(),
			m.Object(),
		}

	case crc_invokeAfterMsgs:
		r = TL_invokeAfterMsgs{
			m.VectorLong(),
			m.Object(),
		}

	case crc_auth_checkPhone:
		r = TL_auth_checkPhone{
			m.String(),
		}

	case crc_auth_sendCode:
		r = TL_auth_sendCode{
			m.String(),
			m.Int(),
			m.Int(),
			m.String(),
			m.String(),
		}

	case crc_auth_sendCall:
		r = TL_auth_sendCall{
			m.String(),
			m.String(),
		}

	case crc_auth_signUp:
		r = TL_auth_signUp{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_auth_signIn:
		r = TL_auth_signIn{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_auth_logOut:
		r = TL_auth_logOut{}

	case crc_auth_resetAuthorizations:
		r = TL_auth_resetAuthorizations{}

	case crc_auth_sendInvites:
		r = TL_auth_sendInvites{
			m.VectorString(),
			m.String(),
		}

	case crc_auth_exportAuthorization:
		r = TL_auth_exportAuthorization{
			m.Int(),
		}

	case crc_auth_importAuthorization:
		r = TL_auth_importAuthorization{
			m.Int(),
			m.StringBytes(),
		}

	case crc_auth_bindTempAuthKey:
		r = TL_auth_bindTempAuthKey{
			m.Long(),
			m.Long(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_account_registerDevice:
		r = TL_account_registerDevice{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
			m.String(),
		}

	case crc_account_unregisterDevice:
		r = TL_account_unregisterDevice{
			m.Int(),
			m.String(),
		}

	case crc_account_updateNotifySettings:
		r = TL_account_updateNotifySettings{
			m.Object(),
			m.Object(),
		}

	case crc_account_getNotifySettings:
		r = TL_account_getNotifySettings{
			m.Object(),
		}

	case crc_account_resetNotifySettings:
		r = TL_account_resetNotifySettings{}

	case crc_account_updateProfile:
		r = TL_account_updateProfile{
			m.String(),
			m.String(),
		}

	case crc_account_updateStatus:
		r = TL_account_updateStatus{
			m.Object(),
		}

	case crc_account_getWallPapers:
		r = TL_account_getWallPapers{}

	case crc_users_getUsers:
		r = TL_users_getUsers{
			m.Vector(),
		}

	case crc_users_getFullUser:
		r = TL_users_getFullUser{
			m.Object(),
		}

	case crc_contacts_getStatuses:
		r = TL_contacts_getStatuses{}

	case crc_contacts_getContacts:
		r = TL_contacts_getContacts{
			m.String(),
		}

	case crc_contacts_importContacts:
		r = TL_contacts_importContacts{
			m.Vector(),
			m.Object(),
		}

	case crc_contacts_getSuggested:
		r = TL_contacts_getSuggested{
			m.Int(),
		}

	case crc_contacts_deleteContact:
		r = TL_contacts_deleteContact{
			m.Object(),
		}

	case crc_contacts_deleteContacts:
		r = TL_contacts_deleteContacts{
			m.Vector(),
		}

	case crc_contacts_block:
		r = TL_contacts_block{
			m.Object(),
		}

	case crc_contacts_unblock:
		r = TL_contacts_unblock{
			m.Object(),
		}

	case crc_contacts_getBlocked:
		r = TL_contacts_getBlocked{
			m.Int(),
			m.Int(),
		}

	case crc_contacts_exportCard:
		r = TL_contacts_exportCard{}

	case crc_contacts_importCard:
		r = TL_contacts_importCard{
			m.VectorInt(),
		}

	case crc_messages_getMessages:
		r = TL_messages_getMessages{
			m.VectorInt(),
		}

	case crc_messages_getDialogs:
		r = TL_messages_getDialogs{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_getHistory:
		r = TL_messages_getHistory{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_search:
		r = TL_messages_search{
			m.Object(),
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_readHistory:
		r = TL_messages_readHistory{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Object(),
		}

	case crc_messages_deleteHistory:
		r = TL_messages_deleteHistory{
			m.Object(),
			m.Int(),
		}

	case crc_messages_deleteMessages:
		r = TL_messages_deleteMessages{
			m.VectorInt(),
		}

	case crc_messages_receivedMessages:
		r = TL_messages_receivedMessages{
			m.Int(),
		}

	case crc_messages_setTyping:
		r = TL_messages_setTyping{
			m.Object(),
			m.Object(),
		}

	case crc_messages_sendMessage:
		r = TL_messages_sendMessage{
			m.Object(),
			m.String(),
			m.Long(),
		}

	case crc_messages_sendMedia:
		r = TL_messages_sendMedia{
			m.Object(),
			m.Object(),
			m.Long(),
		}

	case crc_messages_forwardMessages:
		r = TL_messages_forwardMessages{
			m.Object(),
			m.VectorInt(),
		}

	case crc_messages_getChats:
		r = TL_messages_getChats{
			m.VectorInt(),
		}

	case crc_messages_getFullChat:
		r = TL_messages_getFullChat{
			m.Int(),
		}

	case crc_messages_editChatTitle:
		r = TL_messages_editChatTitle{
			m.Int(),
			m.String(),
		}

	case crc_messages_editChatPhoto:
		r = TL_messages_editChatPhoto{
			m.Int(),
			m.Object(),
		}

	case crc_messages_addChatUser:
		r = TL_messages_addChatUser{
			m.Int(),
			m.Object(),
			m.Int(),
		}

	case crc_messages_deleteChatUser:
		r = TL_messages_deleteChatUser{
			m.Int(),
			m.Object(),
		}

	case crc_messages_createChat:
		r = TL_messages_createChat{
			m.Vector(),
			m.String(),
		}

	case crc_updates_getState:
		r = TL_updates_getState{}

	case crc_updates_getDifference:
		r = TL_updates_getDifference{
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_photos_updateProfilePhoto:
		r = TL_photos_updateProfilePhoto{
			m.Object(),
			m.Object(),
		}

	case crc_photos_uploadProfilePhoto:
		r = TL_photos_uploadProfilePhoto{
			m.Object(),
			m.String(),
			m.Object(),
			m.Object(),
		}

	case crc_photos_deletePhotos:
		r = TL_photos_deletePhotos{
			m.Vector(),
		}

	case crc_upload_saveFilePart:
		r = TL_upload_saveFilePart{
			m.Long(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_upload_getFile:
		r = TL_upload_getFile{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_help_getConfig:
		r = TL_help_getConfig{}

	case crc_help_getNearestDc:
		r = TL_help_getNearestDc{}

	case crc_help_getAppUpdate:
		r = TL_help_getAppUpdate{
			m.String(),
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_help_saveAppLog:
		r = TL_help_saveAppLog{
			m.Vector(),
		}

	case crc_help_getInviteText:
		r = TL_help_getInviteText{
			m.String(),
		}

	case crc_photos_getUserPhotos:
		r = TL_photos_getUserPhotos{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_messages_forwardMessage:
		r = TL_messages_forwardMessage{
			m.Object(),
			m.Int(),
			m.Long(),
		}

	case crc_messages_sendBroadcast:
		r = TL_messages_sendBroadcast{
			m.Vector(),
			m.String(),
			m.Object(),
		}

	case crc_geochats_getLocated:
		r = TL_geochats_getLocated{
			m.Object(),
			m.Int(),
			m.Int(),
		}

	case crc_geochats_getRecents:
		r = TL_geochats_getRecents{
			m.Int(),
			m.Int(),
		}

	case crc_geochats_checkin:
		r = TL_geochats_checkin{
			m.Object(),
		}

	case crc_geochats_getFullChat:
		r = TL_geochats_getFullChat{
			m.Object(),
		}

	case crc_geochats_editChatTitle:
		r = TL_geochats_editChatTitle{
			m.Object(),
			m.String(),
			m.String(),
		}

	case crc_geochats_editChatPhoto:
		r = TL_geochats_editChatPhoto{
			m.Object(),
			m.Object(),
		}

	case crc_geochats_search:
		r = TL_geochats_search{
			m.Object(),
			m.String(),
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_geochats_getHistory:
		r = TL_geochats_getHistory{
			m.Object(),
			m.Int(),
			m.Int(),
			m.Int(),
		}

	case crc_geochats_setTyping:
		r = TL_geochats_setTyping{
			m.Object(),
			m.Object(),
		}

	case crc_geochats_sendMessage:
		r = TL_geochats_sendMessage{
			m.Object(),
			m.String(),
			m.Long(),
		}

	case crc_geochats_sendMedia:
		r = TL_geochats_sendMedia{
			m.Object(),
			m.Object(),
			m.Long(),
		}

	case crc_geochats_createGeoChat:
		r = TL_geochats_createGeoChat{
			m.String(),
			m.Object(),
			m.String(),
			m.String(),
		}

	case crc_messages_getDhConfig:
		r = TL_messages_getDhConfig{
			m.Int(),
			m.Int(),
		}

	case crc_messages_requestEncryption:
		r = TL_messages_requestEncryption{
			m.Object(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_messages_acceptEncryption:
		r = TL_messages_acceptEncryption{
			m.Object(),
			m.StringBytes(),
			m.Long(),
		}

	case crc_messages_discardEncryption:
		r = TL_messages_discardEncryption{
			m.Int(),
		}

	case crc_messages_setEncryptedTyping:
		r = TL_messages_setEncryptedTyping{
			m.Object(),
			m.Object(),
		}

	case crc_messages_readEncryptedHistory:
		r = TL_messages_readEncryptedHistory{
			m.Object(),
			m.Int(),
		}

	case crc_messages_sendEncrypted:
		r = TL_messages_sendEncrypted{
			m.Object(),
			m.Long(),
			m.StringBytes(),
		}

	case crc_messages_sendEncryptedFile:
		r = TL_messages_sendEncryptedFile{
			m.Object(),
			m.Long(),
			m.StringBytes(),
			m.Object(),
		}

	case crc_messages_sendEncryptedService:
		r = TL_messages_sendEncryptedService{
			m.Object(),
			m.Long(),
			m.StringBytes(),
		}

	case crc_messages_receivedQueue:
		r = TL_messages_receivedQueue{
			m.Int(),
		}

	case crc_upload_saveBigFilePart:
		r = TL_upload_saveBigFilePart{
			m.Long(),
			m.Int(),
			m.Int(),
			m.StringBytes(),
		}

	case crc_initConnection:
		r = TL_initConnection{
			m.Int(),
			m.String(),
			m.String(),
			m.String(),
			m.String(),
			m.Object(),
		}

	case crc_help_getSupport:
		r = TL_help_getSupport{}

	case crc_auth_sendSms:
		r = TL_auth_sendSms{
			m.String(),
			m.String(),
		}

	case crc_messages_readMessageContents:
		r = TL_messages_readMessageContents{
			m.VectorInt(),
		}

	case crc_account_checkUsername:
		r = TL_account_checkUsername{
			m.String(),
		}

	case crc_account_updateUsername:
		r = TL_account_updateUsername{
			m.String(),
		}

	case crc_contacts_search:
		r = TL_contacts_search{
			m.String(),
			m.Int(),
		}

	case crc_account_getPrivacy:
		r = TL_account_getPrivacy{
			m.Object(),
		}

	case crc_account_setPrivacy:
		r = TL_account_setPrivacy{
			m.Object(),
			m.Vector(),
		}

	case crc_account_deleteAccount:
		r = TL_account_deleteAccount{
			m.String(),
		}

	case crc_account_getAccountTTL:
		r = TL_account_getAccountTTL{}

	case crc_account_setAccountTTL:
		r = TL_account_setAccountTTL{
			m.Object(),
		}

	case crc_invokeWithLayer:
		r = TL_invokeWithLayer{
			m.Int(),
			m.Object(),
		}

	case crc_contacts_resolveUsername:
		r = TL_contacts_resolveUsername{
			m.String(),
		}

	case crc_account_sendChangePhoneCode:
		r = TL_account_sendChangePhoneCode{
			m.String(),
		}

	case crc_account_changePhone:
		r = TL_account_changePhone{
			m.String(),
			m.String(),
			m.String(),
		}

	case crc_messages_getStickers:
		r = TL_messages_getStickers{
			m.String(),
			m.String(),
		}

	case crc_messages_getAllStickers:
		r = TL_messages_getAllStickers{
			m.String(),
		}

	case crc_account_updateDeviceLocked:
		r = TL_account_updateDeviceLocked{
			m.Int(),
		}

	default:
		m.err = fmt.Errorf("Unknown constructor: %08x", constructor)
		return nil

	}

	if m.err != nil {
		return nil
	}

	return
}
