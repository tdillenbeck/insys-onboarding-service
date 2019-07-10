package wmetricsconfig

import "weavelab.xyz/monorail/shared/wlib/config"

const (
	//WMetricsDPortCfg is the command-line flag that controls what port wmetricsd should be running on
	WMetricsDPortCfg = "wmetricsDPort"
	//PacketSizeCfg is the command-line flag that controls the packet size that can be sent to/received by wmetricsd
	PacketSizeCfg = "packetSize"
)

func init() {
	config.Add(WMetricsDPortCfg, "3051", "UDP port to listen on. Set using environment variable WMETRICSD_PORT to change both wmetricsd and the wmetrics lib. (Set to 0 to disable)", "WMETRICSD_PORT")

	// For cross internet the max packetSize should be 508 because according to http://stackoverflow.com/questions/1098897/what-is-the-largest-safe-udp-packet-size-on-the-internet
	// this is the number of bytes that is guaranteed to be deliverable (though not guaranteed to be delivered).

	config.Add(PacketSizeCfg, "8192", "The maximum size in bytes of each UDP packet. Set using environment variable WMETRICSD_UDP_PACKET_SIZE to change both wmetricsd and the wmetrics lib.", "WMETRICSD_UDP_PACKET_SIZE")
}
