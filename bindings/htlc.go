// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package HTLC

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// HTLCABI is the input ABI used to generate the binding from.
const HTLCABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_preimage\",\"type\":\"bytes32\"}],\"name\":\"complete\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_preimage\",\"type\":\"bytes32\"}],\"name\":\"reclaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_recipient\",\"type\":\"address\"},{\"name\":\"_image\",\"type\":\"bytes32\"},{\"name\":\"_expirationTime\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_expires\",\"type\":\"uint256\"}],\"name\":\"Initiated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Completed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Expired\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"Reclaimed\",\"type\":\"event\"}]"

// HTLCBin is the compiled bytecode used for deploying new contracts.
const HTLCBin = `608060408190526000805460ff191690556060806105238339810160409081528151602083015191909201516000805461010060a860020a031916336101000217815560018054600160a060020a031916600160a060020a038616178155600284905542830160035560048054909160ff199091169083021790555060005460015460035460408051600160a060020a03610100909504851681529290931660208301523482840152606082015290517f37c1773c79ca6b1053e12c2f683cc8c77717acdaaef4fac7e70dd578971d783a9181900360800190a1505050610438806100eb6000396000f30060806040526004361061004b5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166383ccfb84811461005057806396afb3651461006a575b600080fd5b34801561005c57600080fd5b50610068600435610082565b005b34801561007657600080fd5b5061006860043561022a565b6000805460ff161561009357600080fd5b6000805460ff191660011790556002546100ac836103bc565b146100b657600080fd5b60015473ffffffffffffffffffffffffffffffffffffffff1633146100da57600080fd5b600060045460ff1660038111156100ed57fe5b146100f757600080fd5b60035442116101aa57506040513080319133913180156108fc02916000818181858888f19350505050158015610131573d6000803e3d6000fd5b506004805460ff1916600190811790915560005490546040805161010090930473ffffffffffffffffffffffffffffffffffffffff90811684529091166020830152818101839052517fcdb9f4e58f1d1744a76d4003dc0c940c04eb85e7c04b3ce13f230594318c995d916060908290030190a161021c565b6004805460ff191660021790556000546001546040805161010090930473ffffffffffffffffffffffffffffffffffffffff90811684529091166020830152303182820152517f4f250de368f70919105efcd7254a862cc87a72d732d1fa5760792925064f4eb9916060908290030190a15b50506000805460ff19169055565b6000805460ff161561023b57600080fd5b6000805460ff19166001179055600254610254836103bc565b1461025e57600080fd5b600054610100900473ffffffffffffffffffffffffffffffffffffffff16331461028757600080fd5b600260045460ff16600381111561029a57fe5b14806102b65750600060045460ff1660038111156102b457fe5b145b15156102c157600080fd5b600260045460ff1660038111156102d457fe5b14806102fc5750600060045460ff1660038111156102ee57fe5b1480156102fc575060035442115b1561004b57506000805460405130319261010090920473ffffffffffffffffffffffffffffffffffffffff169183156108fc02918491818181858888f1935050505015801561034f573d6000803e3d6000fd5b5060048054600360ff199091161790556000546040805161010090920473ffffffffffffffffffffffffffffffffffffffff1682526020820183905280517fcab347be18c8c37ef504626a027a12ea6725f491b5f5b154778d8d0c54f7841e9281900390910190a161021c565b60408051828152905160009160029160208083019290919081900382018186865af11580156103ef573d6000803e3d6000fd5b5050506040513d602081101561040457600080fd5b5051929150505600a165627a7a723058201fef37879b471865a86daaf54ac0947ac7e5714193b67dd52f8cc7e9b88c6f780029`

// DeployHTLC deploys a new Ethereum contract, binding an instance of HTLC to it.
func DeployHTLC(auth *bind.TransactOpts, backend bind.ContractBackend, _recipient common.Address, _image [32]byte, _expirationTime *big.Int) (common.Address, *types.Transaction, *HTLC, error) {
	parsed, err := abi.JSON(strings.NewReader(HTLCABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HTLCBin), backend, _recipient, _image, _expirationTime)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &HTLC{HTLCCaller: HTLCCaller{contract: contract}, HTLCTransactor: HTLCTransactor{contract: contract}, HTLCFilterer: HTLCFilterer{contract: contract}}, nil
}

