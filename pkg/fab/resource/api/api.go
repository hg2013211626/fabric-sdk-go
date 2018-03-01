/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

package api

import (
	"github.com/hyperledger/fabric-sdk-go/pkg/context"
	"github.com/hyperledger/fabric-sdk-go/pkg/context/api/fab"
	common "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/common"
	pb "github.com/hyperledger/fabric-sdk-go/third_party/github.com/hyperledger/fabric/protos/peer"
)

// Resource is a client that provides access to fabric resources such as chaincode.
type Resource interface {
	CreateChannel(request CreateChannelRequest) (fab.TransactionID, error)
	InstallChaincode(request InstallChaincodeRequest) ([]*fab.TransactionProposalResponse, fab.TransactionID, error)
	QueryInstalledChaincodes(peer fab.ProposalProcessor) (*pb.ChaincodeQueryResponse, error)
	QueryChannels(peer fab.ProposalProcessor) (*pb.ChannelQueryResponse, error)
	GenesisBlockFromOrderer(channelName string, orderer fab.Orderer) (*common.Block, error)
	LastConfigFromOrderer(channelName string, orderer fab.Orderer) (*common.ConfigEnvelope, error)
	JoinChannel(request JoinChannelRequest) error
	SignChannelConfig(config []byte, signer context.IdentityContext) (*common.ConfigSignature, error)
}

// CreateChannelRequest requests channel creation on the network
type CreateChannelRequest struct {
	// required - The name of the new channel
	Name string
	// required - The Orderer to send the update request
	Orderer fab.Orderer
	// optional - the envelope object containing all
	// required settings and signatures to initialize this channel.
	// This envelope would have been created by the command
	// line tool "configtx"
	Envelope []byte
	// optional - ConfigUpdate object built by the
	// buildChannelConfig() method of this package
	Config []byte
	// optional - the list of collected signatures
	// required by the channel create policy when using the `apiconfig` parameter.
	// see signChannelConfig() method of this package
	Signatures []*common.ConfigSignature
}

// InstallChaincodeRequest requests chaincode installation on the network
type InstallChaincodeRequest struct {
	// required - name of the chaincode
	Name string
	// required - path to the location of chaincode sources (path from GOPATH/src folder)
	Path string
	// chaincodeVersion: required - version of the chaincode
	Version string
	// required - package (chaincode package type and bytes)
	Package *CCPackage
	// required - proposal processor list
	Targets []fab.ProposalProcessor
}

// JoinChannelRequest allows a set of peers to transact on a channel on the network
type JoinChannelRequest struct {
	// The name of the channel to be joined.
	Name         string
	GenesisBlock *common.Block
	Targets      []fab.ProposalProcessor
}

// CCPackage contains package type and bytes required to create CDS
type CCPackage struct {
	Type pb.ChaincodeSpec_Type
	Code []byte
}
