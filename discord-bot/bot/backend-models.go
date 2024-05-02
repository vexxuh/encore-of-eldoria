package discordbot

type BackendResponse struct {
	Inventory      BackendInventory `json:"inventory"`
	Username       string           `json:"username"`
	User           string           `json:"user"`
	C_level        int              `json:"c-level"`
	C_health       int              `json:"c-health"`
	M_health       int              `json:"m-health"`
	B_health       int              `json:"b-health"`
	S_strength     int              `json:"s-strength"`
	S_agility      int              `json:"s-agility"`
	S_constitution int              `json:"s-constitution"`
	S_intelligence int              `json:"s-intelligence"`
	S_wisdom       int              `json:"s-wisdom"`
	W_s_sword      int              `json:"w-s-sword"`
	W_s_axe        int              `json:"w-s-axe"`
	W_s_spear      int              `json:"w-s-spear"`
	P_state        string           `json:"p-state"`
	C_area         string           `json:"c-area"`
}

type BackendInventory struct {
	I_apple      int `json:"i-apple"`
	I_potion     int `json:"i-potion"`
	I_potionPlus int `json:"i-potion-plus"`
	C_gold       int `json:"c-gold"`
	B_gold       int `json:"b-gold"`
}
