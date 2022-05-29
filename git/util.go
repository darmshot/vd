package git

func PushDevelopMasterTags() error {
	_, err := gitPushMaster()
	if err != nil {
		return err
	}

	_, err = gitPushDevelop()
	if err != nil {
		return err

	}

	_, err = gitPushTags()
	if err != nil {
		return err
	}

	return nil
}
