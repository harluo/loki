package config

type wrapper struct {
	Logging *Logging `json:"logging,omitempty" validate:"required"`
}
