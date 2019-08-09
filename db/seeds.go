package db

// FillSeedsInformation <function>
// is used to fill information that needed before usage of news endpoint
func (s Service) FillSeedsInformation() {
	s.fillNewsSources()
	s.fillNewsTypes()
}
