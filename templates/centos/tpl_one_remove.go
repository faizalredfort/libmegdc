/*
** Copyright [2013-2015] [Megam Systems]
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

package centos

import (
	"github.com/megamsys/libmegdc/templates"
	"github.com/megamsys/urknall"

)

var centosoneremove *CentOsOneRemove

func init() {
	centosoneremove = &CentOsOneRemove{}
	templates.Register("CentOsOneRemove", centosoneremove)
}

type CentOsOneRemove struct{}

func (tpl *CentOsOneRemove) Render(p urknall.Package) {
	p.AddTemplate("one", &CentOsOneRemoveTemplate{})
}

func (tpl *CentOsOneRemove) Options(t *templates.Template) {
}

func (tpl *CentOsOneRemove) Run(target urknall.Target) error {
	return urknall.Run(target, &CentOsOneRemove{})
}

type CentOsOneRemoveTemplate struct{}

func (m *CentOsOneRemoveTemplate) Render(pkg urknall.Package) {
	pkg.AddCommands("one",
		RemovePackage("opennebula opennebula-sunstone"),
		RemovePackages(""),

	)
	pkg.AddCommands("one-clean",
		Shell("rm -rf /var/lib/urknall/one.*"),
	)
}
