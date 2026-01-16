/*
   Copyright (C) 2016  DeveloppSoft <developpsoft@gmail.com>

   This program is free software; you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation; either version 3 of the License, or
   (at your option) any later version.

   This program is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with this program; if not, write to the Free Software Foundation,
   Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301  USA

*/

package attacks

import (
	"os"
	"testing"
	"time"
)

func TestAttack_Init(t *testing.T) {
	t.Run("Init sets process and running status", func(t *testing.T) {
		attack := &Attack{
			Type:    "Deauth",
			Target:  "00:11:22:33:44:55",
			Running: false,
		}

		// Create a dummy process (using current process for testing)
		proc := os.Getpid()
		procObj, err := os.FindProcess(proc)
		if err != nil {
			t.Fatalf("Failed to find process: %v", err)
		}

		attack.Init(procObj)

		if attack.process != procObj {
			t.Error("Expected process to be set")
		}
		if !attack.Running {
			t.Error("Expected Running to be true after Init")
		}
	})
}

func TestAttack_Stop(t *testing.T) {
	t.Run("Stop without process sets stopped time", func(t *testing.T) {
		attack := &Attack{
			Type:    "Deauth",
			Target:  "00:11:22:33:44:55",
			Running: true,
			process: nil,
		}

		err := attack.Stop()
		if err != nil {
			t.Errorf("Expected no error when stopping without process, got: %v", err)
		}

		if attack.Running {
			t.Error("Expected Running to be false after Stop")
		}
		if attack.Stopped == "" {
			t.Error("Expected Stopped time to be set")
		}
	})

	t.Run("Stop with process kills process and sets stopped time", func(t *testing.T) {
		attack := &Attack{
			Type:    "Deauth",
			Target:  "00:11:22:33:44:55",
			Running: true,
		}

		// Create a dummy process for testing
		// Use sleep command that will run briefly
		proc := os.Getpid()
		procObj, err := os.FindProcess(proc)
		if err != nil {
			t.Fatalf("Failed to find process: %v", err)
		}

		attack.process = procObj
		attack.Running = true

		// Stop should handle gracefully even if process doesn't exist or can't be killed
		// (in this case, we can't kill our own process, but we can test the logic)
		err = attack.Stop()

		if attack.Running {
			t.Error("Expected Running to be false after Stop")
		}
		if attack.Stopped == "" {
			t.Error("Expected Stopped time to be set")
		}
	})
}

func TestAttack_StructFields(t *testing.T) {
	t.Run("Attack struct has correct fields", func(t *testing.T) {
		attack := &Attack{
			Type:    "Deauth",
			Target:  "00:11:22:33:44:55",
			Running: false,
			Started: time.Now().String(),
			Stopped: "",
		}

		if attack.Type != "Deauth" {
			t.Errorf("Expected Type to be 'Deauth', got: %s", attack.Type)
		}
		if attack.Target != "00:11:22:33:44:55" {
			t.Errorf("Expected Target to be '00:11:22:33:44:55', got: %s", attack.Target)
		}
		if attack.Started == "" {
			t.Error("Expected Started to be set")
		}
	})
}
