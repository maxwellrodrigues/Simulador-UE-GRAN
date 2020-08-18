package ngapType

// Need to import "free5gc/lib/aper" if it uses "aper"

type ProtocolIEID struct {
	Value int64 `aper:"valueLB:0,valueUB:65535"`
}

const ProtocolIEIDAllowedNSSAI int64 = 0
const ProtocolIEIDAMFName int64 = 1
const ProtocolIEIDAMFOverloadResponse int64 = 2
const ProtocolIEIDAMFSetID int64 = 3
const ProtocolIEIDAMFTNLAssociationFailedToSetupList int64 = 4
const ProtocolIEIDAMFTNLAssociationSetupList int64 = 5
const ProtocolIEIDAMFTNLAssociationToAddList int64 = 6
const ProtocolIEIDAMFTNLAssociationToRemoveList int64 = 7
const ProtocolIEIDAMFTNLAssociationToUpdateList int64 = 8
const ProtocolIEIDAMFTrafficLoadReductionIndication int64 = 9
const ProtocolIEIDAMFUENGAPID int64 = 10
const ProtocolIEIDAssistanceDataForPaging int64 = 11
const ProtocolIEIDBroadcastCancelledAreaList int64 = 12
const ProtocolIEIDBroadcastCompletedAreaList int64 = 13
const ProtocolIEIDCancelAllWarningMessages int64 = 14
const ProtocolIEIDCause int64 = 15
const ProtocolIEIDCellIDListForRestart int64 = 16
const ProtocolIEIDConcurrentWarningMessageInd int64 = 17
const ProtocolIEIDCoreNetworkAssistanceInformation int64 = 18
const ProtocolIEIDCriticalityDiagnostics int64 = 19
const ProtocolIEIDDataCodingScheme int64 = 20
const ProtocolIEIDDefaultPagingDRX int64 = 21
const ProtocolIEIDDirectForwardingPathAvailability int64 = 22
const ProtocolIEIDEmergencyAreaIDListForRestart int64 = 23
const ProtocolIEIDEmergencyFallbackIndicator int64 = 24
const ProtocolIEIDEUTRACGI int64 = 25
const ProtocolIEIDFiveGSTMSI int64 = 26
const ProtocolIEIDGlobalRANNodeID int64 = 27
const ProtocolIEIDGUAMI int64 = 28
const ProtocolIEIDHandoverType int64 = 29
const ProtocolIEIDIMSVoiceSupportIndicator int64 = 30
const ProtocolIEIDIndexToRFSP int64 = 31
const ProtocolIEIDInfoOnRecommendedCellsAndRANNodesForPaging int64 = 32
const ProtocolIEIDLocationReportingRequestType int64 = 33
const ProtocolIEIDMaskedIMEISV int64 = 34
const ProtocolIEIDMessageIdentifier int64 = 35
const ProtocolIEIDMobilityRestrictionList int64 = 36
const ProtocolIEIDNASC int64 = 37
const ProtocolIEIDNASPDU int64 = 38
const ProtocolIEIDNASSecurityParametersFromNGRAN int64 = 39
const ProtocolIEIDNewAMFUENGAPID int64 = 40
const ProtocolIEIDNewSecurityContextInd int64 = 41
const ProtocolIEIDNGAPMessage int64 = 42
const ProtocolIEIDNGRANCGI int64 = 43
const ProtocolIEIDNGRANTraceID int64 = 44
const ProtocolIEIDNRCGI int64 = 45
const ProtocolIEIDNRPPaPDU int64 = 46
const ProtocolIEIDNumberOfBroadcastsRequested int64 = 47
const ProtocolIEIDOldAMF int64 = 48
const ProtocolIEIDOverloadStartNSSAIList int64 = 49
const ProtocolIEIDPagingDRX int64 = 50
const ProtocolIEIDPagingOrigin int64 = 51
const ProtocolIEIDPagingPriority int64 = 52
const ProtocolIEIDPDUSessionResourceAdmittedList int64 = 53
const ProtocolIEIDPDUSessionResourceFailedToModifyListModRes int64 = 54
const ProtocolIEIDPDUSessionResourceFailedToSetupListCxtRes int64 = 55
const ProtocolIEIDPDUSessionResourceFailedToSetupListHOAck int64 = 56
const ProtocolIEIDPDUSessionResourceFailedToSetupListPSReq int64 = 57
const ProtocolIEIDPDUSessionResourceFailedToSetupListSURes int64 = 58
const ProtocolIEIDPDUSessionResourceHandoverList int64 = 59
const ProtocolIEIDPDUSessionResourceListCxtRelCpl int64 = 60
const ProtocolIEIDPDUSessionResourceListHORqd int64 = 61
const ProtocolIEIDPDUSessionResourceModifyListModCfm int64 = 62
const ProtocolIEIDPDUSessionResourceModifyListModInd int64 = 63
const ProtocolIEIDPDUSessionResourceModifyListModReq int64 = 64
const ProtocolIEIDPDUSessionResourceModifyListModRes int64 = 65
const ProtocolIEIDPDUSessionResourceNotifyList int64 = 66
const ProtocolIEIDPDUSessionResourceReleasedListNot int64 = 67
const ProtocolIEIDPDUSessionResourceReleasedListPSAck int64 = 68
const ProtocolIEIDPDUSessionResourceReleasedListPSFail int64 = 69
const ProtocolIEIDPDUSessionResourceReleasedListRelRes int64 = 70
const ProtocolIEIDPDUSessionResourceSetupListCxtReq int64 = 71
const ProtocolIEIDPDUSessionResourceSetupListCxtRes int64 = 72
const ProtocolIEIDPDUSessionResourceSetupListHOReq int64 = 73
const ProtocolIEIDPDUSessionResourceSetupListSUReq int64 = 74
const ProtocolIEIDPDUSessionResourceSetupListSURes int64 = 75
const ProtocolIEIDPDUSessionResourceToBeSwitchedDLList int64 = 76
const ProtocolIEIDPDUSessionResourceSwitchedList int64 = 77
const ProtocolIEIDPDUSessionResourceToReleaseListHOCmd int64 = 78
const ProtocolIEIDPDUSessionResourceToReleaseListRelCmd int64 = 79
const ProtocolIEIDPLMNSupportList int64 = 80
const ProtocolIEIDPWSFailedCellIDList int64 = 81
const ProtocolIEIDRANNodeName int64 = 82
const ProtocolIEIDRANPagingPriority int64 = 83
const ProtocolIEIDRANStatusTransferTransparentContainer int64 = 84
const ProtocolIEIDRANUENGAPID int64 = 85
const ProtocolIEIDRelativeAMFCapacity int64 = 86
const ProtocolIEIDRepetitionPeriod int64 = 87
const ProtocolIEIDResetType int64 = 88
const ProtocolIEIDRoutingID int64 = 89
const ProtocolIEIDRRCEstablishmentCause int64 = 90
const ProtocolIEIDRRCInactiveTransitionReportRequest int64 = 91
const ProtocolIEIDRRCState int64 = 92
const ProtocolIEIDSecurityContext int64 = 93
const ProtocolIEIDSecurityKey int64 = 94
const ProtocolIEIDSerialNumber int64 = 95
const ProtocolIEIDServedGUAMIList int64 = 96
const ProtocolIEIDSliceSupportList int64 = 97
const ProtocolIEIDSONConfigurationTransferDL int64 = 98
const ProtocolIEIDSONConfigurationTransferUL int64 = 99
const ProtocolIEIDSourceAMFUENGAPID int64 = 100
const ProtocolIEIDSourceToTargetTransparentContainer int64 = 101
const ProtocolIEIDSupportedTAList int64 = 102
const ProtocolIEIDTAIListForPaging int64 = 103
const ProtocolIEIDTAIListForRestart int64 = 104
const ProtocolIEIDTargetID int64 = 105
const ProtocolIEIDTargetToSourceTransparentContainer int64 = 106
const ProtocolIEIDTimeToWait int64 = 107
const ProtocolIEIDTraceActivation int64 = 108
const ProtocolIEIDTraceCollectionEntityIPAddress int64 = 109
const ProtocolIEIDUEAggregateMaximumBitRate int64 = 110
const ProtocolIEIDUEAssociatedLogicalNGConnectionList int64 = 111
const ProtocolIEIDUEContextRequest int64 = 112
const ProtocolIEIDUENGAPIDs int64 = 114
const ProtocolIEIDUEPagingIdentity int64 = 115
const ProtocolIEIDUEPresenceInAreaOfInterestList int64 = 116
const ProtocolIEIDUERadioCapability int64 = 117
const ProtocolIEIDUERadioCapabilityForPaging int64 = 118
const ProtocolIEIDUESecurityCapabilities int64 = 119
const ProtocolIEIDUnavailableGUAMIList int64 = 120
const ProtocolIEIDUserLocationInformation int64 = 121
const ProtocolIEIDWarningAreaList int64 = 122
const ProtocolIEIDWarningMessageContents int64 = 123
const ProtocolIEIDWarningSecurityInfo int64 = 124
const ProtocolIEIDWarningType int64 = 125
const ProtocolIEIDAdditionalULNGUUPTNLInformation int64 = 126
const ProtocolIEIDDataForwardingNotPossible int64 = 127
const ProtocolIEIDDLNGUUPTNLInformation int64 = 128
const ProtocolIEIDNetworkInstance int64 = 129
const ProtocolIEIDPDUSessionAggregateMaximumBitRate int64 = 130
const ProtocolIEIDPDUSessionResourceFailedToModifyListModCfm int64 = 131
const ProtocolIEIDPDUSessionResourceFailedToSetupListCxtFail int64 = 132
const ProtocolIEIDPDUSessionResourceListCxtRelReq int64 = 133
const ProtocolIEIDPDUSessionType int64 = 134
const ProtocolIEIDQosFlowAddOrModifyRequestList int64 = 135
const ProtocolIEIDQosFlowSetupRequestList int64 = 136
const ProtocolIEIDQosFlowToReleaseList int64 = 137
const ProtocolIEIDSecurityIndication int64 = 138
const ProtocolIEIDULNGUUPTNLInformation int64 = 139
const ProtocolIEIDULNGUUPTNLModifyList int64 = 140
const ProtocolIEIDWarningAreaCoordinates int64 = 141