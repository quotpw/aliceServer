package kuzya

// volume
func getVolume(args *Args) map[string]any {
	volume := 0
	if args.WsDevice != nil {
		volume = args.WsDevice.GetVolume()
	}
	return map[string]any{
		"value": volume,
	}
}

func setVolume(args *Args) map[string]any {
	if args.WsDevice != nil && args.Json != nil {
		if value, ok := (*args.Json)["value"]; ok {
			args.WsDevice.SetVolume(int(value.(float64)))
		}
	}
	return getVolume(args)
}

// power
func powerState(args *Args) map[string]any {
	return map[string]any{
		"value": args.WsDevice != nil,
	}
}

func powerTurn(args *Args) map[string]any {
	if args.Json != nil {
		if value, ok := (*args.Json)["value"]; ok {
			if value.(float64) == 0 && args.WsDevice != nil { // turn off
				args.WsDevice.OsSystem("shutdown /s /f /t 1")
			} else if value.(float64) == 1 && args.DbDevice != nil { // turn on
				args.DbDevice.WOL()
			}
		}
	}
	return powerState(args)
}

// mute
func getMute(args *Args) map[string]any {
	mute := false
	if args.WsDevice != nil {
		mute = args.WsDevice.GetMute()
	}
	return map[string]any{
		"value": mute,
	}
}

func setMute(args *Args) map[string]any {
	if args.WsDevice != nil && args.Json != nil {
		if value, ok := (*args.Json)["value"]; ok {
			args.WsDevice.SetMute(value.(float64) == 1)
		}
	}
	return getMute(args)
}
