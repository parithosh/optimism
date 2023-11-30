// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"

	"github.com/ethereum-optimism/optimism/op-bindings/solc"
)

const OutputBisectionGameStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"src/dispute/OutputBisectionGame.sol:OutputBisectionGame\",\"label\":\"createdAt\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_userDefinedValueType(Timestamp)1017\"},{\"astId\":1001,\"contract\":\"src/dispute/OutputBisectionGame.sol:OutputBisectionGame\",\"label\":\"resolvedAt\",\"offset\":8,\"slot\":\"0\",\"type\":\"t_userDefinedValueType(Timestamp)1017\"},{\"astId\":1002,\"contract\":\"src/dispute/OutputBisectionGame.sol:OutputBisectionGame\",\"label\":\"status\",\"offset\":16,\"slot\":\"0\",\"type\":\"t_enum(GameStatus)1010\"},{\"astId\":1003,\"contract\":\"src/dispute/OutputBisectionGame.sol:OutputBisectionGame\",\"label\":\"bondManager\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_contract(IBondManager)1009\"},{\"astId\":1004,\"contract\":\"src/dispute/OutputBisectionGame.sol:OutputBisectionGame\",\"label\":\"l1Head\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_userDefinedValueType(Hash)1015\"},{\"astId\":1005,\"contract\":\"src/dispute/OutputBisectionGame.sol:OutputBisectionGame\",\"label\":\"claimData\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_array(t_struct(ClaimData)1011_storage)dyn_storage\"},{\"astId\":1006,\"contract\":\"src/dispute/OutputBisectionGame.sol:OutputBisectionGame\",\"label\":\"claims\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_mapping(t_userDefinedValueType(ClaimHash)1013,t_bool)\"},{\"astId\":1007,\"contract\":\"src/dispute/OutputBisectionGame.sol:OutputBisectionGame\",\"label\":\"subgames\",\"offset\":0,\"slot\":\"5\",\"type\":\"t_mapping(t_uint256,t_array(t_uint256)dyn_storage)\"},{\"astId\":1008,\"contract\":\"src/dispute/OutputBisectionGame.sol:OutputBisectionGame\",\"label\":\"subgameAtRootResolved\",\"offset\":0,\"slot\":\"6\",\"type\":\"t_bool\"}],\"types\":{\"t_array(t_struct(ClaimData)1011_storage)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"struct IOutputBisectionGame.ClaimData[]\",\"numberOfBytes\":\"32\",\"base\":\"t_struct(ClaimData)1011_storage\"},\"t_array(t_uint256)dyn_storage\":{\"encoding\":\"dynamic_array\",\"label\":\"uint256[]\",\"numberOfBytes\":\"32\",\"base\":\"t_uint256\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_contract(IBondManager)1009\":{\"encoding\":\"inplace\",\"label\":\"contract IBondManager\",\"numberOfBytes\":\"20\"},\"t_enum(GameStatus)1010\":{\"encoding\":\"inplace\",\"label\":\"enum GameStatus\",\"numberOfBytes\":\"1\"},\"t_mapping(t_uint256,t_array(t_uint256)dyn_storage)\":{\"encoding\":\"mapping\",\"label\":\"mapping(uint256 =\u003e uint256[])\",\"numberOfBytes\":\"32\",\"key\":\"t_uint256\",\"value\":\"t_array(t_uint256)dyn_storage\"},\"t_mapping(t_userDefinedValueType(ClaimHash)1013,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(ClaimHash =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_userDefinedValueType(ClaimHash)1013\",\"value\":\"t_bool\"},\"t_struct(ClaimData)1011_storage\":{\"encoding\":\"inplace\",\"label\":\"struct IOutputBisectionGame.ClaimData\",\"numberOfBytes\":\"96\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"},\"t_uint32\":{\"encoding\":\"inplace\",\"label\":\"uint32\",\"numberOfBytes\":\"4\"},\"t_userDefinedValueType(Claim)1012\":{\"encoding\":\"inplace\",\"label\":\"Claim\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(ClaimHash)1013\":{\"encoding\":\"inplace\",\"label\":\"ClaimHash\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(Clock)1014\":{\"encoding\":\"inplace\",\"label\":\"Clock\",\"numberOfBytes\":\"16\"},\"t_userDefinedValueType(Hash)1015\":{\"encoding\":\"inplace\",\"label\":\"Hash\",\"numberOfBytes\":\"32\"},\"t_userDefinedValueType(Position)1016\":{\"encoding\":\"inplace\",\"label\":\"Position\",\"numberOfBytes\":\"16\"},\"t_userDefinedValueType(Timestamp)1017\":{\"encoding\":\"inplace\",\"label\":\"Timestamp\",\"numberOfBytes\":\"8\"}}}"

