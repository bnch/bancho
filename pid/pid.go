// Package pid is simply a list of constants being the bancho commands.
package pid

// These are all the commands us humans are currently aware of.
// (as a general rule: packet ID = line number - 7)
const (
	OsuSendUserState              = iota // update osu about the user state
	OsuSendIRCMessage                    // receive message from IRC
	OsuExit                              // osu closes
	OsuRequestStatusUpdate               // update player stats
	OsuPong                              // ping callback
	BanchoLoginReply                     // login
	BanchoCommandError                   // reply to an error
	BanchoSendMessage                    // i have no idea how is this handled
	BanchoPing                           // ping request
	BanchoHandleIRCUsernameChange        // someone changes name in irc
	BanchoHandleIRCQuit                  // someone logged out
	BanchoHandleUserUpdate               // someone else's stats updated
	BanchoHandleUserQuit                 // user quit bancho entirely (not irc)
	BanchoSpectatorJoined                // new spec
	BanchoSpectatorLeft                  // spectator left
	BanchoSpectateFrames                 // spectator frames chunks
	OsuStartSpectating                   // request to spectate someone
	OsuStopSpectating                    // stop spectating
	OsuSpectateFrames                    // spectator frames (client packet not from bancho unlike BanchoSpectateFrames)
	BanchoVersionUpdate                  // check for updates
	OsuErrorReport                       // report error to osu.ppy.sh
	OsuCantSpectate                      // can't spectate the host for whatever reason
	BanchoSpectatorCantSpectate          // can't spectate because no map
	BanchoGetAttention                   // make osu popup
	BanchoAnnounce                       // announcement popup
	OsuSendIRCMessagePrivate             // not sure
	BanchoMatchUpdate                    // update match details
	BanchoMatchNew                       // new match
	BanchoMatchDisband                   // close room
	OsuLobbyPart                         // client left lobby
	OsuLobbyJoin                         // client joined lobby
	OsuMatchCreate                       // client created a new lobby
	OsuMatchJoin                         // sends a request to bancho (join lobby)
	OsuLobbySomething                    // i can't figure out this
	BanchoLobbyJoinOBSOLETE              // according to the mid-2014 decompiled code this is when bancho informs a client about a new player that joins a lobby this is obsolete now.
	BanchoLobbyPartOBSOLETE              // according to the mid-2014 decompiled code this is when bancho informs a client about a new player that joins a lobby this is obsolete now.
	BanchoMatchJoinSuccess
	BanchoMatchJoinFail
	OsuMatchChangeSlot
	OsuMatchReady
	OsuMatchLock
	OsuMatchChangeSettings
	BanchoFellowSpectatorJoined
	BanchoFellowSpectatorLeft
	OsuMatchStart
	AllPlayersLoaded // no one is missing beatmap
	BanchoMatchStart
	OsuMatchScoreUpdate
	BanchoMatchScoreUpdate
	OsuMatchComplete
	BanchoMatchTransferHost
	OsuMatchChangeMods
	OsuMatchLoadComplete
	BanchoMatchAllPlayersLoaded
	OsuMatchNoBeatmap
	OsuMatchNotReady
	OsuMatchFailed
	BanchoMatchPlayerFailed
	BanchoMatchComplete
	OsuMatchHasBeatmap
	OsuMatchSkipRequest
	BanchoMatchSkip
	BanchoUnauthorised
	OsuChannelJoin
	BanchoChannelJoinSuccess
	BanchoChannelAvailable
	BanchoChannelRevoked
	BanchoChannelAvailableAutojoin
	OsuBeatmapInfoRequest
	BanchoBeatmapInfoReply
	OsuMatchTransferHost
	BanchoLoginPermissions
	BanchoFriendList
	OsuFriendAdd
	OsuFriendRemove
	BanchoProtocolVersion
	BanchoTitleUpdate
	OsuMatchChangeTeam
	OsuChannelLeave
	OsuReceiveUpdates
	BanchoMonitor
	BanchoMatchPlayerSkipped
	OsuSetIrcAwayMessage
	BanchoUserPresence
	IRCOnly
	OsuUserStatsRequest
	BanchoRestart
	OsuInvite
	BanchoInvite
	BanchoChannelListingComplete
	OsuMatchChangePassword
	BanchoMatchChangePassword
	BanchoBanInfo
	OsuSpecialMatchInfoRequest
	BanchoUserSilenced
	BanchoUserPresenceSingle
	BanchoUserPresenceBundle
	OsuUserPresenceRequest
	OsuUserPresenceRequestAll
	OsuUserToggleBlockNonFriendPM
	BanchoUserPMBlocked
	BanchoTargetIsSilenced
	BanchoVersionUpdateForced // force client update
	BanchoSwitchServer
	BanchoAccountRestricted
	BanchoRTX // pops up spooky message on your screen
	OsuMatchAbort
	BanchoSwitchTourneyServer
	OsuSpecialJoinMatchChannel  // force a client to join lobby (this is what OsuSQL uses afaik)
	OsuSpecialLeaveMatchChannel // force a client to leave lobby
)
