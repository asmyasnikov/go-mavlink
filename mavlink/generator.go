package mavlink

//go:generate mavgen -f ../schemas/public/mavlink-upstream/message_definitions/v1.0/all.xml -p
//go:generate mavgen -f ../schemas/public/mavlink-upstream/message_definitions/v1.0/ardupilotmega.xml
//go:generate mavgen -f ../schemas/public/mavlink-upstream/message_definitions/v1.0/ASLUAV.xml
//go:generate mavgen -f ../schemas/public/mavlink-upstream/message_definitions/v1.0/autoquad.xml -p
//go:generate mavgen -f ../schemas/public/mavlink-upstream/message_definitions/v1.0/common.xml -p
//go:generate mavgen -f ../schemas/public/mavlink-upstream/message_definitions/v1.0/icarous.xml -p
//go:generate mavgen -f ../schemas/public/mavlink-upstream/message_definitions/v1.0/matrixpilot.xml
//go:generate mavgen -f ../schemas/public/mavlink-upstream/message_definitions/v1.0/minimal.xml -p
//go:generate mavgen -f ../schemas/public/mavlink-upstream/message_definitions/v1.0/paparazzi.xml -p
//go:generate mavgen -f ../schemas/public/mavlink-upstream/message_definitions/v1.0/standard.xml -p
//go:generate mavgen -f ../schemas/public/mavlink-upstream/message_definitions/v1.0/ualberta.xml
//go:generate mavgen -f ../schemas/public/mavlink-upstream/message_definitions/v1.0/uAvionix.xml

func main() {
}