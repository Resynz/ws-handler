/**
 * @Author: Resynz
 * @Date: 2021/7/20 17:36
 */
package ws_handler

type PlatformType uint8

const (
	PlatformTypeUnknown PlatformType = iota
	PlatformPC
	PlatformH5
	PlatformIOS
	PlatformAndroid
)
