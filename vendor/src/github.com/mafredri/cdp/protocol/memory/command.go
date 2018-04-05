// Code generated by cdpgen. DO NOT EDIT.

package memory

// GetDOMCountersReply represents the return values for GetDOMCounters in the Memory domain.
type GetDOMCountersReply struct {
	Documents        int `json:"documents"`        // No description.
	Nodes            int `json:"nodes"`            // No description.
	JsEventListeners int `json:"jsEventListeners"` // No description.
}

// SetPressureNotificationsSuppressedArgs represents the arguments for SetPressureNotificationsSuppressed in the Memory domain.
type SetPressureNotificationsSuppressedArgs struct {
	Suppressed bool `json:"suppressed"` // If true, memory pressure notifications will be suppressed.
}

// NewSetPressureNotificationsSuppressedArgs initializes SetPressureNotificationsSuppressedArgs with the required arguments.
func NewSetPressureNotificationsSuppressedArgs(suppressed bool) *SetPressureNotificationsSuppressedArgs {
	args := new(SetPressureNotificationsSuppressedArgs)
	args.Suppressed = suppressed
	return args
}

// SimulatePressureNotificationArgs represents the arguments for SimulatePressureNotification in the Memory domain.
type SimulatePressureNotificationArgs struct {
	Level PressureLevel `json:"level"` // Memory pressure level of the notification.
}

// NewSimulatePressureNotificationArgs initializes SimulatePressureNotificationArgs with the required arguments.
func NewSimulatePressureNotificationArgs(level PressureLevel) *SimulatePressureNotificationArgs {
	args := new(SimulatePressureNotificationArgs)
	args.Level = level
	return args
}
