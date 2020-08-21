package convai_package_sdk

import (
	"testing"

	ctypes "github.com/datomar-labs-inc/convai-types"
	"github.com/google/uuid"
)

func TestRunnablePackage_GetNode(t *testing.T) {
	id := uuid.Must(uuid.NewRandom()).String()
	id2 := uuid.Must(uuid.NewRandom()).String()

	node := RunnableNode{
		PackageNode: ctypes.PackageNode{
			Name:    "Test Node",
			TypeID:      id,
			Version: "0.1.0",
		},
	}

	node2 := RunnableNode{
		PackageNode: ctypes.PackageNode{
			Name:    "Test Node",
			TypeID:      id2,
			Version: "0.1.0",
		},
	}

	p := RunnablePackage{
		Nodes: []RunnableNode{node, node2},
	}

	gn := p.GetNode(id)

	if gn == nil {
		t.Error("node was not returned correctly")
	} else {
		ndr := *gn

		if ndr.TypeID != gn.TypeID {
			t.Error("incorrect node was returned")
		}
	}

	randomID := uuid.Must(uuid.NewRandom()).String()

	gn2 := p.GetNode(randomID)

	if gn2 != nil {
		t.Error("expected nil node returned but got node instead")
	}
}

func TestRunnablePackage_GetLink(t *testing.T) {
	id := uuid.Must(uuid.NewRandom()).String()
	id2 := uuid.Must(uuid.NewRandom()).String()

	link := RunnableLink{
		PackageLink: ctypes.PackageLink{
			Name:    "Test Node",
			TypeID:      id,
			Version: "0.1.0",
		},
	}

	link2 := RunnableLink{
		PackageLink: ctypes.PackageLink{
			Name:    "Test Node",
			TypeID:      id2,
			Version: "0.1.0",
		},
	}

	p := RunnablePackage{
		Links: []RunnableLink{link, link2},
	}

	gl := p.GetLink(id)

	if gl == nil {
		t.Error("link was not returned correctly")
	} else {
		ndr := *gl

		if ndr.TypeID != gl.TypeID {
			t.Error("incorrect link was returned")
		}
	}

	randomID := uuid.Must(uuid.NewRandom()).String()

	gl2 := p.GetLink(randomID)

	if gl2 != nil {
		t.Error("expected nil link returned but got link instead")
	}
}

func TestRunnablePackage_GetDispatch(t *testing.T) {
	id := uuid.Must(uuid.NewRandom()).String()
	id2 := uuid.Must(uuid.NewRandom()).String()

	dispatch := RunnableDispatch{
		PackageDispatch: ctypes.PackageDispatch{
			Name: "Test Node",
			ID:   id,
		},
	}

	dispatch2 := RunnableDispatch{
		PackageDispatch: ctypes.PackageDispatch{
			Name: "Test Node",
			ID:   id2,
		},
	}

	p := RunnablePackage{
		Dispatches: []RunnableDispatch{dispatch, dispatch2},
	}

	gd := p.GetDispatch(id)

	if gd == nil {
		t.Error("dispatch was not returned correctly")
	} else {
		ndr := *gd

		if ndr.ID != gd.ID {
			t.Error("incorrect dispatch was returned")
		}
	}

	randomID := uuid.Must(uuid.NewRandom()).String()

	gl2 := p.GetDispatch(randomID)

	if gl2 != nil {
		t.Error("expected nil dispatch returned but got dispatch instead")
	}
}