var OutputBisectionGameStorageLayout = new(solc.StorageLayout)

var OutputBisectionGameDeployedBin = "0x6080604052600436106101a15760003560e01c80638980e0cc116100e1578063c55cd0c71161008a578063d8cc1a3c11610064578063d8cc1a3c146105cb578063f8f43ff6146105eb578063fa24f7431461060b578063fdffbb281461062f57600080fd5b8063c55cd0c714610533578063c6f0308c14610546578063cf09e0d0146105aa57600080fd5b8063bbdc02db116100bb578063bbdc02db14610484578063bcef3b55146104c2578063c31b29ce146104ff57600080fd5b80638980e0cc146103fb5780638b85902b14610410578063929312981461045057600080fd5b8063363cc4271161014e578063609d333411610128578063609d3334146103a8578063632247ea146103bd5780636361506d146103d05780638129fc1c146103e657600080fd5b8063363cc427146102cc5780634778efe81461031e57806354fd4d501461035257600080fd5b8063266198f91161017f578063266198f91461026e5780632810e1d6146102a257806335fef567146102b757600080fd5b806319effeb4146101a6578063200d2ed2146101f157806324185bc61461022c575b600080fd5b3480156101b257600080fd5b506000546101d39068010000000000000000900467ffffffffffffffff1681565b60405167ffffffffffffffff90911681526020015b60405180910390f35b3480156101fd57600080fd5b5060005461021f90700100000000000000000000000000000000900460ff1681565b6040516101e89190612561565b34801561023857600080fd5b506102607f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016101e8565b34801561027a57600080fd5b506102607f000000000000000000000000000000000000000000000000000000000000000081565b3480156102ae57600080fd5b5061021f610642565b6102ca6102c53660046125a2565b610813565b005b3480156102d857600080fd5b506001546102f99073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101e8565b34801561032a57600080fd5b506102607f000000000000000000000000000000000000000000000000000000000000000081565b34801561035e57600080fd5b5061039b6040518060400160405280600681526020017f302e302e3133000000000000000000000000000000000000000000000000000081525081565b6040516101e8919061262f565b3480156103b457600080fd5b5061039b610823565b6102ca6103cb36600461265e565b610835565b3480156103dc57600080fd5b5061026060025481565b3480156103f257600080fd5b506102ca610f2e565b34801561040757600080fd5b50600354610260565b34801561041c57600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c900360200135610260565b34801561045c57600080fd5b506102f97f000000000000000000000000000000000000000000000000000000000000000081565b34801561049057600080fd5b5060405160ff7f00000000000000000000000000000000000000000000000000000000000000001681526020016101e8565b3480156104ce57600080fd5b50367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c900335610260565b34801561050b57600080fd5b506101d37f000000000000000000000000000000000000000000000000000000000000000081565b6102ca6105413660046125a2565b61107f565b34801561055257600080fd5b50610566610561366004612693565b61108b565b6040805163ffffffff90961686529315156020860152928401919091526fffffffffffffffffffffffffffffffff908116606084015216608082015260a0016101e8565b3480156105b657600080fd5b506000546101d39067ffffffffffffffff1681565b3480156105d757600080fd5b506102ca6105e63660046126f5565b6110fc565b3480156105f757600080fd5b506102ca61060636600461277f565b611698565b34801561061757600080fd5b50610620611b3b565b6040516101e8939291906127ab565b6102ca61063d366004612693565b611b98565b600080600054700100000000000000000000000000000000900460ff16600281111561067057610670612532565b146106a7576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60065460ff166106e3576040517f9a07664600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b60036000815481106106f7576106f76127d6565b6000918252602090912060039091020154640100000000900460ff1661071e576002610721565b60015b6000805467ffffffffffffffff421668010000000000000000027fffffffffffffffffffffffffffffffff0000000000000000ffffffffffffffff82168117835592935083927fffffffffffffffffffffffffffffff00ffffffffffffffffffffffffffffffff167fffffffffffffffffffffffffffffff000000000000000000ffffffffffffffff909116177001000000000000000000000000000000008360028111156107d2576107d2612532565b0217905560028111156107e7576107e7612532565b6040517f5e186f09b9c93491f14e277eea7faa5de6a2d4bda75a79af7a3684fbfb42da6090600090a290565b61081f82826000610835565b5050565b6060610830602080611ed2565b905090565b60008054700100000000000000000000000000000000900460ff16600281111561086157610861612532565b14610898576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b821580156108a4575080155b156108db576040517fa42637bc00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000600384815481106108f0576108f06127d6565b600091825260208083206040805160a081018252600394909402909101805463ffffffff808216865264010000000090910460ff16151593850193909352600181015491840191909152600201546fffffffffffffffffffffffffffffffff8082166060850181905270010000000000000000000000000000000090920416608084015291935061098491908590611f6916565b90507f0000000000000000000000000000000000000000000000000000000000000000610a43826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff161115610a85576040517f56f57b2b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b610ab07f00000000000000000000000000000000000000000000000000000000000000006001612834565b610b4c826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff1603610b6557610b658486611f71565b815160009063ffffffff90811614610bc5576003836000015163ffffffff1681548110610b9457610b946127d6565b906000526020600020906003020160020160109054906101000a90046fffffffffffffffffffffffffffffffff1690505b608083015160009067ffffffffffffffff1667ffffffffffffffff1642610bfe846fffffffffffffffffffffffffffffffff1660401c90565b67ffffffffffffffff16610c129190612834565b610c1c919061284c565b9050677fffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000060011c1667ffffffffffffffff82161115610c8f576040517f3381d11400000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000604082901b42176000888152608086901b6fffffffffffffffffffffffffffffffff8b1617602052604081209192509060008181526004602052604090205490915060ff1615610d0d576040517f80497e3b00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600081815260046020908152604080832080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001908117909155815160a08101835263ffffffff808f1682529381018581529281018d81526fffffffffffffffffffffffffffffffff808c1660608401908152898216608085019081526003805480880182559981905294519885027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b8101805498511515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff00000000009099169a909916999099179690961790965590517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c8701559351925184167001000000000000000000000000000000000292909316919091177fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85d9093019290925580548b908110610e8357610e836127d6565b600091825260208083206003928302018054941515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff909516949094179093558b82526005909252604090209054610ee49060019061284c565b8154600181018355600092835260208320015560405133918a918c917f9b3245740ec3b155098a55be84957a4da13eaf7f14a8bc6f53126c0b9350f2be91a4505050505050505050565b600080547fffffffffffffffffffffffffffffffffffffffffffffffff0000000000000000164267ffffffffffffffff161781556040805160a08101825263ffffffff81526020810192909252600391908101610fb37ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe369081013560f01c90033590565b815260016020820152604001426fffffffffffffffffffffffffffffffff908116909152825460018181018555600094855260209485902084516003909302018054958501511515640100000000027fffffffffffffffffffffffffffffffffffffffffffffffffffffff000000000090961663ffffffff9093169290921794909417815560408301518185015560608301516080909301518216700100000000000000000000000000000000029290911691909117600290910155611079904361284c565b40600255565b61081f82826001610835565b6003818154811061109b57600080fd5b600091825260209091206003909102018054600182015460029092015463ffffffff8216935064010000000090910460ff1691906fffffffffffffffffffffffffffffffff8082169170010000000000000000000000000000000090041685565b60008054700100000000000000000000000000000000900460ff16600281111561112857611128612532565b1461115f576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060038781548110611174576111746127d6565b6000918252602082206003919091020160028101549092506fffffffffffffffffffffffffffffffff16908715821760011b90506111d37f00000000000000000000000000000000000000000000000000000000000000006001612834565b61126f826fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff16146112b0576040517f5f53dd9800000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000808915611333576112d4836fffffffffffffffffffffffffffffffff16611fca565b67ffffffffffffffff1615611307576112fe6112f1600186612863565b865463ffffffff16612070565b60010154611329565b7f00000000000000000000000000000000000000000000000000000000000000005b915084905061134d565b8460010154915061134a8460016112f19190612894565b90505b600882901b60088a8a6040516113649291906128c8565b6040518091039020901b146113a5576040517f696550ff00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000806113b18d61210d565b600182810154835491830154835460408051602081019490945263ffffffff948516908401526060830191909152919091166080820152919350915060009060a001604051602081830303815290604052805190602001209050600084600101547f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663e14ced328f8f8f8f886040518663ffffffff1660e01b8152600401611473959493929190612921565b6020604051808303816000875af1158015611492573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114b6919061295b565b600287810154929091149250600091611561906fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b6115fd8b6fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b6116079190612974565b6116119190612995565b67ffffffffffffffff161590508115158103611659576040517ffb4e40dd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b505087547fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff166401000000001790975550505050505050505050505050565b60008054700100000000000000000000000000000000900460ff1660028111156116c4576116c4612532565b146116fb576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000806117078461210d565b600182810154835491830154835460408051602081019490945263ffffffff948516908401526060830191909152919091166080820152919350915060009060a00160405160208183030381529060405280519060200120905060007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16637dc0d1d06040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117ce573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906117f291906129e3565b9050866001036118b4576002546040517f52f0f3ad00000000000000000000000000000000000000000000000000000000815260048101899052602481018490526044810191909152602060648201526084810186905273ffffffffffffffffffffffffffffffffffffffff8216906352f0f3ad9060a4015b6020604051808303816000875af115801561188a573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118ae919061295b565b50611b32565b866002036119315760018401546040517f52f0f3ad00000000000000000000000000000000000000000000000000000000815260048101899052602481018490526044810191909152602060648201526084810186905273ffffffffffffffffffffffffffffffffffffffff8216906352f0f3ad9060a40161186b565b866003036119ae5760018301546040517f52f0f3ad00000000000000000000000000000000000000000000000000000000815260048101899052602481018490526044810191909152602060648201526084810186905273ffffffffffffffffffffffffffffffffffffffff8216906352f0f3ad9060a40161186b565b86600403611a8757600284015473ffffffffffffffffffffffffffffffffffffffff8216906352f0f3ad908990859060c0906119fb906fffffffffffffffffffffffffffffffff16611fca565b611a2f9067ffffffffffffffff167f0000000000000000000000000000000000000000000000000000000000000000612834565b6040517fffffffff0000000000000000000000000000000000000000000000000000000060e087901b168152600481019490945260248401929092521b6044820152600860648201526084810188905260a40161186b565b86600503611b00576040517f52f0f3ad00000000000000000000000000000000000000000000000000000000815260048101889052602481018390524660c01b6044820152600860648201526084810186905273ffffffffffffffffffffffffffffffffffffffff8216906352f0f3ad9060a40161186b565b6040517fff137e6500000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505050505050565b7f0000000000000000000000000000000000000000000000000000000000000000367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c9003356060611b91610823565b9050909192565b60008054700100000000000000000000000000000000900460ff166002811115611bc457611bc4612532565b14611bfb576040517f67fe195000000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060038281548110611c1057611c106127d6565b60009182526020909120600260039092020190810154909150677fffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000060011c1690611c8090700100000000000000000000000000000000900467ffffffffffffffff164261284c565b6002830154611cb09190700100000000000000000000000000000000900460401c67ffffffffffffffff16612834565b11611ce7576040517ff2440b5300000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600082815260056020526040902082158015611d05575060065460ff165b15611d3c576040517ff1a9458100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b8054158015611d4a57508215155b15611d81576040517ff1a9458100000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000805b8254811015611e4f576000838281548110611da257611da26127d6565b6000918252602080832090910154808352600590915260409091205490915015611df8576040517f9a07664600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b600060038281548110611e0d57611e0d6127d6565b600091825260209091206003909102018054909150640100000000900460ff16611e3c57600193505050611e4f565b505080611e4890612a19565b9050611d85565b5082547fffffffffffffffffffffffffffffffffffffffffffffffffffffff00ffffffff16640100000000821515021783556000848152600560205260408120611e98916124f8565b83600003611ecc57600680547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001660011790555b50505050565b60606000611f0984367ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe81013560f01c9003612834565b90508267ffffffffffffffff1667ffffffffffffffff811115611f2e57611f2e612a51565b6040519080825280601f01601f191660200182016040528015611f58576020820181803683370190505b509150828160208401375092915050565b151760011b90565b600082901a6001811480611f88575060ff81166002145b611fc5576040517ff40239db0000000000000000000000000000000000000000000000000000000081526004810184905260240160405180910390fd5b505050565b600080612057837e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b600167ffffffffffffffff919091161b90920392915050565b60008061208e846fffffffffffffffffffffffffffffffff1661244c565b9050600383815481106120a3576120a36127d6565b906000526020600020906003020191505b60028201546fffffffffffffffffffffffffffffffff82811691161461210657815460038054909163ffffffff169081106120f1576120f16127d6565b906000526020600020906003020191506120b4565b5092915050565b600080600083905060006003828154811061212a5761212a6127d6565b600091825260209091206002600390920201908101549091507f000000000000000000000000000000000000000000000000000000000000000090612201906fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff1611612242576040517fb34b5c2200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6000815b60028301547f000000000000000000000000000000000000000000000000000000000000000090612309906fffffffffffffffffffffffffffffffff167e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff169250821461238457825463ffffffff1661234e7f00000000000000000000000000000000000000000000000000000000000000006001612834565b8303612358578391505b6003818154811061236b5761236b6127d6565b9060005260206000209060030201935080945050612246565b600280820154908401546fffffffffffffffffffffffffffffffff91821691166000816123b18460011c90565b6fffffffffffffffffffffffffffffffff16149050801561240a576123e06123da600184612863565b88612070565b9850600387815481106123f5576123f56127d6565b90600052602060002090600302019750612440565b6003878154811061241d5761241d6127d6565b9060005260206000209060030201985061243d8260016123da9190612894565b97505b50505050505050915091565b600081196001830116816124e0827e09010a0d15021d0b0e10121619031e080c141c0f111807131b17061a05041f7f07c4acdd0000000000000000000000000000000000000000000000000000000067ffffffffffffffff831160061b83811c63ffffffff1060051b1792831c600181901c17600281901c17600481901c17600881901c17601081901c170260fb1c1a1790565b67ffffffffffffffff169390931c8015179392505050565b50805460008255906000526020600020908101906125169190612519565b50565b5b8082111561252e576000815560010161251a565b5090565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b602081016003831061259c577f4e487b7100000000000000000000000000000000000000000000000000000000600052602160045260246000fd5b91905290565b600080604083850312156125b557600080fd5b50508035926020909101359150565b6000815180845260005b818110156125ea576020818501810151868301820152016125ce565b818111156125fc576000602083870101525b50601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061264260208301846125c4565b9392505050565b8035801515811461265957600080fd5b919050565b60008060006060848603121561267357600080fd5b833592506020840135915061268a60408501612649565b90509250925092565b6000602082840312156126a557600080fd5b5035919050565b60008083601f8401126126be57600080fd5b50813567ffffffffffffffff8111156126d657600080fd5b6020830191508360208285010111156126ee57600080fd5b9250929050565b6000806000806000806080878903121561270e57600080fd5b8635955061271e60208801612649565b9450604087013567ffffffffffffffff8082111561273b57600080fd5b6127478a838b016126ac565b9096509450606089013591508082111561276057600080fd5b5061276d89828a016126ac565b979a9699509497509295939492505050565b60008060006060848603121561279457600080fd5b505081359360208301359350604090920135919050565b60ff841681528260208201526060604082015260006127cd60608301846125c4565b95945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000821982111561284757612847612805565b500190565b60008282101561285e5761285e612805565b500390565b60006fffffffffffffffffffffffffffffffff8381169083168181101561288c5761288c612805565b039392505050565b60006fffffffffffffffffffffffffffffffff8083168185168083038211156128bf576128bf612805565b01949350505050565b8183823760009101908152919050565b8183528181602085013750600060208284010152600060207fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0601f840116840101905092915050565b6060815260006129356060830187896128d8565b82810360208401526129488186886128d8565b9150508260408301529695505050505050565b60006020828403121561296d57600080fd5b5051919050565b600067ffffffffffffffff8381169083168181101561288c5761288c612805565b600067ffffffffffffffff808416806129d7577f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b92169190910692915050565b6000602082840312156129f557600080fd5b815173ffffffffffffffffffffffffffffffffffffffff8116811461264257600080fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203612a4a57612a4a612805565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(OutputBisectionGameStorageLayoutJSON), OutputBisectionGameStorageLayout); err != nil {
		panic(err)
	}

	layouts["OutputBisectionGame"] = OutputBisectionGameStorageLayout
	deployedBytecodes["OutputBisectionGame"] = OutputBisectionGameDeployedBin
}