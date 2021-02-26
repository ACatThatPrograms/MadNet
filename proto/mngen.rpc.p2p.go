// Code generated. DO NOT EDIT.

package proto

import (
	"context"
	"errors"
	"sync"
)

// P2PStatusHandler is an interface class that only contains
// the method HandleP2PStatus
// The class that implements this method MUST handle the RPC call for
// the method Status of the RPC service P2P
type P2PStatusHandler interface {
	HandleP2PStatus(context.Context, *StatusRequest) (*StatusResponse, error)
}

// P2PGetBlockHeadersHandler is an interface class that only contains
// the method HandleP2PGetBlockHeaders
// The class that implements this method MUST handle the RPC call for
// the method GetBlockHeaders of the RPC service P2P
type P2PGetBlockHeadersHandler interface {
	HandleP2PGetBlockHeaders(context.Context, *GetBlockHeadersRequest) (*GetBlockHeadersResponse, error)
}

// P2PGetMinedTxsHandler is an interface class that only contains
// the method HandleP2PGetMinedTxs
// The class that implements this method MUST handle the RPC call for
// the method GetMinedTxs of the RPC service P2P
type P2PGetMinedTxsHandler interface {
	HandleP2PGetMinedTxs(context.Context, *GetMinedTxsRequest) (*GetMinedTxsResponse, error)
}

// P2PGetPendingTxsHandler is an interface class that only contains
// the method HandleP2PGetPendingTxs
// The class that implements this method MUST handle the RPC call for
// the method GetPendingTxs of the RPC service P2P
type P2PGetPendingTxsHandler interface {
	HandleP2PGetPendingTxs(context.Context, *GetPendingTxsRequest) (*GetPendingTxsResponse, error)
}

// P2PGetSnapShotNodeHandler is an interface class that only contains
// the method HandleP2PGetSnapShotNode
// The class that implements this method MUST handle the RPC call for
// the method GetSnapShotNode of the RPC service P2P
type P2PGetSnapShotNodeHandler interface {
	HandleP2PGetSnapShotNode(context.Context, *GetSnapShotNodeRequest) (*GetSnapShotNodeResponse, error)
}

// P2PGetSnapShotStateDataHandler is an interface class that only contains
// the method HandleP2PGetSnapShotStateData
// The class that implements this method MUST handle the RPC call for
// the method GetSnapShotStateData of the RPC service P2P
type P2PGetSnapShotStateDataHandler interface {
	HandleP2PGetSnapShotStateData(context.Context, *GetSnapShotStateDataRequest) (*GetSnapShotStateDataResponse, error)
}

// P2PGetSnapShotHdrNodeHandler is an interface class that only contains
// the method HandleP2PGetSnapShotHdrNode
// The class that implements this method MUST handle the RPC call for
// the method GetSnapShotHdrNode of the RPC service P2P
type P2PGetSnapShotHdrNodeHandler interface {
	HandleP2PGetSnapShotHdrNode(context.Context, *GetSnapShotHdrNodeRequest) (*GetSnapShotHdrNodeResponse, error)
}

// P2PGossipTransactionHandler is an interface class that only contains
// the method HandleP2PGossipTransaction
// The class that implements this method MUST handle the RPC call for
// the method GossipTransaction of the RPC service P2P
type P2PGossipTransactionHandler interface {
	HandleP2PGossipTransaction(context.Context, *GossipTransactionMessage) (*GossipTransactionAck, error)
}

// P2PGossipProposalHandler is an interface class that only contains
// the method HandleP2PGossipProposal
// The class that implements this method MUST handle the RPC call for
// the method GossipProposal of the RPC service P2P
type P2PGossipProposalHandler interface {
	HandleP2PGossipProposal(context.Context, *GossipProposalMessage) (*GossipProposalAck, error)
}

// P2PGossipPreVoteHandler is an interface class that only contains
// the method HandleP2PGossipPreVote
// The class that implements this method MUST handle the RPC call for
// the method GossipPreVote of the RPC service P2P
type P2PGossipPreVoteHandler interface {
	HandleP2PGossipPreVote(context.Context, *GossipPreVoteMessage) (*GossipPreVoteAck, error)
}

// P2PGossipPreVoteNilHandler is an interface class that only contains
// the method HandleP2PGossipPreVoteNil
// The class that implements this method MUST handle the RPC call for
// the method GossipPreVoteNil of the RPC service P2P
type P2PGossipPreVoteNilHandler interface {
	HandleP2PGossipPreVoteNil(context.Context, *GossipPreVoteNilMessage) (*GossipPreVoteNilAck, error)
}

// P2PGossipPreCommitHandler is an interface class that only contains
// the method HandleP2PGossipPreCommit
// The class that implements this method MUST handle the RPC call for
// the method GossipPreCommit of the RPC service P2P
type P2PGossipPreCommitHandler interface {
	HandleP2PGossipPreCommit(context.Context, *GossipPreCommitMessage) (*GossipPreCommitAck, error)
}

// P2PGossipPreCommitNilHandler is an interface class that only contains
// the method HandleP2PGossipPreCommitNil
// The class that implements this method MUST handle the RPC call for
// the method GossipPreCommitNil of the RPC service P2P
type P2PGossipPreCommitNilHandler interface {
	HandleP2PGossipPreCommitNil(context.Context, *GossipPreCommitNilMessage) (*GossipPreCommitNilAck, error)
}

// P2PGossipNextRoundHandler is an interface class that only contains
// the method HandleP2PGossipNextRound
// The class that implements this method MUST handle the RPC call for
// the method GossipNextRound of the RPC service P2P
type P2PGossipNextRoundHandler interface {
	HandleP2PGossipNextRound(context.Context, *GossipNextRoundMessage) (*GossipNextRoundAck, error)
}

