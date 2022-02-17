package types

type charMap map[string]struct {
	upper rune
	lower rune
}

func (c charMap) ToLower(s string) rune { return c[s].lower }
func (c charMap) ToUpper(s string) rune { return c[s].upper }

var Greek = charMap{
	// Name, Uppercase, Lowercase
	"Alpha":   {'Α', 'α'},
	"Beta":    {'Β', 'β'},
	"Gamma":   {'Γ', 'γ'},
	"Delta":   {'Δ', 'δ'},
	"Epsilon": {'Ε', 'ε'},
	"Zeta":    {'Ζ', 'ζ'},
	"Eta":     {'Η', 'η'},
	"Theta":   {'Θ', 'θ'},
	"Iota":    {'Ι', 'ι'},
	"Kappa":   {'Κ', 'κ'},
	"Lambda":  {'Λ', 'λ'},
	"Mu":      {'Μ', 'μ'},
	"Nu":      {'Ν', 'ν'},
	"Xi":      {'Ξ', 'ξ'},
	"Omicron": {'Ο', 'ο'},
	"Pi":      {'Π', 'π'},
	"Rho":     {'Ρ', 'ρ'},
	"Sigma":   {'Σ', 'σ'},
	"Tau":     {'Τ', 'τ'},
	"Upsilon": {'Υ', 'υ'},
	"Phi":     {'Φ', 'φ'},
	"Chi":     {'Χ', 'χ'},
	"Psi":     {'Ψ', 'ψ'},
	"Omega":   {'Ω', 'ω'},
}
