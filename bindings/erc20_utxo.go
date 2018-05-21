// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ERC20UTXO

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

// ERC20UTXOABI is the input ABI used to generate the binding from.
const ERC20UTXOABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"create\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"utxos\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"},{\"name\":\"createdBy\",\"type\":\"bytes32\"},{\"name\":\"id\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"}],\"name\":\"getUtxo\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_id\",\"type\":\"bytes32\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"spend\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"LogCreate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"oldId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"newId\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"LogSpend\",\"type\":\"event\"}]"

// ERC20UTXOBin is the compiled bytecode used for deploying new contracts.
const ERC20UTXOBin = `608060405234801561001057600080fd5b506105e1806100206000396000f3006080604052600436106100615763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630ecaea7381146100665780633cf18ab91461008c578063c972949b146100d4578063de86ff1314610114575b600080fd5b34801561007257600080fd5b5061008a600160a060020a036004351660243561013b565b005b34801561009857600080fd5b506100a4600435610231565b60408051600160a060020a0390951685526020850193909352838301919091526060830152519081900360800190f35b3480156100e057600080fd5b506100ec600435610262565b60408051600160a060020a039094168452602084019290925282820152519081900360600190f35b34801561012057600080fd5b5061008a600435602435600160a060020a03604435166102c1565b600061014561058e565b5050604080514381526c01000000000000000000000000338102602080840191909152600160a060020a03808716928302603485015284519384900360480184206080850186528385528483018781526000868801818152606088018481528483528287529189902088518154961673ffffffffffffffffffffffffffffffffffffffff1990961695909517855591516001858101919091559151600285015551600390930192909255815487019091558451868152945190948593927fa1116f2f79e89469171f448796729e5d361ce252f2757cda0497962f9ac6338492918290030190a350505050565b6000602081905290815260409020805460018201546002830154600390930154600160a060020a0390921692909184565b600080600061026f61058e565b505050600091825250602081815260409182902082516080810184528154600160a060020a03168082526001830154938201849052600283015494820185905260039092015460609091015292909190565b6102c961058e565b60006102d361058e565b60006102dd61058e565b600088815260208190526040902054600160a060020a0316331461030057600080fd5b60008881526020819052604090206001015487111561031e57600080fd5b60008881526020818152604080832081516080810183528154600160a060020a0381168252600183018054838701526002840180549584019590955260038401805460608501528f88529587905273ffffffffffffffffffffffffffffffffffffffff199091169092559084905590839055919055945061039f8689610549565b60408051608081018252600160a060020a0389811680835260208084018d81528486018f8152606080870189815260008a81528086528990208851815473ffffffffffffffffffffffffffffffffffffffff19169816979097178755925160018701559051600286015590516003909401939093558a83015185519081529081018690528085018d90529351949850919650909233927fa2acb8475c115dd9b6b4f331c3fa77c14342cf09a9ea4167e72bf9a60cd76d1892908290030190a3846020015187101561053f576104773360018a18610549565b6040805160808101825233808252602089810180518d90038285019081528486018f8152606080870189815260008a81528087528990208851815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a03918216178255945160018201559251600284015551600390920191909155808d015192518751938452938301889052928e9003828701529451959750929550928a169390927fa2acb8475c115dd9b6b4f331c3fa77c14342cf09a9ea4167e72bf9a60cd76d18928290030190a35b5050505050505050565b604080514381526c010000000000000000000000003381026020830152600160a060020a03851602603482015260488101839052905190819003606801902092915050565b604080516080810182526000808252602082018190529181018290526060810191909152905600a165627a7a72305820eab6875e24a03f816af00029181b6467f019852b725e76ee91b99a9cfa74229f0029`

// DeployERC20UTXO deploys a new Ethereum contract, binding an instance of ERC20UTXO to it.
func DeployERC20UTXO(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC20UTXO, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20UTXOABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC20UTXOBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC20UTXO{ERC20UTXOCaller: ERC20UTXOCaller{contract: contract}, ERC20UTXOTransactor: ERC20UTXOTransactor{contract: contract}, ERC20UTXOFilterer: ERC20UTXOFilterer{contract: contract}}, nil
}