// P2PGossipNextHeightHandler is an interface class that only contains
// the method HandleP2PGossipNextHeight
// The class that implements this method MUST handle the RPC call for
// the method GossipNextHeight of the RPC service P2P
type P2PGossipNextHeightHandler interface {
	HandleP2PGossipNextHeight(context.Context, *GossipNextHeightMessage) (*GossipNextHeightAck, error)
}

// P2PGossipBlockHeaderHandler is an interface class that only contains
// the method HandleP2PGossipBlockHeader
// The class that implements this method MUST handle the RPC call for
// the method GossipBlockHeader of the RPC service P2P
type P2PGossipBlockHeaderHandler interface {
	HandleP2PGossipBlockHeader(context.Context, *GossipBlockHeaderMessage) (*GossipBlockHeaderAck, error)
}

// P2PGetPeersHandler is an interface class that only contains
// the method HandleP2PGetPeers
// The class that implements this method MUST handle the RPC call for
// the method GetPeers of the RPC service P2P
type P2PGetPeersHandler interface {
	HandleP2PGetPeers(context.Context, *GetPeersRequest) (*GetPeersResponse, error)
}

// DiscoveryGetPeersHandler is an interface class that only contains
// the method HandleDiscoveryGetPeers
// The class that implements this method MUST handle the RPC call for
// the method GetPeers of the RPC service Discovery
type DiscoveryGetPeersHandler interface {
	HandleDiscoveryGetPeers(context.Context, *GetPeersRequest) (*GetPeersResponse, error)
}

// inboundRPCDispatch allows handlers to be registered for all RPC methods
// using the Register<Service><Name> methods.
// After registration, the inboundRPCDispatch struct will dispatch calls
// to an rpc method via the methods named as <Service><Name>(...)
type inboundRPCDispatch struct {
	sync.Mutex
  //	handlerP2PStatus is the registered handler for the
	//  Status RPC method of service P2P
	handlerP2PStatus P2PStatusHandler
	// waitChanP2PStatus will cause a caller of the RPC
	// method Status on service P2P to block until the
	// method has been registered.
	waitChanP2PStatus chan struct{}
  //	handlerP2PGetBlockHeaders is the registered handler for the
	//  GetBlockHeaders RPC method of service P2P
	handlerP2PGetBlockHeaders P2PGetBlockHeadersHandler
	// waitChanP2PGetBlockHeaders will cause a caller of the RPC
	// method GetBlockHeaders on service P2P to block until the
	// method has been registered.
	waitChanP2PGetBlockHeaders chan struct{}
  //	handlerP2PGetMinedTxs is the registered handler for the
	//  GetMinedTxs RPC method of service P2P
	handlerP2PGetMinedTxs P2PGetMinedTxsHandler
	// waitChanP2PGetMinedTxs will cause a caller of the RPC
	// method GetMinedTxs on service P2P to block until the
	// method has been registered.
	waitChanP2PGetMinedTxs chan struct{}
  //	handlerP2PGetPendingTxs is the registered handler for the
	//  GetPendingTxs RPC method of service P2P
	handlerP2PGetPendingTxs P2PGetPendingTxsHandler
	// waitChanP2PGetPendingTxs will cause a caller of the RPC
	// method GetPendingTxs on service P2P to block until the
	// method has been registered.
	waitChanP2PGetPendingTxs chan struct{}
  //	handlerP2PGetSnapShotNode is the registered handler for the
	//  GetSnapShotNode RPC method of service P2P
	handlerP2PGetSnapShotNode P2PGetSnapShotNodeHandler
	// waitChanP2PGetSnapShotNode will cause a caller of the RPC
	// method GetSnapShotNode on service P2P to block until the
	// method has been registered.
	waitChanP2PGetSnapShotNode chan struct{}
  //	handlerP2PGetSnapShotStateData is the registered handler for the
	//  GetSnapShotStateData RPC method of service P2P
	handlerP2PGetSnapShotStateData P2PGetSnapShotStateDataHandler
	// waitChanP2PGetSnapShotStateData will cause a caller of the RPC
	// method GetSnapShotStateData on service P2P to block until the
	// method has been registered.
	waitChanP2PGetSnapShotStateData chan struct{}
  //	handlerP2PGetSnapShotHdrNode is the registered handler for the
	//  GetSnapShotHdrNode RPC method of service P2P
	handlerP2PGetSnapShotHdrNode P2PGetSnapShotHdrNodeHandler
	// waitChanP2PGetSnapShotHdrNode will cause a caller of the RPC
	// method GetSnapShotHdrNode on service P2P to block until the
	// method has been registered.
	waitChanP2PGetSnapShotHdrNode chan struct{}
  //	handlerP2PGossipTransaction is the registered handler for the
	//  GossipTransaction RPC method of service P2P
	handlerP2PGossipTransaction P2PGossipTransactionHandler
	// waitChanP2PGossipTransaction will cause a caller of the RPC
	// method GossipTransaction on service P2P to block until the
	// method has been registered.
	waitChanP2PGossipTransaction chan struct{}
  //	handlerP2PGossipProposal is the registered handler for the
	//  GossipProposal RPC method of service P2P
	handlerP2PGossipProposal P2PGossipProposalHandler
	// waitChanP2PGossipProposal will cause a caller of the RPC
	// method GossipProposal on service P2P to block until the
	// method has been registered.
	waitChanP2PGossipProposal chan struct{}
  //	handlerP2PGossipPreVote is the registered handler for the
	//  GossipPreVote RPC method of service P2P
	handlerP2PGossipPreVote P2PGossipPreVoteHandler
	// waitChanP2PGossipPreVote will cause a caller of the RPC
	// method GossipPreVote on service P2P to block until the
	// method has been registered.
	waitChanP2PGossipPreVote chan struct{}
  //	handlerP2PGossipPreVoteNil is the registered handler for the
	//  GossipPreVoteNil RPC method of service P2P
	handlerP2PGossipPreVoteNil P2PGossipPreVoteNilHandler
	// waitChanP2PGossipPreVoteNil will cause a caller of the RPC
	// method GossipPreVoteNil on service P2P to block until the
	// method has been registered.
	waitChanP2PGossipPreVoteNil chan struct{}
  //	handlerP2PGossipPreCommit is the registered handler for the
	//  GossipPreCommit RPC method of service P2P
	handlerP2PGossipPreCommit P2PGossipPreCommitHandler
	// waitChanP2PGossipPreCommit will cause a caller of the RPC
	// method GossipPreCommit on service P2P to block until the
	// method has been registered.
	waitChanP2PGossipPreCommit chan struct{}
  //	handlerP2PGossipPreCommitNil is the registered handler for the
	//  GossipPreCommitNil RPC method of service P2P
	handlerP2PGossipPreCommitNil P2PGossipPreCommitNilHandler
	// waitChanP2PGossipPreCommitNil will cause a caller of the RPC
	// method GossipPreCommitNil on service P2P to block until the
	// method has been registered.
	waitChanP2PGossipPreCommitNil chan struct{}
  //	handlerP2PGossipNextRound is the registered handler for the
	//  GossipNextRound RPC method of service P2P
	handlerP2PGossipNextRound P2PGossipNextRoundHandler
	// waitChanP2PGossipNextRound will cause a caller of the RPC
	// method GossipNextRound on service P2P to block until the
	// method has been registered.
	waitChanP2PGossipNextRound chan struct{}
  //	handlerP2PGossipNextHeight is the registered handler for the
	//  GossipNextHeight RPC method of service P2P
	handlerP2PGossipNextHeight P2PGossipNextHeightHandler
	// waitChanP2PGossipNextHeight will cause a caller of the RPC
	// method GossipNextHeight on service P2P to block until the
	// method has been registered.
	waitChanP2PGossipNextHeight chan struct{}
  //	handlerP2PGossipBlockHeader is the registered handler for the
	//  GossipBlockHeader RPC method of service P2P
	handlerP2PGossipBlockHeader P2PGossipBlockHeaderHandler
	// waitChanP2PGossipBlockHeader will cause a caller of the RPC
	// method GossipBlockHeader on service P2P to block until the
	// method has been registered.
	waitChanP2PGossipBlockHeader chan struct{}
  //	handlerP2PGetPeers is the registered handler for the
	//  GetPeers RPC method of service P2P
	handlerP2PGetPeers P2PGetPeersHandler
	// waitChanP2PGetPeers will cause a caller of the RPC
	// method GetPeers on service P2P to block until the
	// method has been registered.
	waitChanP2PGetPeers chan struct{}
  //	handlerDiscoveryGetPeers is the registered handler for the
	//  GetPeers RPC method of service Discovery
	handlerDiscoveryGetPeers DiscoveryGetPeersHandler
	// waitChanDiscoveryGetPeers will cause a caller of the RPC
	// method GetPeers on service Discovery to block until the
	// method has been registered.
	waitChanDiscoveryGetPeers chan struct{}
}

