/*
** Copyright [2013-2016] [Megam Systems]
**
** Licensed under the Apache License, Version 2.0 (the "License");
** you may not use this file except in compliance with the License.
** You may obtain a copy of the License at
**
** http://www.apache.org/licenses/LICENSE-2.0
**
** Unless required by applicable law or agreed to in writing, software
** distributed under the License is distributed on an "AS IS" BASIS,
** WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
** See the License for the specific language governing permissions and
** limitations under the License.
 */
package ubuntu

import (
	//"fmt"
	//"strings"
	"testing"

	"gopkg.in/check.v1"
)

type S struct{}

var _ = check.Suite(&S{})

func Test(t *testing.T) { check.TestingT(t) }

/*func (s *S) TestSubnetMaskLocalSize(c *check.C) {
  a, _ := IPNet().Mask.Size()
	c.Assert(a, check.Equals, 22)
}

func (s *S) TestSubnetMaskIP(c *check.C) {
	st := strings.Split(IP(""), ".")
	p := st[0 : len(st)-1]
	p = append(p, "0")
	si := fmt.Sprintf("%s/%d", strings.Join(p, "."), 24)
	c.Assert(si, check.Equals, "192.168.1.0/24")
}
*/
