package ipmi

import "fmt"

// 22.18 Set Session Privilege Level Command
type SetSessionPrivilegeLevelRequest struct {
	PrivilegeLevel PrivilegeLevel
}

type SetSessionPrivilegeLevelResponse struct {
	// New Privilege Level (or present level if 'return present privilege level' was selected.)
	PrivilegeLevel uint8
}

func (req *SetSessionPrivilegeLevelRequest) Command() Command {
	return CommandSetSessionPrivilegeLevel
}

func (req *SetSessionPrivilegeLevelRequest) Pack() []byte {
	var msg = make([]byte, 1)
	packUint8(uint8(req.PrivilegeLevel), msg, 0)
	return msg
}

func (res *SetSessionPrivilegeLevelResponse) Unpack(msg []byte) error {
	if len(msg) < 1 {
		return ErrUnpackedDataTooShort
	}
	res.PrivilegeLevel = msg[0]
	return nil
}

func (*SetSessionPrivilegeLevelResponse) CompletionCodes() map[uint8]string {
	return map[uint8]string{
		0x80: "Requested level not available for this user",
		0x81: "Requested level exceeds Channel and/or User Privilege Limit",
		0x82: "Cannot disable User Level authentication",
	}
}

func (res *SetSessionPrivilegeLevelResponse) Format() string {
	return fmt.Sprintf("%v", res)
}

func (c *Client) SetSessionPrivilegeLevel(privilegeLevel PrivilegeLevel) (*SetSessionPrivilegeLevelResponse, error) {
	req := &SetSessionPrivilegeLevelRequest{
		PrivilegeLevel: privilegeLevel,
	}
	res := &SetSessionPrivilegeLevelResponse{}
	err := c.Exchange(req, res)
	return res, err
}