// ERC20UTXO is an auto generated Go binding around an Ethereum contract.
type ERC20UTXO struct {
	ERC20UTXOCaller     // Read-only binding to the contract
	ERC20UTXOTransactor // Write-only binding to the contract
	ERC20UTXOFilterer   // Log filterer for contract events
}

// ERC20UTXOCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20UTXOCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20UTXOTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20UTXOTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20UTXOFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20UTXOFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20UTXOSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20UTXOSession struct {
	Contract     *ERC20UTXO        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20UTXOCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20UTXOCallerSession struct {
	Contract *ERC20UTXOCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ERC20UTXOTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20UTXOTransactorSession struct {
	Contract     *ERC20UTXOTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ERC20UTXORaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20UTXORaw struct {
	Contract *ERC20UTXO // Generic contract binding to access the raw methods on
}

// ERC20UTXOCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20UTXOCallerRaw struct {
	Contract *ERC20UTXOCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20UTXOTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20UTXOTransactorRaw struct {
	Contract *ERC20UTXOTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20UTXO creates a new instance of ERC20UTXO, bound to a specific deployed contract.
func NewERC20UTXO(address common.Address, backend bind.ContractBackend) (*ERC20UTXO, error) {
	contract, err := bindERC20UTXO(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20UTXO{ERC20UTXOCaller: ERC20UTXOCaller{contract: contract}, ERC20UTXOTransactor: ERC20UTXOTransactor{contract: contract}, ERC20UTXOFilterer: ERC20UTXOFilterer{contract: contract}}, nil
}

// NewERC20UTXOCaller creates a new read-only instance of ERC20UTXO, bound to a specific deployed contract.
func NewERC20UTXOCaller(address common.Address, caller bind.ContractCaller) (*ERC20UTXOCaller, error) {
	contract, err := bindERC20UTXO(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20UTXOCaller{contract: contract}, nil
}

// NewERC20UTXOTransactor creates a new write-only instance of ERC20UTXO, bound to a specific deployed contract.
func NewERC20UTXOTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20UTXOTransactor, error) {
	contract, err := bindERC20UTXO(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20UTXOTransactor{contract: contract}, nil
}

// NewERC20UTXOFilterer creates a new log filterer instance of ERC20UTXO, bound to a specific deployed contract.
func NewERC20UTXOFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20UTXOFilterer, error) {
	contract, err := bindERC20UTXO(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20UTXOFilterer{contract: contract}, nil
}

// bindERC20UTXO binds a generic wrapper to an already deployed contract.
func bindERC20UTXO(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20UTXOABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20UTXO *ERC20UTXORaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20UTXO.Contract.ERC20UTXOCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20UTXO *ERC20UTXORaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20UTXO.Contract.ERC20UTXOTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20UTXO *ERC20UTXORaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20UTXO.Contract.ERC20UTXOTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20UTXO *ERC20UTXOCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20UTXO.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20UTXO *ERC20UTXOTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20UTXO.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20UTXO *ERC20UTXOTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20UTXO.Contract.contract.Transact(opts, method, params...)
}

// GetUtxo is a free data retrieval call binding the contract method 0xc972949b.
//
// Solidity: function getUtxo(_id bytes32) constant returns(address, uint256, bytes32)
func (_ERC20UTXO *ERC20UTXOCaller) GetUtxo(opts *bind.CallOpts, _id [32]byte) (common.Address, *big.Int, [32]byte, error) {
	var (
		ret0 = new(common.Address)
		ret1 = new(*big.Int)
		ret2 = new([32]byte)
	)
	out := &[]interface{}{
		ret0,
		ret1,
		ret2,
	}
	err := _ERC20UTXO.contract.Call(opts, out, "getUtxo", _id)
	return *ret0, *ret1, *ret2, err
}

// GetUtxo is a free data retrieval call binding the contract method 0xc972949b.
//
// Solidity: function getUtxo(_id bytes32) constant returns(address, uint256, bytes32)
func (_ERC20UTXO *ERC20UTXOSession) GetUtxo(_id [32]byte) (common.Address, *big.Int, [32]byte, error) {
	return _ERC20UTXO.Contract.GetUtxo(&_ERC20UTXO.CallOpts, _id)
}

// GetUtxo is a free data retrieval call binding the contract method 0xc972949b.
//
// Solidity: function getUtxo(_id bytes32) constant returns(address, uint256, bytes32)
func (_ERC20UTXO *ERC20UTXOCallerSession) GetUtxo(_id [32]byte) (common.Address, *big.Int, [32]byte, error) {
	return _ERC20UTXO.Contract.GetUtxo(&_ERC20UTXO.CallOpts, _id)
}

// Utxos is a free data retrieval call binding the contract method 0x3cf18ab9.
//
// Solidity: function utxos( bytes32) constant returns(owner address, value uint256, createdBy bytes32, id bytes32)
func (_ERC20UTXO *ERC20UTXOCaller) Utxos(opts *bind.CallOpts, arg0 [32]byte) (struct {
	Owner     common.Address
	Value     *big.Int
	CreatedBy [32]byte
	Id        [32]byte
}, error) {
	ret := new(struct {
		Owner     common.Address
		Value     *big.Int
		CreatedBy [32]byte
		Id        [32]byte
	})
	out := ret
	err := _ERC20UTXO.contract.Call(opts, out, "utxos", arg0)
	return *ret, err
}

// Utxos is a free data retrieval call binding the contract method 0x3cf18ab9.
//
// Solidity: function utxos( bytes32) constant returns(owner address, value uint256, createdBy bytes32, id bytes32)
func (_ERC20UTXO *ERC20UTXOSession) Utxos(arg0 [32]byte) (struct {
	Owner     common.Address
	Value     *big.Int
	CreatedBy [32]byte
	Id        [32]byte
}, error) {
	return _ERC20UTXO.Contract.Utxos(&_ERC20UTXO.CallOpts, arg0)
}

// Utxos is a free data retrieval call binding the contract method 0x3cf18ab9.
//
// Solidity: function utxos( bytes32) constant returns(owner address, value uint256, createdBy bytes32, id bytes32)
func (_ERC20UTXO *ERC20UTXOCallerSession) Utxos(arg0 [32]byte) (struct {
	Owner     common.Address
	Value     *big.Int
	CreatedBy [32]byte
	Id        [32]byte
}, error) {
	return _ERC20UTXO.Contract.Utxos(&_ERC20UTXO.CallOpts, arg0)
}

// Create is a paid mutator transaction binding the contract method 0x0ecaea73.
//
// Solidity: function create(_to address, _value uint256) returns()
func (_ERC20UTXO *ERC20UTXOTransactor) Create(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20UTXO.contract.Transact(opts, "create", _to, _value)
}

// Create is a paid mutator transaction binding the contract method 0x0ecaea73.
//
// Solidity: function create(_to address, _value uint256) returns()
func (_ERC20UTXO *ERC20UTXOSession) Create(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20UTXO.Contract.Create(&_ERC20UTXO.TransactOpts, _to, _value)
}

// Create is a paid mutator transaction binding the contract method 0x0ecaea73.
//
// Solidity: function create(_to address, _value uint256) returns()
func (_ERC20UTXO *ERC20UTXOTransactorSession) Create(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20UTXO.Contract.Create(&_ERC20UTXO.TransactOpts, _to, _value)
}

// Spend is a paid mutator transaction binding the contract method 0xde86ff13.
//
// Solidity: function spend(_id bytes32, _amount uint256, _to address) returns()
func (_ERC20UTXO *ERC20UTXOTransactor) Spend(opts *bind.TransactOpts, _id [32]byte, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _ERC20UTXO.contract.Transact(opts, "spend", _id, _amount, _to)
}

// Spend is a paid mutator transaction binding the contract method 0xde86ff13.
//
// Solidity: function spend(_id bytes32, _amount uint256, _to address) returns()
func (_ERC20UTXO *ERC20UTXOSession) Spend(_id [32]byte, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _ERC20UTXO.Contract.Spend(&_ERC20UTXO.TransactOpts, _id, _amount, _to)
}

// Spend is a paid mutator transaction binding the contract method 0xde86ff13.
//
// Solidity: function spend(_id bytes32, _amount uint256, _to address) returns()
func (_ERC20UTXO *ERC20UTXOTransactorSession) Spend(_id [32]byte, _amount *big.Int, _to common.Address) (*types.Transaction, error) {
	return _ERC20UTXO.Contract.Spend(&_ERC20UTXO.TransactOpts, _id, _amount, _to)
}

// ERC20UTXOLogCreateIterator is returned from FilterLogCreate and is used to iterate over the raw logs and unpacked data for LogCreate events raised by the ERC20UTXO contract.
type ERC20UTXOLogCreateIterator struct {
	Event *ERC20UTXOLogCreate // Event containing the contract specifics and raw log

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
func (it *ERC20UTXOLogCreateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20UTXOLogCreate)
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
		it.Event = new(ERC20UTXOLogCreate)
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
func (it *ERC20UTXOLogCreateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20UTXOLogCreateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20UTXOLogCreate represents a LogCreate event raised by the ERC20UTXO contract.
type ERC20UTXOLogCreate struct {
	Owner common.Address
	Id    [32]byte
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterLogCreate is a free log retrieval operation binding the contract event 0xa1116f2f79e89469171f448796729e5d361ce252f2757cda0497962f9ac63384.
//
// Solidity: e LogCreate(owner indexed address, id indexed bytes32, value uint256)
func (_ERC20UTXO *ERC20UTXOFilterer) FilterLogCreate(opts *bind.FilterOpts, owner []common.Address, id [][32]byte) (*ERC20UTXOLogCreateIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ERC20UTXO.contract.FilterLogs(opts, "LogCreate", ownerRule, idRule)
	if err != nil {
		return nil, err
	}
	return &ERC20UTXOLogCreateIterator{contract: _ERC20UTXO.contract, event: "LogCreate", logs: logs, sub: sub}, nil
}

// WatchLogCreate is a free log subscription operation binding the contract event 0xa1116f2f79e89469171f448796729e5d361ce252f2757cda0497962f9ac63384.
//
// Solidity: e LogCreate(owner indexed address, id indexed bytes32, value uint256)
func (_ERC20UTXO *ERC20UTXOFilterer) WatchLogCreate(opts *bind.WatchOpts, sink chan<- *ERC20UTXOLogCreate, owner []common.Address, id [][32]byte) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ERC20UTXO.contract.WatchLogs(opts, "LogCreate", ownerRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20UTXOLogCreate)
				if err := _ERC20UTXO.contract.UnpackLog(event, "LogCreate", log); err != nil {
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

// ERC20UTXOLogSpendIterator is returned from FilterLogSpend and is used to iterate over the raw logs and unpacked data for LogSpend events raised by the ERC20UTXO contract.
type ERC20UTXOLogSpendIterator struct {
	Event *ERC20UTXOLogSpend // Event containing the contract specifics and raw log

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
func (it *ERC20UTXOLogSpendIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20UTXOLogSpend)
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
		it.Event = new(ERC20UTXOLogSpend)
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
func (it *ERC20UTXOLogSpendIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20UTXOLogSpendIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20UTXOLogSpend represents a LogSpend event raised by the ERC20UTXO contract.
type ERC20UTXOLogSpend struct {
	From     common.Address
	To       common.Address
	OldId    [32]byte
	NewId    [32]byte
	NewValue *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterLogSpend is a free log retrieval operation binding the contract event 0xa2acb8475c115dd9b6b4f331c3fa77c14342cf09a9ea4167e72bf9a60cd76d18.
//
// Solidity: e LogSpend(from indexed address, to indexed address, oldId bytes32, newId bytes32, newValue uint256)
func (_ERC20UTXO *ERC20UTXOFilterer) FilterLogSpend(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20UTXOLogSpendIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20UTXO.contract.FilterLogs(opts, "LogSpend", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20UTXOLogSpendIterator{contract: _ERC20UTXO.contract, event: "LogSpend", logs: logs, sub: sub}, nil
}

// WatchLogSpend is a free log subscription operation binding the contract event 0xa2acb8475c115dd9b6b4f331c3fa77c14342cf09a9ea4167e72bf9a60cd76d18.
//
// Solidity: e LogSpend(from indexed address, to indexed address, oldId bytes32, newId bytes32, newValue uint256)
func (_ERC20UTXO *ERC20UTXOFilterer) WatchLogSpend(opts *bind.WatchOpts, sink chan<- *ERC20UTXOLogSpend, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20UTXO.contract.WatchLogs(opts, "LogSpend", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20UTXOLogSpend)
				if err := _ERC20UTXO.contract.UnpackLog(event, "LogSpend", log); err != nil {
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