// RegisterP2PStatus will register the object 't' as the service
// handler for the RPC method Status from service P2P
func (d *inboundRPCDispatch) RegisterP2PStatus(t P2PStatusHandler) {
	d.Lock()
	defer d.Unlock()
	// double registration is not allowed
	if d.handlerP2PStatus != nil {
		panic("double registration of P2PStatus")
	}
	// register the service handler
	d.handlerP2PStatus = t
	// close the wait channel to signal that the method is ready to use
	close(d.waitChanP2PStatus)
}

// P2PStatus will invoke the handler for the RPC method
// Status from service P2P
func (d *inboundRPCDispatch) P2PStatus(ctx context.Context, r *StatusRequest) (*StatusResponse, error) {
	// wait for registration to complete or context to be canceled
	select {
	case <-ctx.Done():
		return nil, errors.New("context canceled")
	case <-d.waitChanP2PStatus:
		// return the invoked methods response
		return d.handlerP2PStatus.HandleP2PStatus(ctx, r)
	}
}

// RegisterP2PGetBlockHeaders will register the object 't' as the service
// handler for the RPC method GetBlockHeaders from service P2P
func (d *inboundRPCDispatch) RegisterP2PGetBlockHeaders(t P2PGetBlockHeadersHandler) {
	d.Lock()
	defer d.Unlock()
	// double registration is not allowed
	if d.handlerP2PGetBlockHeaders != nil {
		panic("double registration of P2PGetBlockHeaders")
	}
	// register the service handler
	d.handlerP2PGetBlockHeaders = t
	// close the wait channel to signal that the method is ready to use
	close(d.waitChanP2PGetBlockHeaders)
}

// P2PGetBlockHeaders will invoke the handler for the RPC method
// GetBlockHeaders from service P2P
func (d *inboundRPCDispatch) P2PGetBlockHeaders(ctx context.Context, r *GetBlockHeadersRequest) (*GetBlockHeadersResponse, error) {
	// wait for registration to complete or context to be canceled
	select {
	case <-ctx.Done():
		return nil, errors.New("context canceled")
	case <-d.waitChanP2PGetBlockHeaders:
		// return the invoked methods response
		return d.handlerP2PGetBlockHeaders.HandleP2PGetBlockHeaders(ctx, r)
	}
}

// RegisterP2PGetMinedTxs will register the object 't' as the service
// handler for the RPC method GetMinedTxs from service P2P
func (d *inboundRPCDispatch) RegisterP2PGetMinedTxs(t P2PGetMinedTxsHandler) {
	d.Lock()
	defer d.Unlock()
	// double registration is not allowed
	if d.handlerP2PGetMinedTxs != nil {
		panic("double registration of P2PGetMinedTxs")
	}
	// register the service handler
	d.handlerP2PGetMinedTxs = t
	// close the wait channel to signal that the method is ready to use
	close(d.waitChanP2PGetMinedTxs)
}

