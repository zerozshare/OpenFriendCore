/*
 * OpenFriend — Minecraft Java Edition Friends List bridge.
 * Copyright (c) 2026 ZSHARE (https://zpw.jp). Licensed under the MIT License.
 *
 * "Minecraft", "Xbox", "Xbox Live", "Microsoft", and "Mojang" are trademarks
 * of their respective owners. OpenFriend is not affiliated with, endorsed by,
 * sponsored by, or otherwise officially connected to Microsoft Corporation,
 * Mojang AB, or the Xbox brand. See LICENSE for the full notice.
 */
package skinimg

import (
	"bytes"
	"image"
	"image/color"
	"image/png"
	"os"
	"path/filepath"
	"testing"
)

func writeFixture(t *testing.T, dir string, size int) string {
	t.Helper()
	img := image.NewNRGBA(image.Rect(0, 0, size, size))
	for y := 0; y < size; y++ {
		for x := 0; x < size; x++ {
			img.SetNRGBA(x, y, color.NRGBA{R: uint8(x), G: uint8(y), B: 128, A: 255})
		}
	}
	path := filepath.Join(dir, "face.png")
	f, err := os.Create(path)
	if err != nil {
		t.Fatalf("create: %v", err)
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		t.Fatalf("encode: %v", err)
	}
	return path
}

func TestPrepareFace(t *testing.T) {
	dir := t.TempDir()
	path := writeFixture(t, dir, 32)
	res, err := Prepare(path)
	if err != nil {
		t.Fatalf("Prepare: %v", err)
	}
	if !res.Processed {
		t.Fatalf("expected processed=true for 32x32 input")
	}
	img, err := png.Decode(bytes.NewReader(res.PNG))
	if err != nil {
		t.Fatalf("decode: %v", err)
	}
	b := img.Bounds()
	if b.Dx() != 64 || b.Dy() != 64 {
		t.Fatalf("expected 64x64, got %dx%d", b.Dx(), b.Dy())
	}
	_, _, _, a := img.At(8, 8).RGBA()
	if a == 0 {
		t.Fatalf("face area (8,8) should be opaque, got alpha=0")
	}
	_, _, _, a = img.At(0, 0).RGBA()
	if a != 0 {
		t.Fatalf("non-face area (0,0) should be transparent, got alpha=%d", a)
	}
}

func TestPreparePassesThroughFullSkin(t *testing.T) {
	dir := t.TempDir()
	path := writeFixture(t, dir, 64)
	res, err := Prepare(path)
	if err != nil {
		t.Fatalf("Prepare: %v", err)
	}
	if res.Processed {
		t.Fatalf("expected processed=false for 64x64 input (treated as full skin)")
	}
}
