// Package mgmtapi provides primitives to interact with the openapi HTTP API.
//
// Code generated by unknown module path version unknown version DO NOT EDIT.
package mgmtapi

import (
	"time"
)

// Defines values for LogLevelLevel.
const (
	Debug LogLevelLevel = "debug"
	Error LogLevelLevel = "error"
	Info  LogLevelLevel = "info"
)

// Certificate defines model for Certificate.
type Certificate struct {
	DistinguishedName string       `json:"distinguished_name"`
	IsdAs             IsdAs        `json:"isd_as"`
	SubjectKeyAlgo    string       `json:"subject_key_algo"`
	SubjectKeyId      SubjectKeyID `json:"subject_key_id"`
	Validity          Validity     `json:"validity"`
}

// Chain defines model for Chain.
type Chain struct {
	Issuer  Certificate `json:"issuer"`
	Subject Certificate `json:"subject"`
}

// ChainBrief defines model for ChainBrief.
type ChainBrief struct {
	Id       ChainID  `json:"id"`
	Issuer   IsdAs    `json:"issuer"`
	Subject  IsdAs    `json:"subject"`
	Validity Validity `json:"validity"`
}

// ChainID defines model for ChainID.
type ChainID = string

// Hop defines model for Hop.
type Hop struct {
	Interface int   `json:"interface"`
	IsdAs     IsdAs `json:"isd_as"`
}

// IsdAs defines model for IsdAs.
type IsdAs = string

// LogLevel defines model for LogLevel.
type LogLevel struct {
	// Level Logging level
	Level LogLevelLevel `json:"level"`
}

// LogLevelLevel Logging level
type LogLevelLevel string

// Problem defines model for Problem.
type Problem struct {
	// Detail A human readable explanation specific to this occurrence of the problem that is helpful to locate the problem and give advice on how to proceed. Written in English and readable for engineers, usually not suited for non technical stakeholders and not localized.
	Detail *string `json:"detail,omitempty"`

	// Instance A URI reference that identifies the specific occurrence of the problem, e.g. by adding a fragment identifier or sub-path to the problem type. May be used to locate the root of this problem in the source code.
	Instance *string `json:"instance,omitempty"`

	// Status The HTTP status code generated by the origin server for this occurrence of the problem.
	Status int `json:"status"`

	// Title A short summary of the problem type. Written in English and readable for engineers, usually not suited for non technical stakeholders and not localized.
	Title string `json:"title"`

	// Type A URI reference that uniquely identifies the problem type only in the context of the provided API. Opposed to the specification in RFC-7807, it is neither recommended to be dereferencable and point to a human-readable documentation nor globally unique for the problem type.
	Type *string `json:"type,omitempty"`
}

// Segment defines model for Segment.
type Segment struct {
	Expiration  time.Time `json:"expiration"`
	Hops        []Hop     `json:"hops"`
	Id          SegmentID `json:"id"`
	LastUpdated time.Time `json:"last_updated"`
	Timestamp   time.Time `json:"timestamp"`
}

// SegmentBrief defines model for SegmentBrief.
type SegmentBrief struct {
	EndIsdAs IsdAs     `json:"end_isd_as"`
	Id       SegmentID `json:"id"`

	// Length Length of the segment.
	Length     int   `json:"length"`
	StartIsdAs IsdAs `json:"start_isd_as"`
}

// SegmentID defines model for SegmentID.
type SegmentID = string

// StandardError defines model for StandardError.
type StandardError struct {
	// Error Error message
	Error string `json:"error"`
}

// SubjectKeyID defines model for SubjectKeyID.
type SubjectKeyID = string

// TRC defines model for TRC.
type TRC struct {
	AuthoritativeAses []IsdAs  `json:"authoritative_ases"`
	CoreAses          []IsdAs  `json:"core_ases"`
	Description       string   `json:"description"`
	Id                TRCID    `json:"id"`
	Validity          Validity `json:"validity"`
}

// TRCBrief defines model for TRCBrief.
type TRCBrief struct {
	Id TRCID `json:"id"`
}

// TRCID defines model for TRCID.
type TRCID struct {
	BaseNumber   int `json:"base_number"`
	Isd          int `json:"isd"`
	SerialNumber int `json:"serial_number"`
}

// Validity defines model for Validity.
type Validity struct {
	NotAfter  time.Time `json:"not_after"`
	NotBefore time.Time `json:"not_before"`
}

// BadRequest defines model for BadRequest.
type BadRequest = StandardError

// GetCertificatesParams defines parameters for GetCertificates.
type GetCertificatesParams struct {
	IsdAs   *IsdAs     `form:"isd_as,omitempty" json:"isd_as,omitempty"`
	ValidAt *time.Time `form:"valid_at,omitempty" json:"valid_at,omitempty"`
	All     *bool      `form:"all,omitempty" json:"all,omitempty"`
}

// GetSegmentsParams defines parameters for GetSegments.
type GetSegmentsParams struct {
	// StartIsdAs Start ISD-AS of segment.
	StartIsdAs *IsdAs `form:"start_isd_as,omitempty" json:"start_isd_as,omitempty"`

	// EndIsdAs Terminal AS of segment.
	EndIsdAs *IsdAs `form:"end_isd_as,omitempty" json:"end_isd_as,omitempty"`
}

// GetTrcsParams defines parameters for GetTrcs.
type GetTrcsParams struct {
	Isd *[]int `form:"isd,omitempty" json:"isd,omitempty"`
	All *bool  `form:"all,omitempty" json:"all,omitempty"`
}

// SetLogLevelJSONRequestBody defines body for SetLogLevel for application/json ContentType.
type SetLogLevelJSONRequestBody = LogLevel