// P2PGetMinedTxs will invoke the handler for the RPC method
// GetMinedTxs from service P2P
func (d *inboundRPCDispatch) P2PGetMinedTxs(ctx context.Context, r *GetMinedTxsRequest) (*GetMinedTxsResponse, error) {
	// wait for registration to complete or context to be canceled
	select {
	case <-ctx.Done():
		return nil, errors.New("context canceled")
	case <-d.waitChanP2PGetMinedTxs:
		// return the invoked methods response
		return d.handlerP2PGetMinedTxs.HandleP2PGetMinedTxs(ctx, r)
	}
}

// RegisterP2PGetPendingTxs will register the object 't' as the service
// handler for the RPC method GetPendingTxs from service P2P
func (d *inboundRPCDispatch) RegisterP2PGetPendingTxs(t P2PGetPendingTxsHandler) {
	d.Lock()
	defer d.Unlock()
	// double registration is not allowed
	if d.handlerP2PGetPendingTxs != nil {
		panic("double registration of P2PGetPendingTxs")
	}
	// register the service handler
	d.handlerP2PGetPendingTxs = t
	// close the wait channel to signal that the method is ready to use
	close(d.waitChanP2PGetPendingTxs)
}

// P2PGetPendingTxs will invoke the handler for the RPC method
// GetPendingTxs from service P2P
func (d *inboundRPCDispatch) P2PGetPendingTxs(ctx context.Context, r *GetPendingTxsRequest) (*GetPendingTxsResponse, error) {
	// wait for registration to complete or context to be canceled
	select {
	case <-ctx.Done():
		return nil, errors.New("context canceled")
	case <-d.waitChanP2PGetPendingTxs:
		// return the invoked methods response
		return d.handlerP2PGetPendingTxs.HandleP2PGetPendingTxs(ctx, r)
	}
}

// RegisterP2PGetSnapShotNode will register the object 't' as the service
// handler for the RPC method GetSnapShotNode from service P2P
func (d *inboundRPCDispatch) RegisterP2PGetSnapShotNode(t P2PGetSnapShotNodeHandler) {
	d.Lock()
	defer d.Unlock()
	// double registration is not allowed
	if d.handlerP2PGetSnapShotNode != nil {
		panic("double registration of P2PGetSnapShotNode")
	}
	// register the service handler
	d.handlerP2PGetSnapShotNode = t
	// close the wait channel to signal that the method is ready to use
	close(d.waitChanP2PGetSnapShotNode)
}

// P2PGetSnapShotNode will invoke the handler for the RPC method
// GetSnapShotNode from service P2P
func (d *inboundRPCDispatch) P2PGetSnapShotNode(ctx context.Context, r *GetSnapShotNodeRequest) (*GetSnapShotNodeResponse, error) {
	// wait for registration to complete or context to be canceled
	select {
	case <-ctx.Done():
		return nil, errors.New("context canceled")
	case <-d.waitChanP2PGetSnapShotNode:
		// return the invoked methods response
		return d.handlerP2PGetSnapShotNode.HandleP2PGetSnapShotNode(ctx, r)
	}
}

// RegisterP2PGetSnapShotStateData will register the object 't' as the service
// handler for the RPC method GetSnapShotStateData from service P2P
func (d *inboundRPCDispatch) RegisterP2PGetSnapShotStateData(t P2PGetSnapShotStateDataHandler) {
	d.Lock()
	defer d.Unlock()
	// double registration is not allowed
	if d.handlerP2PGetSnapShotStateData != nil {
		panic("double registration of P2PGetSnapShotStateData")
	}
	// register the service handler
	d.handlerP2PGetSnapShotStateData = t
	// close the wait channel to signal that the method is ready to use
	close(d.waitChanP2PGetSnapShotStateData)
}

// P2PGetSnapShotStateData will invoke the handler for the RPC method
// GetSnapShotStateData from service P2P
func (d *inboundRPCDispatch) P2PGetSnapShotStateData(ctx context.Context, r *GetSnapShotStateDataRequest) (*GetSnapShotStateDataResponse, error) {
	// wait for registration to complete or context to be canceled
	select {
	case <-ctx.Done():
		return nil, errors.New("context canceled")
	case <-d.waitChanP2PGetSnapShotStateData:
		// return the invoked methods response
		return d.handlerP2PGetSnapShotStateData.HandleP2PGetSnapShotStateData(ctx, r)
	}
}

// RegisterP2PGetSnapShotHdrNode will register the object 't' as the service
// handler for the RPC method GetSnapShotHdrNode from service P2P
func (d *inboundRPCDispatch) RegisterP2PGetSnapShotHdrNode(t P2PGetSnapShotHdrNodeHandler) {
	d.Lock()
	defer d.Unlock()
	// double registration is not allowed
	if d.handlerP2PGetSnapShotHdrNode != nil {
		panic("double registration of P2PGetSnapShotHdrNode")
	}
	// register the service handler
	d.handlerP2PGetSnapShotHdrNode = t
	// close the wait channel to signal that the method is ready to use
	close(d.waitChanP2PGetSnapShotHdrNode)
}

// P2PGetSnapShotHdrNode will invoke the handler for the RPC method
// GetSnapShotHdrNode from service P2P
func (d *inboundRPCDispatch) P2PGetSnapShotHdrNode(ctx context.Context, r *GetSnapShotHdrNodeRequest) (*GetSnapShotHdrNodeResponse, error) {
	// wait for registration to complete or context to be canceled
	select {
	case <-ctx.Done():
		return nil, errors.New("context canceled")
	case <-d.waitChanP2PGetSnapShotHdrNode:
		// return the invoked methods response
		return d.handlerP2PGetSnapShotHdrNode.HandleP2PGetSnapShotHdrNode(ctx, r)
	}
}

