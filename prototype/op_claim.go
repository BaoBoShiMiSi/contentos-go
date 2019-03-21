package prototype


func (m *ClaimOperation) GetSigner(auths *map[string]bool) {
	(*auths)[m.Account.Value] = true
}

func (m *ClaimOperation) Validate() error {
	return nil
}

func (m *ClaimOperation) GetAffectedProps(props *map[string]bool) {

}


func (m *ClaimAllOperation) GetSigner(auths *map[string]bool) {
	(*auths)[m.Account.Value] = true
}

func (m *ClaimAllOperation) Validate() error {
	return nil
}

func (m *ClaimAllOperation) GetAffectedProps(props *map[string]bool) {

}