// HTLC is an auto generated Go binding around an Ethereum contract.
type HTLC struct {
	HTLCCaller     // Read-only binding to the contract
	HTLCTransactor // Write-only binding to the contract
	HTLCFilterer   // Log filterer for contract events
}

// HTLCCaller is an auto generated read-only Go binding around an Ethereum contract.
type HTLCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HTLCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HTLCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HTLCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HTLCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HTLCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HTLCSession struct {
	Contract     *HTLC             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HTLCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HTLCCallerSession struct {
	Contract *HTLCCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HTLCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HTLCTransactorSession struct {
	Contract     *HTLCTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HTLCRaw is an auto generated low-level Go binding around an Ethereum contract.
type HTLCRaw struct {
	Contract *HTLC // Generic contract binding to access the raw methods on
}

// HTLCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HTLCCallerRaw struct {
	Contract *HTLCCaller // Generic read-only contract binding to access the raw methods on
}

// HTLCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HTLCTransactorRaw struct {
	Contract *HTLCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHTLC creates a new instance of HTLC, bound to a specific deployed contract.
func NewHTLC(address common.Address, backend bind.ContractBackend) (*HTLC, error) {
	contract, err := bindHTLC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &HTLC{HTLCCaller: HTLCCaller{contract: contract}, HTLCTransactor: HTLCTransactor{contract: contract}, HTLCFilterer: HTLCFilterer{contract: contract}}, nil
}

// NewHTLCCaller creates a new read-only instance of HTLC, bound to a specific deployed contract.
func NewHTLCCaller(address common.Address, caller bind.ContractCaller) (*HTLCCaller, error) {
	contract, err := bindHTLC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HTLCCaller{contract: contract}, nil
}

// NewHTLCTransactor creates a new write-only instance of HTLC, bound to a specific deployed contract.
func NewHTLCTransactor(address common.Address, transactor bind.ContractTransactor) (*HTLCTransactor, error) {
	contract, err := bindHTLC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HTLCTransactor{contract: contract}, nil
}

// NewHTLCFilterer creates a new log filterer instance of HTLC, bound to a specific deployed contract.
func NewHTLCFilterer(address common.Address, filterer bind.ContractFilterer) (*HTLCFilterer, error) {
	contract, err := bindHTLC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HTLCFilterer{contract: contract}, nil
}

// bindHTLC binds a generic wrapper to an already deployed contract.
func bindHTLC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HTLCABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HTLC *HTLCRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HTLC.Contract.HTLCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HTLC *HTLCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HTLC.Contract.HTLCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HTLC *HTLCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HTLC.Contract.HTLCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_HTLC *HTLCCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _HTLC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_HTLC *HTLCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _HTLC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_HTLC *HTLCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _HTLC.Contract.contract.Transact(opts, method, params...)
}

// Complete is a paid mutator transaction binding the contract method 0x83ccfb84.
//
// Solidity: function complete(_preimage bytes32) returns()
func (_HTLC *HTLCTransactor) Complete(opts *bind.TransactOpts, _preimage [32]byte) (*types.Transaction, error) {
	return _HTLC.contract.Transact(opts, "complete", _preimage)
}

// Complete is a paid mutator transaction binding the contract method 0x83ccfb84.
//
// Solidity: function complete(_preimage bytes32) returns()
func (_HTLC *HTLCSession) Complete(_preimage [32]byte) (*types.Transaction, error) {
	return _HTLC.Contract.Complete(&_HTLC.TransactOpts, _preimage)
}

// Complete is a paid mutator transaction binding the contract method 0x83ccfb84.
//
// Solidity: function complete(_preimage bytes32) returns()
func (_HTLC *HTLCTransactorSession) Complete(_preimage [32]byte) (*types.Transaction, error) {
	return _HTLC.Contract.Complete(&_HTLC.TransactOpts, _preimage)
}

// Reclaim is a paid mutator transaction binding the contract method 0x96afb365.
//
// Solidity: function reclaim(_preimage bytes32) returns()
func (_HTLC *HTLCTransactor) Reclaim(opts *bind.TransactOpts, _preimage [32]byte) (*types.Transaction, error) {
	return _HTLC.contract.Transact(opts, "reclaim", _preimage)
}

// Reclaim is a paid mutator transaction binding the contract method 0x96afb365.
//
// Solidity: function reclaim(_preimage bytes32) returns()
func (_HTLC *HTLCSession) Reclaim(_preimage [32]byte) (*types.Transaction, error) {
	return _HTLC.Contract.Reclaim(&_HTLC.TransactOpts, _preimage)
}

// Reclaim is a paid mutator transaction binding the contract method 0x96afb365.
//
// Solidity: function reclaim(_preimage bytes32) returns()
func (_HTLC *HTLCTransactorSession) Reclaim(_preimage [32]byte) (*types.Transaction, error) {
	return _HTLC.Contract.Reclaim(&_HTLC.TransactOpts, _preimage)
}

// HTLCCompletedIterator is returned from FilterCompleted and is used to iterate over the raw logs and unpacked data for Completed events raised by the HTLC contract.
type HTLCCompletedIterator struct {
	Event *HTLCCompleted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HTLCCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HTLCCompleted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HTLCCompleted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HTLCCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HTLCCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HTLCCompleted represents a Completed event raised by the HTLC contract.
type HTLCCompleted struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCompleted is a free log retrieval operation binding the contract event 0xcdb9f4e58f1d1744a76d4003dc0c940c04eb85e7c04b3ce13f230594318c995d.
//
// Solidity: e Completed(_sender address, _recipient address, _amount uint256)
func (_HTLC *HTLCFilterer) FilterCompleted(opts *bind.FilterOpts) (*HTLCCompletedIterator, error) {

	logs, sub, err := _HTLC.contract.FilterLogs(opts, "Completed")
	if err != nil {
		return nil, err
	}
	return &HTLCCompletedIterator{contract: _HTLC.contract, event: "Completed", logs: logs, sub: sub}, nil
}

// WatchCompleted is a free log subscription operation binding the contract event 0xcdb9f4e58f1d1744a76d4003dc0c940c04eb85e7c04b3ce13f230594318c995d.
//
// Solidity: e Completed(_sender address, _recipient address, _amount uint256)
func (_HTLC *HTLCFilterer) WatchCompleted(opts *bind.WatchOpts, sink chan<- *HTLCCompleted) (event.Subscription, error) {

	logs, sub, err := _HTLC.contract.WatchLogs(opts, "Completed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HTLCCompleted)
				if err := _HTLC.contract.UnpackLog(event, "Completed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// HTLCExpiredIterator is returned from FilterExpired and is used to iterate over the raw logs and unpacked data for Expired events raised by the HTLC contract.
type HTLCExpiredIterator struct {
	Event *HTLCExpired // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HTLCExpiredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HTLCExpired)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HTLCExpired)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HTLCExpiredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HTLCExpiredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HTLCExpired represents a Expired event raised by the HTLC contract.
type HTLCExpired struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterExpired is a free log retrieval operation binding the contract event 0x4f250de368f70919105efcd7254a862cc87a72d732d1fa5760792925064f4eb9.
//
// Solidity: e Expired(_sender address, _recipient address, _amount uint256)
func (_HTLC *HTLCFilterer) FilterExpired(opts *bind.FilterOpts) (*HTLCExpiredIterator, error) {

	logs, sub, err := _HTLC.contract.FilterLogs(opts, "Expired")
	if err != nil {
		return nil, err
	}
	return &HTLCExpiredIterator{contract: _HTLC.contract, event: "Expired", logs: logs, sub: sub}, nil
}

// WatchExpired is a free log subscription operation binding the contract event 0x4f250de368f70919105efcd7254a862cc87a72d732d1fa5760792925064f4eb9.
//
// Solidity: e Expired(_sender address, _recipient address, _amount uint256)
func (_HTLC *HTLCFilterer) WatchExpired(opts *bind.WatchOpts, sink chan<- *HTLCExpired) (event.Subscription, error) {

	logs, sub, err := _HTLC.contract.WatchLogs(opts, "Expired")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HTLCExpired)
				if err := _HTLC.contract.UnpackLog(event, "Expired", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// HTLCInitiatedIterator is returned from FilterInitiated and is used to iterate over the raw logs and unpacked data for Initiated events raised by the HTLC contract.
type HTLCInitiatedIterator struct {
	Event *HTLCInitiated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HTLCInitiatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HTLCInitiated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HTLCInitiated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HTLCInitiatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HTLCInitiatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HTLCInitiated represents a Initiated event raised by the HTLC contract.
type HTLCInitiated struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Expires   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInitiated is a free log retrieval operation binding the contract event 0x37c1773c79ca6b1053e12c2f683cc8c77717acdaaef4fac7e70dd578971d783a.
//
// Solidity: e Initiated(_sender address, _recipient address, _amount uint256, _expires uint256)
func (_HTLC *HTLCFilterer) FilterInitiated(opts *bind.FilterOpts) (*HTLCInitiatedIterator, error) {

	logs, sub, err := _HTLC.contract.FilterLogs(opts, "Initiated")
	if err != nil {
		return nil, err
	}
	return &HTLCInitiatedIterator{contract: _HTLC.contract, event: "Initiated", logs: logs, sub: sub}, nil
}

// WatchInitiated is a free log subscription operation binding the contract event 0x37c1773c79ca6b1053e12c2f683cc8c77717acdaaef4fac7e70dd578971d783a.
//
// Solidity: e Initiated(_sender address, _recipient address, _amount uint256, _expires uint256)
func (_HTLC *HTLCFilterer) WatchInitiated(opts *bind.WatchOpts, sink chan<- *HTLCInitiated) (event.Subscription, error) {

	logs, sub, err := _HTLC.contract.WatchLogs(opts, "Initiated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HTLCInitiated)
				if err := _HTLC.contract.UnpackLog(event, "Initiated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// HTLCReclaimedIterator is returned from FilterReclaimed and is used to iterate over the raw logs and unpacked data for Reclaimed events raised by the HTLC contract.
type HTLCReclaimedIterator struct {
	Event *HTLCReclaimed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *HTLCReclaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HTLCReclaimed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(HTLCReclaimed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *HTLCReclaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HTLCReclaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HTLCReclaimed represents a Reclaimed event raised by the HTLC contract.
type HTLCReclaimed struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReclaimed is a free log retrieval operation binding the contract event 0xcab347be18c8c37ef504626a027a12ea6725f491b5f5b154778d8d0c54f7841e.
//
// Solidity: e Reclaimed(_sender address, _amount uint256)
func (_HTLC *HTLCFilterer) FilterReclaimed(opts *bind.FilterOpts) (*HTLCReclaimedIterator, error) {

	logs, sub, err := _HTLC.contract.FilterLogs(opts, "Reclaimed")
	if err != nil {
		return nil, err
	}
	return &HTLCReclaimedIterator{contract: _HTLC.contract, event: "Reclaimed", logs: logs, sub: sub}, nil
}

// WatchReclaimed is a free log subscription operation binding the contract event 0xcab347be18c8c37ef504626a027a12ea6725f491b5f5b154778d8d0c54f7841e.
//
// Solidity: e Reclaimed(_sender address, _amount uint256)
func (_HTLC *HTLCFilterer) WatchReclaimed(opts *bind.WatchOpts, sink chan<- *HTLCReclaimed) (event.Subscription, error) {

	logs, sub, err := _HTLC.contract.WatchLogs(opts, "Reclaimed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HTLCReclaimed)
				if err := _HTLC.contract.UnpackLog(event, "Reclaimed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
