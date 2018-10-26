package predefined

//FragmentAction holds additional feature of a resource
type FragmentAction struct {
	Platform PlatformFragment `json:"platform"`
}

//PlatformFragment holds the fragment of platform
type PlatformFragment struct {
	PasswordPolicy string `json:"password_policy"`
}
