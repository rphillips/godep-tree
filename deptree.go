/*
	Licensed under the Apache License, Version 2.0 (the "License");
	you may not use this file except in compliance with the License.
	You may obtain a copy of the License at

	http://www.apache.org/licenses/LICENSE-2.0

	Unless required by applicable law or agreed to in writing, software
	distributed under the License is distributed on an "AS-IS" BASIS,
	WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
	See the License for the specific language governing permissions and
	limitations under the License.
*/
package deptree

type DepTree struct {
	parents map[string][]string
}

func NewTree() *DepTree {
	d := &DepTree{}
	d.parents = make(map[string][]string)
	return d
}

func (d *DepTree) Add(parent string, child string) {
	if _, ok := d.parents[child]; !ok {
		d.parents[child] = []string{parent}
	} else {
		d.parents[child] = append(d.parents[child], parent)
	}
}

func appendIfMissing(slice []string, elem string) []string {
	for _, ele := range slice {
		if ele == elem {
			return slice
		}
	}
	return append(slice, elem)
}

func (d *DepTree) Solve(child string) []string {
	if _, ok := d.parents[child]; !ok {
		return nil
	}

	results := []string{child}
	for k := range d.parents[child] {
		results = appendIfMissing(results, d.parents[child][k])
		ch := d.Solve(d.parents[child][k])
		for c := range ch {
			results = appendIfMissing(results, ch[c])
		}
	}

	return results
}