// RegisterP2PGossipTransaction will register the object 't' as the service
// handler for the RPC method GossipTransaction from service P2P
func (d *inboundRPCDispatch) RegisterP2PGossipTransaction(t P2PGossipTransactionHandler) {
	d.Lock()
	defer d.Unlock()
	// double registration is not allowed
	if d.handlerP2PGossipTransaction != nil {
		panic("double registration of P2PGossipTransaction")
	}
	// register the service handler
	d.handlerP2PGossipTransaction = t
	// close the wait channel to signal that the method is ready to use
	close(d.waitChanP2PGossipTransaction)
}

// P2PGossipTransaction will invoke the handler for the RPC method
// GossipTransaction from service P2P
func (d *inboundRPCDispatch) P2PGossipTransaction(ctx context.Context, r *GossipTransactionMessage) (*GossipTransactionAck, error) {
	// wait for registration to complete or context to be canceled
	select {
	case <-ctx.Done():
		return nil, errors.New("context canceled")
	case <-d.waitChanP2PGossipTransaction:
		// return the invoked methods response
		return d.handlerP2PGossipTransaction.HandleP2PGossipTransaction(ctx, r)
	}
}

// RegisterP2PGossipProposal will register the object 't' as the service
// handler for the RPC method GossipProposal from service P2P
func (d *inboundRPCDispatch) RegisterP2PGossipProposal(t P2PGossipProposalHandler) {
	d.Lock()
	defer d.Unlock()
	// double registration is not allowed
	if d.handlerP2PGossipProposal != nil {
		panic("double registration of P2PGossipProposal")
	}
	// register the service handler
	d.handlerP2PGossipProposal = t
	// close the wait channel to signal that the method is ready to use
	close(d.waitChanP2PGossipProposal)
}

// P2PGossipProposal will invoke the handler for the RPC method
// GossipProposal from service P2P
func (d *inboundRPCDispatch) P2PGossipProposal(ctx context.Context, r *GossipProposalMessage) (*GossipProposalAck, error) {
	// wait for registration to complete or context to be canceled
	select {
	case <-ctx.Done():
		return nil, errors.New("context canceled")
	case <-d.waitChanP2PGossipProposal:
		// return the invoked methods response
		return d.handlerP2PGossipProposal.HandleP2PGossipProposal(ctx, r)
	}
}

// RegisterP2PGossipPreVote will register the object 't' as the service
// handler for the RPC method GossipPreVote from service P2P
func (d *inboundRPCDispatch) RegisterP2PGossipPreVote(t P2PGossipPreVoteHandler) {
	d.Lock()
	defer d.Unlock()
	// double registration is not allowed
	if d.handlerP2PGossipPreVote != nil {
		panic("double registration of P2PGossipPreVote")
	}
	// register the service handler
	d.handlerP2PGossipPreVote = t
	// close the wait channel to signal that the method is ready to use
	close(d.waitChanP2PGossipPreVote)
}

// P2PGossipPreVote will invoke the handler for the RPC method
// GossipPreVote from service P2P
func (d *inboundRPCDispatch) P2PGossipPreVote(ctx context.Context, r *GossipPreVoteMessage) (*GossipPreVoteAck, error) {
	// wait for registration to complete or context to be canceled
	select {
	case <-ctx.Done():
		return nil, errors.New("context canceled")
	case <-d.waitChanP2PGossipPreVote:
		// return the invoked methods response
		return d.handlerP2PGossipPreVote.HandleP2PGossipPreVote(ctx, r)
	}
}

// RegisterP2PGossipPreVoteNil will register the object 't' as the service
// handler for the RPC method GossipPreVoteNil from service P2P
func (d *inboundRPCDispatch) RegisterP2PGossipPreVoteNil(t P2PGossipPreVoteNilHandler) {
	d.Lock()
	defer d.Unlock()
	// double registration is not allowed
	if d.handlerP2PGossipPreVoteNil != nil {
		panic("double registration of P2PGossipPreVoteNil")
	}
	// register the service handler
	d.handlerP2PGossipPreVoteNil = t
	// close the wait channel to signal that the method is ready to use
	close(d.waitChanP2PGossipPreVoteNil)
}

// P2PGossipPreVoteNil will invoke the handler for the RPC method
// GossipPreVoteNil from service P2P
func (d *inboundRPCDispatch) P2PGossipPreVoteNil(ctx context.Context, r *GossipPreVoteNilMessage) (*GossipPreVoteNilAck, error) {
	// wait for registration to complete or context to be canceled
	select {
	case <-ctx.Done():
		return nil, errors.New("context canceled")
	case <-d.waitChanP2PGossipPreVoteNil:
		// return the invoked methods response
		return d.handlerP2PGossipPreVoteNil.HandleP2PGossipPreVoteNil(ctx, r)
	}
}

// RegisterP2PGossipPreCommit will register the object 't' as the service
// handler for the RPC method GossipPreCommit from service P2P
func (d *inboundRPCDispatch) RegisterP2PGossipPreCommit(t P2PGossipPreCommitHandler) {
	d.Lock()
	defer d.Unlock()
	// double registration is not allowed
	if d.handlerP2PGossipPreCommit != nil {
		panic("double registration of P2PGossipPreCommit")
	}
	// register the service handler
	d.handlerP2PGossipPreCommit = t
	// close the wait channel to signal that the method is ready to use
	close(d.waitChanP2PGossipPreCommit)
}

// P2PGossipPreCommit will invoke the handler for the RPC method
// GossipPreCommit from service P2P
func (d *inboundRPCDispatch) P2PGossipPreCommit(ctx context.Context, r *GossipPreCommitMessage) (*GossipPreCommitAck, error) {
	// wait for registration to complete or context to be canceled
	select {
	case <-ctx.Done():
		return nil, errors.New("context canceled")
	case <-d.waitChanP2PGossipPreCommit:
		// return the invoked methods response
		return d.handlerP2PGossipPreCommit.HandleP2PGossipPreCommit(ctx, r)
	}
}

