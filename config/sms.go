package config

import "fmt"

func ConfigSMS(contact string, code int) string {
	return fmt.Sprintf("http://kazinfoteh.org:9507?action=sendmessage&username=4sala1&password=741tVZXxd&recipient=%s&messagetype=SMS:TEXT&originator=KiT_Notify&messagedata=Keruen auth code: %d", contact, code)
}
