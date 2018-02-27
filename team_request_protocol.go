package kr

import "encoding/json"

type CreateTeamRequest struct {
	TeamInfo TeamInfo `json:"team_info"`
}

type TeamCheckpoint struct {
	PublicKey       []byte          `json:"public_key"`
	TeamPublicKey   []byte          `json:"team_public_key"`
	LastBlockHash   []byte          `json:"last_block_hash"`
	ServerEndpoints ServerEndpoints `json:"server_endpoints"`
}
type ServerEndpoints struct {
	ApiHost     string `json:"api_host"`
	BillingHost string `json:"billing_host"`
}
type TeamOperationRequest struct {
	Operation RequestableTeamOperation `json:"operation"`
}

type TeamOperationResponse struct {
	PostedBlockHash *[]byte                    `json:"posted_block_hash"`
	Data            *TeamOperationResponseData `json:"data,omitempty"`
	Error           *string                    `json:"error,omitempty"`
}

type TeamOperationResponseData struct {
	InviteLink *string `json:"invite_link,omitempty"`
}

type ReadTeamRequest struct {
	PublicKey []byte `json:"public_key"`
}

type ReadTeamResponse struct {
	SignerPublicKey []byte `json:"public_key"`
	Token           string `json:"message"`
	Signature       []byte `json:"signature"`
}

type ReadToken struct {
	Time *TimeToken `json:"time,omitempty"`
}

type TimeToken struct {
	PublicKey  []byte `json:"public_key"`
	Expiration uint64 `json:"expiration"`
}

type RequestableTeamOperation struct {
	DirectInvite   *json.RawMessage `json:"direct_invite,omitempty"`
	IndirectInvite *json.RawMessage `json:"indirect_invite,omitempty"`

	CloseInvitations *json.RawMessage `json:"close_invitations,omitempty"`

	Remove *json.RawMessage `json:"remove,omitempty"`

	SetPolicy   *Policy   `json:"set_policy,omitempty"`
	SetTeamInfo *TeamInfo `json:"set_team_info,omitempty"`

	PinHostKey   *SSHHostKey `json:"pin_host_key,omitempty"`
	UnpinHostKey *SSHHostKey `json:"unpin_host_key,omitempty"`

	AddLoggingEndpoint    *LoggingEndpoint `json:"add_logging_endpoint,omitempty"`
	RemoveLoggingEndpoint *LoggingEndpoint `json:"remove_logging_endpoint,omitempty"`

	Promote *json.RawMessage `json:"promote,omitempty"`
	Demote  *json.RawMessage `json:"demote,omitempty"`
}

type LogDecryptionResponse struct {
	LogDecryptionKey []byte `json:"log_decryption_key"`
}

type Policy struct {
	TemporaryApprovalSeconds *uint64 `json:"temporary_approval_seconds,omitempty"`
}

type TeamInfo struct {
	Name string `json:"name,omitempty"`
}

type SSHHostKey struct {
	Host      string `json:"host"`
	PublicKey []byte `json:"public_key"`
}

type LoggingEndpoint struct {
	CommandEncrypted *struct{} `json:"command_encrypted,omitempty"`
}
