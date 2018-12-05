package models

type CyberpowerResponse struct {
	Status struct {
		CommunicationAvaiable bool `json:"communicationAvaiable"`
		OnlyPhaseArch         bool `json:"onlyPhaseArch"`
		Utility               struct {
			State        string      `json:"state"`
			StateWarning bool        `json:"stateWarning"`
			Voltage      string      `json:"voltage"`
			Frequency    interface{} `json:"frequency"`
			Voltages     interface{} `json:"voltages"`
			Currents     interface{} `json:"currents"`
			Frequencies  interface{} `json:"frequencies"`
			PowerFactors interface{} `json:"powerFactors"`
		} `json:"utility"`
		Bypass struct {
			State        string      `json:"state"`
			StateWarning bool        `json:"stateWarning"`
			Voltage      interface{} `json:"voltage"`
			Current      interface{} `json:"current"`
			Frequency    interface{} `json:"frequency"`
			Voltages     interface{} `json:"voltages"`
			Currents     interface{} `json:"currents"`
			Frequencies  interface{} `json:"frequencies"`
			PowerFactors interface{} `json:"powerFactors"`
		} `json:"bypass"`
		Output struct {
			State             string      `json:"state"`
			StateWarning      bool        `json:"stateWarning"`
			Voltage           string      `json:"voltage"`
			Frequency         interface{} `json:"frequency"`
			Load              int         `json:"load"`
			Watt              int         `json:"watt"`
			Current           interface{} `json:"current"`
			OutputLoadWarning bool        `json:"outputLoadWarning"`
			Outlet1           interface{} `json:"outlet1"`
			Outlet2           interface{} `json:"outlet2"`
			ActivePower       interface{} `json:"activePower"`
			ApparentPower     interface{} `json:"apparentPower"`
			ReactivePower     interface{} `json:"reactivePower"`
			Voltages          interface{} `json:"voltages"`
			Currents          interface{} `json:"currents"`
			Frequencies       interface{} `json:"frequencies"`
			PowerFactors      interface{} `json:"powerFactors"`
			Loads             interface{} `json:"loads"`
			ActivePowers      interface{} `json:"activePowers"`
			ApparentPowers    interface{} `json:"apparentPowers"`
			ReactivePowers    interface{} `json:"reactivePowers"`
			EmergencyOff      interface{} `json:"emergencyOff"`
			BatteryExhausted  interface{} `json:"batteryExhausted"`
		} `json:"output"`
		Battery struct {
			State                string      `json:"state"`
			StateWarning         bool        `json:"stateWarning"`
			Voltage              string      `json:"voltage"`
			Capacity             int         `json:"capacity"`
			RuntimeFormat        int         `json:"runtimeFormat"`
			RuntimeFormatWarning bool        `json:"runtimeFormatWarning"`
			RuntimeHour          int         `json:"runtimeHour"`
			RuntimeMinute        int         `json:"runtimeMinute"`
			ChargetimeFormat     interface{} `json:"chargetimeFormat"`
			ChargetimeHour       interface{} `json:"chargetimeHour"`
			ChargetimeMinute     interface{} `json:"chargetimeMinute"`
			TemperatureCelsius   interface{} `json:"temperatureCelsius"`
			HighVoltage          interface{} `json:"highVoltage"`
			LowVoltage           interface{} `json:"lowVoltage"`
			HighCurrent          interface{} `json:"highCurrent"`
			LowCurrent           interface{} `json:"lowCurrent"`
		} `json:"battery"`
		UpsSystem struct {
			State                     string      `json:"state"`
			StateWarning              bool        `json:"stateWarning"`
			TemperatureCelsius        interface{} `json:"temperatureCelsius"`
			TemperatureFahrenheit     interface{} `json:"temperatureFahrenheit"`
			MaintenanceBreak          interface{} `json:"maintenanceBreak"`
			SystemFaultDueBypass      interface{} `json:"systemFaultDueBypass"`
			SystemFaultDueBypassFan   interface{} `json:"systemFaultDueBypassFan"`
			OriginalHardwareFaultCode string      `json:"originalHardwareFaultCode"`
		} `json:"upsSystem"`
		Modules  interface{} `json:"modules"`
		DeviceID int         `json:"deviceId"`
	} `json:"status"`
}
