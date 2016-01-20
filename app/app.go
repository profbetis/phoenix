// Copyright 2016 SynapseGarden
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package app

import (
	"github.com/go-phoenix/phoenix/graphics"
	"github.com/go-phoenix/phoenix/windower"
	"github.com/go-phoenix/phoenix/world"
)

// App blah blah.
type App struct {
	world.World
	graphics.Graphics
	windower.Windower
}

// Run starts the main loop of an App.
func (a *App) Run() error {
	if err := a.Windower.Setup(); err != nil {
		return err
	}

	for !a.finished() {
		a.World.Update()
		a.Graphics.Update()
	}

	return nil
}

// finished returns true if the App is finished running.
func (a *App) finished() bool {
	return false
}
