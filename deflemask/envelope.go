package deflemask

// Envelope describes STD discrete envelope.
type Envelope struct {
	Values []int
	Loop   int
}

// MakeAutomaton generates FSM that acts like working envelope.
func (env *Envelope) MakeAutomaton() *EnvAutomaton {
	return &EnvAutomaton{Env: env}
}

// EnvAutomaton describes finite state machine that goes through the states of
// STD envelope, replicating its behaviour.
type EnvAutomaton struct {
	Env   *Envelope
	Phase int
}

// State returns state of envelope.
func (ep *EnvAutomaton) State() int {
	return ep.Env.Values[ep.Phase]
}

// Iterate advances envelope further in time.
func (ep *EnvAutomaton) Iterate() {
	ep.Phase++
	if ep.Phase < len(ep.Env.Values) {
		return
	}

	if ep.Env.Loop == -1 {
		ep.Phase--
	} else {
		ep.Phase = ep.Env.Loop
	}
}

// Reset resets envelope.
func (ep *EnvAutomaton) Reset() {
	ep.Phase = 0
}
