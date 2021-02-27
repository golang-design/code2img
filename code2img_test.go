// Copyright 2021 The golang.design Initiative authors.
// All rights reserved. Use of this source code is governed
// by a GNU GPL-3.0 license that can be found in the LICENSE file.
//
// Written by Changkun Ou <changkun.de>

package code2img_test

import (
	"bytes"
	"context"
	"image/png"
	"math"
	"os"
	"testing"
	"time"

	"golang.design/x/code2img"
)

func TestRender(t *testing.T) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*10)
	defer cancel()

	got, err := code2img.Render(ctx, `import "golang.design/x/code2img`)
	if err != nil {
		t.Fatalf("render failed: %v", err)
	}

	imgGot, err := png.Decode(bytes.NewReader(got))
	if err != nil {
		t.Fatalf("cannot read rendered image: %v", err)
	}

	want, err := os.ReadFile("testdata/want.png")
	if err != nil {
		t.Fatalf("cannot read gold test file: %v", err)
	}

	imgWant, err := png.Decode(bytes.NewReader(want))
	if err != nil {
		t.Fatalf("cannot read gold image: %v", err)
	}

	if math.Abs(float64(imgGot.Bounds().Dx()-imgWant.Bounds().Dx())) > 5 ||
		math.Abs(float64(imgGot.Bounds().Dy()-imgWant.Bounds().Dy())) > 5 {

		err := os.WriteFile("testdata/got.png", got, os.ModePerm)
		if err != nil {
			t.Errorf("failed to write image: %v", err)
		}

		t.Fatalf("image size does not match: got %+v, want %+v", imgGot.Bounds(), imgWant.Bounds())
	}
}
