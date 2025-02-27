package ipmi

// 22.19
type CloseSessionRequest struct {
	// For IPMI v2.0/RMCP+ this is the Managed System Session ID value that was generated by the BMC, not the ID from the remote console. If Session ID = 0000_0000h then an implementation can optionally enable this command to take an additional byte of parameter data that allows a Session handle to be used to close a Session.
	SessionID uint32

	// Session Handle. (only present if Session ID = 0000_0000h)
	SessionHandle uint8
}

type CloseSessionResponse struct {
}

func (req *CloseSessionRequest) Pack() []byte {
	msg := make([]byte, 4)
	packUint32L(req.SessionID, msg, 0)
	if req.SessionID == 0 {
		msg = append(msg, 0)
		packUint8(req.SessionHandle, msg, 4)
	}
	return msg
}

func (req *CloseSessionRequest) Command() Command {
	return CommandCloseSession
}

func (res *CloseSessionResponse) Unpack(msg []byte) error {
	return nil
}

func (res *CloseSessionResponse) CompletionCodes() map[uint8]string {
	return map[uint8]string{
		0x87: "Invalid Session id",
		0x88: "Invalid Session handle",
	}
}

func (res *CloseSessionResponse) Format() string {
	return ""
}

func (c *Client) CloseSession(request *CloseSessionRequest) (response *CloseSessionResponse, err error) {
	response = &CloseSessionResponse{}
	err = c.Exchange(request, response)
	return
}
