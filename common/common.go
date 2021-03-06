// Copyright 2014 CoreOS, Inc.
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

// Package common defines values shared by different parts
// of rkt (e.g. stage0 and stage1)
package common

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"

	"github.com/coreos/rocket/Godeps/_workspace/src/github.com/appc/spec/aci"
	"github.com/coreos/rocket/Godeps/_workspace/src/github.com/appc/spec/schema/types"
)

const (
	stage1Dir = "/stage1"
	stage2Dir = "/opt/stage2"

	EnvLockFd               = "RKT_LOCK_FD"
	Stage1IDFilename        = "stage1ID"
	OverlayPreparedFilename = "overlay-prepared"

	MetadataServiceIP      = "169.254.169.255"
	MetadataServicePubPort = 80
	MetadataServicePrvPort = 2375
)

// Stage1ImagePath returns the path where the stage1 app image (unpacked ACI) is rooted,
// (i.e. where its contents are extracted during stage0).
func Stage1ImagePath(root string) string {
	return filepath.Join(root, stage1Dir)
}

// Stage1RootfsPath returns the path to the stage1 rootfs
func Stage1RootfsPath(root string) string {
	return filepath.Join(Stage1ImagePath(root), aci.RootfsDir)
}

// Stage1ManifestPath returns the path to the stage1's manifest file inside the expanded ACI.
func Stage1ManifestPath(root string) string {
	return filepath.Join(Stage1ImagePath(root), aci.ManifestFile)
}

// PodManifestPath returns the path in root to the Pod Manifest
func PodManifestPath(root string) string {
	return filepath.Join(root, "container")
}

// AppImagesPath returns the path where the app images live
func AppImagesPath(root string) string {
	return filepath.Join(Stage1RootfsPath(root), stage2Dir)
}

// AppImagePath returns the path where an app image (i.e. unpacked ACI) is rooted (i.e.
// where its contents are extracted during stage0), based on the app image ID.
func AppImagePath(root string, imageID types.Hash) string {
	return filepath.Join(AppImagesPath(root), types.ShortHash(imageID.String()))
}

// AppRootfsPath returns the path to an app's rootfs.
// imageID should be the app image ID.
func AppRootfsPath(root string, imageID types.Hash) string {
	return filepath.Join(AppImagePath(root, imageID), aci.RootfsDir)
}

// RelAppImagePath returns the path of an application image relative to the
// stage1 chroot
func RelAppImagePath(imageID types.Hash) string {
	return filepath.Join(stage2Dir, types.ShortHash(imageID.String()))
}

// RelAppImagePath returns the path of an application's rootfs relative to the
// stage1 chroot
func RelAppRootfsPath(imageID types.Hash) string {
	return filepath.Join(RelAppImagePath(imageID), aci.RootfsDir)
}

// ImageManifestPath returns the path to the app's manifest file inside the expanded ACI.
// id should be the app image ID.
func ImageManifestPath(root string, imageID types.Hash) string {
	return filepath.Join(AppImagePath(root, imageID), aci.ManifestFile)
}

// MetadataServicePrivateURL returns the private URL used to host the metadata service
func MetadataServicePrivateURL() string {
	return fmt.Sprintf("http://127.0.0.1:%v", MetadataServicePrvPort)
}

// MetadataServicePublicURL returns the public URL used to host the metadata service
func MetadataServicePublicURL() string {
	return fmt.Sprintf("http://%v:%v", MetadataServiceIP, MetadataServicePubPort)
}

func GetRktLockFD() (int, error) {
	if v := os.Getenv(EnvLockFd); v != "" {
		fd, err := strconv.ParseUint(v, 10, 32)
		if err != nil {
			return -1, err
		}
		return int(fd), nil
	}
	return -1, fmt.Errorf("%v env var is not set", EnvLockFd)
}
