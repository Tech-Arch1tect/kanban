package repository

import (
	"log"
)

type linkType string

const (
	LinkTypeDepends      linkType = "depends_on"
	LinkTypeBlocks       linkType = "blocks"
	LinkTypeFixes        linkType = "fixes"
	LinkTypeBlockedBy    linkType = "blocked_by"
	LinkTypeFixedBy      linkType = "fixed_by"
	LinkTypeDependedOnBy linkType = "depended_on_by"
)

var LinkTypeMap = map[string]linkType{
	"depends_on": LinkTypeDepends,
	"blocks":     LinkTypeBlocks,
	"fixes":      LinkTypeFixes,
}

var InverseLinkTypeMap = map[string]linkType{
	"depended_on_by": LinkTypeDepends,
	"blocked_by":     LinkTypeBlocks,
	"fixed_by":       LinkTypeFixes,
	"depends_on":     LinkTypeDependedOnBy,
	"blocks":         LinkTypeBlockedBy,
	"fixes":          LinkTypeFixedBy,
}

func (t linkType) GetInverse() string {
	switch t {
	case LinkTypeDepends:
		return string(LinkTypeDependedOnBy)
	case LinkTypeBlocks:
		return string(LinkTypeBlockedBy)
	case LinkTypeFixes:
		return string(LinkTypeFixedBy)
	case LinkTypeBlockedBy:
		return string(LinkTypeBlocks)
	case LinkTypeFixedBy:
		return string(LinkTypeFixes)
	case LinkTypeDependedOnBy:
		return string(LinkTypeDepends)
	default:
		log.Println("couldnt_find_inverse_for_" + string(t))
		return "couldnt_find_inverse_for_" + string(t)
	}
}
