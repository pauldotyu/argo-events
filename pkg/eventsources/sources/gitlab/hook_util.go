package gitlab

import (
	gitlab "gitlab.com/gitlab-org/api/client-go"
)

func getProjectHook(hooks []*gitlab.ProjectHook, url string) *gitlab.ProjectHook {
	for _, h := range hooks {
		if h.URL != url {
			continue
		}
		return h
	}
	return nil
}

func getGroupHook(hooks []*gitlab.GroupHook, url string) *gitlab.GroupHook {
	for _, h := range hooks {
		if h.URL != url {
			continue
		}
		return h
	}
	return nil
}