// RegisterP2PGossipPreCommitNil will register the object 't' as the service
// handler for the RPC method GossipPreCommitNil from service P2P
func (d *inboundRPCDispatch) RegisterP2PGossipPreCommitNil(t P2PGossipPreCommitNilHandler) {
	d.Lock()
	defer d.Unlock()
	// double registration is not allowed
	if d.handlerP2PGossipPreCommitNil != nil {
		panic("double registration of P2PGossipPreCommitNil")
	}
	// register the service handler
	d.handlerP2PGossipPreCommitNil = t
	// close the wait channel to signal that the method is ready to use
	close(d.waitChanP2PGossipPreCommitNil)
}

// P2PGossipPreCommitNil will invoke the handler for the RPC method
// GossipPreCommitNil from service P2P
func (d *inboundRPCDispatch) P2PGossipPreCommitNil(ctx context.Context, r *GossipPreCommitNilMessage) (*GossipPreCommitNilAck, error) {
	// wait for registration to complete or context to be canceled
	select {
	case <-ctx.Done():
		return nil, errors.New("context canceled")
	case <-d.waitChanP2PGossipPreCommitNil:
		// return the invoked methods response
		return d.handlerP2PGossipPreCommitNil.HandleP2PGossipPreCommitNil(ctx, r)
	}
}

// RegisterP2PGossipNextRound will register the object 't' as the service
// handler for the RPC method GossipNextRound from service P2P
func (d *inboundRPCDispatch) RegisterP2PGossipNextRound(t P2PGossipNextRoundHandler) {
	d.Lock()
	defer d.Unlock()
	// double registration is not allowed
	if d.handlerP2PGossipNextRound != nil {
		panic("double registration of P2PGossipNextRound")
	}
	// register the service handler
	d.handlerP2PGossipNextRound = t
	// close the wait channel to signal that the method is ready to use
	close(d.waitChanP2PGossipNextRound)
}

// P2PGossipNextRound will invoke the handler for the RPC method
// GossipNextRound from service P2P
func (d *inboundRPCDispatch) P2PGossipNextRound(ctx context.Context, r *GossipNextRoundMessage) (*GossipNextRoundAck, error) {
	// wait for registration to complete or context to be canceled
	select {
	case <-ctx.Done():
		return nil, errors.New("context canceled")
	case <-d.waitChanP2PGossipNextRound:
		// return the invoked methods response
		return d.handlerP2PGossipNextRound.HandleP2PGossipNextRound(ctx, r)
	}
}

// RegisterP2PGossipNextHeight will register the object 't' as the service
// handler for the RPC method GossipNextHeight from service P2P
func (d *inboundRPCDispatch) RegisterP2PGossipNextHeight(t P2PGossipNextHeightHandler) {
	d.Lock()
	defer d.Unlock()
	// double registration is not allowed
	if d.handlerP2PGossipNextHeight != nil {
		panic("double registration of P2PGossipNextHeight")
	}
	// register the service handler
	d.handlerP2PGossipNextHeight = t
	// close the wait channel to signal that the method is ready to use
	close(d.waitChanP2PGossipNextHeight)
}

// P2PGossipNextHeight will invoke the handler for the RPC method
// GossipNextHeight from service P2P
func (d *inboundRPCDispatch) P2PGossipNextHeight(ctx context.Context, r *GossipNextHeightMessage) (*GossipNextHeightAck, error) {
	// wait for registration to complete or context to be canceled
	select {
	case <-ctx.Done():
		return nil, errors.New("context canceled")
	case <-d.waitChanP2PGossipNextHeight:
		// return the invoked methods response
		return d.handlerP2PGossipNextHeight.HandleP2PGossipNextHeight(ctx, r)
	}
}

// RegisterP2PGossipBlockHeader will register the object 't' as the service
// handler for the RPC method GossipBlockHeader from service P2P
func (d *inboundRPCDispatch) RegisterP2PGossipBlockHeader(t P2PGossipBlockHeaderHandler) {
	d.Lock()
	defer d.Unlock()
	// double registration is not allowed
	if d.handlerP2PGossipBlockHeader != nil {
		panic("double registration of P2PGossipBlockHeader")
	}
	// register the service handler
	d.handlerP2PGossipBlockHeader = t
	// close the wait channel to signal that the method is ready to use
	close(d.waitChanP2PGossipBlockHeader)
}

// P2PGossipBlockHeader will invoke the handler for the RPC method
// GossipBlockHeader from service P2P
func (d *inboundRPCDispatch) P2PGossipBlockHeader(ctx context.Context, r *GossipBlockHeaderMessage) (*GossipBlockHeaderAck, error) {
	// wait for registration to complete or context to be canceled
	select {
	case <-ctx.Done():
		return nil, errors.New("context canceled")
	case <-d.waitChanP2PGossipBlockHeader:
		// return the invoked methods response
		return d.handlerP2PGossipBlockHeader.HandleP2PGossipBlockHeader(ctx, r)
	}
}

// RegisterP2PGetPeers will register the object 't' as the service
// handler for the RPC method GetPeers from service P2P
func (d *inboundRPCDispatch) RegisterP2PGetPeers(t P2PGetPeersHandler) {
	d.Lock()
	defer d.Unlock()
	// double registration is not allowed
	if d.handlerP2PGetPeers != nil {
		panic("double registration of P2PGetPeers")
	}
	// register the service handler
	d.handlerP2PGetPeers = t
	// close the wait channel to signal that the method is ready to use
	close(d.waitChanP2PGetPeers)
}

