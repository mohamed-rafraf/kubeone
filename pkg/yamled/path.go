/*
Copyright 2019 The KubeOne Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package yamled

import (
	"fmt"
	"strconv"
	"strings"
)

type Step any

type Path []Step

func (p Path) Parent() Path {
	if len(p) < 1 {
		return nil
	}

	return p[0 : len(p)-1]
}

func (p Path) Tail() Step {
	if len(p) == 0 {
		return nil
	}

	return p[len(p)-1]
}

func (p Path) String() string {
	items := []string{}

	for _, item := range p {
		if s, ok := item.(string); ok {
			if strings.Contains(s, ".") {
				s = fmt.Sprintf(`"%s"`, s)
			}

			items = append(items, s)
		} else if i, ok := item.(int); ok {
			items = append(items, strconv.Itoa(i))
		}
	}

	return strings.Join(items, ".")
}
