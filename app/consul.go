package app

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	consul "github.com/hashicorp/consul/api"
)

type Service struct {
	Title       string
	Description string
	URL         string
	IconURL     string
	Status      string
}

type Group struct {
	Title    string
	Services []*Service
}

func fetchServices() []Group {
	groups := map[string]*Group{"default": {Title: "Default"}}

	services, _, err := consulClient.Catalog().Services(&consul.QueryOptions{})
	if err != nil {
		panic(err)
	}

	enabledServices := make(map[string][]string)

	for name, tags := range services {
		for _, tag := range tags {
			// fmt.Println(tag)
			if tag == fmt.Sprintf("%s.enable=true", config.TagPrefix) {
				enabledServices[name] = tags
			}
		}
	}

	// fmt.Printf("%#+v\n", enabledServices)

	for service, tags := range enabledServices {
		serviceObj := &Service{Title: service}
		serviceGroup := "default"

		for _, tag := range tags {
			if !strings.HasPrefix(tag, fmt.Sprintf("%s.", config.TagPrefix)) {
				continue
			}
			tagParts := strings.Split(tag, "=")
			name := strings.TrimPrefix(tagParts[0], fmt.Sprintf("%s.", config.TagPrefix))
			value := tagParts[1]
			unquoted, err := strconv.Unquote(value)
			if err == nil {
				value = unquoted
			}

			switch name {
			case "title":
				serviceObj.Title = value
			case "description":
				serviceObj.Description = value
			case "url":
				serviceObj.URL = value
			case "icon_url":
				serviceObj.IconURL = value
			case "group":
				serviceGroup = value
			}
		}

		group, exists := groups[serviceGroup]
		if !exists {
			group = &Group{Title: serviceGroup}
			groups[serviceGroup] = group
		}
		group.Services = append(group.Services, serviceObj)
	}

	groupList := []Group{}
	for _, group := range groups {
		sort.Slice(group.Services, func(i, j int) bool {
			return group.Services[i].Title < group.Services[j].Title
		})
		if len(group.Services) != 0 {
			groupList = append(groupList, *group)
		}
	}

	sort.Slice(groupList, func(i, j int) bool {
		return groupList[i].Title < groupList[j].Title
	})

	return groupList

	// fmt.Printf("%#+v\n%#+v\n%#+v\n", services, metadata, err)
	// fmt.Printf("%#+v\n", enabledServices)
}