// P2PGetPeers will invoke the handler for the RPC method
// GetPeers from service P2P
func (d *inboundRPCDispatch) P2PGetPeers(ctx context.Context, r *GetPeersRequest) (*GetPeersResponse, error) {
	// wait for registration to complete or context to be canceled
	select {
	case <-ctx.Done():
		return nil, errors.New("context canceled")
	case <-d.waitChanP2PGetPeers:
		// return the invoked methods response
		return d.handlerP2PGetPeers.HandleP2PGetPeers(ctx, r)
	}
}

// RegisterDiscoveryGetPeers will register the object 't' as the service
// handler for the RPC method GetPeers from service Discovery
func (d *inboundRPCDispatch) RegisterDiscoveryGetPeers(t DiscoveryGetPeersHandler) {
	d.Lock()
	defer d.Unlock()
	// double registration is not allowed
	if d.handlerDiscoveryGetPeers != nil {
		panic("double registration of DiscoveryGetPeers")
	}
	// register the service handler
	d.handlerDiscoveryGetPeers = t
	// close the wait channel to signal that the method is ready to use
	close(d.waitChanDiscoveryGetPeers)
}

// DiscoveryGetPeers will invoke the handler for the RPC method
// GetPeers from service Discovery
func (d *inboundRPCDispatch) DiscoveryGetPeers(ctx context.Context, r *GetPeersRequest) (*GetPeersResponse, error) {
	// wait for registration to complete or context to be canceled
	select {
	case <-ctx.Done():
		return nil, errors.New("context canceled")
	case <-d.waitChanDiscoveryGetPeers:
		// return the invoked methods response
		return d.handlerDiscoveryGetPeers.HandleDiscoveryGetPeers(ctx, r)
	}
}

// NewInboundRPCDispatch will construct a new inboundRPCDispatcher with all fields properly
// initialized.
func NewInboundRPCDispatch() *inboundRPCDispatch {
	return &inboundRPCDispatch{ 
		// initialize the wait channel for method Status on service P2P
		waitChanP2PStatus: make(chan struct{}),
		// initialize the wait channel for method GetBlockHeaders on service P2P
		waitChanP2PGetBlockHeaders: make(chan struct{}),
		// initialize the wait channel for method GetMinedTxs on service P2P
		waitChanP2PGetMinedTxs: make(chan struct{}),
		// initialize the wait channel for method GetPendingTxs on service P2P
		waitChanP2PGetPendingTxs: make(chan struct{}),
		// initialize the wait channel for method GetSnapShotNode on service P2P
		waitChanP2PGetSnapShotNode: make(chan struct{}),
		// initialize the wait channel for method GetSnapShotStateData on service P2P
		waitChanP2PGetSnapShotStateData: make(chan struct{}),
		// initialize the wait channel for method GetSnapShotHdrNode on service P2P
		waitChanP2PGetSnapShotHdrNode: make(chan struct{}),
		// initialize the wait channel for method GossipTransaction on service P2P
		waitChanP2PGossipTransaction: make(chan struct{}),
		// initialize the wait channel for method GossipProposal on service P2P
		waitChanP2PGossipProposal: make(chan struct{}),
		// initialize the wait channel for method GossipPreVote on service P2P
		waitChanP2PGossipPreVote: make(chan struct{}),
		// initialize the wait channel for method GossipPreVoteNil on service P2P
		waitChanP2PGossipPreVoteNil: make(chan struct{}),
		// initialize the wait channel for method GossipPreCommit on service P2P
		waitChanP2PGossipPreCommit: make(chan struct{}),
		// initialize the wait channel for method GossipPreCommitNil on service P2P
		waitChanP2PGossipPreCommitNil: make(chan struct{}),
		// initialize the wait channel for method GossipNextRound on service P2P
		waitChanP2PGossipNextRound: make(chan struct{}),
		// initialize the wait channel for method GossipNextHeight on service P2P
		waitChanP2PGossipNextHeight: make(chan struct{}),
		// initialize the wait channel for method GossipBlockHeader on service P2P
		waitChanP2PGossipBlockHeader: make(chan struct{}),
		// initialize the wait channel for method GetPeers on service P2P
		waitChanP2PGetPeers: make(chan struct{}),
		// initialize the wait channel for method GetPeers on service Discovery
		waitChanDiscoveryGetPeers: make(chan struct{}),
	}
}

// GeneratedP2PServer implements the P2P service as a gRPC
// server. GeneratedP2PServer invokes methods on the services
// through the inboundRPCDispatch handlers.
type GeneratedP2PServer struct {
	dispatch *inboundRPCDispatch
}

// Status will invoke the method Status on the RPC service P2P
// using the inboundRPCDispatch handler.
func (s *GeneratedP2PServer) Status(ctx context.Context, r *StatusRequest) (*StatusResponse, error) {
	return s.dispatch.P2PStatus(ctx, r)
}


// GetBlockHeaders will invoke the method GetBlockHeaders on the RPC service P2P
// using the inboundRPCDispatch handler.
func (s *GeneratedP2PServer) GetBlockHeaders(ctx context.Context, r *GetBlockHeadersRequest) (*GetBlockHeadersResponse, error) {
	return s.dispatch.P2PGetBlockHeaders(ctx, r)
}


// GetMinedTxs will invoke the method GetMinedTxs on the RPC service P2P
// using the inboundRPCDispatch handler.
func (s *GeneratedP2PServer) GetMinedTxs(ctx context.Context, r *GetMinedTxsRequest) (*GetMinedTxsResponse, error) {
	return s.dispatch.P2PGetMinedTxs(ctx, r)
}


// GetPendingTxs will invoke the method GetPendingTxs on the RPC service P2P
// using the inboundRPCDispatch handler.
func (s *GeneratedP2PServer) GetPendingTxs(ctx context.Context, r *GetPendingTxsRequest) (*GetPendingTxsResponse, error) {
	return s.dispatch.P2PGetPendingTxs(ctx, r)
}


// GetSnapShotNode will invoke the method GetSnapShotNode on the RPC service P2P
// using the inboundRPCDispatch handler.
func (s *GeneratedP2PServer) GetSnapShotNode(ctx context.Context, r *GetSnapShotNodeRequest) (*GetSnapShotNodeResponse, error) {
	return s.dispatch.P2PGetSnapShotNode(ctx, r)
}


// GetSnapShotStateData will invoke the method GetSnapShotStateData on the RPC service P2P
// using the inboundRPCDispatch handler.
func (s *GeneratedP2PServer) GetSnapShotStateData(ctx context.Context, r *GetSnapShotStateDataRequest) (*GetSnapShotStateDataResponse, error) {
	return s.dispatch.P2PGetSnapShotStateData(ctx, r)
}


// GetSnapShotHdrNode will invoke the method GetSnapShotHdrNode on the RPC service P2P
// using the inboundRPCDispatch handler.
func (s *GeneratedP2PServer) GetSnapShotHdrNode(ctx context.Context, r *GetSnapShotHdrNodeRequest) (*GetSnapShotHdrNodeResponse, error) {
	return s.dispatch.P2PGetSnapShotHdrNode(ctx, r)
}


// GossipTransaction will invoke the method GossipTransaction on the RPC service P2P
// using the inboundRPCDispatch handler.
func (s *GeneratedP2PServer) GossipTransaction(ctx context.Context, r *GossipTransactionMessage) (*GossipTransactionAck, error) {
	return s.dispatch.P2PGossipTransaction(ctx, r)
}


// GossipProposal will invoke the method GossipProposal on the RPC service P2P
// using the inboundRPCDispatch handler.
func (s *GeneratedP2PServer) GossipProposal(ctx context.Context, r *GossipProposalMessage) (*GossipProposalAck, error) {
	return s.dispatch.P2PGossipProposal(ctx, r)
}


// GossipPreVote will invoke the method GossipPreVote on the RPC service P2P
// using the inboundRPCDispatch handler.
func (s *GeneratedP2PServer) GossipPreVote(ctx context.Context, r *GossipPreVoteMessage) (*GossipPreVoteAck, error) {
	return s.dispatch.P2PGossipPreVote(ctx, r)
}


// GossipPreVoteNil will invoke the method GossipPreVoteNil on the RPC service P2P
// using the inboundRPCDispatch handler.
func (s *GeneratedP2PServer) GossipPreVoteNil(ctx context.Context, r *GossipPreVoteNilMessage) (*GossipPreVoteNilAck, error) {
	return s.dispatch.P2PGossipPreVoteNil(ctx, r)
}


// GossipPreCommit will invoke the method GossipPreCommit on the RPC service P2P
// using the inboundRPCDispatch handler.
func (s *GeneratedP2PServer) GossipPreCommit(ctx context.Context, r *GossipPreCommitMessage) (*GossipPreCommitAck, error) {
	return s.dispatch.P2PGossipPreCommit(ctx, r)
}


// GossipPreCommitNil will invoke the method GossipPreCommitNil on the RPC service P2P
// using the inboundRPCDispatch handler.
func (s *GeneratedP2PServer) GossipPreCommitNil(ctx context.Context, r *GossipPreCommitNilMessage) (*GossipPreCommitNilAck, error) {
	return s.dispatch.P2PGossipPreCommitNil(ctx, r)
}


// GossipNextRound will invoke the method GossipNextRound on the RPC service P2P
// using the inboundRPCDispatch handler.
func (s *GeneratedP2PServer) GossipNextRound(ctx context.Context, r *GossipNextRoundMessage) (*GossipNextRoundAck, error) {
	return s.dispatch.P2PGossipNextRound(ctx, r)
}


// GossipNextHeight will invoke the method GossipNextHeight on the RPC service P2P
// using the inboundRPCDispatch handler.
func (s *GeneratedP2PServer) GossipNextHeight(ctx context.Context, r *GossipNextHeightMessage) (*GossipNextHeightAck, error) {
	return s.dispatch.P2PGossipNextHeight(ctx, r)
}


// GossipBlockHeader will invoke the method GossipBlockHeader on the RPC service P2P
// using the inboundRPCDispatch handler.
func (s *GeneratedP2PServer) GossipBlockHeader(ctx context.Context, r *GossipBlockHeaderMessage) (*GossipBlockHeaderAck, error) {
	return s.dispatch.P2PGossipBlockHeader(ctx, r)
}


// GetPeers will invoke the method GetPeers on the RPC service P2P
// using the inboundRPCDispatch handler.
func (s *GeneratedP2PServer) GetPeers(ctx context.Context, r *GetPeersRequest) (*GetPeersResponse, error) {
	return s.dispatch.P2PGetPeers(ctx, r)
}



// NewGeneratedP2PServer constructs a new server for the service.
func NewGeneratedP2PServer(dispatch *inboundRPCDispatch) *GeneratedP2PServer {
  return &GeneratedP2PServer{
    dispatch: dispatch,
  }
}


// GeneratedDiscoveryServer implements the Discovery service as a gRPC
// server. GeneratedDiscoveryServer invokes methods on the services
// through the inboundRPCDispatch handlers.
type GeneratedDiscoveryServer struct {
	dispatch *inboundRPCDispatch
}

// GetPeers will invoke the method GetPeers on the RPC service Discovery
// using the inboundRPCDispatch handler.
func (s *GeneratedDiscoveryServer) GetPeers(ctx context.Context, r *GetPeersRequest) (*GetPeersResponse, error) {
	return s.dispatch.DiscoveryGetPeers(ctx, r)
}



// NewGeneratedDiscoveryServer constructs a new server for the service.
func NewGeneratedDiscoveryServer(dispatch *inboundRPCDispatch) *GeneratedDiscoveryServer {
  return &GeneratedDiscoveryServer{
    dispatch: dispatch,
  }
}
